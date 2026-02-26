package vpnvserver_cspolicy_binding

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
var _ resource.Resource = &VpnvserverCspolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*VpnvserverCspolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*VpnvserverCspolicyBindingResource)(nil)

func NewVpnvserverCspolicyBindingResource() resource.Resource {
	return &VpnvserverCspolicyBindingResource{}
}

// VpnvserverCspolicyBindingResource defines the resource implementation.
type VpnvserverCspolicyBindingResource struct {
	client *service.NitroClient
}

func (r *VpnvserverCspolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VpnvserverCspolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpnvserver_cspolicy_binding"
}

func (r *VpnvserverCspolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VpnvserverCspolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VpnvserverCspolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating vpnvserver_cspolicy_binding resource")

	// vpnvserver_cspolicy_binding := vpnvserver_cspolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Vpnvserver_cspolicy_binding.Type(), &vpnvserver_cspolicy_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create vpnvserver_cspolicy_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("vpnvserver_cspolicy_binding-config")

	tflog.Trace(ctx, "Created vpnvserver_cspolicy_binding resource")

	// Read the updated state back
	r.readVpnvserverCspolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnvserverCspolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VpnvserverCspolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading vpnvserver_cspolicy_binding resource")

	r.readVpnvserverCspolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnvserverCspolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data VpnvserverCspolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating vpnvserver_cspolicy_binding resource")

	// Create API request body from the model
	// vpnvserver_cspolicy_binding := vpnvserver_cspolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Vpnvserver_cspolicy_binding.Type(), &vpnvserver_cspolicy_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update vpnvserver_cspolicy_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated vpnvserver_cspolicy_binding resource")

	// Read the updated state back
	r.readVpnvserverCspolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnvserverCspolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VpnvserverCspolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting vpnvserver_cspolicy_binding resource")

	// For vpnvserver_cspolicy_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted vpnvserver_cspolicy_binding resource from state")
}

// Helper function to read vpnvserver_cspolicy_binding data from API
func (r *VpnvserverCspolicyBindingResource) readVpnvserverCspolicyBindingFromApi(ctx context.Context, data *VpnvserverCspolicyBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Vpnvserver_cspolicy_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read vpnvserver_cspolicy_binding, got error: %s", err))
		return
	}

	vpnvserver_cspolicy_bindingSetAttrFromGet(ctx, data, getResponseData)

}
