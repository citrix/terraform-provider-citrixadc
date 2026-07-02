package sslcertificatechain

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
var _ resource.Resource = &SslcertificatechainResource{}
var _ resource.ResourceWithConfigure = (*SslcertificatechainResource)(nil)
var _ resource.ResourceWithImportState = (*SslcertificatechainResource)(nil)

func NewSslcertificatechainResource() resource.Resource {
	return &SslcertificatechainResource{}
}

// SslcertificatechainResource defines the resource implementation.
type SslcertificatechainResource struct {
	client *service.NitroClient
}

func (r *SslcertificatechainResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SslcertificatechainResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslcertificatechain"
}

func (r *SslcertificatechainResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SslcertificatechainResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SslcertificatechainResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating sslcertificatechain resource")
	sslcertificatechain := sslcertificatechainGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Named resource - use AddResource (POST)
	certkeyname_value := data.Certkeyname.ValueString()
	_, err := r.client.AddResource(service.Sslcertificatechain.Type(), certkeyname_value, &sslcertificatechain)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create sslcertificatechain, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created sslcertificatechain resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Certkeyname.ValueString()))

	// Read the updated state back
	r.readSslcertificatechainFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslcertificatechainResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SslcertificatechainResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading sslcertificatechain resource")

	r.readSslcertificatechainFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslcertificatechainResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state SslcertificatechainResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// NITRO has no update endpoint for sslcertificatechain; all attributes are RequiresReplace.
	tflog.Debug(ctx, "Update is a no-op for sslcertificatechain; all attributes are RequiresReplace")

	// Read the updated state back
	r.readSslcertificatechainFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslcertificatechainResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SslcertificatechainResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting sslcertificatechain resource")
	// NITRO exposes no delete endpoint for sslcertificatechain; remove from state only.
	tflog.Warn(ctx, "sslcertificatechain has no NITRO delete endpoint; removing from Terraform state only (resource remains on the appliance)")
	tflog.Trace(ctx, "Removed sslcertificatechain from Terraform state")
}

// Helper function to read sslcertificatechain data from API
func (r *SslcertificatechainResource) readSslcertificatechainFromApi(ctx context.Context, data *SslcertificatechainResourceModel, diags *diag.Diagnostics) {

	// Case 2: Find with single ID attribute - ID is the plain value
	certkeyname_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Sslcertificatechain.Type(), certkeyname_Name)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read sslcertificatechain, got error: %s", err))
		return
	}

	sslcertificatechainSetAttrFromGet(ctx, data, getResponseData)

}
