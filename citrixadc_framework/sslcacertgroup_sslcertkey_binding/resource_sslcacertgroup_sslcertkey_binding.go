package sslcacertgroup_sslcertkey_binding

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
var _ resource.Resource = &SslcacertgroupSslcertkeyBindingResource{}
var _ resource.ResourceWithConfigure = (*SslcacertgroupSslcertkeyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*SslcacertgroupSslcertkeyBindingResource)(nil)

func NewSslcacertgroupSslcertkeyBindingResource() resource.Resource {
	return &SslcacertgroupSslcertkeyBindingResource{}
}

// SslcacertgroupSslcertkeyBindingResource defines the resource implementation.
type SslcacertgroupSslcertkeyBindingResource struct {
	client *service.NitroClient
}

func (r *SslcacertgroupSslcertkeyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SslcacertgroupSslcertkeyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslcacertgroup_sslcertkey_binding"
}

func (r *SslcacertgroupSslcertkeyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SslcacertgroupSslcertkeyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SslcacertgroupSslcertkeyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating sslcacertgroup_sslcertkey_binding resource")

	// sslcacertgroup_sslcertkey_binding := sslcacertgroup_sslcertkey_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Sslcacertgroup_sslcertkey_binding.Type(), &sslcacertgroup_sslcertkey_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create sslcacertgroup_sslcertkey_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("sslcacertgroup_sslcertkey_binding-config")

	tflog.Trace(ctx, "Created sslcacertgroup_sslcertkey_binding resource")

	// Read the updated state back
	r.readSslcacertgroupSslcertkeyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslcacertgroupSslcertkeyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SslcacertgroupSslcertkeyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading sslcacertgroup_sslcertkey_binding resource")

	r.readSslcacertgroupSslcertkeyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslcacertgroupSslcertkeyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data SslcacertgroupSslcertkeyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating sslcacertgroup_sslcertkey_binding resource")

	// Create API request body from the model
	// sslcacertgroup_sslcertkey_binding := sslcacertgroup_sslcertkey_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Sslcacertgroup_sslcertkey_binding.Type(), &sslcacertgroup_sslcertkey_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update sslcacertgroup_sslcertkey_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated sslcacertgroup_sslcertkey_binding resource")

	// Read the updated state back
	r.readSslcacertgroupSslcertkeyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslcacertgroupSslcertkeyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SslcacertgroupSslcertkeyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting sslcacertgroup_sslcertkey_binding resource")

	// For sslcacertgroup_sslcertkey_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted sslcacertgroup_sslcertkey_binding resource from state")
}

// Helper function to read sslcacertgroup_sslcertkey_binding data from API
func (r *SslcacertgroupSslcertkeyBindingResource) readSslcacertgroupSslcertkeyBindingFromApi(ctx context.Context, data *SslcacertgroupSslcertkeyBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Sslcacertgroup_sslcertkey_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read sslcacertgroup_sslcertkey_binding, got error: %s", err))
		return
	}

	sslcacertgroup_sslcertkey_bindingSetAttrFromGet(ctx, data, getResponseData)

}
