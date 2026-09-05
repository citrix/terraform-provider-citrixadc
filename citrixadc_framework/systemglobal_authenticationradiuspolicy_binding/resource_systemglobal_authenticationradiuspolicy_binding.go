package systemglobal_authenticationradiuspolicy_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &SystemglobalAuthenticationradiuspolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*SystemglobalAuthenticationradiuspolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*SystemglobalAuthenticationradiuspolicyBindingResource)(nil)

func NewSystemglobalAuthenticationradiuspolicyBindingResource() resource.Resource {
	return &SystemglobalAuthenticationradiuspolicyBindingResource{}
}

// SystemglobalAuthenticationradiuspolicyBindingResource defines the resource implementation.
type SystemglobalAuthenticationradiuspolicyBindingResource struct {
	client *service.NitroClient
}

func (r *SystemglobalAuthenticationradiuspolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SystemglobalAuthenticationradiuspolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_systemglobal_authenticationradiuspolicy_binding"
}

func (r *SystemglobalAuthenticationradiuspolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SystemglobalAuthenticationradiuspolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SystemglobalAuthenticationradiuspolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating systemglobal_authenticationradiuspolicy_binding resource")
	systemglobal_authenticationradiuspolicy_binding := systemglobal_authenticationradiuspolicy_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Systemglobal_authenticationradiuspolicy_binding.Type(), &systemglobal_authenticationradiuspolicy_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create systemglobal_authenticationradiuspolicy_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created systemglobal_authenticationradiuspolicy_binding resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Policyname.ValueString()))

	// Read the updated state back
	r.readSystemglobalAuthenticationradiuspolicyBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "systemglobal_authenticationradiuspolicy_binding not found on the ADC immediately after create")
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemglobalAuthenticationradiuspolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SystemglobalAuthenticationradiuspolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading systemglobal_authenticationradiuspolicy_binding resource")

	r.readSystemglobalAuthenticationradiuspolicyBindingFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *SystemglobalAuthenticationradiuspolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state SystemglobalAuthenticationradiuspolicyBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating systemglobal_authenticationradiuspolicy_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false

	if hasChange {
		// Create API request body from the model
		systemglobal_authenticationradiuspolicy_binding := systemglobal_authenticationradiuspolicy_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Systemglobal_authenticationradiuspolicy_binding.Type(), &systemglobal_authenticationradiuspolicy_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update systemglobal_authenticationradiuspolicy_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated systemglobal_authenticationradiuspolicy_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for systemglobal_authenticationradiuspolicy_binding resource, skipping update")
	}

	// Read the updated state back
	r.readSystemglobalAuthenticationradiuspolicyBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "systemglobal_authenticationradiuspolicy_binding not found on the ADC immediately after update")
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemglobalAuthenticationradiuspolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SystemglobalAuthenticationradiuspolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting systemglobal_authenticationradiuspolicy_binding resource")
	// Global binding - delete using DeleteResourceWithArgs with empty resource name
	// Single unique attribute - ID is the plain value
	policyname_value := data.Id.ValueString()
	args := []string{
		fmt.Sprintf("policyname:%s", policyname_value),
	}

	err := r.client.DeleteResourceWithArgs(service.Systemglobal_authenticationradiuspolicy_binding.Type(), "", args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete systemglobal_authenticationradiuspolicy_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted systemglobal_authenticationradiuspolicy_binding binding")
}

// Helper function to read systemglobal_authenticationradiuspolicy_binding data from API
func (r *SystemglobalAuthenticationradiuspolicyBindingResource) readSystemglobalAuthenticationradiuspolicyBindingFromApi(ctx context.Context, data *SystemglobalAuthenticationradiuspolicyBindingResourceModel, diags *diag.Diagnostics) {

	// Single unique attribute -> the ID is the plain policyname value (ParseIdString also accepts
	// a legacy SDK v2 plain value, so the ID is used directly here per Pattern 10).
	policyname := data.Id.ValueString()

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Systemglobal_authenticationradiuspolicy_binding.Type(),
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read systemglobal_authenticationradiuspolicy_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		// Binding (or its parent) no longer exists on the ADC. Signal removal via a null Id
		// (matches SDK v2 d.SetId("")) so the Read caller drops it from state instead of erroring.
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

	// Resource is missing
	if foundIndex == -1 {
		data.Id = types.StringNull()
		return
	}

	systemglobal_authenticationradiuspolicy_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
