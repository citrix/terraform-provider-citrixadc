package tunnelglobal_tunneltrafficpolicy_binding

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
var _ resource.Resource = &TunnelglobalTunneltrafficpolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*TunnelglobalTunneltrafficpolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*TunnelglobalTunneltrafficpolicyBindingResource)(nil)

func NewTunnelglobalTunneltrafficpolicyBindingResource() resource.Resource {
	return &TunnelglobalTunneltrafficpolicyBindingResource{}
}

// TunnelglobalTunneltrafficpolicyBindingResource defines the resource implementation.
type TunnelglobalTunneltrafficpolicyBindingResource struct {
	client *service.NitroClient
}

func (r *TunnelglobalTunneltrafficpolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *TunnelglobalTunneltrafficpolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tunnelglobal_tunneltrafficpolicy_binding"
}

func (r *TunnelglobalTunneltrafficpolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *TunnelglobalTunneltrafficpolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data TunnelglobalTunneltrafficpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating tunnelglobal_tunneltrafficpolicy_binding resource")

	// tunnelglobal_tunneltrafficpolicy_binding := tunnelglobal_tunneltrafficpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Tunnelglobal_tunneltrafficpolicy_binding.Type(), &tunnelglobal_tunneltrafficpolicy_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create tunnelglobal_tunneltrafficpolicy_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("tunnelglobal_tunneltrafficpolicy_binding-config")

	tflog.Trace(ctx, "Created tunnelglobal_tunneltrafficpolicy_binding resource")

	// Read the updated state back
	r.readTunnelglobalTunneltrafficpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TunnelglobalTunneltrafficpolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data TunnelglobalTunneltrafficpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading tunnelglobal_tunneltrafficpolicy_binding resource")

	r.readTunnelglobalTunneltrafficpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TunnelglobalTunneltrafficpolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data TunnelglobalTunneltrafficpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating tunnelglobal_tunneltrafficpolicy_binding resource")

	// Create API request body from the model
	// tunnelglobal_tunneltrafficpolicy_binding := tunnelglobal_tunneltrafficpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Tunnelglobal_tunneltrafficpolicy_binding.Type(), &tunnelglobal_tunneltrafficpolicy_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update tunnelglobal_tunneltrafficpolicy_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated tunnelglobal_tunneltrafficpolicy_binding resource")

	// Read the updated state back
	r.readTunnelglobalTunneltrafficpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TunnelglobalTunneltrafficpolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data TunnelglobalTunneltrafficpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting tunnelglobal_tunneltrafficpolicy_binding resource")

	// For tunnelglobal_tunneltrafficpolicy_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted tunnelglobal_tunneltrafficpolicy_binding resource from state")
}

// Helper function to read tunnelglobal_tunneltrafficpolicy_binding data from API
func (r *TunnelglobalTunneltrafficpolicyBindingResource) readTunnelglobalTunneltrafficpolicyBindingFromApi(ctx context.Context, data *TunnelglobalTunneltrafficpolicyBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Tunnelglobal_tunneltrafficpolicy_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read tunnelglobal_tunneltrafficpolicy_binding, got error: %s", err))
		return
	}

	tunnelglobal_tunneltrafficpolicy_bindingSetAttrFromGet(ctx, data, getResponseData)

}
