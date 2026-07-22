package lbpersistentsessions

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &LbpersistentsessionsClearResource{}
var _ resource.ResourceWithConfigure = (*LbpersistentsessionsClearResource)(nil)

func NewLbpersistentsessionsClearResource() resource.Resource {
	return &LbpersistentsessionsClearResource{}
}

// LbpersistentsessionsClearResource models the NITRO lbpersistentsessions
// `?action=clear` action. clear is a one-shot side-effect action (POST) with no
// GET endpoint and no inverse API, so Read/Update/Delete are no-ops. vserver and
// persistenceparameter are optional clear filters. nodeid is a GET-only cluster
// filter (Pattern 15): it is kept in the schema per the per-action guidance but
// is deliberately excluded from the clear action body (see the payload builder).
type LbpersistentsessionsClearResource struct {
	client *service.NitroClient
}

// LbpersistentsessionsClearResourceModel describes the resource data model.
type LbpersistentsessionsClearResourceModel struct {
	Id                   types.String `tfsdk:"id"`
	Nodeid               types.Int64  `tfsdk:"nodeid"`
	Persistenceparameter types.String `tfsdk:"persistenceparameter"`
	Vserver              types.String `tfsdk:"vserver"`
}

func (r *LbpersistentsessionsClearResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lbpersistentsessions_clear"
}

func (r *LbpersistentsessionsClearResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *LbpersistentsessionsClearResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the lbpersistentsessions_clear resource.",
			},
			// All clear arguments are optional filters. Read is a no-op (clear is an
			// action, the cleared sessions are not a persistent managed object), so
			// these must NOT be Computed or Terraform reports an unknown value after
			// apply (Pattern 13 schema-flag implication).
			"nodeid": schema.Int64Attribute{
				Optional: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Unique number that identifies the cluster node.",
			},
			"persistenceparameter": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The persistence parameter whose persistence sessions are to be flushed.",
			},
			"vserver": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The name of the virtual server.",
			},
		},
	}
}

func (r *LbpersistentsessionsClearResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data LbpersistentsessionsClearResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Clearing lbpersistentsessions (action-only resource)")
	payload := lbpersistentsessions_clearGetThePayloadFromthePlan(ctx, &data)

	// clear is a POST ?action=clear action (Pattern 1). There is no add endpoint;
	// AddResource/UpdateUnnamedResource would target a nonexistent verb. The verb
	// casing is lowercase "clear" per the NITRO URL.
	err := r.client.ActOnResource(service.Lbpersistentsessions.Type(), payload, "clear")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to clear lbpersistentsessions, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Cleared lbpersistentsessions")

	// Synthetic ID for the action-only resource; keeps Read/Delete no-ops
	// addressable by Terraform.
	data.Id = types.StringValue("lbpersistentsessions_clear")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbpersistentsessionsClearResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// clear is a one-shot action. NITRO has no GET endpoint that reports
	// clear-state, so Read is a pure preserve-state no-op.
	var data LbpersistentsessionsClearResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read is a no-op for lbpersistentsessions_clear; NITRO has no query endpoint for clear state")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbpersistentsessionsClearResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// NITRO has no update endpoint for clear; every schema attribute is
	// RequiresReplace, so Terraform never invokes Update for a real change.
	var data, state LbpersistentsessionsClearResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for lbpersistentsessions_clear; NITRO has no update endpoint and all attributes are RequiresReplace")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbpersistentsessionsClearResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// clear is a one-shot side-effect action. There is no inverse NITRO API
	// (no "un-clear"). Delete simply removes the resource from Terraform state.
	tflog.Debug(ctx, "Delete is a no-op for lbpersistentsessions_clear; NITRO has no inverse of the clear action")
}

// lbpersistentsessions_clearGetThePayloadFromthePlan builds the body for the
// clear action. Build a map containing ONLY the clear arguments that are set so
// the action body never includes invalid arguments. nodeid is a GET-only cluster
// filter and is intentionally excluded from the clear payload (Pattern 15).
func lbpersistentsessions_clearGetThePayloadFromthePlan(ctx context.Context, data *LbpersistentsessionsClearResourceModel) map[string]interface{} {
	tflog.Debug(ctx, "In lbpersistentsessions_clearGetThePayloadFromthePlan Function")

	lbpersistentsessions := map[string]interface{}{}
	if !data.Persistenceparameter.IsNull() && !data.Persistenceparameter.IsUnknown() {
		lbpersistentsessions["persistenceparameter"] = data.Persistenceparameter.ValueString()
	}
	if !data.Vserver.IsNull() && !data.Vserver.IsUnknown() {
		lbpersistentsessions["vserver"] = data.Vserver.ValueString()
	}

	return lbpersistentsessions
}
