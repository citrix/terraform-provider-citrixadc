package authenticationvserver_authenticationoauthidppolicy_binding

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
var _ resource.Resource = &AuthenticationvserverAuthenticationoauthidppolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*AuthenticationvserverAuthenticationoauthidppolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*AuthenticationvserverAuthenticationoauthidppolicyBindingResource)(nil)

func NewAuthenticationvserverAuthenticationoauthidppolicyBindingResource() resource.Resource {
	return &AuthenticationvserverAuthenticationoauthidppolicyBindingResource{}
}

// AuthenticationvserverAuthenticationoauthidppolicyBindingResource defines the resource implementation.
type AuthenticationvserverAuthenticationoauthidppolicyBindingResource struct {
	client *service.NitroClient
}

func (r *AuthenticationvserverAuthenticationoauthidppolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AuthenticationvserverAuthenticationoauthidppolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_authenticationvserver_authenticationoauthidppolicy_binding"
}

func (r *AuthenticationvserverAuthenticationoauthidppolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AuthenticationvserverAuthenticationoauthidppolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AuthenticationvserverAuthenticationoauthidppolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating authenticationvserver_authenticationoauthidppolicy_binding resource")

	// authenticationvserver_authenticationoauthidppolicy_binding := authenticationvserver_authenticationoauthidppolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Authenticationvserver_authenticationoauthidppolicy_binding.Type(), &authenticationvserver_authenticationoauthidppolicy_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create authenticationvserver_authenticationoauthidppolicy_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("authenticationvserver_authenticationoauthidppolicy_binding-config")

	tflog.Trace(ctx, "Created authenticationvserver_authenticationoauthidppolicy_binding resource")

	// Read the updated state back
	r.readAuthenticationvserverAuthenticationoauthidppolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationvserverAuthenticationoauthidppolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AuthenticationvserverAuthenticationoauthidppolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading authenticationvserver_authenticationoauthidppolicy_binding resource")

	r.readAuthenticationvserverAuthenticationoauthidppolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationvserverAuthenticationoauthidppolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data AuthenticationvserverAuthenticationoauthidppolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating authenticationvserver_authenticationoauthidppolicy_binding resource")

	// Create API request body from the model
	// authenticationvserver_authenticationoauthidppolicy_binding := authenticationvserver_authenticationoauthidppolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Authenticationvserver_authenticationoauthidppolicy_binding.Type(), &authenticationvserver_authenticationoauthidppolicy_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update authenticationvserver_authenticationoauthidppolicy_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated authenticationvserver_authenticationoauthidppolicy_binding resource")

	// Read the updated state back
	r.readAuthenticationvserverAuthenticationoauthidppolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationvserverAuthenticationoauthidppolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AuthenticationvserverAuthenticationoauthidppolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting authenticationvserver_authenticationoauthidppolicy_binding resource")

	// For authenticationvserver_authenticationoauthidppolicy_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted authenticationvserver_authenticationoauthidppolicy_binding resource from state")
}

// Helper function to read authenticationvserver_authenticationoauthidppolicy_binding data from API
func (r *AuthenticationvserverAuthenticationoauthidppolicyBindingResource) readAuthenticationvserverAuthenticationoauthidppolicyBindingFromApi(ctx context.Context, data *AuthenticationvserverAuthenticationoauthidppolicyBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Authenticationvserver_authenticationoauthidppolicy_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read authenticationvserver_authenticationoauthidppolicy_binding, got error: %s", err))
		return
	}

	authenticationvserver_authenticationoauthidppolicy_bindingSetAttrFromGet(ctx, data, getResponseData)

}
