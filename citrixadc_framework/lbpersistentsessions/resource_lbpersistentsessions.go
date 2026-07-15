package lbpersistentsessions

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &LbpersistentsessionsResource{}
var _ resource.ResourceWithConfigure = (*LbpersistentsessionsResource)(nil)

func NewLbpersistentsessionsResource() resource.Resource {
	return &LbpersistentsessionsResource{}
}

// LbpersistentsessionsResource defines the resource implementation.
// lbpersistentsessions is an ACTION-ONLY resource: NITRO exposes get(all),
// count and clear (POST ?action=clear). There is no add, update/set, or delete
// endpoint.
type LbpersistentsessionsResource struct {
	client *service.NitroClient
}

func (r *LbpersistentsessionsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lbpersistentsessions"
}

func (r *LbpersistentsessionsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *LbpersistentsessionsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data LbpersistentsessionsResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating (clear action) lbpersistentsessions resource")
	payload := lbpersistentsessionsGetThePayloadFromthePlan(ctx, &data)

	// clear is a POST ?action=clear action (Pattern 1). There is no add endpoint;
	// AddResource/UpdateUnnamedResource would target a nonexistent verb. The verb
	// casing is lowercase "clear" per the NITRO URL.
	err := r.client.ActOnResource(service.Lbpersistentsessions.Type(), payload, "clear")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to clear lbpersistentsessions, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Cleared lbpersistentsessions")

	// Set synthetic ID once, here in Create. The cleared sessions are not a
	// queryable managed object, so this ID is purely a Terraform state handle,
	// not a NITRO lookup key.
	data.Id = types.StringValue("lbpersistentsessions")

	// Read is a no-op (no stable GET-backed managed object); state is what the
	// plan supplied plus the synthetic ID.
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbpersistentsessionsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data LbpersistentsessionsResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read is a documented no-op (Pattern 13). lbpersistentsessions.clear flushes
	// persistence sessions; the cleared sessions are not a persistent object keyed
	// by these inputs, so a GET cannot stably re-resolve "this" record.
	// Re-fetching would cause perpetual drift / spurious resource removal.
	// Preserve state.
	tflog.Debug(ctx, "Read is a no-op for lbpersistentsessions (clear action has no stable GET-backed object); preserving state")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbpersistentsessionsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state LbpersistentsessionsResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Update is a no-op (Pattern 5): NITRO exposes no update/set endpoint for
	// lbpersistentsessions, and every attribute is RequiresReplace, so Terraform
	// never reaches here with real diffs.
	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for lbpersistentsessions; no NITRO update endpoint and all attributes are RequiresReplace")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbpersistentsessionsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data LbpersistentsessionsResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Delete is a state-only removal. "clear" is a one-shot action with no inverse
	// API; there is no NITRO delete endpoint for lbpersistentsessions. Removing the
	// resource from Terraform state is the only meaningful operation.
	tflog.Debug(ctx, "Delete is a state-only removal for lbpersistentsessions; clear has no inverse NITRO endpoint")
}
