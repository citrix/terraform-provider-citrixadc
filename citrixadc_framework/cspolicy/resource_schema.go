package cspolicy

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/cs"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// CspolicyResourceModel describes the resource data model.
type CspolicyResourceModel struct {
	Id         types.String `tfsdk:"id"`
	Action     types.String `tfsdk:"action"`
	Logaction  types.String `tfsdk:"logaction"`
	Newname    types.String `tfsdk:"newname"`
	Policyname types.String `tfsdk:"policyname"`
	Rule       types.String `tfsdk:"rule"`
	// SDK v2 convenience attributes that combine cspolicy creation with a
	// csvserver policy binding. These are not NITRO cspolicy fields; they drive
	// BindResource/UnbindResource on the cs vserver, preserving SDK v2 behavior.
	Csvserver       types.String `tfsdk:"csvserver"`
	Targetlbvserver types.String `tfsdk:"targetlbvserver"`
	ForcenewIdSet   types.Set    `tfsdk:"forcenew_id_set"`
	Priority        types.Int64  `tfsdk:"priority"`
}

func (r *CspolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the cspolicy resource.",
			},
			"action": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Content switching action that names the target load balancing virtual server to which the traffic is switched.",
			},
			"logaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The log action associated with the content switching policy",
			},
			// newname is a rename-only attribute (NITRO ?action=rename). It is a pure
			// user input, never echoed by GET, so it must NOT be Computed (avoids
			// known-after-apply churn) and must NOT carry RequiresReplace (the change
			// is handled in-place in Update).
			"newname": schema.StringAttribute{
				Optional:    true,
				Description: "The new name of the content switching policy.",
			},
			// policyname is the primary key. SDK v2 marked it ForceNew -> RequiresReplace.
			"policyname": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the content switching policy. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Cannot be changed after a policy is created.\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my policy\" or 'my policy').",
			},
			// rule is Optional+Computed to match the SDK v2 contract (backward compat).
			"rule": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression, or name of a named expression, against which traffic is evaluated.\nThe following requirements apply only to the Citrix ADC CLI:\n*  If the expression includes one or more spaces, enclose the entire expression in double quotation marks.\n*  If the expression itself includes double quotation marks, escape the quotations by using the  character.\n*  Alternatively, you can use single quotation marks to enclose the rule, in which case you do not have to escape the double quotation marks.",
			},
			// csvserver: SDK v2 Optional + ForceNew -> RequiresReplace. When set, the
			// cspolicy is bound to this content switching vserver on create.
			"csvserver": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The content switching vserver to which the cspolicy should be bound.",
			},
			// targetlbvserver: SDK v2 Optional (updateable) - target lb vserver for the binding.
			"targetlbvserver": schema.StringAttribute{
				Optional:    true,
				Description: "The target load balancing vserver for the csvserver policy binding.",
			},
			// forcenew_id_set: SDK v2 TypeSet + ForceNew -> RequiresReplace. Not a NITRO
			// field; retained for backward compatibility. It only forces recreation.
			"forcenew_id_set": schema.SetAttribute{
				Optional:    true,
				ElementType: types.StringType,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.RequiresReplace(),
				},
				Description: "Auxiliary set attribute used to force recreation of the resource.",
			},
			// priority: SDK v2 TypeInt Optional (updateable) - priority for the binding.
			"priority": schema.Int64Attribute{
				Optional:    true,
				Description: "Priority for the csvserver policy binding.",
			},
		},
	}
}

// cspolicyGetThePayloadFromthePlan builds the NITRO cspolicy add/update payload.
// Only real NITRO cspolicy fields are included. newname is rename-only and is
// excluded here; csvserver/targetlbvserver/priority/forcenew_id_set are
// provider-side binding helpers and are not part of the cspolicy resource.
func cspolicyGetThePayloadFromthePlan(ctx context.Context, data *CspolicyResourceModel) cs.Cspolicy {
	tflog.Debug(ctx, "In cspolicyGetThePayloadFromthePlan Function")

	cspolicy := cs.Cspolicy{}
	if !data.Action.IsNull() && !data.Action.IsUnknown() {
		cspolicy.Action = data.Action.ValueString()
	}
	if !data.Logaction.IsNull() && !data.Logaction.IsUnknown() {
		cspolicy.Logaction = data.Logaction.ValueString()
	}
	if !data.Policyname.IsNull() && !data.Policyname.IsUnknown() {
		cspolicy.Policyname = data.Policyname.ValueString()
	}
	if !data.Rule.IsNull() && !data.Rule.IsUnknown() {
		cspolicy.Rule = data.Rule.ValueString()
	}
	// newname is intentionally excluded - it is rename-only.

	return cspolicy
}

func cspolicySetAttrFromGet(ctx context.Context, data *CspolicyResourceModel, getResponseData map[string]interface{}) *CspolicyResourceModel {
	tflog.Debug(ctx, "In cspolicySetAttrFromGet Function")

	// policyname is the configured key. Preserve the existing value and only adopt
	// the GET value on import (when the model has no policyname yet). This prevents
	// clobbering the user's configured key after a rename.
	if data.Policyname.IsNull() || data.Policyname.ValueString() == "" {
		if val, ok := getResponseData["policyname"]; ok && val != nil {
			data.Policyname = types.StringValue(val.(string))
		}
	}
	if val, ok := getResponseData["action"]; ok && val != nil {
		data.Action = types.StringValue(val.(string))
	} else {
		data.Action = types.StringNull()
	}
	if val, ok := getResponseData["logaction"]; ok && val != nil {
		data.Logaction = types.StringValue(val.(string))
	} else {
		data.Logaction = types.StringNull()
	}
	if val, ok := getResponseData["rule"]; ok && val != nil {
		data.Rule = types.StringValue(val.(string))
	} else {
		data.Rule = types.StringNull()
	}
	// newname, csvserver, targetlbvserver, priority and forcenew_id_set are not
	// returned by the single-resource GET; their values are preserved as-is. The
	// csvserver value is refreshed from the policy binding in readCspolicyFromApi.

	// ID tracks the live policy name from GET (this also makes rename work); fall
	// back to the configured policyname when GET does not return it.
	if val, ok := getResponseData["policyname"]; ok && val != nil {
		data.Id = types.StringValue(val.(string))
	} else {
		data.Id = types.StringValue(data.Policyname.ValueString())
	}

	return data
}
