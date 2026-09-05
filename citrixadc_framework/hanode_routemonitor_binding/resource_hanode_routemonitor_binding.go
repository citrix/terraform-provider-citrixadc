package hanode_routemonitor_binding

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
var _ resource.Resource = &HanodeRoutemonitorBindingResource{}
var _ resource.ResourceWithConfigure = (*HanodeRoutemonitorBindingResource)(nil)
var _ resource.ResourceWithImportState = (*HanodeRoutemonitorBindingResource)(nil)

func NewHanodeRoutemonitorBindingResource() resource.Resource {
	return &HanodeRoutemonitorBindingResource{}
}

// HanodeRoutemonitorBindingResource defines the resource implementation.
type HanodeRoutemonitorBindingResource struct {
	client *service.NitroClient
}

// legacyIdAttrOrder matches the SDK v2 comma-separated ID order so that imported
// SDK v2 state IDs (e.g. "0,10.222.74.128") still parse correctly.
var legacyIdAttrOrder = []string{"hanode_id", "routemonitor"}

func (r *HanodeRoutemonitorBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *HanodeRoutemonitorBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_hanode_routemonitor_binding"
}

func (r *HanodeRoutemonitorBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *HanodeRoutemonitorBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data HanodeRoutemonitorBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating hanode_routemonitor_binding resource")
	hanode_routemonitor_binding := hanode_routemonitor_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Hanode_routemonitor_binding.Type(), &hanode_routemonitor_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create hanode_routemonitor_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created hanode_routemonitor_binding resource")

	// Set ID for the resource before reading state.
	// Composite key:UrlEncode(value) pairs in the SDK v2 legacy attribute order
	// (hanode_id,routemonitor).
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("hanode_id:%s", utils.UrlEncode(fmt.Sprintf("%v", data.HanodeId.ValueInt64()))))
	idParts = append(idParts, fmt.Sprintf("routemonitor:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Routemonitor.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readHanodeRoutemonitorBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "hanode_routemonitor_binding not found on the ADC immediately after create")
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *HanodeRoutemonitorBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data HanodeRoutemonitorBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading hanode_routemonitor_binding resource")

	r.readHanodeRoutemonitorBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	// Binding is gone on the ADC (readFromApi nulled the Id): drop it from state so a
	// subsequent apply recreates it, matching the SDK v2 provider's behaviour.
	if data.Id.IsNull() {
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *HanodeRoutemonitorBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state HanodeRoutemonitorBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// No NITRO update endpoint exists for this binding; every attribute uses
	// RequiresReplace, so Terraform never reaches Update with a real change.
	tflog.Debug(ctx, "Update is a no-op for hanode_routemonitor_binding; all attributes are RequiresReplace")

	// Read the updated state back
	r.readHanodeRoutemonitorBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "hanode_routemonitor_binding not found on the ADC immediately after update")
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *HanodeRoutemonitorBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data HanodeRoutemonitorBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting hanode_routemonitor_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs.
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), legacyIdAttrOrder, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	hanodeId, ok := idMap["hanode_id"]
	if !ok {
		resp.Diagnostics.AddError("Parse Error", "Parent attribute 'hanode_id' not found in ID")
		return
	}

	// netmask is not part of the ID; the DELETE endpoint requires it as an arg.
	// Read it from state (matching the SDK v2 behaviour) and URL-encode it.
	args := make([]string, 0)
	if val, ok := idMap["routemonitor"]; ok && val != "" {
		args = append(args, fmt.Sprintf("routemonitor:%s", utils.UrlEncode(val)))
	}
	if !data.Netmask.IsNull() && data.Netmask.ValueString() != "" {
		args = append(args, fmt.Sprintf("netmask:%s", utils.UrlEncode(data.Netmask.ValueString())))
	}

	err = r.client.DeleteResourceWithArgs(service.Hanode_routemonitor_binding.Type(), hanodeId, args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete hanode_routemonitor_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted hanode_routemonitor_binding binding")
}

// Helper function to read hanode_routemonitor_binding data from API
func (r *HanodeRoutemonitorBindingResource) readHanodeRoutemonitorBindingFromApi(ctx context.Context, data *HanodeRoutemonitorBindingResourceModel, diags *diag.Diagnostics) {

	// Array filter with parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), legacyIdAttrOrder, nil)
	if err != nil {
		diags.AddError("Parse Error", fmt.Sprintf("Unable to parse ID: %s", err))
		return
	}

	hanodeId, ok := idMap["hanode_id"]
	if !ok {
		diags.AddError("Parse Error", "ID attribute 'hanode_id' not found in ID string")
		return
	}

	var dataArr []map[string]interface{}

	findParams := service.FindParams{
		ResourceType:             service.Hanode_routemonitor_binding.Type(),
		ResourceName:             hanodeId,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read hanode_routemonitor_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		// Binding (or its parent) no longer exists on the ADC. Signal removal via a null Id
		// (matches SDK v2 d.SetId("")) so the Read caller drops it from state instead of erroring.
		data.Id = types.StringNull()
		return
	}

	// Iterate through results to find the one with the matching routemonitor
	foundIndex := -1
	for i, v := range dataArr {
		if idVal, ok := idMap["routemonitor"]; ok {
			if val, ok := v["routemonitor"].(string); ok && val == idVal {
				foundIndex = i
				break
			}
		}
	}

	//  Resource is missing
	if foundIndex == -1 {
		data.Id = types.StringNull()
		return
	}

	hanode_routemonitor_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
