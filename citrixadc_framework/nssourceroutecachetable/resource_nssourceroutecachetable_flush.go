package nssourceroutecachetable

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/adc-nitro-go/service"
)

// nssourceroutecachetableResourceType is the NITRO resource-type string.
// nssourceroutecachetable is not present in the adc-nitro-go service enum, so the
// type is declared here. It is also referenced by the datasource in this package.
const nssourceroutecachetableResourceType = "nssourceroutecachetable"

// NssourceroutecachetableFlushResource models the NITRO nssourceroutecachetable
// `?action=flush` action.
//
//   - NITRO exposes flush as POST /config/nssourceroutecachetable?action=flush.
//     There is no add/set/delete endpoint. The read side (get(all)/count) is
//     exposed separately via the citrixadc_nssourceroutecachetable datasource.
//   - Create performs the flush action. Read/Update/Delete are no-ops: flushing
//     the cache table has no persistent object to reconcile or remove.
//   - flush accepts NO attributes (empty payload); the model carries only the
//     synthetic id.
var _ resource.Resource = &NssourceroutecachetableFlushResource{}
var _ resource.ResourceWithConfigure = (*NssourceroutecachetableFlushResource)(nil)

func NewNssourceroutecachetableFlushResource() resource.Resource {
	return &NssourceroutecachetableFlushResource{}
}

// NssourceroutecachetableFlushResource defines the resource implementation.
type NssourceroutecachetableFlushResource struct {
	client *service.NitroClient
}

// NssourceroutecachetableFlushResourceModel describes the resource data model.
//
// flush is a ZERO-ATTRIBUTE, action-only resource. The model carries only the
// synthetic id.
type NssourceroutecachetableFlushResourceModel struct {
	Id types.String `tfsdk:"id"`
}

func (r *NssourceroutecachetableFlushResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nssourceroutecachetable_flush"
}

func (r *NssourceroutecachetableFlushResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NssourceroutecachetableFlushResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nssourceroutecachetable_flush resource. It is a synthetic value (nssourceroutecachetable_flush).",
			},
		},
	}
}

func (r *NssourceroutecachetableFlushResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NssourceroutecachetableFlushResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Flushing nssourceroutecachetable (action-only resource)")
	payload := nssourceroutecachetable_flushGetThePayloadFromthePlan(ctx, &data)

	// NITRO exposes flush as POST ?action=flush. Use ActOnResource with the
	// case-sensitive "flush" verb (lower-case per the NITRO URL).
	err := r.client.ActOnResource(nssourceroutecachetableResourceType, &payload, "flush")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to flush nssourceroutecachetable, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Flushed nssourceroutecachetable")

	// Synthetic ID for the action-only resource; keeps Read/Delete no-ops
	// addressable by Terraform.
	data.Id = types.StringValue("nssourceroutecachetable_flush")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NssourceroutecachetableFlushResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// flush is a one-shot action. NITRO has no GET endpoint that reports
	// flush-state, so Read is a pure preserve-state no-op.
	var data NssourceroutecachetableFlushResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read is a no-op for nssourceroutecachetable_flush; NITRO has no query endpoint for flush state")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NssourceroutecachetableFlushResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// NITRO has no update endpoint for flush; the resource has no read/write
	// attributes, so Terraform never invokes Update for a real change.
	var data, state NssourceroutecachetableFlushResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for nssourceroutecachetable_flush; it has no read/write attributes")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NssourceroutecachetableFlushResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// flush is a one-shot side-effect action. There is no inverse NITRO API
	// (no "un-flush"). Delete simply removes the resource from Terraform state.
	tflog.Debug(ctx, "Delete is a no-op for nssourceroutecachetable_flush; NITRO has no inverse of the flush action")
}

// nssourceroutecachetable_flushGetThePayloadFromthePlan builds the (empty) NITRO
// payload for the flush action. flush has no read/write attributes.
func nssourceroutecachetable_flushGetThePayloadFromthePlan(ctx context.Context, data *NssourceroutecachetableFlushResourceModel) map[string]interface{} {
	tflog.Debug(ctx, "In nssourceroutecachetable_flushGetThePayloadFromthePlan Function")
	nssourceroutecachetable := make(map[string]interface{})
	return nssourceroutecachetable
}
