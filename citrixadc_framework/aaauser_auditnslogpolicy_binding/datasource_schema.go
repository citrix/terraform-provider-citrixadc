package aaauser_auditnslogpolicy_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AaauserAuditnslogpolicyBindingDataSourceModel is the data-source-specific
// model, decoupled from AaauserAuditnslogpolicyBindingResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes AND the
// read-only attributes that the resource deliberately omits (acttype). Every
// non-key attribute is Computed.
type AaauserAuditnslogpolicyBindingDataSourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Policy                 types.String `tfsdk:"policy"`
	Priority               types.Int64  `tfsdk:"priority"`
	Type                   types.String `tfsdk:"type"`
	Username               types.String `tfsdk:"username"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/aaauser_auditnslogpolicy_binding.json). Never settable;
	// populated from GET, null when the appliance omits them.
	Acttype types.Int64 `tfsdk:"acttype"`
}

func AaauserAuditnslogpolicyBindingDataSourceSchema() schema.Schema {
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
			"policy": schema.StringAttribute{
				Required:    true,
				Description: "The policy Name.",
			},
			"priority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Integer specifying the priority of the policy.  A lower number indicates a higher priority. Policies are evaluated in the order of their priority numbers. Maximum value for default syntax policies is 2147483647 and for classic policies max priority is 64000.",
			},
			"type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Bindpoint to which the policy is bound.",
			},
			"username": schema.StringAttribute{
				Required:    true,
				Description: "User account to which to bind the policy.",
			},

			// Read-only (GET-only) attributes surfaced by the data source.
			"acttype": schema.Int64Attribute{
				Computed:    true,
				Description: "Action type of the bound policy. Read-only; returned by the appliance on a GET.",
			},
		},
	}
}

// aaauser_auditnslogpolicy_bindingDataSourceSetAttrFromGet projects a NITRO
// aaauser_auditnslogpolicy_binding GET response onto the data-source model.
// Attributes are simply filled from the GET (or left Null when the GET omits
// them) via the shared utils.MapGet* helpers.
func aaauser_auditnslogpolicy_bindingDataSourceSetAttrFromGet(ctx context.Context, data *AaauserAuditnslogpolicyBindingDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In aaauser_auditnslogpolicy_bindingDataSourceSetAttrFromGet Function")

	data.Gotopriorityexpression = utils.MapGetString(g, "gotopriorityexpression")
	data.Priority = utils.MapGetInt64(g, "priority")

	// policy / username are the required lookup keys; preserve the config value
	// when the GET omits them.
	if v, ok := g["policy"]; ok && v != nil {
		data.Policy = types.StringValue(utils.AnyToString(v))
	}
	if v, ok := g["username"]; ok && v != nil {
		data.Username = types.StringValue(utils.AnyToString(v))
	}

	// NITRO does not echo back "type" (bindpoint) in the binding GET response,
	// so preserve the configured value instead of nulling it.
	if v, ok := g["type"]; ok && v != nil {
		data.Type = types.StringValue(utils.AnyToString(v))
	} else if data.Type.IsUnknown() {
		data.Type = types.StringNull()
	}

	// Read-only attribute.
	data.Acttype = utils.MapGetInt64(g, "acttype")

	// Set ID (composite key: policy + username), matching the resource getter.
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("policy:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Policy.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("username:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Username.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))
}
