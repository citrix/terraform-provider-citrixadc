package rnatsession

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
var _ resource.Resource = &RnatsessionResource{}
var _ resource.ResourceWithConfigure = (*RnatsessionResource)(nil)
var _ resource.ResourceWithImportState = (*RnatsessionResource)(nil)

func NewRnatsessionResource() resource.Resource {
	return &RnatsessionResource{}
}

// RnatsessionResource defines the resource implementation.
type RnatsessionResource struct {
	client *service.NitroClient
}

func (r *RnatsessionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *RnatsessionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_rnatsession"
}

func (r *RnatsessionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *RnatsessionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data RnatsessionResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Flushing rnatsession (action-only resource)")
	rnatsession := rnatsessionGetThePayloadFromthePlan(ctx, &data)

	// Action-only resource: only the flush action exists (NITRO POST ?action=flush).
	err := r.client.ActOnResource(service.Rnatsession.Type(), &rnatsession, "flush")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to flush rnatsession, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Flushed rnatsession")

	// Synthetic ID for the action-only resource
	data.Id = types.StringValue("rnatsession-config")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RnatsessionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data RnatsessionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read is a no-op for rnatsession: NITRO exposes no GET endpoint for this
	// action-only resource, so there is nothing to reconcile. Preserve state.
	tflog.Debug(ctx, "Read is a no-op for rnatsession; no GET endpoint on NITRO side")

	// Save (unchanged) data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RnatsessionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state RnatsessionResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// Update is a no-op for rnatsession; all filter attributes are RequiresReplace,
	// so any change re-triggers the flush action via Create/Delete instead of Update.
	tflog.Debug(ctx, "Update is a no-op for rnatsession; all attributes are RequiresReplace")

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RnatsessionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data RnatsessionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Action-only resource: flush has no inverse API, so Delete just removes from state.
	tflog.Debug(ctx, "Delete is a no-op for rnatsession; removing from Terraform state")
}
