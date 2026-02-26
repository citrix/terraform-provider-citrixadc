package aaagroup_authorizationpolicy_binding

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
var _ resource.Resource = &AaagroupAuthorizationpolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*AaagroupAuthorizationpolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*AaagroupAuthorizationpolicyBindingResource)(nil)

func NewAaagroupAuthorizationpolicyBindingResource() resource.Resource {
	return &AaagroupAuthorizationpolicyBindingResource{}
}

// AaagroupAuthorizationpolicyBindingResource defines the resource implementation.
type AaagroupAuthorizationpolicyBindingResource struct {
	client *service.NitroClient
}

func (r *AaagroupAuthorizationpolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AaagroupAuthorizationpolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_aaagroup_authorizationpolicy_binding"
}

func (r *AaagroupAuthorizationpolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AaagroupAuthorizationpolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AaagroupAuthorizationpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating aaagroup_authorizationpolicy_binding resource")

	// aaagroup_authorizationpolicy_binding := aaagroup_authorizationpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Aaagroup_authorizationpolicy_binding.Type(), &aaagroup_authorizationpolicy_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create aaagroup_authorizationpolicy_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("aaagroup_authorizationpolicy_binding-config")

	tflog.Trace(ctx, "Created aaagroup_authorizationpolicy_binding resource")

	// Read the updated state back
	r.readAaagroupAuthorizationpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AaagroupAuthorizationpolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AaagroupAuthorizationpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading aaagroup_authorizationpolicy_binding resource")

	r.readAaagroupAuthorizationpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AaagroupAuthorizationpolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data AaagroupAuthorizationpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating aaagroup_authorizationpolicy_binding resource")

	// Create API request body from the model
	// aaagroup_authorizationpolicy_binding := aaagroup_authorizationpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Aaagroup_authorizationpolicy_binding.Type(), &aaagroup_authorizationpolicy_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update aaagroup_authorizationpolicy_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated aaagroup_authorizationpolicy_binding resource")

	// Read the updated state back
	r.readAaagroupAuthorizationpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AaagroupAuthorizationpolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AaagroupAuthorizationpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting aaagroup_authorizationpolicy_binding resource")

	// For aaagroup_authorizationpolicy_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted aaagroup_authorizationpolicy_binding resource from state")
}

// Helper function to read aaagroup_authorizationpolicy_binding data from API
func (r *AaagroupAuthorizationpolicyBindingResource) readAaagroupAuthorizationpolicyBindingFromApi(ctx context.Context, data *AaagroupAuthorizationpolicyBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Aaagroup_authorizationpolicy_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read aaagroup_authorizationpolicy_binding, got error: %s", err))
		return
	}

	aaagroup_authorizationpolicy_bindingSetAttrFromGet(ctx, data, getResponseData)

}
