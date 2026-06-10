package lsnrtspalgsession

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
var _ resource.Resource = &LsnrtspalgsessionResource{}
var _ resource.ResourceWithConfigure = (*LsnrtspalgsessionResource)(nil)
var _ resource.ResourceWithImportState = (*LsnrtspalgsessionResource)(nil)

func NewLsnrtspalgsessionResource() resource.Resource {
	return &LsnrtspalgsessionResource{}
}

// LsnrtspalgsessionResource defines the resource implementation.
// lsnrtspalgsession is an ACTION-ONLY runtime resource: NITRO exposes get(all),
// get(by sessionid), count and flush (POST ?action=flush). There is no add,
// update/set, or delete endpoint.
type LsnrtspalgsessionResource struct {
	client *service.NitroClient
}

func (r *LsnrtspalgsessionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *LsnrtspalgsessionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lsnrtspalgsession"
}

func (r *LsnrtspalgsessionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *LsnrtspalgsessionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data LsnrtspalgsessionResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating (flush action) lsnrtspalgsession resource")
	payload := lsnrtspalgsessionGetThePayloadFromthePlan(ctx, &data)

	// flush is a POST ?action=flush action (Pattern 1). There is no add endpoint;
	// AddResource/UpdateUnnamedResource would target a nonexistent verb.
	err := r.client.ActOnResource(service.Lsnrtspalgsession.Type(), payload, "flush")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to flush lsnrtspalgsession, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Flushed lsnrtspalgsession")

	// Set synthetic ID once, here in Create (Pattern 6). The flushed runtime
	// session is not a stable managed object, so this ID is purely a Terraform
	// state handle, not a NITRO lookup key.
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Sessionid.ValueString()))

	// Read is a no-op (no stable GET-backed managed object); state is what the
	// plan supplied plus the synthetic ID.
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsnrtspalgsessionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data LsnrtspalgsessionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read is a documented no-op (Pattern 13). lsnrtspalgsession.flush clears a
	// transient runtime RTSP-ALG session; that session is not a persistent object
	// stably keyed by these inputs, and may already be gone. Re-fetching would
	// cause perpetual drift / spurious resource removal. Preserve state.
	tflog.Debug(ctx, "Read is a no-op for lsnrtspalgsession (flush action has no stable GET-backed object); preserving state")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsnrtspalgsessionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state LsnrtspalgsessionResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Update is a no-op (Pattern 5): NITRO exposes no update/set endpoint for
	// lsnrtspalgsession, and every attribute is RequiresReplace, so Terraform
	// never reaches here with real diffs.
	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for lsnrtspalgsession; no NITRO update endpoint and all attributes are RequiresReplace")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsnrtspalgsessionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data LsnrtspalgsessionResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Delete is a state-only removal. "flush" is a one-shot action with no inverse
	// API; there is no NITRO delete endpoint for lsnrtspalgsession. Removing the
	// resource from Terraform state is the only meaningful operation.
	tflog.Debug(ctx, "Delete is a state-only removal for lsnrtspalgsession; flush has no inverse NITRO endpoint")
}
