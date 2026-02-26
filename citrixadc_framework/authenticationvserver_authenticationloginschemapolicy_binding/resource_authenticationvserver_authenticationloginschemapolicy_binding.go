package authenticationvserver_authenticationloginschemapolicy_binding

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
var _ resource.Resource = &AuthenticationvserverAuthenticationloginschemapolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*AuthenticationvserverAuthenticationloginschemapolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*AuthenticationvserverAuthenticationloginschemapolicyBindingResource)(nil)

func NewAuthenticationvserverAuthenticationloginschemapolicyBindingResource() resource.Resource {
	return &AuthenticationvserverAuthenticationloginschemapolicyBindingResource{}
}

// AuthenticationvserverAuthenticationloginschemapolicyBindingResource defines the resource implementation.
type AuthenticationvserverAuthenticationloginschemapolicyBindingResource struct {
	client *service.NitroClient
}

func (r *AuthenticationvserverAuthenticationloginschemapolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AuthenticationvserverAuthenticationloginschemapolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_authenticationvserver_authenticationloginschemapolicy_binding"
}

func (r *AuthenticationvserverAuthenticationloginschemapolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AuthenticationvserverAuthenticationloginschemapolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AuthenticationvserverAuthenticationloginschemapolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating authenticationvserver_authenticationloginschemapolicy_binding resource")

	// authenticationvserver_authenticationloginschemapolicy_binding := authenticationvserver_authenticationloginschemapolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Authenticationvserver_authenticationloginschemapolicy_binding.Type(), &authenticationvserver_authenticationloginschemapolicy_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create authenticationvserver_authenticationloginschemapolicy_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("authenticationvserver_authenticationloginschemapolicy_binding-config")

	tflog.Trace(ctx, "Created authenticationvserver_authenticationloginschemapolicy_binding resource")

	// Read the updated state back
	r.readAuthenticationvserverAuthenticationloginschemapolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationvserverAuthenticationloginschemapolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AuthenticationvserverAuthenticationloginschemapolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading authenticationvserver_authenticationloginschemapolicy_binding resource")

	r.readAuthenticationvserverAuthenticationloginschemapolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationvserverAuthenticationloginschemapolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data AuthenticationvserverAuthenticationloginschemapolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating authenticationvserver_authenticationloginschemapolicy_binding resource")

	// Create API request body from the model
	// authenticationvserver_authenticationloginschemapolicy_binding := authenticationvserver_authenticationloginschemapolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Authenticationvserver_authenticationloginschemapolicy_binding.Type(), &authenticationvserver_authenticationloginschemapolicy_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update authenticationvserver_authenticationloginschemapolicy_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated authenticationvserver_authenticationloginschemapolicy_binding resource")

	// Read the updated state back
	r.readAuthenticationvserverAuthenticationloginschemapolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationvserverAuthenticationloginschemapolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AuthenticationvserverAuthenticationloginschemapolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting authenticationvserver_authenticationloginschemapolicy_binding resource")

	// For authenticationvserver_authenticationloginschemapolicy_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted authenticationvserver_authenticationloginschemapolicy_binding resource from state")
}

// Helper function to read authenticationvserver_authenticationloginschemapolicy_binding data from API
func (r *AuthenticationvserverAuthenticationloginschemapolicyBindingResource) readAuthenticationvserverAuthenticationloginschemapolicyBindingFromApi(ctx context.Context, data *AuthenticationvserverAuthenticationloginschemapolicyBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Authenticationvserver_authenticationloginschemapolicy_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read authenticationvserver_authenticationloginschemapolicy_binding, got error: %s", err))
		return
	}

	authenticationvserver_authenticationloginschemapolicy_bindingSetAttrFromGet(ctx, data, getResponseData)

}
