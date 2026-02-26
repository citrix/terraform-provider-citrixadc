package vpnvserver_authenticationoauthidppolicy_binding

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
var _ resource.Resource = &VpnvserverAuthenticationoauthidppolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*VpnvserverAuthenticationoauthidppolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*VpnvserverAuthenticationoauthidppolicyBindingResource)(nil)

func NewVpnvserverAuthenticationoauthidppolicyBindingResource() resource.Resource {
	return &VpnvserverAuthenticationoauthidppolicyBindingResource{}
}

// VpnvserverAuthenticationoauthidppolicyBindingResource defines the resource implementation.
type VpnvserverAuthenticationoauthidppolicyBindingResource struct {
	client *service.NitroClient
}

func (r *VpnvserverAuthenticationoauthidppolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VpnvserverAuthenticationoauthidppolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpnvserver_authenticationoauthidppolicy_binding"
}

func (r *VpnvserverAuthenticationoauthidppolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VpnvserverAuthenticationoauthidppolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VpnvserverAuthenticationoauthidppolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating vpnvserver_authenticationoauthidppolicy_binding resource")

	// vpnvserver_authenticationoauthidppolicy_binding := vpnvserver_authenticationoauthidppolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Vpnvserver_authenticationoauthidppolicy_binding.Type(), &vpnvserver_authenticationoauthidppolicy_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create vpnvserver_authenticationoauthidppolicy_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("vpnvserver_authenticationoauthidppolicy_binding-config")

	tflog.Trace(ctx, "Created vpnvserver_authenticationoauthidppolicy_binding resource")

	// Read the updated state back
	r.readVpnvserverAuthenticationoauthidppolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnvserverAuthenticationoauthidppolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VpnvserverAuthenticationoauthidppolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading vpnvserver_authenticationoauthidppolicy_binding resource")

	r.readVpnvserverAuthenticationoauthidppolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnvserverAuthenticationoauthidppolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data VpnvserverAuthenticationoauthidppolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating vpnvserver_authenticationoauthidppolicy_binding resource")

	// Create API request body from the model
	// vpnvserver_authenticationoauthidppolicy_binding := vpnvserver_authenticationoauthidppolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Vpnvserver_authenticationoauthidppolicy_binding.Type(), &vpnvserver_authenticationoauthidppolicy_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update vpnvserver_authenticationoauthidppolicy_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated vpnvserver_authenticationoauthidppolicy_binding resource")

	// Read the updated state back
	r.readVpnvserverAuthenticationoauthidppolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnvserverAuthenticationoauthidppolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VpnvserverAuthenticationoauthidppolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting vpnvserver_authenticationoauthidppolicy_binding resource")

	// For vpnvserver_authenticationoauthidppolicy_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted vpnvserver_authenticationoauthidppolicy_binding resource from state")
}

// Helper function to read vpnvserver_authenticationoauthidppolicy_binding data from API
func (r *VpnvserverAuthenticationoauthidppolicyBindingResource) readVpnvserverAuthenticationoauthidppolicyBindingFromApi(ctx context.Context, data *VpnvserverAuthenticationoauthidppolicyBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Vpnvserver_authenticationoauthidppolicy_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read vpnvserver_authenticationoauthidppolicy_binding, got error: %s", err))
		return
	}

	vpnvserver_authenticationoauthidppolicy_bindingSetAttrFromGet(ctx, data, getResponseData)

}
