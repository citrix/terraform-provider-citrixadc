package vpnvserver_vpnintranetapplication_binding

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
var _ resource.Resource = &VpnvserverVpnintranetapplicationBindingResource{}
var _ resource.ResourceWithConfigure = (*VpnvserverVpnintranetapplicationBindingResource)(nil)
var _ resource.ResourceWithImportState = (*VpnvserverVpnintranetapplicationBindingResource)(nil)

func NewVpnvserverVpnintranetapplicationBindingResource() resource.Resource {
	return &VpnvserverVpnintranetapplicationBindingResource{}
}

// VpnvserverVpnintranetapplicationBindingResource defines the resource implementation.
type VpnvserverVpnintranetapplicationBindingResource struct {
	client *service.NitroClient
}

func (r *VpnvserverVpnintranetapplicationBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VpnvserverVpnintranetapplicationBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpnvserver_vpnintranetapplication_binding"
}

func (r *VpnvserverVpnintranetapplicationBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VpnvserverVpnintranetapplicationBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VpnvserverVpnintranetapplicationBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating vpnvserver_vpnintranetapplication_binding resource")

	// vpnvserver_vpnintranetapplication_binding := vpnvserver_vpnintranetapplication_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Vpnvserver_vpnintranetapplication_binding.Type(), &vpnvserver_vpnintranetapplication_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create vpnvserver_vpnintranetapplication_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("vpnvserver_vpnintranetapplication_binding-config")

	tflog.Trace(ctx, "Created vpnvserver_vpnintranetapplication_binding resource")

	// Read the updated state back
	r.readVpnvserverVpnintranetapplicationBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnvserverVpnintranetapplicationBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VpnvserverVpnintranetapplicationBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading vpnvserver_vpnintranetapplication_binding resource")

	r.readVpnvserverVpnintranetapplicationBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnvserverVpnintranetapplicationBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data VpnvserverVpnintranetapplicationBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating vpnvserver_vpnintranetapplication_binding resource")

	// Create API request body from the model
	// vpnvserver_vpnintranetapplication_binding := vpnvserver_vpnintranetapplication_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Vpnvserver_vpnintranetapplication_binding.Type(), &vpnvserver_vpnintranetapplication_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update vpnvserver_vpnintranetapplication_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated vpnvserver_vpnintranetapplication_binding resource")

	// Read the updated state back
	r.readVpnvserverVpnintranetapplicationBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnvserverVpnintranetapplicationBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VpnvserverVpnintranetapplicationBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting vpnvserver_vpnintranetapplication_binding resource")

	// For vpnvserver_vpnintranetapplication_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted vpnvserver_vpnintranetapplication_binding resource from state")
}

// Helper function to read vpnvserver_vpnintranetapplication_binding data from API
func (r *VpnvserverVpnintranetapplicationBindingResource) readVpnvserverVpnintranetapplicationBindingFromApi(ctx context.Context, data *VpnvserverVpnintranetapplicationBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Vpnvserver_vpnintranetapplication_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read vpnvserver_vpnintranetapplication_binding, got error: %s", err))
		return
	}

	vpnvserver_vpnintranetapplication_bindingSetAttrFromGet(ctx, data, getResponseData)

}
