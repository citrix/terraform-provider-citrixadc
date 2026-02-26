package vpnvserver_authenticationnegotiatepolicy_binding

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
var _ resource.Resource = &VpnvserverAuthenticationnegotiatepolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*VpnvserverAuthenticationnegotiatepolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*VpnvserverAuthenticationnegotiatepolicyBindingResource)(nil)

func NewVpnvserverAuthenticationnegotiatepolicyBindingResource() resource.Resource {
	return &VpnvserverAuthenticationnegotiatepolicyBindingResource{}
}

// VpnvserverAuthenticationnegotiatepolicyBindingResource defines the resource implementation.
type VpnvserverAuthenticationnegotiatepolicyBindingResource struct {
	client *service.NitroClient
}

func (r *VpnvserverAuthenticationnegotiatepolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VpnvserverAuthenticationnegotiatepolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpnvserver_authenticationnegotiatepolicy_binding"
}

func (r *VpnvserverAuthenticationnegotiatepolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VpnvserverAuthenticationnegotiatepolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VpnvserverAuthenticationnegotiatepolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating vpnvserver_authenticationnegotiatepolicy_binding resource")

	// vpnvserver_authenticationnegotiatepolicy_binding := vpnvserver_authenticationnegotiatepolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Vpnvserver_authenticationnegotiatepolicy_binding.Type(), &vpnvserver_authenticationnegotiatepolicy_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create vpnvserver_authenticationnegotiatepolicy_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("vpnvserver_authenticationnegotiatepolicy_binding-config")

	tflog.Trace(ctx, "Created vpnvserver_authenticationnegotiatepolicy_binding resource")

	// Read the updated state back
	r.readVpnvserverAuthenticationnegotiatepolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnvserverAuthenticationnegotiatepolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VpnvserverAuthenticationnegotiatepolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading vpnvserver_authenticationnegotiatepolicy_binding resource")

	r.readVpnvserverAuthenticationnegotiatepolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnvserverAuthenticationnegotiatepolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data VpnvserverAuthenticationnegotiatepolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating vpnvserver_authenticationnegotiatepolicy_binding resource")

	// Create API request body from the model
	// vpnvserver_authenticationnegotiatepolicy_binding := vpnvserver_authenticationnegotiatepolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Vpnvserver_authenticationnegotiatepolicy_binding.Type(), &vpnvserver_authenticationnegotiatepolicy_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update vpnvserver_authenticationnegotiatepolicy_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated vpnvserver_authenticationnegotiatepolicy_binding resource")

	// Read the updated state back
	r.readVpnvserverAuthenticationnegotiatepolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnvserverAuthenticationnegotiatepolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VpnvserverAuthenticationnegotiatepolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting vpnvserver_authenticationnegotiatepolicy_binding resource")

	// For vpnvserver_authenticationnegotiatepolicy_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted vpnvserver_authenticationnegotiatepolicy_binding resource from state")
}

// Helper function to read vpnvserver_authenticationnegotiatepolicy_binding data from API
func (r *VpnvserverAuthenticationnegotiatepolicyBindingResource) readVpnvserverAuthenticationnegotiatepolicyBindingFromApi(ctx context.Context, data *VpnvserverAuthenticationnegotiatepolicyBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Vpnvserver_authenticationnegotiatepolicy_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read vpnvserver_authenticationnegotiatepolicy_binding, got error: %s", err))
		return
	}

	vpnvserver_authenticationnegotiatepolicy_bindingSetAttrFromGet(ctx, data, getResponseData)

}
