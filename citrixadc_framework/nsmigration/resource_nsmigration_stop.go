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
var _ resource.Resource = &NsmigrationStopResource{}
var _ resource.ResourceWithConfigure = (*NsmigrationStopResource)(nil)
var _ resource.ResourceWithImportState = (*NsmigrationStopResource)(nil)

func NewNsmigrationStopResource() resource.Resource {
	return &NsmigrationStopResource{}
}

// NsmigrationStopResource defines the resource implementation.
type NsmigrationStopResource struct {
	client *service.NitroClient
}

// NsmigrationStopResourceModel describes the resource data model.
//
// This resource wraps the nsmigration `?action=stop` action, which aborts an
// in-progress NetScaler session migration (rollback). The stop action carries an
// EMPTY payload ({"nsmigration":{}}); it requires a migration to be in progress
// (NITRO errorcode 257 "Migration is not in progress" otherwise). There is no
// per-action GET, so Read/Update are no-ops and Delete is a state-only removal.
// See the systemscalablemgmtthreads package for the same multi-action pattern.
type NsmigrationStopResourceModel struct {
	Id types.String `tfsdk:"id"`
}

func (r *NsmigrationStopResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nsmigration_stop"
}

func (r *NsmigrationStopResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NsmigrationStopResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NsmigrationStopResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nsmigration_stop resource.",
			},
		},
	}
}

func (r *NsmigrationStopResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NsmigrationStopResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Stopping nsmigration (action-only resource)")
	payload := ns.Nsmigration{}

	err := r.client.ActOnResource(service.Nsmigration.Type(), &payload, "stop")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to stop nsmigration, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Triggered nsmigration stop")

	data.Id = types.StringValue("nsmigration_stop")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsmigrationStopResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NsmigrationStopResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read is a no-op for nsmigration_stop; NITRO has no per-action query endpoint")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsmigrationStopResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state NsmigrationStopResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for nsmigration_stop; NITRO has no update endpoint")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsmigrationStopResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete is a no-op for nsmigration_stop; the stop action has no NITRO inverse")
}
