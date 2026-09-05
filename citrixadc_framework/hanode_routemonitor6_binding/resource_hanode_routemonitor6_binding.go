package hanode_routemonitor6_binding

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
var _ resource.Resource = &HanodeRoutemonitor6BindingResource{}
var _ resource.ResourceWithConfigure = (*HanodeRoutemonitor6BindingResource)(nil)
var _ resource.ResourceWithImportState = (*HanodeRoutemonitor6BindingResource)(nil)

// legacyIdOrder mirrors resource_id_mapping.json: hanode_id,routemonitor.
var legacyIdOrder = []string{"hanode_id", "routemonitor"}

func NewHanodeRoutemonitor6BindingResource() resource.Resource {
	return &HanodeRoutemonitor6BindingResource{}
}

// HanodeRoutemonitor6BindingResource defines the resource implementation.
type HanodeRoutemonitor6BindingResource struct {
	client *service.NitroClient
}

func (r *HanodeRoutemonitor6BindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *HanodeRoutemonitor6BindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_hanode_routemonitor6_binding"
}

func (r *HanodeRoutemonitor6BindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *HanodeRoutemonitor6BindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data HanodeRoutemonitor6BindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating hanode_routemonitor6_binding resource")
	hanode_routemonitor6_binding := hanode_routemonitor6_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Hanode_routemonitor6_binding.Type(), &hanode_routemonitor6_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create hanode_routemonitor6_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created hanode_routemonitor6_binding resource")

	// Set ID for the resource before reading state.
	// Composite ID uses the legacy SDK v2 key order: hanode_id,routemonitor.
	data.Id = types.StringValue(hanode_routemonitor6_bindingComposeId(&data))

	// Read the updated state back
	r.readHanodeRoutemonitor6BindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "hanode_routemonitor6_binding not found on the ADC immediately after create")
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *HanodeRoutemonitor6BindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data HanodeRoutemonitor6BindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading hanode_routemonitor6_binding resource")

	r.readHanodeRoutemonitor6BindingFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *HanodeRoutemonitor6BindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state HanodeRoutemonitor6BindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// All schema attributes are RequiresReplace; there is no NITRO update endpoint
	// for this binding, so Update is a documented no-op.
	tflog.Debug(ctx, "Update is a no-op for hanode_routemonitor6_binding; all attributes are RequiresReplace")

	// Read the current state back
	r.readHanodeRoutemonitor6BindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "hanode_routemonitor6_binding not found on the ADC immediately after update")
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *HanodeRoutemonitor6BindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data HanodeRoutemonitor6BindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting hanode_routemonitor6_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs.
	// ParseIdString handles both the new key:value form and the legacy
	// positional "hanode_id,routemonitor" form for imported SDK v2 state.
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), legacyIdOrder, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	hanodeId, ok := idMap["hanode_id"]
	if !ok {
		resp.Diagnostics.AddError("Parse Error", "Parent attribute 'hanode_id' not found in ID")
		return
	}

	routemonitor, ok := idMap["routemonitor"]
	if !ok {
		resp.Diagnostics.AddError("Parse Error", "Attribute 'routemonitor' not found in ID")
		return
	}

	// Match SDK v2 delete: name = hanode_id, args = routemonitor (URL-encoded for
	// slashy / special-char values).
	args := []string{fmt.Sprintf("routemonitor:%s", utils.UrlEncode(routemonitor))}

	err = r.client.DeleteResourceWithArgs(service.Hanode_routemonitor6_binding.Type(), hanodeId, args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete hanode_routemonitor6_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted hanode_routemonitor6_binding binding")
}

// Helper function to read hanode_routemonitor6_binding data from API
func (r *HanodeRoutemonitor6BindingResource) readHanodeRoutemonitor6BindingFromApi(ctx context.Context, data *HanodeRoutemonitor6BindingResourceModel, diags *diag.Diagnostics) {

	// Array filter with parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), legacyIdOrder, nil)
	if err != nil {
		diags.AddError("Parse Error", fmt.Sprintf("Unable to parse ID: %s", err))
		return
	}

	hanodeId, ok := idMap["hanode_id"]
	if !ok {
		diags.AddError("Parse Error", "ID attribute 'hanode_id' not found in ID string")
		return
	}

	routemonitor, ok := idMap["routemonitor"]
	if !ok {
		diags.AddError("Parse Error", "ID attribute 'routemonitor' not found in ID string")
		return
	}

	var dataArr []map[string]interface{}

	findParams := service.FindParams{
		ResourceType:             service.Hanode_routemonitor6_binding.Type(),
		ResourceName:             hanodeId,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read hanode_routemonitor6_binding, got error: %s", err))
		return
	}

	// Binding (or its parent) no longer exists on the ADC. Signal removal via a null Id
	// (matches SDK v2 d.SetId("")) so the Read caller drops it from state instead of erroring.
	if len(dataArr) == 0 {
		data.Id = types.StringNull()
		return
	}

	// Iterate through results to find the one with the matching routemonitor.
	foundIndex := -1
	for i, v := range dataArr {
		if val, ok := v["routemonitor"].(string); ok && val == routemonitor {
			foundIndex = i
			break
		}
	}

	// Binding not present in the returned set: signal removal via a null Id (see above).
	if foundIndex == -1 {
		data.Id = types.StringNull()
		return
	}

	hanode_routemonitor6_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
