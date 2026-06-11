package policytracing

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
var _ resource.Resource = &PolicytracingResource{}
var _ resource.ResourceWithConfigure = (*PolicytracingResource)(nil)
var _ resource.ResourceWithImportState = (*PolicytracingResource)(nil)

func NewPolicytracingResource() resource.Resource {
	return &PolicytracingResource{}
}

// PolicytracingResource defines the resource implementation.
// policytracing is an ACTION-ONLY resource: NITRO exposes get(all), count and
// clear (POST ?action=clear). There is no add, update/set, or delete endpoint
// (confirmed by NITRO doc Operations section and live CLI -- `set policytracing`
// does not exist). The only write capability is the clear reset action.
type PolicytracingResource struct {
	client *service.NitroClient
}

func (r *PolicytracingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *PolicytracingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_policytracing"
}

func (r *PolicytracingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *PolicytracingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data PolicytracingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating (clear action) policytracing resource")
	payload := policytracingGetThePayloadFromthePlan(ctx, &data)

	// clear is a POST ?action=clear action (Pattern 1). There is no add endpoint;
	// AddResource/UpdateUnnamedResource would target a nonexistent verb. The clear
	// body is empty per the NITRO doc ({"policytracing":{}}).
	err := r.client.ActOnResource(service.Policytracing.Type(), payload, "clear")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to clear policytracing, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Cleared policytracing")

	// Set synthetic constant ID once, here in Create (Pattern 6). clear is a
	// one-shot reset action; the resource is not a queryable managed object, so
	// this ID is purely a Terraform state handle, not a NITRO lookup key.
	data.Id = types.StringValue("policytracing")

	// Read is a no-op (no stable GET-backed managed object); state is the
	// synthetic ID only.
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PolicytracingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data PolicytracingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read is a documented no-op (Pattern 13). policytracing.clear resets the
	// captured policy-trace records; clear is an action with no persistent object
	// keyed by these inputs, so a GET cannot stably re-resolve "this" record.
	// Re-fetching would cause perpetual drift. Preserve state.
	tflog.Debug(ctx, "Read is a no-op for policytracing (clear action has no stable GET-backed object); preserving state")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PolicytracingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state PolicytracingResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Update is a no-op (Pattern 5): NITRO exposes no update/set endpoint for
	// policytracing. The schema has no writable attributes, so Terraform never
	// reaches here with real diffs.
	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for policytracing; no NITRO update endpoint")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PolicytracingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data PolicytracingResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Delete is a state-only removal. "clear" is a one-shot action with no inverse
	// API; there is no NITRO delete endpoint for policytracing. Removing the
	// resource from Terraform state is the only meaningful operation.
	tflog.Debug(ctx, "Delete is a state-only removal for policytracing; clear has no inverse NITRO endpoint")
}
