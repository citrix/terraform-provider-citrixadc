package nsmigration

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/ns"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &NsmigrationStartResource{}
var _ resource.ResourceWithConfigure = (*NsmigrationStartResource)(nil)
var _ resource.ResourceWithImportState = (*NsmigrationStartResource)(nil)

func NewNsmigrationStartResource() resource.Resource {
	return &NsmigrationStartResource{}
}

// NsmigrationStartResource defines the resource implementation.
type NsmigrationStartResource struct {
	client *service.NitroClient
}

// NsmigrationStartResourceModel describes the resource data model.
//
// nsmigration is a NITRO object that supports multiple actions (start / stop /
// complete) plus a get. Mirroring the systemscalablemgmtthreads package, each
// action is modelled as its own action-only resource. This resource wraps the
// `?action=start` action, which begins a NetScaler session migration.
//
// The start action carries an EMPTY payload ({"nsmigration":{}}): dumpsession
// (the only read/write property) is a GET-only field and is rejected in the
// action payload (NITRO errorcode 278 "Invalid argument [dumpsession]"), so it is
// intentionally excluded here. There is no per-action GET, so Read/Update are
// no-ops; the live migration state is queryable via the citrixadc_nsmigration
// data source. Delete is a state-only removal (the inverse of start is the
// separate citrixadc_nsmigration_stop resource).
type NsmigrationStartResourceModel struct {
	Id types.String `tfsdk:"id"`
}

func (r *NsmigrationStartResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nsmigration_start"
}

func (r *NsmigrationStartResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NsmigrationStartResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NsmigrationStartResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nsmigration_start resource.",
			},
			// The start action accepts no attributes: the NITRO payload is empty and
			// dumpsession is rejected in the action body (errorcode 278).
		},
	}
}

func (r *NsmigrationStartResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NsmigrationStartResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Starting nsmigration (action-only resource)")
	// The start action carries an empty payload.
	payload := ns.Nsmigration{}

	// NITRO exposes start as POST ?action=start. Verb casing matches the URL.
	err := r.client.ActOnResource(service.Nsmigration.Type(), &payload, "start")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to start nsmigration, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Triggered nsmigration start")

	// Synthetic ID for the action-only resource.
	data.Id = types.StringValue("nsmigration_start")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsmigrationStartResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// start is a one-shot action with no per-action GET; Read is a preserve-state
	// no-op. Query live migration state via the data source instead.
	var data NsmigrationStartResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read is a no-op for nsmigration_start; NITRO has no per-action query endpoint")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsmigrationStartResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// NITRO has no update endpoint for the start action; there are no config
	// attributes to change, so Terraform never invokes Update for a real change.
	var data, state NsmigrationStartResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for nsmigration_start; NITRO has no update endpoint")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsmigrationStartResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// start is a one-shot action. Its inverse is the separate
	// citrixadc_nsmigration_stop resource, so Delete simply removes the resource
	// from Terraform state (it does NOT auto-stop the migration).
	tflog.Debug(ctx, "Delete is a no-op for nsmigration_start; use citrixadc_nsmigration_stop to stop the migration")
}
