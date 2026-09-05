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
var _ resource.Resource = &NsmigrationCompleteResource{}
var _ resource.ResourceWithConfigure = (*NsmigrationCompleteResource)(nil)
var _ resource.ResourceWithImportState = (*NsmigrationCompleteResource)(nil)

func NewNsmigrationCompleteResource() resource.Resource {
	return &NsmigrationCompleteResource{}
}

// NsmigrationCompleteResource defines the resource implementation.
type NsmigrationCompleteResource struct {
	client *service.NitroClient
}

// NsmigrationCompleteResourceModel describes the resource data model.
//
// This resource wraps the nsmigration `?action=complete` action, which finalizes
// an in-progress NetScaler session migration. The complete action carries an
// EMPTY payload ({"nsmigration":{}}); it requires a migration to be in progress
// (NITRO errorcode 257 "Migration is not in progress" otherwise). There is no
// per-action GET, so Read/Update are no-ops and Delete is a state-only removal.
// See the systemscalablemgmtthreads package for the same multi-action pattern.
type NsmigrationCompleteResourceModel struct {
	Id types.String `tfsdk:"id"`
}

func (r *NsmigrationCompleteResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nsmigration_complete"
}

func (r *NsmigrationCompleteResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NsmigrationCompleteResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NsmigrationCompleteResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nsmigration_complete resource.",
			},
		},
	}
}

func (r *NsmigrationCompleteResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NsmigrationCompleteResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Completing nsmigration (action-only resource)")
	payload := ns.Nsmigration{}

	err := r.client.ActOnResource(service.Nsmigration.Type(), &payload, "complete")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to complete nsmigration, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Triggered nsmigration complete")

	data.Id = types.StringValue("nsmigration_complete")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsmigrationCompleteResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NsmigrationCompleteResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read is a no-op for nsmigration_complete; NITRO has no per-action query endpoint")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsmigrationCompleteResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state NsmigrationCompleteResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for nsmigration_complete; NITRO has no update endpoint")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsmigrationCompleteResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete is a no-op for nsmigration_complete; the complete action has no NITRO inverse")
}
