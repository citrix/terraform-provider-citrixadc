package vpnvserver_vpnnexthopserver_binding

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
var _ resource.Resource = &VpnvserverVpnnexthopserverBindingResource{}
var _ resource.ResourceWithConfigure = (*VpnvserverVpnnexthopserverBindingResource)(nil)
var _ resource.ResourceWithImportState = (*VpnvserverVpnnexthopserverBindingResource)(nil)

func NewVpnvserverVpnnexthopserverBindingResource() resource.Resource {
	return &VpnvserverVpnnexthopserverBindingResource{}
}

// VpnvserverVpnnexthopserverBindingResource defines the resource implementation.
type VpnvserverVpnnexthopserverBindingResource struct {
	client *service.NitroClient
}

func (r *VpnvserverVpnnexthopserverBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VpnvserverVpnnexthopserverBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpnvserver_vpnnexthopserver_binding"
}

func (r *VpnvserverVpnnexthopserverBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VpnvserverVpnnexthopserverBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VpnvserverVpnnexthopserverBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating vpnvserver_vpnnexthopserver_binding resource")

	// vpnvserver_vpnnexthopserver_binding := vpnvserver_vpnnexthopserver_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Vpnvserver_vpnnexthopserver_binding.Type(), &vpnvserver_vpnnexthopserver_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create vpnvserver_vpnnexthopserver_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("vpnvserver_vpnnexthopserver_binding-config")

	tflog.Trace(ctx, "Created vpnvserver_vpnnexthopserver_binding resource")

	// Read the updated state back
	r.readVpnvserverVpnnexthopserverBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnvserverVpnnexthopserverBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VpnvserverVpnnexthopserverBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading vpnvserver_vpnnexthopserver_binding resource")

	r.readVpnvserverVpnnexthopserverBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnvserverVpnnexthopserverBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data VpnvserverVpnnexthopserverBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating vpnvserver_vpnnexthopserver_binding resource")

	// Create API request body from the model
	// vpnvserver_vpnnexthopserver_binding := vpnvserver_vpnnexthopserver_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Vpnvserver_vpnnexthopserver_binding.Type(), &vpnvserver_vpnnexthopserver_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update vpnvserver_vpnnexthopserver_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated vpnvserver_vpnnexthopserver_binding resource")

	// Read the updated state back
	r.readVpnvserverVpnnexthopserverBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnvserverVpnnexthopserverBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VpnvserverVpnnexthopserverBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting vpnvserver_vpnnexthopserver_binding resource")

	// For vpnvserver_vpnnexthopserver_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted vpnvserver_vpnnexthopserver_binding resource from state")
}

// Helper function to read vpnvserver_vpnnexthopserver_binding data from API
func (r *VpnvserverVpnnexthopserverBindingResource) readVpnvserverVpnnexthopserverBindingFromApi(ctx context.Context, data *VpnvserverVpnnexthopserverBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Vpnvserver_vpnnexthopserver_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read vpnvserver_vpnnexthopserver_binding, got error: %s", err))
		return
	}

	vpnvserver_vpnnexthopserver_bindingSetAttrFromGet(ctx, data, getResponseData)

}
