package systemglobal_authenticationldappolicy_binding

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
var _ resource.Resource = &SystemglobalAuthenticationldappolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*SystemglobalAuthenticationldappolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*SystemglobalAuthenticationldappolicyBindingResource)(nil)

func NewSystemglobalAuthenticationldappolicyBindingResource() resource.Resource {
	return &SystemglobalAuthenticationldappolicyBindingResource{}
}

// SystemglobalAuthenticationldappolicyBindingResource defines the resource implementation.
type SystemglobalAuthenticationldappolicyBindingResource struct {
	client *service.NitroClient
}

func (r *SystemglobalAuthenticationldappolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SystemglobalAuthenticationldappolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_systemglobal_authenticationldappolicy_binding"
}

func (r *SystemglobalAuthenticationldappolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SystemglobalAuthenticationldappolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SystemglobalAuthenticationldappolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating systemglobal_authenticationldappolicy_binding resource")
	systemglobal_authenticationldappolicy_binding := systemglobal_authenticationldappolicy_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Systemglobal_authenticationldappolicy_binding.Type(), &systemglobal_authenticationldappolicy_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create systemglobal_authenticationldappolicy_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created systemglobal_authenticationldappolicy_binding resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Policyname.ValueString()))

	// Read the updated state back
	r.readSystemglobalAuthenticationldappolicyBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "systemglobal_authenticationldappolicy_binding not found on the ADC immediately after create")
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemglobalAuthenticationldappolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SystemglobalAuthenticationldappolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading systemglobal_authenticationldappolicy_binding resource")

	r.readSystemglobalAuthenticationldappolicyBindingFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *SystemglobalAuthenticationldappolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state SystemglobalAuthenticationldappolicyBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating systemglobal_authenticationldappolicy_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false

	if hasChange {
		// Create API request body from the model
		systemglobal_authenticationldappolicy_binding := systemglobal_authenticationldappolicy_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Systemglobal_authenticationldappolicy_binding.Type(), &systemglobal_authenticationldappolicy_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update systemglobal_authenticationldappolicy_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated systemglobal_authenticationldappolicy_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for systemglobal_authenticationldappolicy_binding resource, skipping update")
	}

	// Read the updated state back
	r.readSystemglobalAuthenticationldappolicyBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "systemglobal_authenticationldappolicy_binding not found on the ADC immediately after update")
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemglobalAuthenticationldappolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SystemglobalAuthenticationldappolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting systemglobal_authenticationldappolicy_binding resource")
	// Global binding - delete using DeleteResourceWithArgs with empty resource name
	// Single unique attribute - ID is the plain value
	policyname_value := data.Id.ValueString()
	args := []string{
		fmt.Sprintf("policyname:%s", policyname_value),
	}

	err := r.client.DeleteResourceWithArgs(service.Systemglobal_authenticationldappolicy_binding.Type(), "", args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete systemglobal_authenticationldappolicy_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted systemglobal_authenticationldappolicy_binding binding")
}

// Helper function to read systemglobal_authenticationldappolicy_binding data from API
func (r *SystemglobalAuthenticationldappolicyBindingResource) readSystemglobalAuthenticationldappolicyBindingFromApi(ctx context.Context, data *SystemglobalAuthenticationldappolicyBindingResourceModel, diags *diag.Diagnostics) {

	// Case 3: Array filter without parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"policyname"}, nil)
	if err != nil {
		diags.AddError("Parse Error", fmt.Sprintf("Unable to parse ID: %s", err))
		return
	}

	var dataArr []map[string]interface{}

	findParams := service.FindParams{
		ResourceType:             service.Systemglobal_authenticationldappolicy_binding.Type(),
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read systemglobal_authenticationldappolicy_binding, got error: %s", err))
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
		match := true

		// Check policyname
		if idVal, ok := idMap["policyname"]; ok {
			if val, ok := v["policyname"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["policyname"].(string); ok {
			match = false
			continue
		}

		if match {
			foundIndex = i
			break
		}
	}

	// Binding not present in the returned set: signal removal via a null Id (see above).
	if foundIndex == -1 {
		data.Id = types.StringNull()
		return
	}

	systemglobal_authenticationldappolicy_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
