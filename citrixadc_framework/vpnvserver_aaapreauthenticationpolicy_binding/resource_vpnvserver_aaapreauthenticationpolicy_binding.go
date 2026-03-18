package vpnvserver_aaapreauthenticationpolicy_binding

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
var _ resource.Resource = &VpnvserverAaapreauthenticationpolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*VpnvserverAaapreauthenticationpolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*VpnvserverAaapreauthenticationpolicyBindingResource)(nil)

func NewVpnvserverAaapreauthenticationpolicyBindingResource() resource.Resource {
	return &VpnvserverAaapreauthenticationpolicyBindingResource{}
}

// VpnvserverAaapreauthenticationpolicyBindingResource defines the resource implementation.
type VpnvserverAaapreauthenticationpolicyBindingResource struct {
	client *service.NitroClient
}

func (r *VpnvserverAaapreauthenticationpolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VpnvserverAaapreauthenticationpolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpnvserver_aaapreauthenticationpolicy_binding"
}

func (r *VpnvserverAaapreauthenticationpolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VpnvserverAaapreauthenticationpolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VpnvserverAaapreauthenticationpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating vpnvserver_aaapreauthenticationpolicy_binding resource")

	// vpnvserver_aaapreauthenticationpolicy_binding := vpnvserver_aaapreauthenticationpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Vpnvserver_aaapreauthenticationpolicy_binding.Type(), &vpnvserver_aaapreauthenticationpolicy_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create vpnvserver_aaapreauthenticationpolicy_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("vpnvserver_aaapreauthenticationpolicy_binding-config")

	tflog.Trace(ctx, "Created vpnvserver_aaapreauthenticationpolicy_binding resource")

	// Read the updated state back
	r.readVpnvserverAaapreauthenticationpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnvserverAaapreauthenticationpolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VpnvserverAaapreauthenticationpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading vpnvserver_aaapreauthenticationpolicy_binding resource")

	r.readVpnvserverAaapreauthenticationpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnvserverAaapreauthenticationpolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data VpnvserverAaapreauthenticationpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating vpnvserver_aaapreauthenticationpolicy_binding resource")

	// Create API request body from the model
	// vpnvserver_aaapreauthenticationpolicy_binding := vpnvserver_aaapreauthenticationpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Vpnvserver_aaapreauthenticationpolicy_binding.Type(), &vpnvserver_aaapreauthenticationpolicy_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update vpnvserver_aaapreauthenticationpolicy_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated vpnvserver_aaapreauthenticationpolicy_binding resource")

	// Read the updated state back
	r.readVpnvserverAaapreauthenticationpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnvserverAaapreauthenticationpolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VpnvserverAaapreauthenticationpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting vpnvserver_aaapreauthenticationpolicy_binding resource")

	// For vpnvserver_aaapreauthenticationpolicy_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted vpnvserver_aaapreauthenticationpolicy_binding resource from state")
}

// Helper function to read vpnvserver_aaapreauthenticationpolicy_binding data from API
func (r *VpnvserverAaapreauthenticationpolicyBindingResource) readVpnvserverAaapreauthenticationpolicyBindingFromApi(ctx context.Context, data *VpnvserverAaapreauthenticationpolicyBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Vpnvserver_aaapreauthenticationpolicy_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read vpnvserver_aaapreauthenticationpolicy_binding, got error: %s", err))
		return
	}

	vpnvserver_aaapreauthenticationpolicy_bindingSetAttrFromGet(ctx, data, getResponseData)

}
