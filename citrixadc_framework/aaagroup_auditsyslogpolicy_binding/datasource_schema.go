package aaagroup_auditsyslogpolicy_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AaagroupAuditsyslogpolicyBindingDataSourceModel is the data-source-specific
// model, decoupled from the resource model so the data source can expose the
// FULL GET projection: the existing lookup/config attributes (as Computed
// outputs) PLUS the read-only attributes the resource intentionally omits.
type AaagroupAuditsyslogpolicyBindingDataSourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Groupname              types.String `tfsdk:"groupname"` // Required lookup key
	Policy                 types.String `tfsdk:"policy"`    // Required lookup key
	Priority               types.Int64  `tfsdk:"priority"`
	Type                   types.String `tfsdk:"type"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/aaagroup_auditsyslogpolicy_binding.json).
	Acttype types.Int64 `tfsdk:"acttype"`
}

func AaagroupAuditsyslogpolicyBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"gotopriorityexpression": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression or other value specifying the next policy to evaluate if the current policy evaluates to TRUE.  Specify one of the following values:\n* NEXT - Evaluate the policy with the next higher priority number.\n* END - End policy evaluation.\n* USE_INVOCATION_RESULT - Applicable if this policy invokes another policy label. If the final goto in the invoked policy label has a value of END, the evaluation stops. If the final goto is anything other than END, the current policy label performs a NEXT.\n* An expression that evaluates to a number.\nIf you specify an expression, the number to which it evaluates determines the next policy to evaluate, as follows:\n*  If the expression evaluates to a higher numbered priority, the policy with that priority is evaluated next.\n* If the expression evaluates to the priority of the current policy, the policy with the next higher numbered priority is evaluated next.\n* If the expression evaluates to a number that is larger than the largest numbered priority, policy evaluation ends.\nAn UNDEF event is triggered if:\n* The expression is invalid.\n* The expression evaluates to a priority number that is numerically lower than the current policy's priority.\n* The expression evaluates to a priority number that is between the current policy's priority number (say, 30) and the highest priority number (say, 100), but does not match any configured priority number (for example, the expression evaluates to the number 85). This example assumes that the priority number increments by 10 for every successive policy, and therefore a priority number of 85 does not exist in the policy label.",
			},
			"groupname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the group that you are binding.",
			},
			"policy": schema.StringAttribute{
				Required:    true,
				Description: "The policy name.",
			},
			"priority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Integer specifying the priority of the policy. A lower number indicates a higher priority. Policies are evaluated in the order of their priority numbers. Maximum value for default syntax policies is 2147483647 and for classic policies is 64000.",
			},
			"type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Bindpoint to which the policy is bound.",
			},

			// Read-only (GET-only) attributes surfaced by the data source.
			"acttype": schema.Int64Attribute{
				Computed:    true,
				Description: "Bound policy action type.",
			},
		},
	}
}

// aaagroup_auditsyslogpolicy_bindingDataSourceSetAttrFromGet projects a NITRO GET
// response onto the data-source model. Attributes are simply filled from the GET
// (or left Null when the GET omits them) via the shared utils.MapGet* helpers.
func aaagroup_auditsyslogpolicy_bindingDataSourceSetAttrFromGet(ctx context.Context, data *AaagroupAuditsyslogpolicyBindingDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In aaagroup_auditsyslogpolicy_bindingDataSourceSetAttrFromGet Function")

	data.Gotopriorityexpression = utils.MapGetString(g, "gotopriorityexpression")
	// groupname is the parent lookup key; preserve the config value when the
	// binding GET response does not echo it back.
	if v := utils.MapGetString(g, "groupname"); !v.IsNull() {
		data.Groupname = v
	}
	data.Policy = utils.MapGetString(g, "policy")
	data.Priority = utils.MapGetInt64(g, "priority")
	data.Type = utils.MapGetString(g, "type")

	// Read-only attributes.
	data.Acttype = utils.MapGetInt64(g, "acttype")

	// Set ID: comma-separated key:UrlEncode(value) pairs.
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("groupname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Groupname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("policy:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Policy.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))
}
