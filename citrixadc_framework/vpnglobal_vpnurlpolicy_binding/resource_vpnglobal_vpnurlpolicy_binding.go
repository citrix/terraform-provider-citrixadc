package vpnglobal_vpnurlpolicy_binding

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
var _ resource.Resource = &VpnglobalVpnurlpolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*VpnglobalVpnurlpolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*VpnglobalVpnurlpolicyBindingResource)(nil)

func NewVpnglobalVpnurlpolicyBindingResource() resource.Resource {
	return &VpnglobalVpnurlpolicyBindingResource{}
}

// VpnglobalVpnurlpolicyBindingResource defines the resource implementation.
type VpnglobalVpnurlpolicyBindingResource struct {
	client *service.NitroClient
}

func (r *VpnglobalVpnurlpolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VpnglobalVpnurlpolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpnglobal_vpnurlpolicy_binding"
}

func (r *VpnglobalVpnurlpolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VpnglobalVpnurlpolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VpnglobalVpnurlpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating vpnglobal_vpnurlpolicy_binding resource")

	// vpnglobal_vpnurlpolicy_binding := vpnglobal_vpnurlpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Vpnglobal_vpnurlpolicy_binding.Type(), &vpnglobal_vpnurlpolicy_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create vpnglobal_vpnurlpolicy_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("vpnglobal_vpnurlpolicy_binding-config")

	tflog.Trace(ctx, "Created vpnglobal_vpnurlpolicy_binding resource")

	// Read the updated state back
	r.readVpnglobalVpnurlpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnglobalVpnurlpolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VpnglobalVpnurlpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading vpnglobal_vpnurlpolicy_binding resource")

	r.readVpnglobalVpnurlpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnglobalVpnurlpolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data VpnglobalVpnurlpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating vpnglobal_vpnurlpolicy_binding resource")

	// Create API request body from the model
	// vpnglobal_vpnurlpolicy_binding := vpnglobal_vpnurlpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Vpnglobal_vpnurlpolicy_binding.Type(), &vpnglobal_vpnurlpolicy_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update vpnglobal_vpnurlpolicy_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated vpnglobal_vpnurlpolicy_binding resource")

	// Read the updated state back
	r.readVpnglobalVpnurlpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnglobalVpnurlpolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VpnglobalVpnurlpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting vpnglobal_vpnurlpolicy_binding resource")

	// For vpnglobal_vpnurlpolicy_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted vpnglobal_vpnurlpolicy_binding resource from state")
}

// Helper function to read vpnglobal_vpnurlpolicy_binding data from API
func (r *VpnglobalVpnurlpolicyBindingResource) readVpnglobalVpnurlpolicyBindingFromApi(ctx context.Context, data *VpnglobalVpnurlpolicyBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Vpnglobal_vpnurlpolicy_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read vpnglobal_vpnurlpolicy_binding, got error: %s", err))
		return
	}

	vpnglobal_vpnurlpolicy_bindingSetAttrFromGet(ctx, data, getResponseData)

}
