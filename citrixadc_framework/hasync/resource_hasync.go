package hasync

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
var _ resource.Resource = &HasyncResource{}
var _ resource.ResourceWithConfigure = (*HasyncResource)(nil)
var _ resource.ResourceWithImportState = (*HasyncResource)(nil)

func NewHasyncResource() resource.Resource {
	return &HasyncResource{}
}

// HasyncResource defines the resource implementation.
type HasyncResource struct {
	client *service.NitroClient
}

func (r *HasyncResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *HasyncResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_hasync"
}

func (r *HasyncResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *HasyncResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data HasyncResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating hasync resource")

	// hasync exposes only the POST ?action=Force action on NITRO (capital F,
	// case-sensitive). There is no add/get/update/delete endpoint. Use
	// ActOnResource with the exact "Force" verb.
	payload := hasyncGetThePayloadFromthePlan(ctx, &data)

	err := r.client.ActOnResource(service.Hasync.Type(), payload, "Force")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to Force sync hasync, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Force synced hasync resource")

	// Synthetic constant ID - there is no NITRO identity for this action resource.
	data.Id = types.StringValue("hasync")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *HasyncResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data HasyncResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read is a no-op for hasync: NITRO exposes no GET endpoint for this
	// action-only resource, so drift detection is impossible. Preserve state.
	tflog.Debug(ctx, "Read is a no-op for hasync; no GET endpoint on NITRO side")

	// Save (unchanged) data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *HasyncResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state HasyncResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// Update is a no-op for hasync; force/save are RequiresReplace, so
	// Terraform re-creates (re-Force-syncs) on change instead.
	tflog.Debug(ctx, "Update is a no-op for hasync; all attributes are RequiresReplace")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *HasyncResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data HasyncResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Delete is a no-op for hasync: NITRO exposes no DELETE endpoint for
	// this action-only resource. Just remove from Terraform state.
	tflog.Debug(ctx, "Delete is a no-op for hasync; no DELETE endpoint on NITRO side")
	tflog.Trace(ctx, "Removed hasync from Terraform state")
}
