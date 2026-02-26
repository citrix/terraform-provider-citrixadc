package vpnvserver_authenticationcertpolicy_binding

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
var _ resource.Resource = &VpnvserverAuthenticationcertpolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*VpnvserverAuthenticationcertpolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*VpnvserverAuthenticationcertpolicyBindingResource)(nil)

func NewVpnvserverAuthenticationcertpolicyBindingResource() resource.Resource {
	return &VpnvserverAuthenticationcertpolicyBindingResource{}
}

// VpnvserverAuthenticationcertpolicyBindingResource defines the resource implementation.
type VpnvserverAuthenticationcertpolicyBindingResource struct {
	client *service.NitroClient
}

func (r *VpnvserverAuthenticationcertpolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VpnvserverAuthenticationcertpolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpnvserver_authenticationcertpolicy_binding"
}

func (r *VpnvserverAuthenticationcertpolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VpnvserverAuthenticationcertpolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VpnvserverAuthenticationcertpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating vpnvserver_authenticationcertpolicy_binding resource")

	// vpnvserver_authenticationcertpolicy_binding := vpnvserver_authenticationcertpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Vpnvserver_authenticationcertpolicy_binding.Type(), &vpnvserver_authenticationcertpolicy_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create vpnvserver_authenticationcertpolicy_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("vpnvserver_authenticationcertpolicy_binding-config")

	tflog.Trace(ctx, "Created vpnvserver_authenticationcertpolicy_binding resource")

	// Read the updated state back
	r.readVpnvserverAuthenticationcertpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnvserverAuthenticationcertpolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VpnvserverAuthenticationcertpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading vpnvserver_authenticationcertpolicy_binding resource")

	r.readVpnvserverAuthenticationcertpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnvserverAuthenticationcertpolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data VpnvserverAuthenticationcertpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating vpnvserver_authenticationcertpolicy_binding resource")

	// Create API request body from the model
	// vpnvserver_authenticationcertpolicy_binding := vpnvserver_authenticationcertpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Vpnvserver_authenticationcertpolicy_binding.Type(), &vpnvserver_authenticationcertpolicy_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update vpnvserver_authenticationcertpolicy_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated vpnvserver_authenticationcertpolicy_binding resource")

	// Read the updated state back
	r.readVpnvserverAuthenticationcertpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnvserverAuthenticationcertpolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VpnvserverAuthenticationcertpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting vpnvserver_authenticationcertpolicy_binding resource")

	// For vpnvserver_authenticationcertpolicy_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted vpnvserver_authenticationcertpolicy_binding resource from state")
}

// Helper function to read vpnvserver_authenticationcertpolicy_binding data from API
func (r *VpnvserverAuthenticationcertpolicyBindingResource) readVpnvserverAuthenticationcertpolicyBindingFromApi(ctx context.Context, data *VpnvserverAuthenticationcertpolicyBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Vpnvserver_authenticationcertpolicy_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read vpnvserver_authenticationcertpolicy_binding, got error: %s", err))
		return
	}

	vpnvserver_authenticationcertpolicy_bindingSetAttrFromGet(ctx, data, getResponseData)

}
