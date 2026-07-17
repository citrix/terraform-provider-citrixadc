package lsnsipalgcall

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/lsn"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &LsnsipalgcallFlushResource{}
var _ resource.ResourceWithConfigure = (*LsnsipalgcallFlushResource)(nil)
var _ resource.ResourceWithImportState = (*LsnsipalgcallFlushResource)(nil)

func NewLsnsipalgcallFlushResource() resource.Resource {
	return &LsnsipalgcallFlushResource{}
}

// LsnsipalgcallFlushResource defines the resource implementation.
type LsnsipalgcallFlushResource struct {
	client *service.NitroClient
}

// LsnsipalgcallFlushResourceModel describes the resource data model.
//
// This resource models the NITRO lsnsipalgcall `?action=flush` action. flush is
// a one-shot runtime side-effect action with no add/update/delete endpoint, so
// Read/Update/Delete are no-ops. The flush action terminates the SIP ALG call
// identified by callid.
type LsnsipalgcallFlushResourceModel struct {
	Id     types.String `tfsdk:"id"`
	Callid types.String `tfsdk:"callid"`
	Nodeid types.Int64  `tfsdk:"nodeid"`
}

func (r *LsnsipalgcallFlushResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *LsnsipalgcallFlushResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lsnsipalgcall_flush"
}

func (r *LsnsipalgcallFlushResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *LsnsipalgcallFlushResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the lsnsipalgcall_flush resource.",
			},
			"callid": schema.StringAttribute{
				// NITRO flush marks callid mandatory; it is the only field in the
				// flush payload. Required (not Optional+Computed). Read is a no-op,
				// so it must not be Computed.
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Call ID for the SIP call.",
			},
			"nodeid": schema.Int64Attribute{
				// nodeid is a GET-only filter argument (it appears only in the get
				// args=callid:...,nodeid:... list, NOT in the flush payload).
				// Kept Optional, excluded from the flush payload (Pattern 15).
				// Read is a no-op, so it must not be Computed.
				Optional:    true,
				Description: "Unique number that identifies the cluster node.",
			},
		},
	}
}

// Create performs the flush action. lsnsipalgcall_flush is an ACTION-ONLY
// runtime resource: NITRO exposes flush as POST ?action=flush. There is no
// add/update/delete. The flush action terminates the SIP ALG call identified by
// callid.
func (r *LsnsipalgcallFlushResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data LsnsipalgcallFlushResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Flushing lsnsipalgcall (action-only resource)")
	payload := lsnsipalgcall_flushGetThePayloadFromthePlan(ctx, &data)

	// NITRO exposes flush as POST ?action=flush. Use ActOnResource with the
	// case-sensitive "flush" verb (lower-case per the NITRO URL).
	err := r.client.ActOnResource(service.Lsnsipalgcall.Type(), &payload, "flush")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to flush lsnsipalgcall, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Flushed lsnsipalgcall")

	// Synthetic ID for the action-only resource; keeps Read/Delete no-ops
	// addressable by Terraform.
	data.Id = types.StringValue("lsnsipalgcall_flush")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Read is a no-op for this action-only runtime resource. The flushed SIP ALG
// call is transient runtime state; NITRO has no query endpoint that reports
// flush-state, so Read is a pure preserve-state no-op.
func (r *LsnsipalgcallFlushResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data LsnsipalgcallFlushResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read is a no-op for lsnsipalgcall_flush; NITRO has no query endpoint for flush state")

	// Save unchanged data back into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Update is a no-op. NITRO has no update endpoint for flush; every schema
// attribute is RequiresReplace, so Terraform never invokes Update for a real
// change.
func (r *LsnsipalgcallFlushResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state LsnsipalgcallFlushResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for lsnsipalgcall_flush; NITRO has no update endpoint and all attributes are RequiresReplace")

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Delete is state-only. flush is a one-shot side-effect action with no inverse
// NITRO API (no "un-flush"). Delete simply removes the resource from Terraform
// state.
func (r *LsnsipalgcallFlushResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete is a no-op for lsnsipalgcall_flush; NITRO has no inverse of the flush action")
}

func lsnsipalgcall_flushGetThePayloadFromthePlan(ctx context.Context, data *LsnsipalgcallFlushResourceModel) lsn.Lsnsipalgcall {
	tflog.Debug(ctx, "In lsnsipalgcall_flushGetThePayloadFromthePlan Function")

	// Create API request body from the model.
	// Only callid is accepted by the flush action. nodeid is a GET-only filter
	// (Pattern 15) and is intentionally excluded from the flush payload.
	lsnsipalgcall := lsn.Lsnsipalgcall{}
	if !data.Callid.IsNull() && !data.Callid.IsUnknown() {
		lsnsipalgcall.Callid = data.Callid.ValueString()
	}

	return lsnsipalgcall
}
