package vpnvserver_authenticationsamlidppolicy_binding

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
var _ resource.Resource = &VpnvserverAuthenticationsamlidppolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*VpnvserverAuthenticationsamlidppolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*VpnvserverAuthenticationsamlidppolicyBindingResource)(nil)

func NewVpnvserverAuthenticationsamlidppolicyBindingResource() resource.Resource {
	return &VpnvserverAuthenticationsamlidppolicyBindingResource{}
}

// VpnvserverAuthenticationsamlidppolicyBindingResource defines the resource implementation.
type VpnvserverAuthenticationsamlidppolicyBindingResource struct {
	client *service.NitroClient
}

func (r *VpnvserverAuthenticationsamlidppolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VpnvserverAuthenticationsamlidppolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpnvserver_authenticationsamlidppolicy_binding"
}

func (r *VpnvserverAuthenticationsamlidppolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VpnvserverAuthenticationsamlidppolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VpnvserverAuthenticationsamlidppolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating vpnvserver_authenticationsamlidppolicy_binding resource")

	// vpnvserver_authenticationsamlidppolicy_binding := vpnvserver_authenticationsamlidppolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Vpnvserver_authenticationsamlidppolicy_binding.Type(), &vpnvserver_authenticationsamlidppolicy_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create vpnvserver_authenticationsamlidppolicy_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("vpnvserver_authenticationsamlidppolicy_binding-config")

	tflog.Trace(ctx, "Created vpnvserver_authenticationsamlidppolicy_binding resource")

	// Read the updated state back
	r.readVpnvserverAuthenticationsamlidppolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnvserverAuthenticationsamlidppolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VpnvserverAuthenticationsamlidppolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading vpnvserver_authenticationsamlidppolicy_binding resource")

	r.readVpnvserverAuthenticationsamlidppolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnvserverAuthenticationsamlidppolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data VpnvserverAuthenticationsamlidppolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating vpnvserver_authenticationsamlidppolicy_binding resource")

	// Create API request body from the model
	// vpnvserver_authenticationsamlidppolicy_binding := vpnvserver_authenticationsamlidppolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Vpnvserver_authenticationsamlidppolicy_binding.Type(), &vpnvserver_authenticationsamlidppolicy_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update vpnvserver_authenticationsamlidppolicy_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated vpnvserver_authenticationsamlidppolicy_binding resource")

	// Read the updated state back
	r.readVpnvserverAuthenticationsamlidppolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnvserverAuthenticationsamlidppolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VpnvserverAuthenticationsamlidppolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting vpnvserver_authenticationsamlidppolicy_binding resource")

	// For vpnvserver_authenticationsamlidppolicy_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted vpnvserver_authenticationsamlidppolicy_binding resource from state")
}

// Helper function to read vpnvserver_authenticationsamlidppolicy_binding data from API
func (r *VpnvserverAuthenticationsamlidppolicyBindingResource) readVpnvserverAuthenticationsamlidppolicyBindingFromApi(ctx context.Context, data *VpnvserverAuthenticationsamlidppolicyBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Vpnvserver_authenticationsamlidppolicy_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read vpnvserver_authenticationsamlidppolicy_binding, got error: %s", err))
		return
	}

	vpnvserver_authenticationsamlidppolicy_bindingSetAttrFromGet(ctx, data, getResponseData)

}
