package sslcertkey_sslocspresponder_binding

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
var _ resource.Resource = &SslcertkeySslocspresponderBindingResource{}
var _ resource.ResourceWithConfigure = (*SslcertkeySslocspresponderBindingResource)(nil)
var _ resource.ResourceWithImportState = (*SslcertkeySslocspresponderBindingResource)(nil)

func NewSslcertkeySslocspresponderBindingResource() resource.Resource {
	return &SslcertkeySslocspresponderBindingResource{}
}

// SslcertkeySslocspresponderBindingResource defines the resource implementation.
type SslcertkeySslocspresponderBindingResource struct {
	client *service.NitroClient
}

func (r *SslcertkeySslocspresponderBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SslcertkeySslocspresponderBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslcertkey_sslocspresponder_binding"
}

func (r *SslcertkeySslocspresponderBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SslcertkeySslocspresponderBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SslcertkeySslocspresponderBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating sslcertkey_sslocspresponder_binding resource")

	// sslcertkey_sslocspresponder_binding := sslcertkey_sslocspresponder_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Sslcertkey_sslocspresponder_binding.Type(), &sslcertkey_sslocspresponder_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create sslcertkey_sslocspresponder_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("sslcertkey_sslocspresponder_binding-config")

	tflog.Trace(ctx, "Created sslcertkey_sslocspresponder_binding resource")

	// Read the updated state back
	r.readSslcertkeySslocspresponderBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslcertkeySslocspresponderBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SslcertkeySslocspresponderBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading sslcertkey_sslocspresponder_binding resource")

	r.readSslcertkeySslocspresponderBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslcertkeySslocspresponderBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data SslcertkeySslocspresponderBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating sslcertkey_sslocspresponder_binding resource")

	// Create API request body from the model
	// sslcertkey_sslocspresponder_binding := sslcertkey_sslocspresponder_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Sslcertkey_sslocspresponder_binding.Type(), &sslcertkey_sslocspresponder_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update sslcertkey_sslocspresponder_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated sslcertkey_sslocspresponder_binding resource")

	// Read the updated state back
	r.readSslcertkeySslocspresponderBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslcertkeySslocspresponderBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SslcertkeySslocspresponderBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting sslcertkey_sslocspresponder_binding resource")

	// For sslcertkey_sslocspresponder_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted sslcertkey_sslocspresponder_binding resource from state")
}

// Helper function to read sslcertkey_sslocspresponder_binding data from API
func (r *SslcertkeySslocspresponderBindingResource) readSslcertkeySslocspresponderBindingFromApi(ctx context.Context, data *SslcertkeySslocspresponderBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Sslcertkey_sslocspresponder_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read sslcertkey_sslocspresponder_binding, got error: %s", err))
		return
	}

	sslcertkey_sslocspresponder_bindingSetAttrFromGet(ctx, data, getResponseData)

}
