package authenticationvserver_authenticationpolicy_binding

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
var _ resource.Resource = &AuthenticationvserverAuthenticationpolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*AuthenticationvserverAuthenticationpolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*AuthenticationvserverAuthenticationpolicyBindingResource)(nil)

func NewAuthenticationvserverAuthenticationpolicyBindingResource() resource.Resource {
	return &AuthenticationvserverAuthenticationpolicyBindingResource{}
}

// AuthenticationvserverAuthenticationpolicyBindingResource defines the resource implementation.
type AuthenticationvserverAuthenticationpolicyBindingResource struct {
	client *service.NitroClient
}

func (r *AuthenticationvserverAuthenticationpolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AuthenticationvserverAuthenticationpolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_authenticationvserver_authenticationpolicy_binding"
}

func (r *AuthenticationvserverAuthenticationpolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AuthenticationvserverAuthenticationpolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AuthenticationvserverAuthenticationpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating authenticationvserver_authenticationpolicy_binding resource")

	// authenticationvserver_authenticationpolicy_binding := authenticationvserver_authenticationpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Authenticationvserver_authenticationpolicy_binding.Type(), &authenticationvserver_authenticationpolicy_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create authenticationvserver_authenticationpolicy_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("authenticationvserver_authenticationpolicy_binding-config")

	tflog.Trace(ctx, "Created authenticationvserver_authenticationpolicy_binding resource")

	// Read the updated state back
	r.readAuthenticationvserverAuthenticationpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationvserverAuthenticationpolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AuthenticationvserverAuthenticationpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading authenticationvserver_authenticationpolicy_binding resource")

	r.readAuthenticationvserverAuthenticationpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationvserverAuthenticationpolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data AuthenticationvserverAuthenticationpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating authenticationvserver_authenticationpolicy_binding resource")

	// Create API request body from the model
	// authenticationvserver_authenticationpolicy_binding := authenticationvserver_authenticationpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Authenticationvserver_authenticationpolicy_binding.Type(), &authenticationvserver_authenticationpolicy_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update authenticationvserver_authenticationpolicy_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated authenticationvserver_authenticationpolicy_binding resource")

	// Read the updated state back
	r.readAuthenticationvserverAuthenticationpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationvserverAuthenticationpolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AuthenticationvserverAuthenticationpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting authenticationvserver_authenticationpolicy_binding resource")

	// For authenticationvserver_authenticationpolicy_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted authenticationvserver_authenticationpolicy_binding resource from state")
}

// Helper function to read authenticationvserver_authenticationpolicy_binding data from API
func (r *AuthenticationvserverAuthenticationpolicyBindingResource) readAuthenticationvserverAuthenticationpolicyBindingFromApi(ctx context.Context, data *AuthenticationvserverAuthenticationpolicyBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Authenticationvserver_authenticationpolicy_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read authenticationvserver_authenticationpolicy_binding, got error: %s", err))
		return
	}

	authenticationvserver_authenticationpolicy_bindingSetAttrFromGet(ctx, data, getResponseData)

}
