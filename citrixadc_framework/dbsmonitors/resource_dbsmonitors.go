package dbsmonitors

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// dbsmonitors is an ACTION-ONLY, ZERO-ATTRIBUTE resource.
//
//   - NITRO exposes only the restart action:
//     POST /nitro/v1/config/dbsmonitors?action=restart, which restarts the
//     database (DBS) monitors.
//   - There is NO add/set/get/delete endpoint, so:
//     Create performs the restart action, Read/Update are no-ops (there is
//     nothing to reconcile), and Delete is a state-only removal.
//   - Because there is no GET endpoint, there is NO datasource for dbsmonitors.
var _ resource.Resource = &DbsmonitorsResource{}
var _ resource.ResourceWithConfigure = (*DbsmonitorsResource)(nil)
var _ resource.ResourceWithImportState = (*DbsmonitorsResource)(nil)

func NewDbsmonitorsResource() resource.Resource {
	return &DbsmonitorsResource{}
}

// DbsmonitorsResource defines the resource implementation.
type DbsmonitorsResource struct {
	client *service.NitroClient
}

func (r *DbsmonitorsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *DbsmonitorsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dbsmonitors"
}

func (r *DbsmonitorsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *DbsmonitorsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data DbsmonitorsResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating dbsmonitors resource (restart action)")
	dbsmonitors := dbsmonitorsGetThePayloadFromthePlan(ctx, &data)

	// Action-only resource - NITRO exposes only POST ?action=restart
	err := r.client.ActOnResource(service.Dbsmonitors.Type(), &dbsmonitors, "restart")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to restart dbsmonitors, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Restarted dbsmonitors")

	// Synthetic ID - no GET endpoint exists to derive it from
	data.Id = types.StringValue("dbsmonitors-config")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Read is a no-op. dbsmonitors has no GET endpoint; there is nothing to reconcile.
func (r *DbsmonitorsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data DbsmonitorsResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read is a no-op for dbsmonitors; NITRO exposes no GET endpoint (action=restart only)")

	// Preserve prior state unchanged - no GET endpoint to reconcile against
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Update is a no-op. dbsmonitors has no attributes and no set endpoint.
func (r *DbsmonitorsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state DbsmonitorsResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Update is a no-op for dbsmonitors; it has no attributes and no set endpoint")

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Delete is a no-op. dbsmonitors has no delete endpoint; the action is not
// reversible and there is no persistent object to remove.
func (r *DbsmonitorsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete is a no-op for dbsmonitors; NITRO has no delete endpoint")
}
