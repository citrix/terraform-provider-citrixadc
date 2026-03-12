package vpnvserver_appcontroller_binding

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
var _ resource.Resource = &VpnvserverAppcontrollerBindingResource{}
var _ resource.ResourceWithConfigure = (*VpnvserverAppcontrollerBindingResource)(nil)
var _ resource.ResourceWithImportState = (*VpnvserverAppcontrollerBindingResource)(nil)

func NewVpnvserverAppcontrollerBindingResource() resource.Resource {
	return &VpnvserverAppcontrollerBindingResource{}
}

// VpnvserverAppcontrollerBindingResource defines the resource implementation.
type VpnvserverAppcontrollerBindingResource struct {
	client *service.NitroClient
}

func (r *VpnvserverAppcontrollerBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VpnvserverAppcontrollerBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpnvserver_appcontroller_binding"
}

func (r *VpnvserverAppcontrollerBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VpnvserverAppcontrollerBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VpnvserverAppcontrollerBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating vpnvserver_appcontroller_binding resource")

	// vpnvserver_appcontroller_binding := vpnvserver_appcontroller_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Vpnvserver_appcontroller_binding.Type(), &vpnvserver_appcontroller_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create vpnvserver_appcontroller_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("vpnvserver_appcontroller_binding-config")

	tflog.Trace(ctx, "Created vpnvserver_appcontroller_binding resource")

	// Read the updated state back
	r.readVpnvserverAppcontrollerBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnvserverAppcontrollerBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VpnvserverAppcontrollerBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading vpnvserver_appcontroller_binding resource")

	r.readVpnvserverAppcontrollerBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnvserverAppcontrollerBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data VpnvserverAppcontrollerBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating vpnvserver_appcontroller_binding resource")

	// Create API request body from the model
	// vpnvserver_appcontroller_binding := vpnvserver_appcontroller_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Vpnvserver_appcontroller_binding.Type(), &vpnvserver_appcontroller_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update vpnvserver_appcontroller_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated vpnvserver_appcontroller_binding resource")

	// Read the updated state back
	r.readVpnvserverAppcontrollerBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnvserverAppcontrollerBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VpnvserverAppcontrollerBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting vpnvserver_appcontroller_binding resource")

	// For vpnvserver_appcontroller_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted vpnvserver_appcontroller_binding resource from state")
}

// Helper function to read vpnvserver_appcontroller_binding data from API
func (r *VpnvserverAppcontrollerBindingResource) readVpnvserverAppcontrollerBindingFromApi(ctx context.Context, data *VpnvserverAppcontrollerBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Vpnvserver_appcontroller_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read vpnvserver_appcontroller_binding, got error: %s", err))
		return
	}

	vpnvserver_appcontroller_bindingSetAttrFromGet(ctx, data, getResponseData)

}
