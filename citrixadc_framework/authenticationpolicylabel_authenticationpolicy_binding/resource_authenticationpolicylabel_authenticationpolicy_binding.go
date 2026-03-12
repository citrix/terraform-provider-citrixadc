package authenticationpolicylabel_authenticationpolicy_binding

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
var _ resource.Resource = &AuthenticationpolicylabelAuthenticationpolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*AuthenticationpolicylabelAuthenticationpolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*AuthenticationpolicylabelAuthenticationpolicyBindingResource)(nil)

func NewAuthenticationpolicylabelAuthenticationpolicyBindingResource() resource.Resource {
	return &AuthenticationpolicylabelAuthenticationpolicyBindingResource{}
}

// AuthenticationpolicylabelAuthenticationpolicyBindingResource defines the resource implementation.
type AuthenticationpolicylabelAuthenticationpolicyBindingResource struct {
	client *service.NitroClient
}

func (r *AuthenticationpolicylabelAuthenticationpolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AuthenticationpolicylabelAuthenticationpolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_authenticationpolicylabel_authenticationpolicy_binding"
}

func (r *AuthenticationpolicylabelAuthenticationpolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AuthenticationpolicylabelAuthenticationpolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AuthenticationpolicylabelAuthenticationpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating authenticationpolicylabel_authenticationpolicy_binding resource")

	// authenticationpolicylabel_authenticationpolicy_binding := authenticationpolicylabel_authenticationpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Authenticationpolicylabel_authenticationpolicy_binding.Type(), &authenticationpolicylabel_authenticationpolicy_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create authenticationpolicylabel_authenticationpolicy_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("authenticationpolicylabel_authenticationpolicy_binding-config")

	tflog.Trace(ctx, "Created authenticationpolicylabel_authenticationpolicy_binding resource")

	// Read the updated state back
	r.readAuthenticationpolicylabelAuthenticationpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationpolicylabelAuthenticationpolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AuthenticationpolicylabelAuthenticationpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading authenticationpolicylabel_authenticationpolicy_binding resource")

	r.readAuthenticationpolicylabelAuthenticationpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationpolicylabelAuthenticationpolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data AuthenticationpolicylabelAuthenticationpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating authenticationpolicylabel_authenticationpolicy_binding resource")

	// Create API request body from the model
	// authenticationpolicylabel_authenticationpolicy_binding := authenticationpolicylabel_authenticationpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Authenticationpolicylabel_authenticationpolicy_binding.Type(), &authenticationpolicylabel_authenticationpolicy_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update authenticationpolicylabel_authenticationpolicy_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated authenticationpolicylabel_authenticationpolicy_binding resource")

	// Read the updated state back
	r.readAuthenticationpolicylabelAuthenticationpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationpolicylabelAuthenticationpolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AuthenticationpolicylabelAuthenticationpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting authenticationpolicylabel_authenticationpolicy_binding resource")

	// For authenticationpolicylabel_authenticationpolicy_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted authenticationpolicylabel_authenticationpolicy_binding resource from state")
}

// Helper function to read authenticationpolicylabel_authenticationpolicy_binding data from API
func (r *AuthenticationpolicylabelAuthenticationpolicyBindingResource) readAuthenticationpolicylabelAuthenticationpolicyBindingFromApi(ctx context.Context, data *AuthenticationpolicylabelAuthenticationpolicyBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Authenticationpolicylabel_authenticationpolicy_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read authenticationpolicylabel_authenticationpolicy_binding, got error: %s", err))
		return
	}

	authenticationpolicylabel_authenticationpolicy_bindingSetAttrFromGet(ctx, data, getResponseData)

}
