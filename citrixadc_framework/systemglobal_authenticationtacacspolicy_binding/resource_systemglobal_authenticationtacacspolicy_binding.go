package systemglobal_authenticationtacacspolicy_binding

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
var _ resource.Resource = &SystemglobalAuthenticationtacacspolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*SystemglobalAuthenticationtacacspolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*SystemglobalAuthenticationtacacspolicyBindingResource)(nil)

func NewSystemglobalAuthenticationtacacspolicyBindingResource() resource.Resource {
	return &SystemglobalAuthenticationtacacspolicyBindingResource{}
}

// SystemglobalAuthenticationtacacspolicyBindingResource defines the resource implementation.
type SystemglobalAuthenticationtacacspolicyBindingResource struct {
	client *service.NitroClient
}

func (r *SystemglobalAuthenticationtacacspolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SystemglobalAuthenticationtacacspolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_systemglobal_authenticationtacacspolicy_binding"
}

func (r *SystemglobalAuthenticationtacacspolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SystemglobalAuthenticationtacacspolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SystemglobalAuthenticationtacacspolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating systemglobal_authenticationtacacspolicy_binding resource")
	systemglobal_authenticationtacacspolicy_binding := systemglobal_authenticationtacacspolicy_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Systemglobal_authenticationtacacspolicy_binding.Type(), &systemglobal_authenticationtacacspolicy_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create systemglobal_authenticationtacacspolicy_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created systemglobal_authenticationtacacspolicy_binding resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Policyname.ValueString()))

	// Read the updated state back
	r.readSystemglobalAuthenticationtacacspolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemglobalAuthenticationtacacspolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SystemglobalAuthenticationtacacspolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading systemglobal_authenticationtacacspolicy_binding resource")

	r.readSystemglobalAuthenticationtacacspolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemglobalAuthenticationtacacspolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state SystemglobalAuthenticationtacacspolicyBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating systemglobal_authenticationtacacspolicy_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false

	if hasChange {
		// Create API request body from the model
		systemglobal_authenticationtacacspolicy_binding := systemglobal_authenticationtacacspolicy_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Systemglobal_authenticationtacacspolicy_binding.Type(), &systemglobal_authenticationtacacspolicy_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update systemglobal_authenticationtacacspolicy_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated systemglobal_authenticationtacacspolicy_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for systemglobal_authenticationtacacspolicy_binding resource, skipping update")
	}

	// Read the updated state back
	r.readSystemglobalAuthenticationtacacspolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemglobalAuthenticationtacacspolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SystemglobalAuthenticationtacacspolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting systemglobal_authenticationtacacspolicy_binding resource")
	// Global binding - delete using DeleteResourceWithArgs with empty resource name
	// Single unique attribute - ID is the plain value
	policyname_value := data.Id.ValueString()
	args := []string{
		fmt.Sprintf("policyname:%s", policyname_value),
	}

	err := r.client.DeleteResourceWithArgs(service.Systemglobal_authenticationtacacspolicy_binding.Type(), "", args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete systemglobal_authenticationtacacspolicy_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted systemglobal_authenticationtacacspolicy_binding binding")
}

// Helper function to read systemglobal_authenticationtacacspolicy_binding data from API
func (r *SystemglobalAuthenticationtacacspolicyBindingResource) readSystemglobalAuthenticationtacacspolicyBindingFromApi(ctx context.Context, data *SystemglobalAuthenticationtacacspolicyBindingResourceModel, diags *diag.Diagnostics) {

	// Case 3: Array filter without parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"policyname"}, nil)
	if err != nil {
		diags.AddError("Parse Error", fmt.Sprintf("Unable to parse ID: %s", err))
		return
	}

	var dataArr []map[string]interface{}

	findParams := service.FindParams{
		ResourceType:             service.Systemglobal_authenticationtacacspolicy_binding.Type(),
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read systemglobal_authenticationtacacspolicy_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		diags.AddError("Client Error", "systemglobal_authenticationtacacspolicy_binding returned empty array")
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

	// Resource is missing
	if foundIndex == -1 {
		diags.AddError("Client Error", fmt.Sprintf("systemglobal_authenticationtacacspolicy_binding not found with the provided ID attributes"))
		return
	}

	systemglobal_authenticationtacacspolicy_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
