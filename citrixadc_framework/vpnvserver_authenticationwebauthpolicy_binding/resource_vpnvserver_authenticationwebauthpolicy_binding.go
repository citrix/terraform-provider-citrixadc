package vpnvserver_authenticationwebauthpolicy_binding

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
var _ resource.Resource = &VpnvserverAuthenticationwebauthpolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*VpnvserverAuthenticationwebauthpolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*VpnvserverAuthenticationwebauthpolicyBindingResource)(nil)

func NewVpnvserverAuthenticationwebauthpolicyBindingResource() resource.Resource {
	return &VpnvserverAuthenticationwebauthpolicyBindingResource{}
}

// VpnvserverAuthenticationwebauthpolicyBindingResource defines the resource implementation.
type VpnvserverAuthenticationwebauthpolicyBindingResource struct {
	client *service.NitroClient
}

func (r *VpnvserverAuthenticationwebauthpolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VpnvserverAuthenticationwebauthpolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpnvserver_authenticationwebauthpolicy_binding"
}

func (r *VpnvserverAuthenticationwebauthpolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VpnvserverAuthenticationwebauthpolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VpnvserverAuthenticationwebauthpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating vpnvserver_authenticationwebauthpolicy_binding resource")

	// vpnvserver_authenticationwebauthpolicy_binding := vpnvserver_authenticationwebauthpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Vpnvserver_authenticationwebauthpolicy_binding.Type(), &vpnvserver_authenticationwebauthpolicy_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create vpnvserver_authenticationwebauthpolicy_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("vpnvserver_authenticationwebauthpolicy_binding-config")

	tflog.Trace(ctx, "Created vpnvserver_authenticationwebauthpolicy_binding resource")

	// Read the updated state back
	r.readVpnvserverAuthenticationwebauthpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnvserverAuthenticationwebauthpolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VpnvserverAuthenticationwebauthpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading vpnvserver_authenticationwebauthpolicy_binding resource")

	r.readVpnvserverAuthenticationwebauthpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnvserverAuthenticationwebauthpolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data VpnvserverAuthenticationwebauthpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating vpnvserver_authenticationwebauthpolicy_binding resource")

	// Create API request body from the model
	// vpnvserver_authenticationwebauthpolicy_binding := vpnvserver_authenticationwebauthpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Vpnvserver_authenticationwebauthpolicy_binding.Type(), &vpnvserver_authenticationwebauthpolicy_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update vpnvserver_authenticationwebauthpolicy_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated vpnvserver_authenticationwebauthpolicy_binding resource")

	// Read the updated state back
	r.readVpnvserverAuthenticationwebauthpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnvserverAuthenticationwebauthpolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VpnvserverAuthenticationwebauthpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting vpnvserver_authenticationwebauthpolicy_binding resource")

	// For vpnvserver_authenticationwebauthpolicy_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted vpnvserver_authenticationwebauthpolicy_binding resource from state")
}

// Helper function to read vpnvserver_authenticationwebauthpolicy_binding data from API
func (r *VpnvserverAuthenticationwebauthpolicyBindingResource) readVpnvserverAuthenticationwebauthpolicyBindingFromApi(ctx context.Context, data *VpnvserverAuthenticationwebauthpolicyBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Vpnvserver_authenticationwebauthpolicy_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read vpnvserver_authenticationwebauthpolicy_binding, got error: %s", err))
		return
	}

	vpnvserver_authenticationwebauthpolicy_bindingSetAttrFromGet(ctx, data, getResponseData)

}
