package sslservice_sslciphersuite_binding

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
var _ resource.Resource = &SslserviceSslciphersuiteBindingResource{}
var _ resource.ResourceWithConfigure = (*SslserviceSslciphersuiteBindingResource)(nil)
var _ resource.ResourceWithImportState = (*SslserviceSslciphersuiteBindingResource)(nil)

func NewSslserviceSslciphersuiteBindingResource() resource.Resource {
	return &SslserviceSslciphersuiteBindingResource{}
}

// SslserviceSslciphersuiteBindingResource defines the resource implementation.
type SslserviceSslciphersuiteBindingResource struct {
	client *service.NitroClient
}

func (r *SslserviceSslciphersuiteBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SslserviceSslciphersuiteBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslservice_sslciphersuite_binding"
}

func (r *SslserviceSslciphersuiteBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SslserviceSslciphersuiteBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SslserviceSslciphersuiteBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating sslservice_sslciphersuite_binding resource")

	// sslservice_sslciphersuite_binding := sslservice_sslciphersuite_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Sslservice_sslciphersuite_binding.Type(), &sslservice_sslciphersuite_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create sslservice_sslciphersuite_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("sslservice_sslciphersuite_binding-config")

	tflog.Trace(ctx, "Created sslservice_sslciphersuite_binding resource")

	// Read the updated state back
	r.readSslserviceSslciphersuiteBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslserviceSslciphersuiteBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SslserviceSslciphersuiteBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading sslservice_sslciphersuite_binding resource")

	r.readSslserviceSslciphersuiteBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslserviceSslciphersuiteBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data SslserviceSslciphersuiteBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating sslservice_sslciphersuite_binding resource")

	// Create API request body from the model
	// sslservice_sslciphersuite_binding := sslservice_sslciphersuite_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Sslservice_sslciphersuite_binding.Type(), &sslservice_sslciphersuite_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update sslservice_sslciphersuite_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated sslservice_sslciphersuite_binding resource")

	// Read the updated state back
	r.readSslserviceSslciphersuiteBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslserviceSslciphersuiteBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SslserviceSslciphersuiteBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting sslservice_sslciphersuite_binding resource")

	// For sslservice_sslciphersuite_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted sslservice_sslciphersuite_binding resource from state")
}

// Helper function to read sslservice_sslciphersuite_binding data from API
func (r *SslserviceSslciphersuiteBindingResource) readSslserviceSslciphersuiteBindingFromApi(ctx context.Context, data *SslserviceSslciphersuiteBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Sslservice_sslciphersuite_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read sslservice_sslciphersuite_binding, got error: %s", err))
		return
	}

	sslservice_sslciphersuite_bindingSetAttrFromGet(ctx, data, getResponseData)

}
