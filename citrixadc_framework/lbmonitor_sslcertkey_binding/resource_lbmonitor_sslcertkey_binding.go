package lbmonitor_sslcertkey_binding

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
var _ resource.Resource = &LbmonitorSslcertkeyBindingResource{}
var _ resource.ResourceWithConfigure = (*LbmonitorSslcertkeyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*LbmonitorSslcertkeyBindingResource)(nil)

func NewLbmonitorSslcertkeyBindingResource() resource.Resource {
	return &LbmonitorSslcertkeyBindingResource{}
}

// LbmonitorSslcertkeyBindingResource defines the resource implementation.
type LbmonitorSslcertkeyBindingResource struct {
	client *service.NitroClient
}

func (r *LbmonitorSslcertkeyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *LbmonitorSslcertkeyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lbmonitor_sslcertkey_binding"
}

func (r *LbmonitorSslcertkeyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *LbmonitorSslcertkeyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data LbmonitorSslcertkeyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating lbmonitor_sslcertkey_binding resource")

	// lbmonitor_sslcertkey_binding := lbmonitor_sslcertkey_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Lbmonitor_sslcertkey_binding.Type(), &lbmonitor_sslcertkey_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create lbmonitor_sslcertkey_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("lbmonitor_sslcertkey_binding-config")

	tflog.Trace(ctx, "Created lbmonitor_sslcertkey_binding resource")

	// Read the updated state back
	r.readLbmonitorSslcertkeyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbmonitorSslcertkeyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data LbmonitorSslcertkeyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading lbmonitor_sslcertkey_binding resource")

	r.readLbmonitorSslcertkeyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbmonitorSslcertkeyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data LbmonitorSslcertkeyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating lbmonitor_sslcertkey_binding resource")

	// Create API request body from the model
	// lbmonitor_sslcertkey_binding := lbmonitor_sslcertkey_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Lbmonitor_sslcertkey_binding.Type(), &lbmonitor_sslcertkey_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update lbmonitor_sslcertkey_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated lbmonitor_sslcertkey_binding resource")

	// Read the updated state back
	r.readLbmonitorSslcertkeyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbmonitorSslcertkeyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data LbmonitorSslcertkeyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting lbmonitor_sslcertkey_binding resource")

	// For lbmonitor_sslcertkey_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted lbmonitor_sslcertkey_binding resource from state")
}

// Helper function to read lbmonitor_sslcertkey_binding data from API
func (r *LbmonitorSslcertkeyBindingResource) readLbmonitorSslcertkeyBindingFromApi(ctx context.Context, data *LbmonitorSslcertkeyBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Lbmonitor_sslcertkey_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read lbmonitor_sslcertkey_binding, got error: %s", err))
		return
	}

	lbmonitor_sslcertkey_bindingSetAttrFromGet(ctx, data, getResponseData)

}
