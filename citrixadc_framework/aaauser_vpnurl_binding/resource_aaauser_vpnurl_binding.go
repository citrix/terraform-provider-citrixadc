package aaauser_vpnurl_binding

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
var _ resource.Resource = &AaauserVpnurlBindingResource{}
var _ resource.ResourceWithConfigure = (*AaauserVpnurlBindingResource)(nil)
var _ resource.ResourceWithImportState = (*AaauserVpnurlBindingResource)(nil)

func NewAaauserVpnurlBindingResource() resource.Resource {
	return &AaauserVpnurlBindingResource{}
}

// AaauserVpnurlBindingResource defines the resource implementation.
type AaauserVpnurlBindingResource struct {
	client *service.NitroClient
}

func (r *AaauserVpnurlBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AaauserVpnurlBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_aaauser_vpnurl_binding"
}

func (r *AaauserVpnurlBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AaauserVpnurlBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AaauserVpnurlBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating aaauser_vpnurl_binding resource")

	// aaauser_vpnurl_binding := aaauser_vpnurl_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Aaauser_vpnurl_binding.Type(), &aaauser_vpnurl_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create aaauser_vpnurl_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("aaauser_vpnurl_binding-config")

	tflog.Trace(ctx, "Created aaauser_vpnurl_binding resource")

	// Read the updated state back
	r.readAaauserVpnurlBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AaauserVpnurlBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AaauserVpnurlBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading aaauser_vpnurl_binding resource")

	r.readAaauserVpnurlBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AaauserVpnurlBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data AaauserVpnurlBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating aaauser_vpnurl_binding resource")

	// Create API request body from the model
	// aaauser_vpnurl_binding := aaauser_vpnurl_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Aaauser_vpnurl_binding.Type(), &aaauser_vpnurl_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update aaauser_vpnurl_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated aaauser_vpnurl_binding resource")

	// Read the updated state back
	r.readAaauserVpnurlBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AaauserVpnurlBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AaauserVpnurlBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting aaauser_vpnurl_binding resource")

	// For aaauser_vpnurl_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted aaauser_vpnurl_binding resource from state")
}

// Helper function to read aaauser_vpnurl_binding data from API
func (r *AaauserVpnurlBindingResource) readAaauserVpnurlBindingFromApi(ctx context.Context, data *AaauserVpnurlBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Aaauser_vpnurl_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read aaauser_vpnurl_binding, got error: %s", err))
		return
	}

	aaauser_vpnurl_bindingSetAttrFromGet(ctx, data, getResponseData)

}
