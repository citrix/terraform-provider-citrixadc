package authenticationvserver_authenticationwebauthpolicy_binding

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
var _ resource.Resource = &AuthenticationvserverAuthenticationwebauthpolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*AuthenticationvserverAuthenticationwebauthpolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*AuthenticationvserverAuthenticationwebauthpolicyBindingResource)(nil)

func NewAuthenticationvserverAuthenticationwebauthpolicyBindingResource() resource.Resource {
	return &AuthenticationvserverAuthenticationwebauthpolicyBindingResource{}
}

// AuthenticationvserverAuthenticationwebauthpolicyBindingResource defines the resource implementation.
type AuthenticationvserverAuthenticationwebauthpolicyBindingResource struct {
	client *service.NitroClient
}

func (r *AuthenticationvserverAuthenticationwebauthpolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AuthenticationvserverAuthenticationwebauthpolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_authenticationvserver_authenticationwebauthpolicy_binding"
}

func (r *AuthenticationvserverAuthenticationwebauthpolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AuthenticationvserverAuthenticationwebauthpolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AuthenticationvserverAuthenticationwebauthpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating authenticationvserver_authenticationwebauthpolicy_binding resource")

	// authenticationvserver_authenticationwebauthpolicy_binding := authenticationvserver_authenticationwebauthpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Authenticationvserver_authenticationwebauthpolicy_binding.Type(), &authenticationvserver_authenticationwebauthpolicy_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create authenticationvserver_authenticationwebauthpolicy_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("authenticationvserver_authenticationwebauthpolicy_binding-config")

	tflog.Trace(ctx, "Created authenticationvserver_authenticationwebauthpolicy_binding resource")

	// Read the updated state back
	r.readAuthenticationvserverAuthenticationwebauthpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationvserverAuthenticationwebauthpolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AuthenticationvserverAuthenticationwebauthpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading authenticationvserver_authenticationwebauthpolicy_binding resource")

	r.readAuthenticationvserverAuthenticationwebauthpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationvserverAuthenticationwebauthpolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data AuthenticationvserverAuthenticationwebauthpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating authenticationvserver_authenticationwebauthpolicy_binding resource")

	// Create API request body from the model
	// authenticationvserver_authenticationwebauthpolicy_binding := authenticationvserver_authenticationwebauthpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Authenticationvserver_authenticationwebauthpolicy_binding.Type(), &authenticationvserver_authenticationwebauthpolicy_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update authenticationvserver_authenticationwebauthpolicy_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated authenticationvserver_authenticationwebauthpolicy_binding resource")

	// Read the updated state back
	r.readAuthenticationvserverAuthenticationwebauthpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationvserverAuthenticationwebauthpolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AuthenticationvserverAuthenticationwebauthpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting authenticationvserver_authenticationwebauthpolicy_binding resource")

	// For authenticationvserver_authenticationwebauthpolicy_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted authenticationvserver_authenticationwebauthpolicy_binding resource from state")
}

// Helper function to read authenticationvserver_authenticationwebauthpolicy_binding data from API
func (r *AuthenticationvserverAuthenticationwebauthpolicyBindingResource) readAuthenticationvserverAuthenticationwebauthpolicyBindingFromApi(ctx context.Context, data *AuthenticationvserverAuthenticationwebauthpolicyBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Authenticationvserver_authenticationwebauthpolicy_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read authenticationvserver_authenticationwebauthpolicy_binding, got error: %s", err))
		return
	}

	authenticationvserver_authenticationwebauthpolicy_bindingSetAttrFromGet(ctx, data, getResponseData)

}
