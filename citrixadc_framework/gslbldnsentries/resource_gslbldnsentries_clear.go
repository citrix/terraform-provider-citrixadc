package gslbldnsentries

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &GslbldnsentriesClearResource{}
var _ resource.ResourceWithConfigure = (*GslbldnsentriesClearResource)(nil)

func NewGslbldnsentriesClearResource() resource.Resource {
	return &GslbldnsentriesClearResource{}
}

// GslbldnsentriesClearResource models the NITRO gslbldnsentries `?action=clear`
// action. clear is a one-shot side-effect action (POST ?action=clear) with no
// add/set/delete endpoint and no GET-backed object, so Read/Update/Delete are
// no-ops. The read/count side of gslbldnsentries lives in the separate
// citrixadc_gslbldnsentries datasource.
type GslbldnsentriesClearResource struct {
	client *service.NitroClient
}

// GslbldnsentriesClearResourceModel describes the resource data model. nodeid is
// retained in the schema per the per-action attribute guidance, but it is a
// GET-only cluster filter (Pattern 15) and is intentionally excluded from the
// clear payload; the CLI rejects `-nodeid` for the clear verb.
type GslbldnsentriesClearResourceModel struct {
	Id     types.String `tfsdk:"id"`
	Nodeid types.Int64  `tfsdk:"nodeid"`
}

func (r *GslbldnsentriesClearResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gslbldnsentries_clear"
}

func (r *GslbldnsentriesClearResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *GslbldnsentriesClearResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the gslbldnsentries_clear resource.",
			},
			// nodeid is a GET-only cluster filter. clear is an action with no
			// GET-backed object (Read is a no-op), so nodeid must NOT be Computed
			// or Terraform reports an unknown value after apply.
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

func (r *GslbldnsentriesClearResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data GslbldnsentriesClearResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Clearing gslbldnsentries (action-only resource)")
	payload := gslbldnsentries_clearGetThePayloadFromthePlan(ctx, &data)

	// NITRO exposes clear as POST ?action=clear (Pattern 1). Use ActOnResource
	// with the case-sensitive "clear" verb (lower-case per the NITRO URL). clear
	// takes no arguments, so the payload is an empty map (nodeid is a GET-only
	// filter, Pattern 15).
	err := r.client.ActOnResource(service.Gslbldnsentries.Type(), payload, "clear")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to clear gslbldnsentries, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Cleared gslbldnsentries")

	// Synthetic ID for the action-only resource; keeps Read/Delete no-ops
	// addressable by Terraform.
	data.Id = types.StringValue("gslbldnsentries_clear")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GslbldnsentriesClearResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// clear is a one-shot action. NITRO has no GET endpoint that reports
	// clear-state, so Read is a pure preserve-state no-op.
	var data GslbldnsentriesClearResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read is a no-op for gslbldnsentries_clear; NITRO has no query endpoint for clear state")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GslbldnsentriesClearResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// NITRO has no update endpoint for clear; every schema attribute is
	// RequiresReplace, so Terraform never invokes Update for a real change.
	var data, state GslbldnsentriesClearResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for gslbldnsentries_clear; NITRO has no update endpoint and all attributes are RequiresReplace")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GslbldnsentriesClearResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// clear is a one-shot side-effect action. There is no inverse NITRO API
	// (no "un-clear"). Delete simply removes the resource from Terraform state.
	tflog.Debug(ctx, "Delete is a no-op for gslbldnsentries_clear; NITRO has no inverse of the clear action")
}

// gslbldnsentries_clearGetThePayloadFromthePlan builds the body for the clear
// action. clear takes no arguments, so the body is always empty. nodeid is a
// GET-only cluster filter and is intentionally excluded from the clear payload
// (Pattern 15). Returns map[string]interface{}, matching the pre-split impl.
func gslbldnsentries_clearGetThePayloadFromthePlan(ctx context.Context, data *GslbldnsentriesClearResourceModel) map[string]interface{} {
	tflog.Debug(ctx, "In gslbldnsentries_clearGetThePayloadFromthePlan Function")

	gslbldnsentries := map[string]interface{}{}

	return gslbldnsentries
}
