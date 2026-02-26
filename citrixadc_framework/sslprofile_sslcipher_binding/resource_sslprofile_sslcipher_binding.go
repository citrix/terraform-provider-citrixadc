package sslprofile_sslcipher_binding

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
var _ resource.Resource = &SslprofileSslcipherBindingResource{}
var _ resource.ResourceWithConfigure = (*SslprofileSslcipherBindingResource)(nil)
var _ resource.ResourceWithImportState = (*SslprofileSslcipherBindingResource)(nil)

func NewSslprofileSslcipherBindingResource() resource.Resource {
	return &SslprofileSslcipherBindingResource{}
}

// SslprofileSslcipherBindingResource defines the resource implementation.
type SslprofileSslcipherBindingResource struct {
	client *service.NitroClient
}

func (r *SslprofileSslcipherBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SslprofileSslcipherBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslprofile_sslcipher_binding"
}

func (r *SslprofileSslcipherBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SslprofileSslcipherBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SslprofileSslcipherBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating sslprofile_sslcipher_binding resource")

	// sslprofile_sslcipher_binding := sslprofile_sslcipher_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Sslprofile_sslcipher_binding.Type(), &sslprofile_sslcipher_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create sslprofile_sslcipher_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("sslprofile_sslcipher_binding-config")

	tflog.Trace(ctx, "Created sslprofile_sslcipher_binding resource")

	// Read the updated state back
	r.readSslprofileSslcipherBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslprofileSslcipherBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SslprofileSslcipherBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading sslprofile_sslcipher_binding resource")

	r.readSslprofileSslcipherBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslprofileSslcipherBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data SslprofileSslcipherBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating sslprofile_sslcipher_binding resource")

	// Create API request body from the model
	// sslprofile_sslcipher_binding := sslprofile_sslcipher_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Sslprofile_sslcipher_binding.Type(), &sslprofile_sslcipher_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update sslprofile_sslcipher_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated sslprofile_sslcipher_binding resource")

	// Read the updated state back
	r.readSslprofileSslcipherBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslprofileSslcipherBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SslprofileSslcipherBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting sslprofile_sslcipher_binding resource")

	// For sslprofile_sslcipher_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted sslprofile_sslcipher_binding resource from state")
}

// Helper function to read sslprofile_sslcipher_binding data from API
func (r *SslprofileSslcipherBindingResource) readSslprofileSslcipherBindingFromApi(ctx context.Context, data *SslprofileSslcipherBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Sslprofile_sslcipher_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read sslprofile_sslcipher_binding, got error: %s", err))
		return
	}

	sslprofile_sslcipher_bindingSetAttrFromGet(ctx, data, getResponseData)

}
