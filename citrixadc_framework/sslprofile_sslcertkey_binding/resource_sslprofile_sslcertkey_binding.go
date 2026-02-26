package sslprofile_sslcertkey_binding

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
var _ resource.Resource = &SslprofileSslcertkeyBindingResource{}
var _ resource.ResourceWithConfigure = (*SslprofileSslcertkeyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*SslprofileSslcertkeyBindingResource)(nil)

func NewSslprofileSslcertkeyBindingResource() resource.Resource {
	return &SslprofileSslcertkeyBindingResource{}
}

// SslprofileSslcertkeyBindingResource defines the resource implementation.
type SslprofileSslcertkeyBindingResource struct {
	client *service.NitroClient
}

func (r *SslprofileSslcertkeyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SslprofileSslcertkeyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslprofile_sslcertkey_binding"
}

func (r *SslprofileSslcertkeyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SslprofileSslcertkeyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SslprofileSslcertkeyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating sslprofile_sslcertkey_binding resource")

	// sslprofile_sslcertkey_binding := sslprofile_sslcertkey_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Sslprofile_sslcertkey_binding.Type(), &sslprofile_sslcertkey_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create sslprofile_sslcertkey_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("sslprofile_sslcertkey_binding-config")

	tflog.Trace(ctx, "Created sslprofile_sslcertkey_binding resource")

	// Read the updated state back
	r.readSslprofileSslcertkeyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslprofileSslcertkeyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SslprofileSslcertkeyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading sslprofile_sslcertkey_binding resource")

	r.readSslprofileSslcertkeyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslprofileSslcertkeyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data SslprofileSslcertkeyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating sslprofile_sslcertkey_binding resource")

	// Create API request body from the model
	// sslprofile_sslcertkey_binding := sslprofile_sslcertkey_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Sslprofile_sslcertkey_binding.Type(), &sslprofile_sslcertkey_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update sslprofile_sslcertkey_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated sslprofile_sslcertkey_binding resource")

	// Read the updated state back
	r.readSslprofileSslcertkeyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslprofileSslcertkeyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SslprofileSslcertkeyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting sslprofile_sslcertkey_binding resource")

	// For sslprofile_sslcertkey_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted sslprofile_sslcertkey_binding resource from state")
}

// Helper function to read sslprofile_sslcertkey_binding data from API
func (r *SslprofileSslcertkeyBindingResource) readSslprofileSslcertkeyBindingFromApi(ctx context.Context, data *SslprofileSslcertkeyBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Sslprofile_sslcertkey_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read sslprofile_sslcertkey_binding, got error: %s", err))
		return
	}

	sslprofile_sslcertkey_bindingSetAttrFromGet(ctx, data, getResponseData)

}
