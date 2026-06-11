package channel_interface_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &ChannelInterfaceBindingResource{}
var _ resource.ResourceWithConfigure = (*ChannelInterfaceBindingResource)(nil)
var _ resource.ResourceWithImportState = (*ChannelInterfaceBindingResource)(nil)

func NewChannelInterfaceBindingResource() resource.Resource {
	return &ChannelInterfaceBindingResource{}
}

// ChannelInterfaceBindingResource defines the resource implementation.
type ChannelInterfaceBindingResource struct {
	client *service.NitroClient
}

func (r *ChannelInterfaceBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *ChannelInterfaceBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_channel_interface_binding"
}

func (r *ChannelInterfaceBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

// channel_interface_bindingComposeId builds the composite resource ID string.
// Format: id:<channel>,ifnum:<interface> (both UrlEncoded; channel and interface
// ids contain '/').
func channel_interface_bindingComposeId(channelid, ifnum string) string {
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("id:%s", utils.UrlEncode(channelid)))
	idParts = append(idParts, fmt.Sprintf("ifnum:%s", utils.UrlEncode(ifnum)))
	return strings.Join(idParts, ",")
}

func (r *ChannelInterfaceBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ChannelInterfaceBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating channel_interface_binding resource")
	channel_interface_binding := channel_interface_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - NITRO add is HTTP PUT (bind), use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Channel_interface_binding.Type(), &channel_interface_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create channel_interface_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created channel_interface_binding resource")

	// Set ID for the resource before reading state. Composite key id:<channel>,ifnum:<intf>.
	var ifnumList []string
	data.Ifnum.ElementsAs(ctx, &ifnumList, false)
	firstIfnum := ""
	if len(ifnumList) > 0 {
		firstIfnum = ifnumList[0]
	}
	data.Id = types.StringValue(channel_interface_bindingComposeId(data.Channelid.ValueString(), firstIfnum))

	// Read the updated state back
	found := r.readChannelInterfaceBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if !found {
		resp.Diagnostics.AddError("Client Error", "channel_interface_binding not found after create")
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ChannelInterfaceBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ChannelInterfaceBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading channel_interface_binding resource")

	found := r.readChannelInterfaceBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	// Binding genuinely absent on the appliance: treat as drift and clear state.
	if !found {
		tflog.Debug(ctx, "channel_interface_binding not found on appliance; removing from state")
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ChannelInterfaceBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state ChannelInterfaceBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// Update is a no-op for channel_interface_binding: NITRO exposes only add (PUT)
	// and delete (no update/change endpoint), and all schema attributes are
	// RequiresReplace, so Terraform recreates the resource on any change rather than
	// calling Update.
	tflog.Debug(ctx, "Update is a no-op for channel_interface_binding; all attributes are RequiresReplace")

	// Read the current state back
	r.readChannelInterfaceBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// channel_interface_bindingAggregateRead queries the AGGREGATE parent endpoint
// (GET /nitro/v1/config/channel_binding/<id>) and flattens the nested
// "channel_interface_binding" arrays into a single slice of binding rows.
//
// On this firmware the direct endpoint
// (GET /nitro/v1/config/channel_interface_binding/<id>) returns a keyless empty
// body, so the bound interfaces are only retrievable via the parent aggregate.
func channel_interface_bindingAggregateRead(client *service.NitroClient, channelid string) ([]map[string]interface{}, error) {
	// The channel id contains a '/' (e.g. "LA/1"). The NITRO client writes
	// ResourceName verbatim into the URL *path*; a path-embedded slash (encoded or
	// not) is intercepted by the front-end web server and never reaches NITRO
	// (Apache 404 / "Invalid query parameters"). The bound interfaces must instead
	// be read with the query-arg form (channel_binding?args=id:LA%2F1), which is what
	// ArgsMap produces. constructQueryMapString writes arg values raw, so the value
	// is URL-encoded here.
	findParams := service.FindParams{
		ResourceType:             "channel_binding",
		ArgsMap:                  map[string]string{"id": utils.UrlEncode(channelid)},
		ResourceMissingErrorCode: 258,
	}
	parentArr, err := client.FindResourceArrayWithParams(findParams)
	if err != nil {
		return nil, err
	}

	rows := make([]map[string]interface{}, 0)
	for _, parent := range parentArr {
		nested, ok := parent["channel_interface_binding"]
		if !ok || nested == nil {
			continue
		}
		nestedArr, ok := nested.([]interface{})
		if !ok {
			continue
		}
		for _, item := range nestedArr {
			if m, ok := item.(map[string]interface{}); ok {
				rows = append(rows, m)
			}
		}
	}
	return rows, nil
}

// channelInterfaceRowHasIfnum reports whether the "ifnum" value from an aggregate
// channel_interface_binding row contains the wanted interface. The NITRO aggregate
// response represents ifnum as a JSON array (e.g. ["1/2"]); this helper also
// tolerates the scalar string form defensively.
func channelInterfaceRowHasIfnum(raw interface{}, want string) bool {
	switch v := raw.(type) {
	case string:
		return v == want
	case []interface{}:
		for _, item := range v {
			if s, ok := item.(string); ok && s == want {
				return true
			}
		}
	}
	return false
}

func (r *ChannelInterfaceBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ChannelInterfaceBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting channel_interface_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs with the parent
	// (channel id) as the resource (URL) name and ifnum passed as the only arg.
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"id", "ifnum"}, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	channelid, ok := idMap["id"]
	if !ok {
		resp.Diagnostics.AddError("Parse Error", "Parent attribute 'id' not found in ID")
		return
	}

	args := make([]string, 0)
	if val, ok := idMap["ifnum"]; ok && val != "" {
		args = append(args, fmt.Sprintf("ifnum:%s", utils.UrlEncode(val)))
	}

	// Pass the RAW channel id (e.g. "LA/1"). DeleteResourceWithArgs / the underlying
	// deleteResourceWithArgs double URL-escapes the resource name itself
	// (QueryEscape(QueryEscape(name)) -> "LA%252F1", which the appliance accepts).
	// Pre-encoding here caused triple-encoding ("LA%25252F1"), which the appliance
	// rejected with errorcode 1074 "Invalid value [interface, LA%2F1]", silently
	// leaving the binding in place.
	err = r.client.DeleteResourceWithArgs(service.Channel_interface_binding.Type(), channelid, args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete channel_interface_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted channel_interface_binding binding")
}

// readChannelInterfaceBindingFromApi reads the binding from the appliance via the
// AGGREGATE parent endpoint (channel_binding/<id>) and matches the row by ifnum.
// It returns true when the binding is found and the model was populated, false when
// the binding is genuinely absent (drift). Hard errors are reported via diags.
func (r *ChannelInterfaceBindingResource) readChannelInterfaceBindingFromApi(ctx context.Context, data *ChannelInterfaceBindingResourceModel, diags *diag.Diagnostics) bool {

	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"id", "ifnum"}, nil)
	if err != nil {
		diags.AddError("Parse Error", fmt.Sprintf("Unable to parse ID: %s", err))
		return false
	}

	channelid, ok := idMap["id"]
	if !ok {
		diags.AddError("Parse Error", "ID attribute 'id' not found in ID string")
		return false
	}

	// The direct channel_interface_binding endpoint returns a keyless empty body on
	// this firmware; read the bound interfaces from the aggregate parent endpoint.
	dataArr, err := channel_interface_bindingAggregateRead(r.client, channelid)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read channel_interface_binding, got error: %s", err))
		return false
	}

	// Binding genuinely absent (parent missing or no nested rows): report drift.
	if len(dataArr) == 0 {
		return false
	}

	// Iterate through results to find the one with the right ifnum.
	//
	// NOTE: the aggregate response represents channel_interface_binding "ifnum" as a
	// JSON ARRAY (e.g. "ifnum": ["1/2"]), not a scalar string. Match against any
	// member of the list (and tolerate the scalar form for safety).
	foundIndex := -1
	for i, v := range dataArr {
		if idVal, ok := idMap["ifnum"]; ok {
			if channelInterfaceRowHasIfnum(v["ifnum"], idVal) {
				foundIndex = i
				break
			}
		}
	}

	// Binding row not present in the aggregate response: drift.
	if foundIndex == -1 {
		return false
	}

	channel_interface_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
	return true
}
