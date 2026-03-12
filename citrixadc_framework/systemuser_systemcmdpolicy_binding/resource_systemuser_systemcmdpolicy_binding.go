package systemuser_systemcmdpolicy_binding

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
var _ resource.Resource = &SystemuserSystemcmdpolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*SystemuserSystemcmdpolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*SystemuserSystemcmdpolicyBindingResource)(nil)

func NewSystemuserSystemcmdpolicyBindingResource() resource.Resource {
	return &SystemuserSystemcmdpolicyBindingResource{}
}

// SystemuserSystemcmdpolicyBindingResource defines the resource implementation.
type SystemuserSystemcmdpolicyBindingResource struct {
	client *service.NitroClient
}

func (r *SystemuserSystemcmdpolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SystemuserSystemcmdpolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_systemuser_systemcmdpolicy_binding"
}

func (r *SystemuserSystemcmdpolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SystemuserSystemcmdpolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SystemuserSystemcmdpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating systemuser_systemcmdpolicy_binding resource")

	// systemuser_systemcmdpolicy_binding := systemuser_systemcmdpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Systemuser_systemcmdpolicy_binding.Type(), &systemuser_systemcmdpolicy_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create systemuser_systemcmdpolicy_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("systemuser_systemcmdpolicy_binding-config")

	tflog.Trace(ctx, "Created systemuser_systemcmdpolicy_binding resource")

	// Read the updated state back
	r.readSystemuserSystemcmdpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemuserSystemcmdpolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SystemuserSystemcmdpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading systemuser_systemcmdpolicy_binding resource")

	r.readSystemuserSystemcmdpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemuserSystemcmdpolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data SystemuserSystemcmdpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating systemuser_systemcmdpolicy_binding resource")

	// Create API request body from the model
	// systemuser_systemcmdpolicy_binding := systemuser_systemcmdpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Systemuser_systemcmdpolicy_binding.Type(), &systemuser_systemcmdpolicy_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update systemuser_systemcmdpolicy_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated systemuser_systemcmdpolicy_binding resource")

	// Read the updated state back
	r.readSystemuserSystemcmdpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemuserSystemcmdpolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SystemuserSystemcmdpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting systemuser_systemcmdpolicy_binding resource")

	// For systemuser_systemcmdpolicy_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted systemuser_systemcmdpolicy_binding resource from state")
}

// Helper function to read systemuser_systemcmdpolicy_binding data from API
func (r *SystemuserSystemcmdpolicyBindingResource) readSystemuserSystemcmdpolicyBindingFromApi(ctx context.Context, data *SystemuserSystemcmdpolicyBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Systemuser_systemcmdpolicy_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read systemuser_systemcmdpolicy_binding, got error: %s", err))
		return
	}

	systemuser_systemcmdpolicy_bindingSetAttrFromGet(ctx, data, getResponseData)

}
