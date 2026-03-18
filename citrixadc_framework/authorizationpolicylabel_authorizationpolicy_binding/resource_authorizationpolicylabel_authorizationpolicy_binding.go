package authorizationpolicylabel_authorizationpolicy_binding

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
var _ resource.Resource = &AuthorizationpolicylabelAuthorizationpolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*AuthorizationpolicylabelAuthorizationpolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*AuthorizationpolicylabelAuthorizationpolicyBindingResource)(nil)

func NewAuthorizationpolicylabelAuthorizationpolicyBindingResource() resource.Resource {
	return &AuthorizationpolicylabelAuthorizationpolicyBindingResource{}
}

// AuthorizationpolicylabelAuthorizationpolicyBindingResource defines the resource implementation.
type AuthorizationpolicylabelAuthorizationpolicyBindingResource struct {
	client *service.NitroClient
}

func (r *AuthorizationpolicylabelAuthorizationpolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AuthorizationpolicylabelAuthorizationpolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_authorizationpolicylabel_authorizationpolicy_binding"
}

func (r *AuthorizationpolicylabelAuthorizationpolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AuthorizationpolicylabelAuthorizationpolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AuthorizationpolicylabelAuthorizationpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating authorizationpolicylabel_authorizationpolicy_binding resource")

	// authorizationpolicylabel_authorizationpolicy_binding := authorizationpolicylabel_authorizationpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Authorizationpolicylabel_authorizationpolicy_binding.Type(), &authorizationpolicylabel_authorizationpolicy_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create authorizationpolicylabel_authorizationpolicy_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("authorizationpolicylabel_authorizationpolicy_binding-config")

	tflog.Trace(ctx, "Created authorizationpolicylabel_authorizationpolicy_binding resource")

	// Read the updated state back
	r.readAuthorizationpolicylabelAuthorizationpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthorizationpolicylabelAuthorizationpolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AuthorizationpolicylabelAuthorizationpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading authorizationpolicylabel_authorizationpolicy_binding resource")

	r.readAuthorizationpolicylabelAuthorizationpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthorizationpolicylabelAuthorizationpolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data AuthorizationpolicylabelAuthorizationpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating authorizationpolicylabel_authorizationpolicy_binding resource")

	// Create API request body from the model
	// authorizationpolicylabel_authorizationpolicy_binding := authorizationpolicylabel_authorizationpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Authorizationpolicylabel_authorizationpolicy_binding.Type(), &authorizationpolicylabel_authorizationpolicy_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update authorizationpolicylabel_authorizationpolicy_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated authorizationpolicylabel_authorizationpolicy_binding resource")

	// Read the updated state back
	r.readAuthorizationpolicylabelAuthorizationpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthorizationpolicylabelAuthorizationpolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AuthorizationpolicylabelAuthorizationpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting authorizationpolicylabel_authorizationpolicy_binding resource")

	// For authorizationpolicylabel_authorizationpolicy_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted authorizationpolicylabel_authorizationpolicy_binding resource from state")
}

// Helper function to read authorizationpolicylabel_authorizationpolicy_binding data from API
func (r *AuthorizationpolicylabelAuthorizationpolicyBindingResource) readAuthorizationpolicylabelAuthorizationpolicyBindingFromApi(ctx context.Context, data *AuthorizationpolicylabelAuthorizationpolicyBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Authorizationpolicylabel_authorizationpolicy_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read authorizationpolicylabel_authorizationpolicy_binding, got error: %s", err))
		return
	}

	authorizationpolicylabel_authorizationpolicy_bindingSetAttrFromGet(ctx, data, getResponseData)

}
