package netbridge_iptunnel_binding

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
var _ resource.Resource = &NetbridgeIptunnelBindingResource{}
var _ resource.ResourceWithConfigure = (*NetbridgeIptunnelBindingResource)(nil)
var _ resource.ResourceWithImportState = (*NetbridgeIptunnelBindingResource)(nil)

func NewNetbridgeIptunnelBindingResource() resource.Resource {
	return &NetbridgeIptunnelBindingResource{}
}

// NetbridgeIptunnelBindingResource defines the resource implementation.
type NetbridgeIptunnelBindingResource struct {
	client *service.NitroClient
}

func (r *NetbridgeIptunnelBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NetbridgeIptunnelBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_netbridge_iptunnel_binding"
}

func (r *NetbridgeIptunnelBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NetbridgeIptunnelBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NetbridgeIptunnelBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating netbridge_iptunnel_binding resource")

	// netbridge_iptunnel_binding := netbridge_iptunnel_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Netbridge_iptunnel_binding.Type(), &netbridge_iptunnel_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create netbridge_iptunnel_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("netbridge_iptunnel_binding-config")

	tflog.Trace(ctx, "Created netbridge_iptunnel_binding resource")

	// Read the updated state back
	r.readNetbridgeIptunnelBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NetbridgeIptunnelBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetbridgeIptunnelBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading netbridge_iptunnel_binding resource")

	r.readNetbridgeIptunnelBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NetbridgeIptunnelBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NetbridgeIptunnelBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating netbridge_iptunnel_binding resource")

	// Create API request body from the model
	// netbridge_iptunnel_binding := netbridge_iptunnel_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Netbridge_iptunnel_binding.Type(), &netbridge_iptunnel_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update netbridge_iptunnel_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated netbridge_iptunnel_binding resource")

	// Read the updated state back
	r.readNetbridgeIptunnelBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NetbridgeIptunnelBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data NetbridgeIptunnelBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting netbridge_iptunnel_binding resource")

	// For netbridge_iptunnel_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted netbridge_iptunnel_binding resource from state")
}

// Helper function to read netbridge_iptunnel_binding data from API
func (r *NetbridgeIptunnelBindingResource) readNetbridgeIptunnelBindingFromApi(ctx context.Context, data *NetbridgeIptunnelBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Netbridge_iptunnel_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read netbridge_iptunnel_binding, got error: %s", err))
		return
	}

	netbridge_iptunnel_bindingSetAttrFromGet(ctx, data, getResponseData)

}
