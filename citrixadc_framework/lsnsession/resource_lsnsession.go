package lsnsession

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
var _ resource.Resource = &LsnsessionResource{}
var _ resource.ResourceWithConfigure = (*LsnsessionResource)(nil)
var _ resource.ResourceWithImportState = (*LsnsessionResource)(nil)

// lsnsessionSyntheticId is a fixed identifier used because lsnsession is an
// action-only (flush) runtime resource with no get-by-name key on NITRO.
const lsnsessionSyntheticId = "lsnsession"

func NewLsnsessionResource() resource.Resource {
	return &LsnsessionResource{}
}

// LsnsessionResource defines the resource implementation.
type LsnsessionResource struct {
	client *service.NitroClient
}

func (r *LsnsessionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *LsnsessionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lsnsession"
}

func (r *LsnsessionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *LsnsessionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data LsnsessionResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Flushing lsnsession resource")
	payload := lsnsessionGetThePayloadFromthePlan(ctx, &data)

	// lsnsession is action-only: the only write verb is POST ?action=flush.
	err := r.client.ActOnResource(service.Lsnsession.Type(), payload, "flush")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to flush lsnsession, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Flushed lsnsession resource")

	// Set the synthetic ID once (Pattern 6) - there is no get-by-name key.
	data.Id = types.StringValue(lsnsessionSyntheticId)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsnsessionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data LsnsessionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read is a no-op for lsnsession: it is an action-only (flush) resource with
	// no get-by-name key on NITRO, so drift detection is impossible by design.
	tflog.Debug(ctx, "Read is a no-op for lsnsession; action-only flush resource with no get-by-name key")

	// Save (unchanged) state back into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsnsessionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state LsnsessionResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// Update is a no-op for lsnsession; all attributes are RequiresReplace and
	// NITRO exposes no update verb for this action-only resource.
	tflog.Debug(ctx, "Update is a no-op for lsnsession; all attributes are RequiresReplace")

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsnsessionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data LsnsessionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Delete is state-only removal for lsnsession: NITRO has no delete verb for
	// this action-only (flush) resource. Terraform removes it from state.
	tflog.Debug(ctx, "Delete is state-only for lsnsession; no NITRO delete verb")
}
