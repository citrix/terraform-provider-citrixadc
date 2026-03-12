package vpnvserver_authenticationtacacspolicy_binding

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
var _ resource.Resource = &VpnvserverAuthenticationtacacspolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*VpnvserverAuthenticationtacacspolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*VpnvserverAuthenticationtacacspolicyBindingResource)(nil)

func NewVpnvserverAuthenticationtacacspolicyBindingResource() resource.Resource {
	return &VpnvserverAuthenticationtacacspolicyBindingResource{}
}

// VpnvserverAuthenticationtacacspolicyBindingResource defines the resource implementation.
type VpnvserverAuthenticationtacacspolicyBindingResource struct {
	client *service.NitroClient
}

func (r *VpnvserverAuthenticationtacacspolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VpnvserverAuthenticationtacacspolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpnvserver_authenticationtacacspolicy_binding"
}

func (r *VpnvserverAuthenticationtacacspolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VpnvserverAuthenticationtacacspolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VpnvserverAuthenticationtacacspolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating vpnvserver_authenticationtacacspolicy_binding resource")

	// vpnvserver_authenticationtacacspolicy_binding := vpnvserver_authenticationtacacspolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Vpnvserver_authenticationtacacspolicy_binding.Type(), &vpnvserver_authenticationtacacspolicy_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create vpnvserver_authenticationtacacspolicy_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("vpnvserver_authenticationtacacspolicy_binding-config")

	tflog.Trace(ctx, "Created vpnvserver_authenticationtacacspolicy_binding resource")

	// Read the updated state back
	r.readVpnvserverAuthenticationtacacspolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnvserverAuthenticationtacacspolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VpnvserverAuthenticationtacacspolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading vpnvserver_authenticationtacacspolicy_binding resource")

	r.readVpnvserverAuthenticationtacacspolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnvserverAuthenticationtacacspolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data VpnvserverAuthenticationtacacspolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating vpnvserver_authenticationtacacspolicy_binding resource")

	// Create API request body from the model
	// vpnvserver_authenticationtacacspolicy_binding := vpnvserver_authenticationtacacspolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Vpnvserver_authenticationtacacspolicy_binding.Type(), &vpnvserver_authenticationtacacspolicy_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update vpnvserver_authenticationtacacspolicy_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated vpnvserver_authenticationtacacspolicy_binding resource")

	// Read the updated state back
	r.readVpnvserverAuthenticationtacacspolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnvserverAuthenticationtacacspolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VpnvserverAuthenticationtacacspolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting vpnvserver_authenticationtacacspolicy_binding resource")

	// For vpnvserver_authenticationtacacspolicy_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted vpnvserver_authenticationtacacspolicy_binding resource from state")
}

// Helper function to read vpnvserver_authenticationtacacspolicy_binding data from API
func (r *VpnvserverAuthenticationtacacspolicyBindingResource) readVpnvserverAuthenticationtacacspolicyBindingFromApi(ctx context.Context, data *VpnvserverAuthenticationtacacspolicyBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Vpnvserver_authenticationtacacspolicy_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read vpnvserver_authenticationtacacspolicy_binding, got error: %s", err))
		return
	}

	vpnvserver_authenticationtacacspolicy_bindingSetAttrFromGet(ctx, data, getResponseData)

}
