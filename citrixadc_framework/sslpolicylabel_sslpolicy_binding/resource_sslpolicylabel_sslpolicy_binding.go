package sslpolicylabel_sslpolicy_binding

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
var _ resource.Resource = &SslpolicylabelSslpolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*SslpolicylabelSslpolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*SslpolicylabelSslpolicyBindingResource)(nil)

func NewSslpolicylabelSslpolicyBindingResource() resource.Resource {
	return &SslpolicylabelSslpolicyBindingResource{}
}

// SslpolicylabelSslpolicyBindingResource defines the resource implementation.
type SslpolicylabelSslpolicyBindingResource struct {
	client *service.NitroClient
}

func (r *SslpolicylabelSslpolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SslpolicylabelSslpolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslpolicylabel_sslpolicy_binding"
}

func (r *SslpolicylabelSslpolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SslpolicylabelSslpolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SslpolicylabelSslpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating sslpolicylabel_sslpolicy_binding resource")

	// sslpolicylabel_sslpolicy_binding := sslpolicylabel_sslpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Sslpolicylabel_sslpolicy_binding.Type(), &sslpolicylabel_sslpolicy_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create sslpolicylabel_sslpolicy_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("sslpolicylabel_sslpolicy_binding-config")

	tflog.Trace(ctx, "Created sslpolicylabel_sslpolicy_binding resource")

	// Read the updated state back
	r.readSslpolicylabelSslpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslpolicylabelSslpolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SslpolicylabelSslpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading sslpolicylabel_sslpolicy_binding resource")

	r.readSslpolicylabelSslpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslpolicylabelSslpolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data SslpolicylabelSslpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating sslpolicylabel_sslpolicy_binding resource")

	// Create API request body from the model
	// sslpolicylabel_sslpolicy_binding := sslpolicylabel_sslpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Sslpolicylabel_sslpolicy_binding.Type(), &sslpolicylabel_sslpolicy_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update sslpolicylabel_sslpolicy_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated sslpolicylabel_sslpolicy_binding resource")

	// Read the updated state back
	r.readSslpolicylabelSslpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslpolicylabelSslpolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SslpolicylabelSslpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting sslpolicylabel_sslpolicy_binding resource")

	// For sslpolicylabel_sslpolicy_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted sslpolicylabel_sslpolicy_binding resource from state")
}

// Helper function to read sslpolicylabel_sslpolicy_binding data from API
func (r *SslpolicylabelSslpolicyBindingResource) readSslpolicylabelSslpolicyBindingFromApi(ctx context.Context, data *SslpolicylabelSslpolicyBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Sslpolicylabel_sslpolicy_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read sslpolicylabel_sslpolicy_binding, got error: %s", err))
		return
	}

	sslpolicylabel_sslpolicy_bindingSetAttrFromGet(ctx, data, getResponseData)

}
