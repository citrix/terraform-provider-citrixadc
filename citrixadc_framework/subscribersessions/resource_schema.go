package subscribersessions

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// SubscribersessionsResourceModel describes the resource data model. The
// resource is an action-only "clear": only the two clear selectors (ip, vlan)
// plus a synthetic id are modelled here. nodeid is a GET-only filter (Pattern
// 15) and is NOT part of the clear payload, so it is not on the resource model.
// The read-only telemetry lives on the datasource model, not here.
type SubscribersessionsResourceModel struct {
	Id   types.String `tfsdk:"id"`
	Ip   types.String `tfsdk:"ip"`
	Vlan types.Int64  `tfsdk:"vlan"`
}

func (r *SubscribersessionsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the subscribersessions resource.",
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

// subscribersessionsGetThePayloadFromthePlan builds the ?action=clear payload.
// Null selectors are omitted so a bare clear (flush the entire subscriber
// session DB) and a narrowed clear (ip/vlan) are both expressible. nodeid is a
// GET-only filter and is intentionally NOT included. Returns
// map[string]interface{} because the clear action needs only the two selectors.
func subscribersessionsGetThePayloadFromthePlan(ctx context.Context, data *SubscribersessionsResourceModel) map[string]interface{} {
	tflog.Debug(ctx, "In subscribersessionsGetThePayloadFromthePlan Function")

	payload := map[string]interface{}{}
	if !data.Ip.IsNull() && !data.Ip.IsUnknown() {
		payload["ip"] = data.Ip.ValueString()
	}
	if !data.Vlan.IsNull() && !data.Vlan.IsUnknown() {
		payload["vlan"] = data.Vlan.ValueInt64()
	}

	return payload
}
