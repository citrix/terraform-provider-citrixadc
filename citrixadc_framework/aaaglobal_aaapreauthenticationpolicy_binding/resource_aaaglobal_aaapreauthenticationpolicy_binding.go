package aaaglobal_aaapreauthenticationpolicy_binding

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

	// aaaglobal_aaapreauthenticationpolicy_binding := aaaglobal_aaapreauthenticationpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.AaaglobalAaapreauthenticationpolicyBinding.Type(), &aaaglobal_aaapreauthenticationpolicy_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create aaaglobal_aaapreauthenticationpolicy_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("aaaglobal_aaapreauthenticationpolicy_binding-config")

	tflog.Trace(ctx, "Created aaaglobal_aaapreauthenticationpolicy_binding resource")

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
	var data AaaglobalAaapreauthenticationpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating aaaglobal_aaapreauthenticationpolicy_binding resource")

	// Create API request body from the model
	// aaaglobal_aaapreauthenticationpolicy_binding := aaaglobal_aaapreauthenticationpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.AaaglobalAaapreauthenticationpolicyBinding.Type(), &aaaglobal_aaapreauthenticationpolicy_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update aaaglobal_aaapreauthenticationpolicy_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated aaaglobal_aaapreauthenticationpolicy_binding resource")

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

	// For aaaglobal_aaapreauthenticationpolicy_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted aaaglobal_aaapreauthenticationpolicy_binding resource from state")
}

// Helper function to read aaaglobal_aaapreauthenticationpolicy_binding data from API
func (r *AaaglobalAaapreauthenticationpolicyBindingResource) readAaaglobalAaapreauthenticationpolicyBindingFromApi(ctx context.Context, data *AaaglobalAaapreauthenticationpolicyBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Aaaglobal_aaapreauthenticationpolicy_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read aaaglobal_aaapreauthenticationpolicy_binding, got error: %s", err))
		return
	}

	aaaglobal_aaapreauthenticationpolicy_bindingSetAttrFromGet(ctx, data, getResponseData)

}
