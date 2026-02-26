package vpnvserver_intranetip6_binding

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
var _ resource.Resource = &VpnvserverIntranetip6BindingResource{}
var _ resource.ResourceWithConfigure = (*VpnvserverIntranetip6BindingResource)(nil)
var _ resource.ResourceWithImportState = (*VpnvserverIntranetip6BindingResource)(nil)

func NewVpnvserverIntranetip6BindingResource() resource.Resource {
	return &VpnvserverIntranetip6BindingResource{}
}

// VpnvserverIntranetip6BindingResource defines the resource implementation.
type VpnvserverIntranetip6BindingResource struct {
	client *service.NitroClient
}

func (r *VpnvserverIntranetip6BindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VpnvserverIntranetip6BindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpnvserver_intranetip6_binding"
}

func (r *VpnvserverIntranetip6BindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VpnvserverIntranetip6BindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VpnvserverIntranetip6BindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating vpnvserver_intranetip6_binding resource")

	// vpnvserver_intranetip6_binding := vpnvserver_intranetip6_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Vpnvserver_intranetip6_binding.Type(), &vpnvserver_intranetip6_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create vpnvserver_intranetip6_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("vpnvserver_intranetip6_binding-config")

	tflog.Trace(ctx, "Created vpnvserver_intranetip6_binding resource")

	// Read the updated state back
	r.readVpnvserverIntranetip6BindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnvserverIntranetip6BindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VpnvserverIntranetip6BindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading vpnvserver_intranetip6_binding resource")

	r.readVpnvserverIntranetip6BindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnvserverIntranetip6BindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data VpnvserverIntranetip6BindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating vpnvserver_intranetip6_binding resource")

	// Create API request body from the model
	// vpnvserver_intranetip6_binding := vpnvserver_intranetip6_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Vpnvserver_intranetip6_binding.Type(), &vpnvserver_intranetip6_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update vpnvserver_intranetip6_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated vpnvserver_intranetip6_binding resource")

	// Read the updated state back
	r.readVpnvserverIntranetip6BindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnvserverIntranetip6BindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VpnvserverIntranetip6BindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting vpnvserver_intranetip6_binding resource")

	// For vpnvserver_intranetip6_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted vpnvserver_intranetip6_binding resource from state")
}

// Helper function to read vpnvserver_intranetip6_binding data from API
func (r *VpnvserverIntranetip6BindingResource) readVpnvserverIntranetip6BindingFromApi(ctx context.Context, data *VpnvserverIntranetip6BindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Vpnvserver_intranetip6_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read vpnvserver_intranetip6_binding, got error: %s", err))
		return
	}

	vpnvserver_intranetip6_bindingSetAttrFromGet(ctx, data, getResponseData)

}
