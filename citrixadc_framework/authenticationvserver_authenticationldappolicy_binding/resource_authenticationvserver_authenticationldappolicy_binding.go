package authenticationvserver_authenticationldappolicy_binding

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
var _ resource.Resource = &AuthenticationvserverAuthenticationldappolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*AuthenticationvserverAuthenticationldappolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*AuthenticationvserverAuthenticationldappolicyBindingResource)(nil)

func NewAuthenticationvserverAuthenticationldappolicyBindingResource() resource.Resource {
	return &AuthenticationvserverAuthenticationldappolicyBindingResource{}
}

// AuthenticationvserverAuthenticationldappolicyBindingResource defines the resource implementation.
type AuthenticationvserverAuthenticationldappolicyBindingResource struct {
	client *service.NitroClient
}

func (r *AuthenticationvserverAuthenticationldappolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AuthenticationvserverAuthenticationldappolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_authenticationvserver_authenticationldappolicy_binding"
}

func (r *AuthenticationvserverAuthenticationldappolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AuthenticationvserverAuthenticationldappolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AuthenticationvserverAuthenticationldappolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating authenticationvserver_authenticationldappolicy_binding resource")

	// authenticationvserver_authenticationldappolicy_binding := authenticationvserver_authenticationldappolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Authenticationvserver_authenticationldappolicy_binding.Type(), &authenticationvserver_authenticationldappolicy_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create authenticationvserver_authenticationldappolicy_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("authenticationvserver_authenticationldappolicy_binding-config")

	tflog.Trace(ctx, "Created authenticationvserver_authenticationldappolicy_binding resource")

	// Read the updated state back
	r.readAuthenticationvserverAuthenticationldappolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationvserverAuthenticationldappolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AuthenticationvserverAuthenticationldappolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading authenticationvserver_authenticationldappolicy_binding resource")

	r.readAuthenticationvserverAuthenticationldappolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationvserverAuthenticationldappolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data AuthenticationvserverAuthenticationldappolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating authenticationvserver_authenticationldappolicy_binding resource")

	// Create API request body from the model
	// authenticationvserver_authenticationldappolicy_binding := authenticationvserver_authenticationldappolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Authenticationvserver_authenticationldappolicy_binding.Type(), &authenticationvserver_authenticationldappolicy_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update authenticationvserver_authenticationldappolicy_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated authenticationvserver_authenticationldappolicy_binding resource")

	// Read the updated state back
	r.readAuthenticationvserverAuthenticationldappolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationvserverAuthenticationldappolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AuthenticationvserverAuthenticationldappolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting authenticationvserver_authenticationldappolicy_binding resource")

	// For authenticationvserver_authenticationldappolicy_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted authenticationvserver_authenticationldappolicy_binding resource from state")
}

// Helper function to read authenticationvserver_authenticationldappolicy_binding data from API
func (r *AuthenticationvserverAuthenticationldappolicyBindingResource) readAuthenticationvserverAuthenticationldappolicyBindingFromApi(ctx context.Context, data *AuthenticationvserverAuthenticationldappolicyBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Authenticationvserver_authenticationldappolicy_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read authenticationvserver_authenticationldappolicy_binding, got error: %s", err))
		return
	}

	authenticationvserver_authenticationldappolicy_bindingSetAttrFromGet(ctx, data, getResponseData)

}
