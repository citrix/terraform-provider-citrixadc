package vpnvserver_authenticationloginschemapolicy_binding

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
var _ resource.Resource = &VpnvserverAuthenticationloginschemapolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*VpnvserverAuthenticationloginschemapolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*VpnvserverAuthenticationloginschemapolicyBindingResource)(nil)

func NewVpnvserverAuthenticationloginschemapolicyBindingResource() resource.Resource {
	return &VpnvserverAuthenticationloginschemapolicyBindingResource{}
}

// VpnvserverAuthenticationloginschemapolicyBindingResource defines the resource implementation.
type VpnvserverAuthenticationloginschemapolicyBindingResource struct {
	client *service.NitroClient
}

func (r *VpnvserverAuthenticationloginschemapolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VpnvserverAuthenticationloginschemapolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpnvserver_authenticationloginschemapolicy_binding"
}

func (r *VpnvserverAuthenticationloginschemapolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VpnvserverAuthenticationloginschemapolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VpnvserverAuthenticationloginschemapolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating vpnvserver_authenticationloginschemapolicy_binding resource")

	// vpnvserver_authenticationloginschemapolicy_binding := vpnvserver_authenticationloginschemapolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Vpnvserver_authenticationloginschemapolicy_binding.Type(), &vpnvserver_authenticationloginschemapolicy_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create vpnvserver_authenticationloginschemapolicy_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("vpnvserver_authenticationloginschemapolicy_binding-config")

	tflog.Trace(ctx, "Created vpnvserver_authenticationloginschemapolicy_binding resource")

	// Read the updated state back
	r.readVpnvserverAuthenticationloginschemapolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnvserverAuthenticationloginschemapolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VpnvserverAuthenticationloginschemapolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading vpnvserver_authenticationloginschemapolicy_binding resource")

	r.readVpnvserverAuthenticationloginschemapolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnvserverAuthenticationloginschemapolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data VpnvserverAuthenticationloginschemapolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating vpnvserver_authenticationloginschemapolicy_binding resource")

	// Create API request body from the model
	// vpnvserver_authenticationloginschemapolicy_binding := vpnvserver_authenticationloginschemapolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Vpnvserver_authenticationloginschemapolicy_binding.Type(), &vpnvserver_authenticationloginschemapolicy_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update vpnvserver_authenticationloginschemapolicy_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated vpnvserver_authenticationloginschemapolicy_binding resource")

	// Read the updated state back
	r.readVpnvserverAuthenticationloginschemapolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnvserverAuthenticationloginschemapolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VpnvserverAuthenticationloginschemapolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting vpnvserver_authenticationloginschemapolicy_binding resource")

	// For vpnvserver_authenticationloginschemapolicy_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted vpnvserver_authenticationloginschemapolicy_binding resource from state")
}

// Helper function to read vpnvserver_authenticationloginschemapolicy_binding data from API
func (r *VpnvserverAuthenticationloginschemapolicyBindingResource) readVpnvserverAuthenticationloginschemapolicyBindingFromApi(ctx context.Context, data *VpnvserverAuthenticationloginschemapolicyBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Vpnvserver_authenticationloginschemapolicy_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read vpnvserver_authenticationloginschemapolicy_binding, got error: %s", err))
		return
	}

	vpnvserver_authenticationloginschemapolicy_bindingSetAttrFromGet(ctx, data, getResponseData)

}
