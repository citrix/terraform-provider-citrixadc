package subscribersessions

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
var _ resource.Resource = &SubscribersessionsClearResource{}
var _ resource.ResourceWithConfigure = (*SubscribersessionsClearResource)(nil)

func NewSubscribersessionsClearResource() resource.Resource {
	return &SubscribersessionsClearResource{}
}

// SubscribersessionsClearResource models the NITRO subscribersessions
// `?action=clear` action (POST). clear is a one-shot side-effect action with no
// GET-by-key endpoint and no inverse API, so Read/Update/Delete are no-ops. The
// clear payload carries the optional selectors ip and vlan; a bare clear (no
// selectors) flushes the entire subscriber session DB. nodeid is a GET-only
// cluster filter (Pattern 15) and is NOT part of the clear payload. The
// read-only telemetry (subscriptionidtype, ttl, flags, ...) is exposed by the
// citrixadc_subscribersessions datasource, not this action resource.
type SubscribersessionsClearResource struct {
	client *service.NitroClient
}

// SubscribersessionsClearResourceModel describes the resource data model. Only
// the two clear selectors (ip, vlan) plus a synthetic id are modelled here.
type SubscribersessionsClearResourceModel struct {
	Id   types.String `tfsdk:"id"`
	Ip   types.String `tfsdk:"ip"`
	Vlan types.Int64  `tfsdk:"vlan"`
}

func (r *SubscribersessionsClearResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_subscribersessions_clear"
}

func (r *SubscribersessionsClearResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SubscribersessionsClearResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the subscribersessions_clear resource.",
			},
			// clear selectors. Both Optional (a bare clear flushes the entire
			// subscriber session DB; ip/vlan narrows to a specific session).
			// Not Computed: Read is a no-op, so no server value ever resolves
			// them. RequiresReplace: any change re-fires the clear.
			"ip": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Subscriber IP Address.",
			},
			"vlan": schema.Int64Attribute{
				Optional: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "The vlan number on which the subscriber is located.",
			},
		},
	}
}

func (r *SubscribersessionsClearResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SubscribersessionsClearResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Clearing subscribersessions (action-only resource, firing ?action=clear)")
	payload := subscribersessions_clearGetThePayloadFromthePlan(ctx, &data)

	// subscribersessions has no add/set endpoint. Fire the clear action.
	// NITRO verb is case-sensitive: "clear" (lower-case per the NITRO URL).
	err := r.client.ActOnResource(service.Subscribersessions.Type(), payload, "clear")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to clear subscribersessions, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Cleared subscribersessions")

	// Synthetic ID for the action-only resource; keeps Read/Delete no-ops
	// addressable by Terraform.
	data.Id = types.StringValue("subscribersessions_clear")

	// Save data into Terraform state (no Read: transient table, no GET-by-key).
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Read is a no-op. clear is a one-shot action on a transient diagnostics table
// with no GET-by-key endpoint, so there is nothing to reconcile. Preserve prior
// state unchanged.
func (r *SubscribersessionsClearResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SubscribersessionsClearResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read is a no-op for subscribersessions_clear; NITRO has no query endpoint for clear state")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Update is a no-op. All attributes (ip, vlan) are RequiresReplace, so Terraform
// never routes a change through Update. There is no set endpoint.
func (r *SubscribersessionsClearResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state SubscribersessionsClearResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for subscribersessions_clear; NITRO has no update endpoint and all attributes are RequiresReplace")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Delete is a no-op. clear is a one-shot side-effect action; there is no inverse
// NITRO API. Removing the resource simply drops it from state.
func (r *SubscribersessionsClearResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete is a no-op for subscribersessions_clear; NITRO has no inverse of the clear action")
}

// subscribersessions_clearGetThePayloadFromthePlan builds the ?action=clear
// payload. Null selectors are omitted so a bare clear (flush the entire
// subscriber session DB) and a narrowed clear (ip/vlan) are both expressible.
// nodeid is a GET-only filter and is intentionally NOT included. Returns
// map[string]interface{} (the same payload type the original impl used) because
// the clear action needs only the two selectors.
func subscribersessions_clearGetThePayloadFromthePlan(ctx context.Context, data *SubscribersessionsClearResourceModel) map[string]interface{} {
	tflog.Debug(ctx, "In subscribersessions_clearGetThePayloadFromthePlan Function")

	payload := map[string]interface{}{}
	if !data.Ip.IsNull() && !data.Ip.IsUnknown() {
		payload["ip"] = data.Ip.ValueString()
	}
	if !data.Vlan.IsNull() && !data.Vlan.IsUnknown() {
		payload["vlan"] = data.Vlan.ValueInt64()
	}

	return payload
}
