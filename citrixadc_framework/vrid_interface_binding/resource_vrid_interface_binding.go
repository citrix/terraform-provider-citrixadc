package vrid_interface_binding

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
var _ resource.Resource = &VridInterfaceBindingResource{}
var _ resource.ResourceWithConfigure = (*VridInterfaceBindingResource)(nil)
var _ resource.ResourceWithImportState = (*VridInterfaceBindingResource)(nil)

func NewVridInterfaceBindingResource() resource.Resource {
	return &VridInterfaceBindingResource{}
}

// VridInterfaceBindingResource defines the resource implementation.
type VridInterfaceBindingResource struct {
	client *service.NitroClient
}

func (r *VridInterfaceBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VridInterfaceBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vrid_interface_binding"
}

func (r *VridInterfaceBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VridInterfaceBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VridInterfaceBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating vrid_interface_binding resource")
	vrid_interface_binding := vrid_interface_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - NITRO add is HTTP PUT (bind), use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Vrid_interface_binding.Type(), &vrid_interface_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create vrid_interface_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created vrid_interface_binding resource")

	// Set ID for the resource before reading state.
	// Composite key: id (the NITRO integer VRID), ifnum.
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("id:%s", utils.UrlEncode(fmt.Sprintf("%v", data.VridId.ValueInt64()))))
	idParts = append(idParts, fmt.Sprintf("ifnum:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Ifnum.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	found := r.readVridInterfaceBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if !found {
		resp.Diagnostics.AddError("Client Error", "vrid_interface_binding not found after create")
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VridInterfaceBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VridInterfaceBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading vrid_interface_binding resource")

	found := r.readVridInterfaceBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	// Binding genuinely absent on the appliance: treat as drift and clear state.
	if !found {
		tflog.Debug(ctx, "vrid_interface_binding not found on appliance; removing from state")
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VridInterfaceBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state VridInterfaceBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// Update is a no-op for vrid_interface_binding: NITRO exposes only add (PUT) and
	// delete (no update/change endpoint), and all schema attributes are
	// RequiresReplace, so Terraform recreates the resource on any change rather
	// than calling Update.
	tflog.Debug(ctx, "Update is a no-op for vrid_interface_binding; all attributes are RequiresReplace")

	// Read the current state back
	r.readVridInterfaceBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VridInterfaceBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VridInterfaceBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting vrid_interface_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs with the parent
	// (the NITRO integer id) as the resource (URL) name and ifnum passed as an arg.
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"id", "ifnum"}, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	id_value, ok := idMap["id"]
	if !ok {
		resp.Diagnostics.AddError("Parse Error", "Parent attribute 'id' not found in ID")
		return
	}

	args := make([]string, 0)
	if val, ok := idMap["ifnum"]; ok && val != "" {
		// Interface ids contain '/' (e.g. 1/3); URL-encode the value.
		args = append(args, fmt.Sprintf("ifnum:%s", utils.UrlEncode(val)))
	}

	err = r.client.DeleteResourceWithArgs(service.Vrid_interface_binding.Type(), id_value, args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete vrid_interface_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted vrid_interface_binding binding")
}

// vrid_interface_bindingAggregateRead queries the AGGREGATE parent endpoint
// (GET /nitro/v1/config/vrid_binding/<id>) and flattens the nested
// "vrid_interface_binding" arrays into a single slice of binding rows. The by-name
// binding endpoint can return a keyless empty body, so the bound members are
// read via the parent aggregate for robustness.
func vrid_interface_bindingAggregateRead(client *service.NitroClient, id string) ([]map[string]interface{}, error) {
	findParams := service.FindParams{
		ResourceType:             "vrid_binding",
		ResourceName:             id,
		ResourceMissingErrorCode: 258,
	}
	parentArr, err := client.FindResourceArrayWithParams(findParams)
	if err != nil {
		return nil, err
	}

	rows := make([]map[string]interface{}, 0)
	for _, parent := range parentArr {
		nested, ok := parent["vrid_interface_binding"]
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

// vrid_interface_bindingRowHasIfnum reports whether an aggregate row's "ifnum"
// matches want. NITRO may represent ifnum as a scalar string or (on some
// firmware) a JSON array; both are handled. Note: on the firmware tested, ifnum
// is not echoed at all -- callers fall back to row presence in that case.
func vrid_interface_bindingRowHasIfnum(raw interface{}, want string) bool {
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

// readVridInterfaceBindingFromApi reads the binding via the AGGREGATE parent
// endpoint (vrid_binding/<id>) and matches the row by ifnum. It returns true when
// the binding is found and the model was populated, false when the binding is
// genuinely absent (drift). Hard errors (parse / transport) are reported via diags.
func (r *VridInterfaceBindingResource) readVridInterfaceBindingFromApi(ctx context.Context, data *VridInterfaceBindingResourceModel, diags *diag.Diagnostics) bool {

	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"id", "ifnum"}, nil)
	if err != nil {
		diags.AddError("Parse Error", fmt.Sprintf("Unable to parse ID: %s", err))
		return false
	}

	id_Name, ok := idMap["id"]
	if !ok {
		diags.AddError("Parse Error", "ID attribute 'id' not found in ID string")
		return false
	}

	dataArr, err := vrid_interface_bindingAggregateRead(r.client, id_Name)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read vrid_interface_binding, got error: %s", err))
		return false
	}

	// Binding genuinely absent (parent missing or no nested rows): report drift.
	if len(dataArr) == 0 {
		return false
	}

	// Iterate through results to find the one with the right ifnum.
	//
	// IMPORTANT (verified live on NS VPX): the aggregate vrid_binding/<id> response
	// represents each bound interface as a vrid_interface_binding row of the form
	// {"id","vlan","flags"} -- it does NOT echo "ifnum". Matching strictly on ifnum
	// therefore never succeeds and the binding appears absent right after create.
	// When the row carries an ifnum, match on it; otherwise (the firmware's actual
	// behavior) fall back to row presence, which is sufficient because the parent
	// vrid id already scopes the result.
	foundIndex := -1
	wantIfnum := idMap["ifnum"]
	for i, v := range dataArr {
		if raw, ok := v["ifnum"]; ok && raw != nil {
			if vrid_interface_bindingRowHasIfnum(raw, wantIfnum) {
				foundIndex = i
				break
			}
			continue
		}
		// ifnum not echoed by the appliance: accept the row by presence.
		foundIndex = i
		break
	}

	// Binding row not present in the aggregate response: drift.
	if foundIndex == -1 {
		return false
	}

	vrid_interface_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
	return true
}
