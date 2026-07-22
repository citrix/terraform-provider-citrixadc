package lldpneighbors

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
var _ resource.Resource = &LldpneighborsClearResource{}
var _ resource.ResourceWithConfigure = (*LldpneighborsClearResource)(nil)

func NewLldpneighborsClearResource() resource.Resource {
	return &LldpneighborsClearResource{}
}

// LldpneighborsClearResource defines the resource implementation.
type LldpneighborsClearResource struct {
	client *service.NitroClient
}

// LldpneighborsClearResourceModel describes the resource data model.
//
// This resource models the NITRO lldpneighbors `?action=clear` action. clear is
// a one-shot side-effect action (POST) with no inverse API, so
// Read/Update/Delete are no-ops. Per the NITRO doc and CLI, the clear payload is
// empty ({"lldpneighbors":{}}) and takes no args; ifnum/nodeid are GET/datasource
// filters only (Pattern 15) and are intentionally NOT sent in the clear payload.
type LldpneighborsClearResourceModel struct {
	Id     types.String `tfsdk:"id"`
	Ifnum  types.String `tfsdk:"ifnum"`
	Nodeid types.Int64  `tfsdk:"nodeid"`
}

func (r *LldpneighborsClearResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lldpneighbors_clear"
}

func (r *LldpneighborsClearResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *LldpneighborsClearResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the lldpneighbors_clear resource.",
			},
			// ifnum/nodeid are GET/datasource filters, not clear-action args, so
			// they are plain Optional inputs here (not Computed - Read is a no-op
			// and could never resolve an "unknown" value at apply time). All
			// config attributes are RequiresReplace on this action-only resource.
			"ifnum": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Interface Name",
			},
			"nodeid": schema.Int64Attribute{
				Optional: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Unique number that identifies the cluster node.",
			},
		},
	}
}

// Create fires the NITRO ?action=clear (POST) with an empty payload.
// lldpneighbors is an action-only resource on the write side: NITRO exposes only
// get(all), get-by-name, count and clear. There is no add/set/update/delete
// endpoint. The clear action takes NO args (bare) with body {"lldpneighbors":{}}.
func (r *LldpneighborsClearResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data LldpneighborsClearResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Clearing lldpneighbors table via ?action=clear")
	payload := lldpneighbors_clearGetThePayloadFromthePlan(ctx, &data)

	// NITRO exposes clear as POST ?action=clear. Use ActOnResource with the
	// case-sensitive "clear" verb (lower-case per the NITRO URL).
	err := r.client.ActOnResource(service.Lldpneighbors.Type(), &payload, "clear")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to clear lldpneighbors, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Cleared lldpneighbors table")

	// Fixed synthetic ID; lldpneighbors_clear is a transient action-only resource.
	data.Id = types.StringValue("lldpneighbors_clear")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Read is a no-op. clear is a one-shot action; NITRO has no query endpoint that
// reports clear-state, so there is nothing to reconcile.
func (r *LldpneighborsClearResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data LldpneighborsClearResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read is a no-op for lldpneighbors_clear; action-only resource with no persistent state")

	// Save prior state back unchanged
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Update is a no-op. lldpneighbors has no set/update endpoint and all config
// attributes are RequiresReplace; Terraform never invokes a real update.
func (r *LldpneighborsClearResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state LldpneighborsClearResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for lldpneighbors_clear; NITRO has no update endpoint and all attributes are RequiresReplace")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Delete is a no-op. clear is a one-shot side-effect action with no inverse
// NITRO API. Removing the resource from Terraform state is sufficient.
func (r *LldpneighborsClearResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete is a no-op for lldpneighbors_clear; NITRO has no inverse of the clear action")
}

// lldpneighbors_clearGetThePayloadFromthePlan builds the clear-action payload.
// clear takes NO args (CLI/NITRO both confirm an empty payload), so ifnum/nodeid
// are deliberately NOT copied here - they are GET/datasource filters (Pattern 15),
// and sending them would risk NITRO errorcode 278 "Invalid argument". The builder
// returns map[string]interface{}, the same payload type the original resource used.
func lldpneighbors_clearGetThePayloadFromthePlan(ctx context.Context, data *LldpneighborsClearResourceModel) map[string]interface{} {
	tflog.Debug(ctx, "In lldpneighbors_clearGetThePayloadFromthePlan Function")

	// clear body is empty: {"lldpneighbors":{}}.
	lldpneighbors := map[string]interface{}{}
	return lldpneighbors
}
