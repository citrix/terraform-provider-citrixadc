package sslcipher_sslciphersuite_binding

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
var _ resource.Resource = &SslcipherSslciphersuiteBindingResource{}
var _ resource.ResourceWithConfigure = (*SslcipherSslciphersuiteBindingResource)(nil)
var _ resource.ResourceWithImportState = (*SslcipherSslciphersuiteBindingResource)(nil)

func NewSslcipherSslciphersuiteBindingResource() resource.Resource {
	return &SslcipherSslciphersuiteBindingResource{}
}

// SslcipherSslciphersuiteBindingResource defines the resource implementation.
type SslcipherSslciphersuiteBindingResource struct {
	client *service.NitroClient
}

func (r *SslcipherSslciphersuiteBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SslcipherSslciphersuiteBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslcipher_sslciphersuite_binding"
}

func (r *SslcipherSslciphersuiteBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SslcipherSslciphersuiteBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SslcipherSslciphersuiteBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating sslcipher_sslciphersuite_binding resource")

	// sslcipher_sslciphersuite_binding := sslcipher_sslciphersuite_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Sslcipher_sslciphersuite_binding.Type(), &sslcipher_sslciphersuite_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create sslcipher_sslciphersuite_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("sslcipher_sslciphersuite_binding-config")

	tflog.Trace(ctx, "Created sslcipher_sslciphersuite_binding resource")

	// Read the updated state back
	r.readSslcipherSslciphersuiteBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslcipherSslciphersuiteBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SslcipherSslciphersuiteBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading sslcipher_sslciphersuite_binding resource")

	r.readSslcipherSslciphersuiteBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslcipherSslciphersuiteBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data SslcipherSslciphersuiteBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating sslcipher_sslciphersuite_binding resource")

	// Create API request body from the model
	// sslcipher_sslciphersuite_binding := sslcipher_sslciphersuite_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Sslcipher_sslciphersuite_binding.Type(), &sslcipher_sslciphersuite_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update sslcipher_sslciphersuite_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated sslcipher_sslciphersuite_binding resource")

	// Read the updated state back
	r.readSslcipherSslciphersuiteBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslcipherSslciphersuiteBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SslcipherSslciphersuiteBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting sslcipher_sslciphersuite_binding resource")

	// For sslcipher_sslciphersuite_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted sslcipher_sslciphersuite_binding resource from state")
}

// Helper function to read sslcipher_sslciphersuite_binding data from API
func (r *SslcipherSslciphersuiteBindingResource) readSslcipherSslciphersuiteBindingFromApi(ctx context.Context, data *SslcipherSslciphersuiteBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Sslcipher_sslciphersuite_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read sslcipher_sslciphersuite_binding, got error: %s", err))
		return
	}

	sslcipher_sslciphersuite_bindingSetAttrFromGet(ctx, data, getResponseData)

}
