package vpnvserver_vpnurlpolicy_binding

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
var _ resource.Resource = &VpnvserverVpnurlpolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*VpnvserverVpnurlpolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*VpnvserverVpnurlpolicyBindingResource)(nil)

func NewVpnvserverVpnurlpolicyBindingResource() resource.Resource {
	return &VpnvserverVpnurlpolicyBindingResource{}
}

// VpnvserverVpnurlpolicyBindingResource defines the resource implementation.
type VpnvserverVpnurlpolicyBindingResource struct {
	client *service.NitroClient
}

func (r *VpnvserverVpnurlpolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VpnvserverVpnurlpolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpnvserver_vpnurlpolicy_binding"
}

func (r *VpnvserverVpnurlpolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VpnvserverVpnurlpolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VpnvserverVpnurlpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating vpnvserver_vpnurlpolicy_binding resource")

	// vpnvserver_vpnurlpolicy_binding := vpnvserver_vpnurlpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Vpnvserver_vpnurlpolicy_binding.Type(), &vpnvserver_vpnurlpolicy_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create vpnvserver_vpnurlpolicy_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("vpnvserver_vpnurlpolicy_binding-config")

	tflog.Trace(ctx, "Created vpnvserver_vpnurlpolicy_binding resource")

	// Read the updated state back
	r.readVpnvserverVpnurlpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnvserverVpnurlpolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VpnvserverVpnurlpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading vpnvserver_vpnurlpolicy_binding resource")

	r.readVpnvserverVpnurlpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnvserverVpnurlpolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data VpnvserverVpnurlpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating vpnvserver_vpnurlpolicy_binding resource")

	// Create API request body from the model
	// vpnvserver_vpnurlpolicy_binding := vpnvserver_vpnurlpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Vpnvserver_vpnurlpolicy_binding.Type(), &vpnvserver_vpnurlpolicy_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update vpnvserver_vpnurlpolicy_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated vpnvserver_vpnurlpolicy_binding resource")

	// Read the updated state back
	r.readVpnvserverVpnurlpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnvserverVpnurlpolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VpnvserverVpnurlpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting vpnvserver_vpnurlpolicy_binding resource")

	// For vpnvserver_vpnurlpolicy_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted vpnvserver_vpnurlpolicy_binding resource from state")
}

// Helper function to read vpnvserver_vpnurlpolicy_binding data from API
func (r *VpnvserverVpnurlpolicyBindingResource) readVpnvserverVpnurlpolicyBindingFromApi(ctx context.Context, data *VpnvserverVpnurlpolicyBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Vpnvserver_vpnurlpolicy_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read vpnvserver_vpnurlpolicy_binding, got error: %s", err))
		return
	}

	vpnvserver_vpnurlpolicy_bindingSetAttrFromGet(ctx, data, getResponseData)

}
