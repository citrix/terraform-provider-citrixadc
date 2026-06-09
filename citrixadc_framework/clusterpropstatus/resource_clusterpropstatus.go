package clusterpropstatus

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
var _ resource.Resource = &ClusterpropstatusResource{}
var _ resource.ResourceWithConfigure = (*ClusterpropstatusResource)(nil)
var _ resource.ResourceWithImportState = (*ClusterpropstatusResource)(nil)

func NewClusterpropstatusResource() resource.Resource {
	return &ClusterpropstatusResource{}
}

// ClusterpropstatusResource defines the resource implementation.
// clusterpropstatus is an ACTION-ONLY resource: NITRO exposes get(all), count
// and clear (POST ?action=clear). There is no add, update/set, or delete
// endpoint (confirmed by NITRO doc and live CLI). The only write capability is
// the clear reset action.
type ClusterpropstatusResource struct {
	client *service.NitroClient
}

func (r *ClusterpropstatusResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *ClusterpropstatusResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_clusterpropstatus"
}

func (r *ClusterpropstatusResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *ClusterpropstatusResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ClusterpropstatusResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating (clear action) clusterpropstatus resource")
	payload := clusterpropstatusGetThePayloadFromthePlan(ctx, &data)

	// clear is a POST ?action=clear action (Pattern 1). There is no add endpoint;
	// AddResource/UpdateUnnamedResource would target a nonexistent verb.
	err := r.client.ActOnResource(service.Clusterpropstatus.Type(), payload, "clear")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to clear clusterpropstatus, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Cleared clusterpropstatus")

	// Set synthetic ID once, here in Create. clear is a one-shot reset action;
	// the resource is not a queryable managed object, so this ID is purely a
	// Terraform state handle, not a NITRO lookup key.
	data.Id = types.StringValue("clusterpropstatus")

	// Read is a no-op (no stable GET-backed managed object); state is what the
	// plan supplied plus the synthetic ID.
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ClusterpropstatusResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ClusterpropstatusResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read is a documented no-op (Pattern 13). clusterpropstatus.clear resets the
	// property-propagation status counters; clear is an action with no persistent
	// object keyed by these inputs, so a GET cannot stably re-resolve "this"
	// record. Re-fetching would cause perpetual drift. Preserve state.
	tflog.Debug(ctx, "Read is a no-op for clusterpropstatus (clear action has no stable GET-backed object); preserving state")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ClusterpropstatusResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state ClusterpropstatusResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Update is a no-op (Pattern 5): NITRO exposes no update/set endpoint for
	// clusterpropstatus, and nodeid is RequiresReplace, so Terraform never
	// reaches here with real diffs.
	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for clusterpropstatus; no NITRO update endpoint and nodeid is RequiresReplace")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ClusterpropstatusResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ClusterpropstatusResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Delete is a state-only removal. "clear" is a one-shot action with no inverse
	// API; there is no NITRO delete endpoint for clusterpropstatus. Removing the
	// resource from Terraform state is the only meaningful operation.
	tflog.Debug(ctx, "Delete is a state-only removal for clusterpropstatus; clear has no inverse NITRO endpoint")
}
