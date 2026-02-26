package sslservicegroup_sslcertkey_binding

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
var _ resource.Resource = &SslservicegroupSslcertkeyBindingResource{}
var _ resource.ResourceWithConfigure = (*SslservicegroupSslcertkeyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*SslservicegroupSslcertkeyBindingResource)(nil)

func NewSslservicegroupSslcertkeyBindingResource() resource.Resource {
	return &SslservicegroupSslcertkeyBindingResource{}
}

// SslservicegroupSslcertkeyBindingResource defines the resource implementation.
type SslservicegroupSslcertkeyBindingResource struct {
	client *service.NitroClient
}

func (r *SslservicegroupSslcertkeyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SslservicegroupSslcertkeyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslservicegroup_sslcertkey_binding"
}

func (r *SslservicegroupSslcertkeyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SslservicegroupSslcertkeyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SslservicegroupSslcertkeyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating sslservicegroup_sslcertkey_binding resource")

	// sslservicegroup_sslcertkey_binding := sslservicegroup_sslcertkey_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Sslservicegroup_sslcertkey_binding.Type(), &sslservicegroup_sslcertkey_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create sslservicegroup_sslcertkey_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("sslservicegroup_sslcertkey_binding-config")

	tflog.Trace(ctx, "Created sslservicegroup_sslcertkey_binding resource")

	// Read the updated state back
	r.readSslservicegroupSslcertkeyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslservicegroupSslcertkeyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SslservicegroupSslcertkeyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading sslservicegroup_sslcertkey_binding resource")

	r.readSslservicegroupSslcertkeyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslservicegroupSslcertkeyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data SslservicegroupSslcertkeyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating sslservicegroup_sslcertkey_binding resource")

	// Create API request body from the model
	// sslservicegroup_sslcertkey_binding := sslservicegroup_sslcertkey_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Sslservicegroup_sslcertkey_binding.Type(), &sslservicegroup_sslcertkey_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update sslservicegroup_sslcertkey_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated sslservicegroup_sslcertkey_binding resource")

	// Read the updated state back
	r.readSslservicegroupSslcertkeyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslservicegroupSslcertkeyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SslservicegroupSslcertkeyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting sslservicegroup_sslcertkey_binding resource")

	// For sslservicegroup_sslcertkey_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted sslservicegroup_sslcertkey_binding resource from state")
}

// Helper function to read sslservicegroup_sslcertkey_binding data from API
func (r *SslservicegroupSslcertkeyBindingResource) readSslservicegroupSslcertkeyBindingFromApi(ctx context.Context, data *SslservicegroupSslcertkeyBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Sslservicegroup_sslcertkey_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read sslservicegroup_sslcertkey_binding, got error: %s", err))
		return
	}

	sslservicegroup_sslcertkey_bindingSetAttrFromGet(ctx, data, getResponseData)

}
