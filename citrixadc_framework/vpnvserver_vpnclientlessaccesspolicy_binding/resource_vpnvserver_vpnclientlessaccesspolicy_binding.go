package vpnvserver_vpnclientlessaccesspolicy_binding

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
var _ resource.Resource = &VpnvserverVpnclientlessaccesspolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*VpnvserverVpnclientlessaccesspolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*VpnvserverVpnclientlessaccesspolicyBindingResource)(nil)

func NewVpnvserverVpnclientlessaccesspolicyBindingResource() resource.Resource {
	return &VpnvserverVpnclientlessaccesspolicyBindingResource{}
}

// VpnvserverVpnclientlessaccesspolicyBindingResource defines the resource implementation.
type VpnvserverVpnclientlessaccesspolicyBindingResource struct {
	client *service.NitroClient
}

func (r *VpnvserverVpnclientlessaccesspolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VpnvserverVpnclientlessaccesspolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpnvserver_vpnclientlessaccesspolicy_binding"
}

func (r *VpnvserverVpnclientlessaccesspolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VpnvserverVpnclientlessaccesspolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VpnvserverVpnclientlessaccesspolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating vpnvserver_vpnclientlessaccesspolicy_binding resource")

	// vpnvserver_vpnclientlessaccesspolicy_binding := vpnvserver_vpnclientlessaccesspolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Vpnvserver_vpnclientlessaccesspolicy_binding.Type(), &vpnvserver_vpnclientlessaccesspolicy_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create vpnvserver_vpnclientlessaccesspolicy_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("vpnvserver_vpnclientlessaccesspolicy_binding-config")

	tflog.Trace(ctx, "Created vpnvserver_vpnclientlessaccesspolicy_binding resource")

	// Read the updated state back
	r.readVpnvserverVpnclientlessaccesspolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnvserverVpnclientlessaccesspolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VpnvserverVpnclientlessaccesspolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading vpnvserver_vpnclientlessaccesspolicy_binding resource")

	r.readVpnvserverVpnclientlessaccesspolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnvserverVpnclientlessaccesspolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data VpnvserverVpnclientlessaccesspolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating vpnvserver_vpnclientlessaccesspolicy_binding resource")

	// Create API request body from the model
	// vpnvserver_vpnclientlessaccesspolicy_binding := vpnvserver_vpnclientlessaccesspolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Vpnvserver_vpnclientlessaccesspolicy_binding.Type(), &vpnvserver_vpnclientlessaccesspolicy_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update vpnvserver_vpnclientlessaccesspolicy_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated vpnvserver_vpnclientlessaccesspolicy_binding resource")

	// Read the updated state back
	r.readVpnvserverVpnclientlessaccesspolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnvserverVpnclientlessaccesspolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VpnvserverVpnclientlessaccesspolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting vpnvserver_vpnclientlessaccesspolicy_binding resource")

	// For vpnvserver_vpnclientlessaccesspolicy_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted vpnvserver_vpnclientlessaccesspolicy_binding resource from state")
}

// Helper function to read vpnvserver_vpnclientlessaccesspolicy_binding data from API
func (r *VpnvserverVpnclientlessaccesspolicyBindingResource) readVpnvserverVpnclientlessaccesspolicyBindingFromApi(ctx context.Context, data *VpnvserverVpnclientlessaccesspolicyBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Vpnvserver_vpnclientlessaccesspolicy_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read vpnvserver_vpnclientlessaccesspolicy_binding, got error: %s", err))
		return
	}

	vpnvserver_vpnclientlessaccesspolicy_bindingSetAttrFromGet(ctx, data, getResponseData)

}
