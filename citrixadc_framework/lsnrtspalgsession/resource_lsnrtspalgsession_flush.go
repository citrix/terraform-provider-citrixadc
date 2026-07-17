package lsnrtspalgsession

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &LsnrtspalgsessionFlushResource{}
var _ resource.ResourceWithConfigure = (*LsnrtspalgsessionFlushResource)(nil)
var _ resource.ResourceWithImportState = (*LsnrtspalgsessionFlushResource)(nil)

func NewLsnrtspalgsessionFlushResource() resource.Resource {
	return &LsnrtspalgsessionFlushResource{}
}

// LsnrtspalgsessionFlushResource defines the resource implementation.
//
// This resource models the NITRO lsnrtspalgsession `?action=flush` action. flush
// is a one-shot side-effect action with no add/update/set/delete endpoint and no
// inverse API, so Read/Update/Delete are no-ops. sessionid is the only flush
// payload argument (mandatory); nodeid is a GET-only cluster filter (Pattern 15)
// and is therefore excluded from the flush payload builder.
type LsnrtspalgsessionFlushResource struct {
	client *service.NitroClient
}

// LsnrtspalgsessionFlushResourceModel describes the resource data model.
type LsnrtspalgsessionFlushResourceModel struct {
	Id        types.String `tfsdk:"id"`
	Nodeid    types.Int64  `tfsdk:"nodeid"`
	Sessionid types.String `tfsdk:"sessionid"`
}

func (r *LsnrtspalgsessionFlushResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *LsnrtspalgsessionFlushResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lsnrtspalgsession_flush"
}

func (r *LsnrtspalgsessionFlushResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *LsnrtspalgsessionFlushResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the lsnrtspalgsession_flush resource.",
			},
			// nodeid is a GET-only cluster filter, not a flush payload argument
			// (Pattern 15). Read is a no-op (flush is an action on a transient
			// session), so it must NOT be Computed or Terraform reports an unknown
			// value after apply (Pattern 13 schema-flag implication).
			"nodeid": schema.Int64Attribute{
				Optional: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Unique number that identifies the cluster node.",
			},
			// sessionid is mandatory for the flush action (NITRO marks it red/bold).
			// Required, not Optional+Computed.
			"sessionid": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Session ID for the RTSP call.",
			},
		},
	}
}

func (r *LsnrtspalgsessionFlushResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data LsnrtspalgsessionFlushResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Flushing lsnrtspalgsession (action-only resource)")
	payload := lsnrtspalgsession_flushGetThePayloadFromthePlan(ctx, &data)

	// flush is a POST ?action=flush action (Pattern 1). There is no add endpoint;
	// AddResource/UpdateUnnamedResource would target a nonexistent verb.
	err := r.client.ActOnResource(service.Lsnrtspalgsession.Type(), payload, "flush")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to flush lsnrtspalgsession, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Flushed lsnrtspalgsession")

	// Synthetic ID for the action-only resource; keeps Read/Delete no-ops
	// addressable by Terraform.
	data.Id = types.StringValue("lsnrtspalgsession_flush")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsnrtspalgsessionFlushResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// flush is a one-shot action. NITRO has no GET endpoint that reports
	// flush-state, so Read is a pure preserve-state no-op.
	var data LsnrtspalgsessionFlushResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read is a no-op for lsnrtspalgsession_flush; NITRO has no query endpoint for flush state")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsnrtspalgsessionFlushResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// NITRO has no update endpoint for flush; every schema attribute is
	// RequiresReplace, so Terraform never invokes Update for a real change.
	var data, state LsnrtspalgsessionFlushResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for lsnrtspalgsession_flush; NITRO has no update endpoint and all attributes are RequiresReplace")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsnrtspalgsessionFlushResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// flush is a one-shot side-effect action. There is no inverse NITRO API
	// (no "un-flush"). Delete simply removes the resource from Terraform state.
	tflog.Debug(ctx, "Delete is a no-op for lsnrtspalgsession_flush; NITRO has no inverse of the flush action")
}

// lsnrtspalgsession_flushGetThePayloadFromthePlan builds the body for the flush
// action. Only sessionid is a valid flush argument. nodeid is a GET-only cluster
// filter and is intentionally excluded from the flush payload (Pattern 15).
func lsnrtspalgsession_flushGetThePayloadFromthePlan(ctx context.Context, data *LsnrtspalgsessionFlushResourceModel) map[string]interface{} {
	tflog.Debug(ctx, "In lsnrtspalgsession_flushGetThePayloadFromthePlan Function")

	lsnrtspalgsession := map[string]interface{}{}
	if !data.Sessionid.IsNull() && !data.Sessionid.IsUnknown() {
		lsnrtspalgsession["sessionid"] = data.Sessionid.ValueString()
	}

	return lsnrtspalgsession
}
