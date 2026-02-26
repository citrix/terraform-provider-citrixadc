package sslservice_sslcertkey_binding

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
var _ resource.Resource = &SslserviceSslcertkeyBindingResource{}
var _ resource.ResourceWithConfigure = (*SslserviceSslcertkeyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*SslserviceSslcertkeyBindingResource)(nil)

func NewSslserviceSslcertkeyBindingResource() resource.Resource {
	return &SslserviceSslcertkeyBindingResource{}
}

// SslserviceSslcertkeyBindingResource defines the resource implementation.
type SslserviceSslcertkeyBindingResource struct {
	client *service.NitroClient
}

func (r *SslserviceSslcertkeyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SslserviceSslcertkeyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslservice_sslcertkey_binding"
}

func (r *SslserviceSslcertkeyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SslserviceSslcertkeyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SslserviceSslcertkeyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating sslservice_sslcertkey_binding resource")

	// sslservice_sslcertkey_binding := sslservice_sslcertkey_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Sslservice_sslcertkey_binding.Type(), &sslservice_sslcertkey_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create sslservice_sslcertkey_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("sslservice_sslcertkey_binding-config")

	tflog.Trace(ctx, "Created sslservice_sslcertkey_binding resource")

	// Read the updated state back
	r.readSslserviceSslcertkeyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslserviceSslcertkeyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SslserviceSslcertkeyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading sslservice_sslcertkey_binding resource")

	r.readSslserviceSslcertkeyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslserviceSslcertkeyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data SslserviceSslcertkeyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating sslservice_sslcertkey_binding resource")

	// Create API request body from the model
	// sslservice_sslcertkey_binding := sslservice_sslcertkey_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Sslservice_sslcertkey_binding.Type(), &sslservice_sslcertkey_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update sslservice_sslcertkey_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated sslservice_sslcertkey_binding resource")

	// Read the updated state back
	r.readSslserviceSslcertkeyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslserviceSslcertkeyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SslserviceSslcertkeyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting sslservice_sslcertkey_binding resource")

	// For sslservice_sslcertkey_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted sslservice_sslcertkey_binding resource from state")
}

// Helper function to read sslservice_sslcertkey_binding data from API
func (r *SslserviceSslcertkeyBindingResource) readSslserviceSslcertkeyBindingFromApi(ctx context.Context, data *SslserviceSslcertkeyBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Sslservice_sslcertkey_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read sslservice_sslcertkey_binding, got error: %s", err))
		return
	}

	sslservice_sslcertkey_bindingSetAttrFromGet(ctx, data, getResponseData)

}
