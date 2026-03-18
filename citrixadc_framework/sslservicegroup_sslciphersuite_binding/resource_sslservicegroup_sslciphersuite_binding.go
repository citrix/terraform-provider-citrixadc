package sslservicegroup_sslciphersuite_binding

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
var _ resource.Resource = &SslservicegroupSslciphersuiteBindingResource{}
var _ resource.ResourceWithConfigure = (*SslservicegroupSslciphersuiteBindingResource)(nil)
var _ resource.ResourceWithImportState = (*SslservicegroupSslciphersuiteBindingResource)(nil)

func NewSslservicegroupSslciphersuiteBindingResource() resource.Resource {
	return &SslservicegroupSslciphersuiteBindingResource{}
}

// SslservicegroupSslciphersuiteBindingResource defines the resource implementation.
type SslservicegroupSslciphersuiteBindingResource struct {
	client *service.NitroClient
}

func (r *SslservicegroupSslciphersuiteBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SslservicegroupSslciphersuiteBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslservicegroup_sslciphersuite_binding"
}

func (r *SslservicegroupSslciphersuiteBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SslservicegroupSslciphersuiteBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SslservicegroupSslciphersuiteBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating sslservicegroup_sslciphersuite_binding resource")

	// sslservicegroup_sslciphersuite_binding := sslservicegroup_sslciphersuite_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Sslservicegroup_sslciphersuite_binding.Type(), &sslservicegroup_sslciphersuite_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create sslservicegroup_sslciphersuite_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("sslservicegroup_sslciphersuite_binding-config")

	tflog.Trace(ctx, "Created sslservicegroup_sslciphersuite_binding resource")

	// Read the updated state back
	r.readSslservicegroupSslciphersuiteBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslservicegroupSslciphersuiteBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SslservicegroupSslciphersuiteBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading sslservicegroup_sslciphersuite_binding resource")

	r.readSslservicegroupSslciphersuiteBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslservicegroupSslciphersuiteBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data SslservicegroupSslciphersuiteBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating sslservicegroup_sslciphersuite_binding resource")

	// Create API request body from the model
	// sslservicegroup_sslciphersuite_binding := sslservicegroup_sslciphersuite_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Sslservicegroup_sslciphersuite_binding.Type(), &sslservicegroup_sslciphersuite_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update sslservicegroup_sslciphersuite_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated sslservicegroup_sslciphersuite_binding resource")

	// Read the updated state back
	r.readSslservicegroupSslciphersuiteBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslservicegroupSslciphersuiteBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SslservicegroupSslciphersuiteBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting sslservicegroup_sslciphersuite_binding resource")

	// For sslservicegroup_sslciphersuite_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted sslservicegroup_sslciphersuite_binding resource from state")
}

// Helper function to read sslservicegroup_sslciphersuite_binding data from API
func (r *SslservicegroupSslciphersuiteBindingResource) readSslservicegroupSslciphersuiteBindingFromApi(ctx context.Context, data *SslservicegroupSslciphersuiteBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Sslservicegroup_sslciphersuite_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read sslservicegroup_sslciphersuite_binding, got error: %s", err))
		return
	}

	sslservicegroup_sslciphersuite_bindingSetAttrFromGet(ctx, data, getResponseData)

}
