package nstrace

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

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &NstraceStopResource{}
var _ resource.ResourceWithConfigure = (*NstraceStopResource)(nil)
var _ resource.ResourceWithImportState = (*NstraceStopResource)(nil)

func NewNstraceStopResource() resource.Resource {
	return &NstraceStopResource{}
}

// NstraceStopResource defines the resource implementation.
type NstraceStopResource struct {
	client *service.NitroClient
}

// NstraceStopResourceModel describes the resource data model.
//
// This resource wraps the nstrace `?action=stop` action, which stops a running
// packet trace. The stop action carries an EMPTY payload ({"nstrace":{}}) and is
// idempotent (stopping when nothing is running succeeds). There is no per-action
// GET, so Read/Update are no-ops and Delete is a state-only removal. See the
// systemscalablemgmtthreads package for the same multi-action pattern.
type NstraceStopResourceModel struct {
	Id types.String `tfsdk:"id"`
}

func (r *NstraceStopResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nstrace_stop"
}

func (r *NstraceStopResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NstraceStopResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NstraceStopResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nstrace_stop resource.",
			},
		},
	}
}

func (r *NstraceStopResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NstraceStopResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Stopping nstrace (action-only resource)")
	payload := basic.Nstrace{}

	err := r.client.ActOnResource(service.Nstrace.Type(), &payload, "stop")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to stop nstrace, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Triggered nstrace stop")

	data.Id = types.StringValue("nstrace_stop")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NstraceStopResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NstraceStopResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read is a no-op for nstrace_stop; NITRO has no per-action query endpoint")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NstraceStopResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state NstraceStopResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for nstrace_stop; NITRO has no update endpoint")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NstraceStopResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete is a no-op for nstrace_stop; the stop action has no NITRO inverse")
}
