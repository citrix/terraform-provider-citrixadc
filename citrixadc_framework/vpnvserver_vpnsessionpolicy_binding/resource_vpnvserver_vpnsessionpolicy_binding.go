package vpnvserver_vpnsessionpolicy_binding

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
var _ resource.Resource = &VpnvserverVpnsessionpolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*VpnvserverVpnsessionpolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*VpnvserverVpnsessionpolicyBindingResource)(nil)

func NewVpnvserverVpnsessionpolicyBindingResource() resource.Resource {
	return &VpnvserverVpnsessionpolicyBindingResource{}
}

// VpnvserverVpnsessionpolicyBindingResource defines the resource implementation.
type VpnvserverVpnsessionpolicyBindingResource struct {
	client *service.NitroClient
}

func (r *VpnvserverVpnsessionpolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VpnvserverVpnsessionpolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpnvserver_vpnsessionpolicy_binding"
}

func (r *VpnvserverVpnsessionpolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VpnvserverVpnsessionpolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VpnvserverVpnsessionpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating vpnvserver_vpnsessionpolicy_binding resource")

	// vpnvserver_vpnsessionpolicy_binding := vpnvserver_vpnsessionpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Vpnvserver_vpnsessionpolicy_binding.Type(), &vpnvserver_vpnsessionpolicy_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create vpnvserver_vpnsessionpolicy_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("vpnvserver_vpnsessionpolicy_binding-config")

	tflog.Trace(ctx, "Created vpnvserver_vpnsessionpolicy_binding resource")

	// Read the updated state back
	r.readVpnvserverVpnsessionpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnvserverVpnsessionpolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VpnvserverVpnsessionpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading vpnvserver_vpnsessionpolicy_binding resource")

	r.readVpnvserverVpnsessionpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnvserverVpnsessionpolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data VpnvserverVpnsessionpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating vpnvserver_vpnsessionpolicy_binding resource")

	// Create API request body from the model
	// vpnvserver_vpnsessionpolicy_binding := vpnvserver_vpnsessionpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Vpnvserver_vpnsessionpolicy_binding.Type(), &vpnvserver_vpnsessionpolicy_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update vpnvserver_vpnsessionpolicy_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated vpnvserver_vpnsessionpolicy_binding resource")

	// Read the updated state back
	r.readVpnvserverVpnsessionpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnvserverVpnsessionpolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VpnvserverVpnsessionpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting vpnvserver_vpnsessionpolicy_binding resource")

	// For vpnvserver_vpnsessionpolicy_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted vpnvserver_vpnsessionpolicy_binding resource from state")
}

// Helper function to read vpnvserver_vpnsessionpolicy_binding data from API
func (r *VpnvserverVpnsessionpolicyBindingResource) readVpnvserverVpnsessionpolicyBindingFromApi(ctx context.Context, data *VpnvserverVpnsessionpolicyBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Vpnvserver_vpnsessionpolicy_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read vpnvserver_vpnsessionpolicy_binding, got error: %s", err))
		return
	}

	vpnvserver_vpnsessionpolicy_bindingSetAttrFromGet(ctx, data, getResponseData)

}
