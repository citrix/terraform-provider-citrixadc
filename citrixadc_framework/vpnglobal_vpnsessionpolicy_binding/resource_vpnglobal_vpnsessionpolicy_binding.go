package vpnglobal_vpnsessionpolicy_binding

import (
	"context"
	"fmt"

	"net/url"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &VpnglobalVpnsessionpolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*VpnglobalVpnsessionpolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*VpnglobalVpnsessionpolicyBindingResource)(nil)

func NewVpnglobalVpnsessionpolicyBindingResource() resource.Resource {
	return &VpnglobalVpnsessionpolicyBindingResource{}
}

// VpnglobalVpnsessionpolicyBindingResource defines the resource implementation.
type VpnglobalVpnsessionpolicyBindingResource struct {
	client *service.NitroClient
}

func (r *VpnglobalVpnsessionpolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VpnglobalVpnsessionpolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpnglobal_vpnsessionpolicy_binding"
}

func (r *VpnglobalVpnsessionpolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VpnglobalVpnsessionpolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VpnglobalVpnsessionpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating vpnglobal_vpnsessionpolicy_binding resource")
	vpnglobal_vpnsessionpolicy_binding := vpnglobal_vpnsessionpolicy_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Vpnglobal_vpnsessionpolicy_binding.Type(), &vpnglobal_vpnsessionpolicy_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create vpnglobal_vpnsessionpolicy_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created vpnglobal_vpnsessionpolicy_binding resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Policyname.ValueString()))

	// Read the updated state back
	r.readVpnglobalVpnsessionpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "vpnglobal_vpnsessionpolicy_binding not found on the ADC immediately after create")
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnglobalVpnsessionpolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VpnglobalVpnsessionpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading vpnglobal_vpnsessionpolicy_binding resource")

	r.readVpnglobalVpnsessionpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *VpnglobalVpnsessionpolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state VpnglobalVpnsessionpolicyBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating vpnglobal_vpnsessionpolicy_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false

	if hasChange {
		// Create API request body from the model
		vpnglobal_vpnsessionpolicy_binding := vpnglobal_vpnsessionpolicy_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Vpnglobal_vpnsessionpolicy_binding.Type(), &vpnglobal_vpnsessionpolicy_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update vpnglobal_vpnsessionpolicy_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated vpnglobal_vpnsessionpolicy_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for vpnglobal_vpnsessionpolicy_binding resource, skipping update")
	}

	// Read the updated state back
	r.readVpnglobalVpnsessionpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "vpnglobal_vpnsessionpolicy_binding not found on the ADC immediately after update")
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnglobalVpnsessionpolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VpnglobalVpnsessionpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting vpnglobal_vpnsessionpolicy_binding resource")
	// Global binding - delete using DeleteResourceWithArgs with empty resource name
	// Single unique attribute - ID is the plain value
	policyname_value := data.Id.ValueString()
	args := []string{
		fmt.Sprintf("policyname:%s", url.QueryEscape(policyname_value)),
	}
	// NITRO delete accepts secondary and groupextraction as disambiguating args
	// (mirrors the SDK v2 resource). Only send them when set in state.
	if !data.Secondary.IsNull() && !data.Secondary.IsUnknown() {
		args = append(args, fmt.Sprintf("secondary:%s", url.QueryEscape(fmt.Sprintf("%t", data.Secondary.ValueBool()))))
	}
	if !data.Groupextraction.IsNull() && !data.Groupextraction.IsUnknown() {
		args = append(args, fmt.Sprintf("groupextraction:%s", url.QueryEscape(fmt.Sprintf("%t", data.Groupextraction.ValueBool()))))
	}

	err := r.client.DeleteResourceWithArgs(service.Vpnglobal_vpnsessionpolicy_binding.Type(), "", args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete vpnglobal_vpnsessionpolicy_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted vpnglobal_vpnsessionpolicy_binding binding")
}

// Helper function to read vpnglobal_vpnsessionpolicy_binding data from API
func (r *VpnglobalVpnsessionpolicyBindingResource) readVpnglobalVpnsessionpolicyBindingFromApi(ctx context.Context, data *VpnglobalVpnsessionpolicyBindingResourceModel, diags *diag.Diagnostics) {

	// Single unique attribute - ID is the plain policyname value (Pattern 10)
	policyname := data.Id.ValueString()

	var dataArr []map[string]interface{}

	findParams := service.FindParams{
		ResourceType:             service.Vpnglobal_vpnsessionpolicy_binding.Type(),
		ResourceMissingErrorCode: 258,
	}
	dataArr, err := r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read vpnglobal_vpnsessionpolicy_binding, got error: %s", err))
		return
	}

	// Binding (or its parent) no longer exists on the ADC. Signal removal via a null Id
	// (matches SDK v2 d.SetId("")) so the Read caller drops it from state instead of erroring.
	if len(dataArr) == 0 {
		data.Id = types.StringNull()
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if val, ok := v["policyname"].(string); ok && val == policyname {
			foundIndex = i
			break
		}
	}

	// Binding not present in the returned set: signal removal via a null Id (see above).
	if foundIndex == -1 {
		data.Id = types.StringNull()
		return
	}

	vpnglobal_vpnsessionpolicy_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
