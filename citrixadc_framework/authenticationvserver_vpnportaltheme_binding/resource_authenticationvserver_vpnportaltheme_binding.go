package authenticationvserver_vpnportaltheme_binding

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
var _ resource.Resource = &AuthenticationvserverVpnportalthemeBindingResource{}
var _ resource.ResourceWithConfigure = (*AuthenticationvserverVpnportalthemeBindingResource)(nil)
var _ resource.ResourceWithImportState = (*AuthenticationvserverVpnportalthemeBindingResource)(nil)

func NewAuthenticationvserverVpnportalthemeBindingResource() resource.Resource {
	return &AuthenticationvserverVpnportalthemeBindingResource{}
}

// AuthenticationvserverVpnportalthemeBindingResource defines the resource implementation.
type AuthenticationvserverVpnportalthemeBindingResource struct {
	client *service.NitroClient
}

func (r *AuthenticationvserverVpnportalthemeBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AuthenticationvserverVpnportalthemeBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_authenticationvserver_vpnportaltheme_binding"
}

func (r *AuthenticationvserverVpnportalthemeBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AuthenticationvserverVpnportalthemeBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AuthenticationvserverVpnportalthemeBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating authenticationvserver_vpnportaltheme_binding resource")

	// authenticationvserver_vpnportaltheme_binding := authenticationvserver_vpnportaltheme_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Authenticationvserver_vpnportaltheme_binding.Type(), &authenticationvserver_vpnportaltheme_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create authenticationvserver_vpnportaltheme_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("authenticationvserver_vpnportaltheme_binding-config")

	tflog.Trace(ctx, "Created authenticationvserver_vpnportaltheme_binding resource")

	// Read the updated state back
	r.readAuthenticationvserverVpnportalthemeBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationvserverVpnportalthemeBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AuthenticationvserverVpnportalthemeBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading authenticationvserver_vpnportaltheme_binding resource")

	r.readAuthenticationvserverVpnportalthemeBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationvserverVpnportalthemeBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data AuthenticationvserverVpnportalthemeBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating authenticationvserver_vpnportaltheme_binding resource")

	// Create API request body from the model
	// authenticationvserver_vpnportaltheme_binding := authenticationvserver_vpnportaltheme_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Authenticationvserver_vpnportaltheme_binding.Type(), &authenticationvserver_vpnportaltheme_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update authenticationvserver_vpnportaltheme_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated authenticationvserver_vpnportaltheme_binding resource")

	// Read the updated state back
	r.readAuthenticationvserverVpnportalthemeBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationvserverVpnportalthemeBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AuthenticationvserverVpnportalthemeBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting authenticationvserver_vpnportaltheme_binding resource")

	// For authenticationvserver_vpnportaltheme_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted authenticationvserver_vpnportaltheme_binding resource from state")
}

// Helper function to read authenticationvserver_vpnportaltheme_binding data from API
func (r *AuthenticationvserverVpnportalthemeBindingResource) readAuthenticationvserverVpnportalthemeBindingFromApi(ctx context.Context, data *AuthenticationvserverVpnportalthemeBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Authenticationvserver_vpnportaltheme_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read authenticationvserver_vpnportaltheme_binding, got error: %s", err))
		return
	}

	authenticationvserver_vpnportaltheme_bindingSetAttrFromGet(ctx, data, getResponseData)

}
