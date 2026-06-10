package gslbldnsentries

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
var _ resource.Resource = &GslbldnsentriesResource{}
var _ resource.ResourceWithConfigure = (*GslbldnsentriesResource)(nil)
var _ resource.ResourceWithImportState = (*GslbldnsentriesResource)(nil)

func NewGslbldnsentriesResource() resource.Resource {
	return &GslbldnsentriesResource{}
}

// GslbldnsentriesResource defines the resource implementation.
// gslbldnsentries is an ACTION-ONLY resource: NITRO exposes get(all), count and
// clear (POST ?action=clear). There is no add, update/set, or delete endpoint.
type GslbldnsentriesResource struct {
	client *service.NitroClient
}

func (r *GslbldnsentriesResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *GslbldnsentriesResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gslbldnsentries"
}

func (r *GslbldnsentriesResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *GslbldnsentriesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data GslbldnsentriesResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating (clear action) gslbldnsentries resource")
	payload := gslbldnsentriesGetThePayloadFromthePlan(ctx, &data)

	// clear is a POST ?action=clear action (Pattern 1). There is no add endpoint;
	// AddResource/UpdateUnnamedResource would target a nonexistent verb. The verb
	// casing is lowercase "clear" per the NITRO URL. clear takes no arguments, so
	// the payload is an empty map (nodeid is a GET-only filter, Pattern 15).
	err := r.client.ActOnResource(service.Gslbldnsentries.Type(), payload, "clear")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to clear gslbldnsentries, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Cleared gslbldnsentries")

	// Set synthetic ID once, here in Create. The cleared LDNS entries are not a
	// queryable managed object, so this ID is purely a Terraform state handle, not
	// a NITRO lookup key.
	data.Id = types.StringValue("gslbldnsentries")

	// Read is a no-op (no stable GET-backed managed object); state is what the
	// plan supplied plus the synthetic ID.
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GslbldnsentriesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data GslbldnsentriesResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read is a documented no-op (Pattern 13). gslbldnsentries.clear flushes the
	// LDNS entries; clear is an action with no persistent object keyed by these
	// inputs, so a GET cannot stably re-resolve "this" record. Re-fetching would
	// cause perpetual drift. Preserve state.
	tflog.Debug(ctx, "Read is a no-op for gslbldnsentries (clear action has no stable GET-backed object); preserving state")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GslbldnsentriesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state GslbldnsentriesResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Update is a no-op (Pattern 5): NITRO exposes no update/set endpoint for
	// gslbldnsentries, and nodeid is RequiresReplace, so Terraform never reaches
	// here with real diffs.
	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for gslbldnsentries; no NITRO update endpoint and nodeid is RequiresReplace")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GslbldnsentriesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data GslbldnsentriesResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Delete is a state-only removal. "clear" is a one-shot action with no inverse
	// API; there is no NITRO delete endpoint for gslbldnsentries. Removing the
	// resource from Terraform state is the only meaningful operation.
	tflog.Debug(ctx, "Delete is a state-only removal for gslbldnsentries; clear has no inverse NITRO endpoint")
}
