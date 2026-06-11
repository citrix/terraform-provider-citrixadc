package linkset_interface_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &LinksetInterfaceBindingResource{}
var _ resource.ResourceWithConfigure = (*LinksetInterfaceBindingResource)(nil)
var _ resource.ResourceWithImportState = (*LinksetInterfaceBindingResource)(nil)

func NewLinksetInterfaceBindingResource() resource.Resource {
	return &LinksetInterfaceBindingResource{}
}

// LinksetInterfaceBindingResource defines the resource implementation.
type LinksetInterfaceBindingResource struct {
	client *service.NitroClient
}

func (r *LinksetInterfaceBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *LinksetInterfaceBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_linkset_interface_binding"
}

func (r *LinksetInterfaceBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *LinksetInterfaceBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data LinksetInterfaceBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating linkset_interface_binding resource")
	linkset_interface_binding := linkset_interface_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - NITRO add is HTTP PUT (bind), use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Linkset_interface_binding.Type(), &linkset_interface_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create linkset_interface_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created linkset_interface_binding resource")

	// Set ID for the resource before reading state. Composite key id:<linkset>,ifnum:<intf>.
	data.Id = types.StringValue(linkset_interface_bindingComposeId(data.Linksetid.ValueString(), data.Ifnum.ValueString()))

	// Read the updated state back
	found := r.readLinksetInterfaceBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if !found {
		resp.Diagnostics.AddError("Client Error", "linkset_interface_binding not found after create")
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LinksetInterfaceBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data LinksetInterfaceBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading linkset_interface_binding resource")

	found := r.readLinksetInterfaceBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	// Binding genuinely absent on the appliance: treat as drift and clear state.
	if !found {
		tflog.Debug(ctx, "linkset_interface_binding not found on appliance; removing from state")
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LinksetInterfaceBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state LinksetInterfaceBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// Update is a no-op for linkset_interface_binding: NITRO exposes only add (PUT)
	// and delete (no update/change endpoint), and all schema attributes are
	// RequiresReplace, so Terraform recreates the resource on any change rather than
	// calling Update.
	tflog.Debug(ctx, "Update is a no-op for linkset_interface_binding; all attributes are RequiresReplace")

	// Read the current state back
	r.readLinksetInterfaceBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// linkset_interface_bindingAggregateRead queries the AGGREGATE parent endpoint
// (GET /nitro/v1/config/linkset_binding/<id>) and flattens the nested
// "linkset_interface_binding" arrays into a single slice of binding rows.
//
// On this firmware the direct endpoint
// (GET /nitro/v1/config/linkset_interface_binding/<id>) returns a keyless empty
// body, so the bound interfaces are only retrievable via the parent aggregate.
func linkset_interface_bindingAggregateRead(client *service.NitroClient, linksetid string) ([]map[string]interface{}, error) {
	// The linkset id contains a '/' (e.g. "LS/1"). The NITRO client writes
	// ResourceName verbatim into the URL *path*; a path-embedded slash (encoded or
	// not) is intercepted by the front-end web server and never reaches NITRO.
	// Read with the query-arg form (linkset_binding?args=id:LS%2F1) via ArgsMap
	// instead. constructQueryMapString writes arg values raw, so URL-encode here.
	// (Mirrors the fix in channel_interface_binding.)
	findParams := service.FindParams{
		ResourceType:             "linkset_binding",
		ArgsMap:                  map[string]string{"id": utils.UrlEncode(linksetid)},
		ResourceMissingErrorCode: 258,
	}
	parentArr, err := client.FindResourceArrayWithParams(findParams)
	if err != nil {
		return nil, err
	}

	rows := make([]map[string]interface{}, 0)
	for _, parent := range parentArr {
		nested, ok := parent["linkset_interface_binding"]
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

func (r *LinksetInterfaceBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data LinksetInterfaceBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting linkset_interface_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs with the parent
	// (linkset id) as the resource (URL) name and ifnum passed as the only arg.
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"id", "ifnum"}, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	linksetid, ok := idMap["id"]
	if !ok {
		resp.Diagnostics.AddError("Parse Error", "Parent attribute 'id' not found in ID")
		return
	}

	args := make([]string, 0)
	if val, ok := idMap["ifnum"]; ok && val != "" {
		args = append(args, fmt.Sprintf("ifnum:%s", utils.UrlEncode(val)))
	}

	// Pass the RAW linkset id (e.g. "LS/1"). DeleteResourceWithArgs double URL-escapes
	// the resource name itself; pre-encoding here would triple-encode it and the
	// appliance would reject the delete, silently leaving the binding in place.
	// (Mirrors the fix in channel_interface_binding.)
	err = r.client.DeleteResourceWithArgs(service.Linkset_interface_binding.Type(), linksetid, args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete linkset_interface_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted linkset_interface_binding binding")
}

// readLinksetInterfaceBindingFromApi reads the binding from the appliance via the
// AGGREGATE parent endpoint (linkset_binding/<id>) and matches the row by ifnum.
// It returns true when the binding is found and the model was populated, false when
// the binding is genuinely absent (drift). Hard errors are reported via diags.
func (r *LinksetInterfaceBindingResource) readLinksetInterfaceBindingFromApi(ctx context.Context, data *LinksetInterfaceBindingResourceModel, diags *diag.Diagnostics) bool {

	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"id", "ifnum"}, nil)
	if err != nil {
		diags.AddError("Parse Error", fmt.Sprintf("Unable to parse ID: %s", err))
		return false
	}

	linksetid, ok := idMap["id"]
	if !ok {
		diags.AddError("Parse Error", "ID attribute 'id' not found in ID string")
		return false
	}

	// The direct linkset_interface_binding endpoint returns a keyless empty body on
	// this firmware; read the bound interfaces from the aggregate parent endpoint.
	dataArr, err := linkset_interface_bindingAggregateRead(r.client, linksetid)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read linkset_interface_binding, got error: %s", err))
		return false
	}

	// Binding genuinely absent (parent missing or no nested rows): report drift.
	if len(dataArr) == 0 {
		return false
	}

	// Iterate through results to find the one with the right ifnum.
	foundIndex := -1
	for i, v := range dataArr {
		if idVal, ok := idMap["ifnum"]; ok {
			if val, ok := v["ifnum"].(string); ok && val == idVal {
				foundIndex = i
				break
			}
		}
	}

	// Binding row not present in the aggregate response: drift.
	if foundIndex == -1 {
		return false
	}

	linkset_interface_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
	return true
}
