package gslbconfig

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &GslbconfigResource{}
var _ resource.ResourceWithConfigure = (*GslbconfigResource)(nil)
var _ resource.ResourceWithImportState = (*GslbconfigResource)(nil)

func NewGslbconfigResource() resource.Resource {
	return &GslbconfigResource{}
}

// GslbconfigResource defines the resource implementation.
type GslbconfigResource struct {
	client *service.NitroClient
}

func (r *GslbconfigResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *GslbconfigResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gslbconfig"
}

func (r *GslbconfigResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *GslbconfigResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data GslbconfigResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating gslbconfig resource")

	// gslbconfig exposes only the POST ?action=sync action on NITRO.
	// There is no add/get/update/delete endpoint. Use ActOnResource with
	// the "sync" verb.
	payload := gslbconfigGetThePayloadFromthePlan(ctx, &data)

	err := r.client.ActOnResource(service.Gslbconfig.Type(), payload, "sync")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to sync gslbconfig, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Synced gslbconfig resource")

	// Synthetic constant ID - there is no NITRO identity for this action resource.
	data.Id = types.StringValue("gslbconfig")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GslbconfigResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data GslbconfigResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read is a no-op for gslbconfig: NITRO exposes no GET endpoint for this
	// action-only resource, so drift detection is impossible. Preserve state.
	tflog.Debug(ctx, "Read is a no-op for gslbconfig; no GET endpoint on NITRO side")

	// Save (unchanged) data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GslbconfigResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state GslbconfigResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// Update is a no-op for gslbconfig; NITRO exposes no update endpoint for
	// this action-only resource (only ?action=sync). Changes re-run sync via
	// Create on the next apply.
	tflog.Debug(ctx, "Update is a no-op for gslbconfig; no update endpoint on NITRO side")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GslbconfigResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data GslbconfigResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Delete is a no-op for gslbconfig: NITRO exposes no DELETE endpoint for
	// this action-only resource. Just remove from Terraform state.
	tflog.Debug(ctx, "Delete is a no-op for gslbconfig; no DELETE endpoint on NITRO side")
	tflog.Trace(ctx, "Removed gslbconfig from Terraform state")
}
