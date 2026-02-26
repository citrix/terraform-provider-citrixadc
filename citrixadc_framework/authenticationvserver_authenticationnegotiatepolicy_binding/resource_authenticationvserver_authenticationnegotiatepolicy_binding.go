package authenticationvserver_authenticationnegotiatepolicy_binding

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
var _ resource.Resource = &AuthenticationvserverAuthenticationnegotiatepolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*AuthenticationvserverAuthenticationnegotiatepolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*AuthenticationvserverAuthenticationnegotiatepolicyBindingResource)(nil)

func NewAuthenticationvserverAuthenticationnegotiatepolicyBindingResource() resource.Resource {
	return &AuthenticationvserverAuthenticationnegotiatepolicyBindingResource{}
}

// AuthenticationvserverAuthenticationnegotiatepolicyBindingResource defines the resource implementation.
type AuthenticationvserverAuthenticationnegotiatepolicyBindingResource struct {
	client *service.NitroClient
}

func (r *AuthenticationvserverAuthenticationnegotiatepolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AuthenticationvserverAuthenticationnegotiatepolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_authenticationvserver_authenticationnegotiatepolicy_binding"
}

func (r *AuthenticationvserverAuthenticationnegotiatepolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AuthenticationvserverAuthenticationnegotiatepolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AuthenticationvserverAuthenticationnegotiatepolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating authenticationvserver_authenticationnegotiatepolicy_binding resource")

	// authenticationvserver_authenticationnegotiatepolicy_binding := authenticationvserver_authenticationnegotiatepolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Authenticationvserver_authenticationnegotiatepolicy_binding.Type(), &authenticationvserver_authenticationnegotiatepolicy_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create authenticationvserver_authenticationnegotiatepolicy_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("authenticationvserver_authenticationnegotiatepolicy_binding-config")

	tflog.Trace(ctx, "Created authenticationvserver_authenticationnegotiatepolicy_binding resource")

	// Read the updated state back
	r.readAuthenticationvserverAuthenticationnegotiatepolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationvserverAuthenticationnegotiatepolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AuthenticationvserverAuthenticationnegotiatepolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading authenticationvserver_authenticationnegotiatepolicy_binding resource")

	r.readAuthenticationvserverAuthenticationnegotiatepolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationvserverAuthenticationnegotiatepolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data AuthenticationvserverAuthenticationnegotiatepolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating authenticationvserver_authenticationnegotiatepolicy_binding resource")

	// Create API request body from the model
	// authenticationvserver_authenticationnegotiatepolicy_binding := authenticationvserver_authenticationnegotiatepolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Authenticationvserver_authenticationnegotiatepolicy_binding.Type(), &authenticationvserver_authenticationnegotiatepolicy_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update authenticationvserver_authenticationnegotiatepolicy_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated authenticationvserver_authenticationnegotiatepolicy_binding resource")

	// Read the updated state back
	r.readAuthenticationvserverAuthenticationnegotiatepolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationvserverAuthenticationnegotiatepolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AuthenticationvserverAuthenticationnegotiatepolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting authenticationvserver_authenticationnegotiatepolicy_binding resource")

	// For authenticationvserver_authenticationnegotiatepolicy_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted authenticationvserver_authenticationnegotiatepolicy_binding resource from state")
}

// Helper function to read authenticationvserver_authenticationnegotiatepolicy_binding data from API
func (r *AuthenticationvserverAuthenticationnegotiatepolicyBindingResource) readAuthenticationvserverAuthenticationnegotiatepolicyBindingFromApi(ctx context.Context, data *AuthenticationvserverAuthenticationnegotiatepolicyBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Authenticationvserver_authenticationnegotiatepolicy_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read authenticationvserver_authenticationnegotiatepolicy_binding, got error: %s", err))
		return
	}

	authenticationvserver_authenticationnegotiatepolicy_bindingSetAttrFromGet(ctx, data, getResponseData)

}
