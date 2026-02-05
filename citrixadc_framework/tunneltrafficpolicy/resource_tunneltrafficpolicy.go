package tunneltrafficpolicy

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
var _ resource.Resource = &TunneltrafficpolicyResource{}
var _ resource.ResourceWithConfigure = (*TunneltrafficpolicyResource)(nil)
var _ resource.ResourceWithImportState = (*TunneltrafficpolicyResource)(nil)

func NewTunneltrafficpolicyResource() resource.Resource {
	return &TunneltrafficpolicyResource{}
}

// TunneltrafficpolicyResource defines the resource implementation.
type TunneltrafficpolicyResource struct {
	client *service.NitroClient
}

func (r *TunneltrafficpolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *TunneltrafficpolicyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tunneltrafficpolicy"
}

func (r *TunneltrafficpolicyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *TunneltrafficpolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data TunneltrafficpolicyResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating tunneltrafficpolicy resource")

	// tunneltrafficpolicy := tunneltrafficpolicyGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Tunneltrafficpolicy.Type(), &tunneltrafficpolicy)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create tunneltrafficpolicy, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("tunneltrafficpolicy-config")

	tflog.Trace(ctx, "Created tunneltrafficpolicy resource")

	// Read the updated state back
	r.readTunneltrafficpolicyFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TunneltrafficpolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data TunneltrafficpolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading tunneltrafficpolicy resource")

	r.readTunneltrafficpolicyFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TunneltrafficpolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data TunneltrafficpolicyResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating tunneltrafficpolicy resource")

	// Create API request body from the model
	// tunneltrafficpolicy := tunneltrafficpolicyGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Tunneltrafficpolicy.Type(), &tunneltrafficpolicy)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update tunneltrafficpolicy, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated tunneltrafficpolicy resource")

	// Read the updated state back
	r.readTunneltrafficpolicyFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TunneltrafficpolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data TunneltrafficpolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting tunneltrafficpolicy resource")

	// For tunneltrafficpolicy, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted tunneltrafficpolicy resource from state")
}

// Helper function to read tunneltrafficpolicy data from API
func (r *TunneltrafficpolicyResource) readTunneltrafficpolicyFromApi(ctx context.Context, data *TunneltrafficpolicyResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Tunneltrafficpolicy.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read tunneltrafficpolicy, got error: %s", err))
		return
	}

	tunneltrafficpolicySetAttrFromGet(ctx, data, getResponseData)

}
