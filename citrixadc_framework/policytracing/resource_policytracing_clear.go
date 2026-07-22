package policytracing

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &PolicytracingClearResource{}
var _ resource.ResourceWithConfigure = (*PolicytracingClearResource)(nil)

func NewPolicytracingClearResource() resource.Resource {
	return &PolicytracingClearResource{}
}

// PolicytracingClearResource models the NITRO policytracing `?action=clear`
// action. clear is a one-shot side-effect action ("remove the collected policy
// tracing data from memory") with no GET endpoint and no inverse API, so
// Read/Update/Delete are no-ops. Per the NITRO doc and the live CLI
// (`clear policy tracing`), clear takes NO arguments -- the request payload is
// empty ({"policytracing":{}}). The resource schema therefore exposes only the
// synthetic id (Pattern 13).
type PolicytracingClearResource struct {
	client *service.NitroClient
}

// PolicytracingClearResourceModel describes the resource data model.
//
// clear takes no arguments, so the model declares ONLY the synthetic id -- the
// Plugin Framework requires the model struct's tfsdk fields to exactly match the
// schema attributes.
type PolicytracingClearResourceModel struct {
	Id types.String `tfsdk:"id"`
}

func (r *PolicytracingClearResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_policytracing_clear"
}

func (r *PolicytracingClearResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *PolicytracingClearResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	// The clear action takes no arguments, so the resource schema is minimal:
	// just the synthetic id (Pattern 13: only id is Computed, no other Computed
	// attrs since Read is a no-op).
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the policytracing_clear resource.",
			},
		},
	}
}

func (r *PolicytracingClearResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data PolicytracingClearResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Clearing policytracing (action-only resource)")
	payload := policytracing_clearGetThePayloadFromthePlan(ctx, &data)

	// clear is a POST ?action=clear action (Pattern 1). There is no add endpoint;
	// AddResource/UpdateUnnamedResource would target a nonexistent verb. The clear
	// body is empty per the NITRO doc ({"policytracing":{}}).
	err := r.client.ActOnResource(service.Policytracing.Type(), payload, "clear")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to clear policytracing, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Cleared policytracing")

	// Set synthetic constant ID once, here in Create (Pattern 6). clear is a
	// one-shot reset action; the resource is not a queryable managed object, so
	// this ID is purely a Terraform state handle, not a NITRO lookup key.
	data.Id = types.StringValue("policytracing_clear")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PolicytracingClearResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// clear is a one-shot action. NITRO has no GET endpoint that reports
	// clear-state, so Read is a pure preserve-state no-op (Pattern 13).
	var data PolicytracingClearResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read is a no-op for policytracing_clear; NITRO has no query endpoint for clear state")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PolicytracingClearResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// NITRO has no update endpoint for clear; the schema has no writable
	// attributes, so Terraform never invokes Update for a real change (Pattern 5).
	var data, state PolicytracingClearResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for policytracing_clear; NITRO has no update endpoint")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PolicytracingClearResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// clear is a one-shot side-effect action. There is no inverse NITRO API
	// (no "un-clear"). Delete simply removes the resource from Terraform state.
	tflog.Debug(ctx, "Delete is a no-op for policytracing_clear; NITRO has no inverse of the clear action")
}

// policytracing_clearGetThePayloadFromthePlan builds the body for the clear
// action. clear takes no arguments (NITRO doc shows an empty body
// {"policytracing":{}}), so the body is always empty. Returns
// map[string]interface{} matching the original implementation (no vendored
// struct fields are populated for this action).
func policytracing_clearGetThePayloadFromthePlan(ctx context.Context, data *PolicytracingClearResourceModel) map[string]interface{} {
	tflog.Debug(ctx, "In policytracing_clearGetThePayloadFromthePlan Function")

	policytracing := map[string]interface{}{}

	return policytracing
}
