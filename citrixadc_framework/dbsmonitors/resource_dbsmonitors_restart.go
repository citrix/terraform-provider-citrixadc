package dbsmonitors

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/basic"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// dbsmonitors_restart is an ACTION-ONLY, ZERO-ATTRIBUTE resource.
//
//   - NITRO exposes only the restart action:
//     POST /nitro/v1/config/dbsmonitors?action=restart, which restarts the
//     database (DBS) monitors.
//   - There is NO add/set/get/delete endpoint, so:
//     Create performs the restart action, Read/Update are no-ops (there is
//     nothing to reconcile), and Delete is a state-only removal.
//   - Because there is no GET endpoint, there is NO datasource for
//     dbsmonitors_restart.
var _ resource.Resource = &DbsmonitorsRestartResource{}
var _ resource.ResourceWithConfigure = (*DbsmonitorsRestartResource)(nil)
var _ resource.ResourceWithImportState = (*DbsmonitorsRestartResource)(nil)

func NewDbsmonitorsRestartResource() resource.Resource {
	return &DbsmonitorsRestartResource{}
}

// DbsmonitorsRestartResource defines the resource implementation.
type DbsmonitorsRestartResource struct {
	client *service.NitroClient
}

// DbsmonitorsRestartResourceModel describes the resource data model.
//
// dbsmonitors_restart is a ZERO-ATTRIBUTE, ACTION-ONLY resource: the NITRO
// "dbsmonitors" object exposes no read/write properties and only the restart
// action (POST /nitro/v1/config/dbsmonitors?action=restart). The model
// therefore carries only the synthetic id.
type DbsmonitorsRestartResourceModel struct {
	Id types.String `tfsdk:"id"`
}

func (r *DbsmonitorsRestartResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *DbsmonitorsRestartResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dbsmonitors_restart"
}

func (r *DbsmonitorsRestartResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *DbsmonitorsRestartResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the dbsmonitors_restart resource.",
			},
		},
	}
}

func (r *DbsmonitorsRestartResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data DbsmonitorsRestartResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating dbsmonitors_restart resource (restart action)")
	dbsmonitors := dbsmonitors_restartGetThePayloadFromthePlan(ctx, &data)

	// Action-only resource - NITRO exposes only POST ?action=restart
	err := r.client.ActOnResource(service.Dbsmonitors.Type(), &dbsmonitors, "restart")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to restart dbsmonitors, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Restarted dbsmonitors")

	// Synthetic ID - no GET endpoint exists to derive it from
	data.Id = types.StringValue("dbsmonitors_restart")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Read is a no-op. dbsmonitors_restart has no GET endpoint; there is nothing to
// reconcile.
func (r *DbsmonitorsRestartResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data DbsmonitorsRestartResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read is a no-op for dbsmonitors_restart; NITRO exposes no GET endpoint (action=restart only)")

	// Preserve prior state unchanged - no GET endpoint to reconcile against
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Update is a no-op. dbsmonitors_restart has no attributes and no set endpoint;
// the synthetic id is RequiresReplace-equivalent (nothing to reconcile).
func (r *DbsmonitorsRestartResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state DbsmonitorsRestartResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Update is a no-op for dbsmonitors_restart; it has no attributes and no set endpoint")

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Delete is a no-op. dbsmonitors_restart has no delete endpoint; the action is
// not reversible and there is no persistent object to remove.
func (r *DbsmonitorsRestartResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete is a no-op for dbsmonitors_restart; NITRO has no delete endpoint")
}

// dbsmonitors_restartGetThePayloadFromthePlan builds the (empty) NITRO payload
// for the restart action. dbsmonitors has no read/write attributes, so the
// payload is an empty basic.Dbsmonitors struct.
func dbsmonitors_restartGetThePayloadFromthePlan(ctx context.Context, data *DbsmonitorsRestartResourceModel) basic.Dbsmonitors {
	tflog.Debug(ctx, "In dbsmonitors_restartGetThePayloadFromthePlan Function")
	dbsmonitors := basic.Dbsmonitors{}
	return dbsmonitors
}
