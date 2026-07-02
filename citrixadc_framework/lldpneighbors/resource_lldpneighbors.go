package lldpneighbors

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
var _ resource.Resource = &LldpneighborsResource{}
var _ resource.ResourceWithConfigure = (*LldpneighborsResource)(nil)
var _ resource.ResourceWithImportState = (*LldpneighborsResource)(nil)

func NewLldpneighborsResource() resource.Resource {
	return &LldpneighborsResource{}
}

// LldpneighborsResource defines the resource implementation.
type LldpneighborsResource struct {
	client *service.NitroClient
}

func (r *LldpneighborsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *LldpneighborsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lldpneighbors"
}

func (r *LldpneighborsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

// Create fires the NITRO ?action=clear (POST) with an empty payload.
// lldpneighbors is an action-only resource: NITRO exposes only get(all),
// get-by-name, count and clear. There is no add/set/update/delete endpoint.
// The clear action takes NO args (bare) with body {"lldpneighbors":{}}.
func (r *LldpneighborsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data LldpneighborsResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Clearing lldpneighbors table via ?action=clear")

	// clear takes NO args; send an empty payload {"lldpneighbors":{}}.
	payload := map[string]interface{}{}
	err := r.client.ActOnResource(service.Lldpneighbors.Type(), payload, "clear")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to clear lldpneighbors, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Cleared lldpneighbors table")

	// Fixed synthetic ID; lldpneighbors is a transient action-only resource.
	data.Id = types.StringValue("lldpneighbors")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Read is a no-op. lldpneighbors is a transient diagnostics table with no
// persistent per-resource state to reconcile; the clear action leaves nothing
// to read back. Drift detection is not meaningful for this action-only resource.
func (r *LldpneighborsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data LldpneighborsResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read is a no-op for lldpneighbors; action-only resource with no persistent state")

	// Save prior state back unchanged
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Update is a no-op. lldpneighbors has no set/update endpoint and all
// attributes are RequiresReplace; Terraform never invokes a real update.
func (r *LldpneighborsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state LldpneighborsResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Update is a no-op for lldpneighbors; no set endpoint")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Delete is a no-op. lldpneighbors has no per-neighbor delete endpoint; only
// the clear action flushes the table. Removing the resource from state is
// sufficient.
func (r *LldpneighborsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data LldpneighborsResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Delete is a no-op for lldpneighbors; only ?action=clear flushes the table")
}
