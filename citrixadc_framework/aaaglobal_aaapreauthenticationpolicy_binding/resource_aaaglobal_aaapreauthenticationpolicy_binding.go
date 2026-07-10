package aaaglobal_aaapreauthenticationpolicy_binding

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
var _ resource.Resource = &AaaglobalAaapreauthenticationpolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*AaaglobalAaapreauthenticationpolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*AaaglobalAaapreauthenticationpolicyBindingResource)(nil)

func NewAaaglobalAaapreauthenticationpolicyBindingResource() resource.Resource {
	return &AaaglobalAaapreauthenticationpolicyBindingResource{}
}

// AaaglobalAaapreauthenticationpolicyBindingResource defines the resource implementation.
type AaaglobalAaapreauthenticationpolicyBindingResource struct {
	client *service.NitroClient
}

func (r *AaaglobalAaapreauthenticationpolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AaaglobalAaapreauthenticationpolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_aaaglobal_aaapreauthenticationpolicy_binding"
}

func (r *AaaglobalAaapreauthenticationpolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AaaglobalAaapreauthenticationpolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AaaglobalAaapreauthenticationpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating aaaglobal_aaapreauthenticationpolicy_binding resource")
	aaaglobal_aaapreauthenticationpolicy_binding := aaaglobal_aaapreauthenticationpolicy_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Aaaglobal_aaapreauthenticationpolicy_binding.Type(), &aaaglobal_aaapreauthenticationpolicy_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create aaaglobal_aaapreauthenticationpolicy_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created aaaglobal_aaapreauthenticationpolicy_binding resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Policy.ValueString()))

	// Read the updated state back
	r.readAaaglobalAaapreauthenticationpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AaaglobalAaapreauthenticationpolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AaaglobalAaapreauthenticationpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading aaaglobal_aaapreauthenticationpolicy_binding resource")

	r.readAaaglobalAaapreauthenticationpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AaaglobalAaapreauthenticationpolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state AaaglobalAaapreauthenticationpolicyBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating aaaglobal_aaapreauthenticationpolicy_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false

	if hasChange {
		// Create API request body from the model
		aaaglobal_aaapreauthenticationpolicy_binding := aaaglobal_aaapreauthenticationpolicy_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Aaaglobal_aaapreauthenticationpolicy_binding.Type(), &aaaglobal_aaapreauthenticationpolicy_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update aaaglobal_aaapreauthenticationpolicy_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated aaaglobal_aaapreauthenticationpolicy_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for aaaglobal_aaapreauthenticationpolicy_binding resource, skipping update")
	}

	// Read the updated state back
	r.readAaaglobalAaapreauthenticationpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AaaglobalAaapreauthenticationpolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AaaglobalAaapreauthenticationpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting aaaglobal_aaapreauthenticationpolicy_binding resource")
	// Global binding - delete using DeleteResourceWithArgs with empty resource name
	// Single unique attribute - ID is the plain value
	policy_value := data.Id.ValueString()
	args := []string{
		fmt.Sprintf("policy:%s", policy_value),
	}

	err := r.client.DeleteResourceWithArgs(service.Aaaglobal_aaapreauthenticationpolicy_binding.Type(), "", args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete aaaglobal_aaapreauthenticationpolicy_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted aaaglobal_aaapreauthenticationpolicy_binding binding")
}

// Helper function to read aaaglobal_aaapreauthenticationpolicy_binding data from API
func (r *AaaglobalAaapreauthenticationpolicyBindingResource) readAaaglobalAaapreauthenticationpolicyBindingFromApi(ctx context.Context, data *AaaglobalAaapreauthenticationpolicyBindingResourceModel, diags *diag.Diagnostics) {

	// Case 3: Array filter without parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"policy"}, nil)
	if err != nil {
		diags.AddError("Parse Error", fmt.Sprintf("Unable to parse ID: %s", err))
		return
	}

	var dataArr []map[string]interface{}

	findParams := service.FindParams{
		ResourceType:             service.Aaaglobal_aaapreauthenticationpolicy_binding.Type(),
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read aaaglobal_aaapreauthenticationpolicy_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		diags.AddError("Client Error", "aaaglobal_aaapreauthenticationpolicy_binding returned empty array")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check policy
		if idVal, ok := idMap["policy"]; ok {
			if val, ok := v["policy"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["policy"].(string); ok {
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
		diags.AddError("Client Error", fmt.Sprintf("aaaglobal_aaapreauthenticationpolicy_binding not found with the provided ID attributes"))
		return
	}

	aaaglobal_aaapreauthenticationpolicy_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
