package appfwglobal_auditsyslogpolicy_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AppfwglobalAuditsyslogpolicyBindingDataSourceModel is the data-source-specific
// model, decoupled from AppfwglobalAuditsyslogpolicyBindingResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attribute the resource deliberately omits
// (policytype). The Framework's per-attribute model <-> schema reflection
// requires this model to have exactly the attributes the data-source schema
// declares.
type AppfwglobalAuditsyslogpolicyBindingDataSourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Invoke                 types.Bool   `tfsdk:"invoke"`
	Labelname              types.String `tfsdk:"labelname"`
	Labeltype              types.String `tfsdk:"labeltype"`
	Policyname             types.String `tfsdk:"policyname"`
	Priority               types.Int64  `tfsdk:"priority"`
	State                  types.String `tfsdk:"state"`
	Type                   types.String `tfsdk:"type"`

	// Read-only (GET-only) attribute from the NITRO doc read-only set
	// (zion73x_readonly/appfwglobal_auditsyslogpolicy_binding.json). Never
	// settable; populated from GET.
	Policytype types.String `tfsdk:"policytype"`
}

func AppfwglobalAuditsyslogpolicyBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"gotopriorityexpression": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression or other value specifying the next policy to evaluate if the current policy evaluates to TRUE.  Specify one of the following values:\n* NEXT - Evaluate the policy with the next higher priority number.\n* END - End policy evaluation.\n* USE_INVOCATION_RESULT - Applicable if this policy invokes another policy label. If the final goto in the invoked policy label has a value of END, the evaluation stops. If the final goto is anything other than END, the current policy label performs a NEXT.\n* An expression that evaluates to a number.\nIf you specify an expression, the number to which it evaluates determines the next policy to evaluate, as follows:\n* If the expression evaluates to a higher numbered priority, the policy with that priority is evaluated next.\n* If the expression evaluates to the priority of the current policy, the policy with the next higher numbered priority is evaluated next.\n* If the expression evaluates to a number that is larger than the largest numbered priority, policy evaluation ends.\n\nAn UNDEF event is triggered if:\n* The expression is invalid.\n* The expression evaluates to a priority number that is smaller than the current policy's priority number.\n* The expression evaluates to a priority number that is between the current policy's priority number (say, 30) and the highest priority number (say, 100), but does not match any configured priority number (for example, the expression evaluates to the number 85). This example assumes that the priority number increments by 10 for every successive policy, and therefore a priority number of 85 does not exist in the policy label.",
			},
			"invoke": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "If the current policy evaluates to TRUE, terminate evaluation of policies bound to the current policy label, and then forward the request to the specified virtual server or evaluate the specified policy label.",
			},
			"labelname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the policy label to invoke if the current policy evaluates to TRUE, the invoke parameter is set, and Label Type is set to Policy Label.",
			},
			"labeltype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of policy label to invoke if the current policy evaluates to TRUE and the invoke parameter is set. Available settings function as follows:\n* reqvserver. Invoke the unnamed policy label associated with the specified request virtual server.\n* policylabel. Invoke the specified user-defined policy label.",
			},
			"policyname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the policy.",
			},
			"priority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The priority of the policy.",
			},
			"state": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable or disable the binding to activate or deactivate the policy. This is applicable to classic policies only.",
			},
			"type": schema.StringAttribute{
				Required:    true,
				Description: "Bind point to which to policy is bound.",
			},

			// Read-only (GET-only) attribute surfaced by the data source (this is
			// intentionally NOT modeled on the resource). Computed.
			"policytype": schema.StringAttribute{
				Computed:    true,
				Description: "Policy type. Possible values: [ Classic Policy, Advanced Policy ]",
			},
		},
	}
}

// appfwglobal_auditsyslogpolicy_bindingDataSourceSetAttrFromGet projects a NITRO
// appfwglobal_auditsyslogpolicy_binding GET response onto the data-source model.
// A data source has no plan/apply reconciliation, so attributes are simply
// filled from the GET (or left Null when the GET omits them) via the shared
// utils.MapGet* helpers. It faithfully copies every field (including 'state')
// and sets the composite ID.
func appfwglobal_auditsyslogpolicy_bindingDataSourceSetAttrFromGet(ctx context.Context, data *AppfwglobalAuditsyslogpolicyBindingDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In appfwglobal_auditsyslogpolicy_bindingDataSourceSetAttrFromGet Function")

	data.Gotopriorityexpression = utils.MapGetString(g, "gotopriorityexpression")
	data.Invoke = utils.MapGetBool(g, "invoke")
	data.Labelname = utils.MapGetString(g, "labelname")
	data.Labeltype = utils.MapGetString(g, "labeltype")
	data.Policyname = utils.MapGetString(g, "policyname")
	data.Priority = utils.MapGetInt64(g, "priority")
	data.State = utils.MapGetString(g, "state")
	// NOTE: 'type' is a Required config-side key that this binding's GET never
	// echoes (it surfaces 'bindpolicytype' instead). Preserve the configured
	// value already present in data rather than nulling it from the GET.

	// Read-only (GET-only) attributes.
	data.Policytype = utils.MapGetString(g, "policytype")

	// Set composite ID for the datasource.
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs.
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("policyname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Policyname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("type:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Type.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))
}
