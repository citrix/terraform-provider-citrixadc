package tmglobal_tmtrafficpolicy_binding

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// TmglobalTmtrafficpolicyBindingDataSourceModel is the data-source-specific
// model, decoupled from TmglobalTmtrafficpolicyBindingResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attribute the resource deliberately omits
// (bindpolicytype). The Framework's per-attribute model <-> schema reflection
// requires this model to have exactly the attributes the data-source schema
// declares.
type TmglobalTmtrafficpolicyBindingDataSourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Globalbindtype         types.String `tfsdk:"globalbindtype"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Policyname             types.String `tfsdk:"policyname"`
	Priority               types.Int64  `tfsdk:"priority"`
	Type                   types.String `tfsdk:"type"`

	// Read-only (GET-only) attribute from the NITRO doc read-only set
	// (zion73x_readonly/tmglobal_tmtrafficpolicy_binding.json). Never settable;
	// populated from GET.
	Bindpolicytype types.Int64 `tfsdk:"bindpolicytype"`
}

func TmglobalTmtrafficpolicyBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"globalbindtype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "0",
			},
			"gotopriorityexpression": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Applicable only to advance tmsession policy. Expression or other value specifying the next policy to be evaluated if the current policy evaluates to TRUE.  Specify one of the following values:\n* NEXT - Evaluate the policy with the next higher priority number.\n* END - End policy evaluation.\n* An expression that evaluates to a number.\nIf you specify an expression, the number to which it evaluates determines the next policy to evaluate, as follows:\n* If the expression evaluates to a higher numbered priority, the policy with that priority is evaluated next.\n* If the expression evaluates to the priority of the current policy, the policy with the next higher numbered priority is evaluated next.\n* If the expression evaluates to a priority number that is numerically higher than the highest numbered priority, policy evaluation ends.\nAn UNDEF event is triggered if:\n* The expression is invalid.\n* The expression evaluates to a priority number that is numerically lower than the current policy's priority.\n* The expression evaluates to a priority number that is between the current policy's priority number (say, 30) and the highest priority number (say, 100), but does not match any configured priority number (for example, the expression evaluates to the number 85). This example assumes that the priority number increments by 10 for every successive policy, and therefore a priority number of 85 does not exist in the policy label.",
			},
			"policyname": schema.StringAttribute{
				Required:    true,
				Description: "The name of the policy.",
			},
			"priority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The priority of the policy.",
			},
			"type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Bindpoint to which the policy is bound",
			},

			// Read-only (GET-only) attribute surfaced by the data source (this is
			// intentionally NOT modeled on the resource). Computed.
			"bindpolicytype": schema.Int64Attribute{
				Computed:    true,
				Description: "Bound policy type.",
			},
		},
	}
}

// tmglobal_tmtrafficpolicy_bindingDataSourceSetAttrFromGet projects a NITRO
// tmglobal_tmtrafficpolicy_binding GET response onto the data-source model. A
// data source has no plan/apply reconciliation, so attributes are simply filled
// from the GET (or left Null when the GET omits them) via the shared
// utils.MapGet* helpers, and the ID is set.
func tmglobal_tmtrafficpolicy_bindingDataSourceSetAttrFromGet(ctx context.Context, data *TmglobalTmtrafficpolicyBindingDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In tmglobal_tmtrafficpolicy_bindingDataSourceSetAttrFromGet Function")

	if v, ok := g["policyname"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Policyname = types.StringValue(utils.AnyToString(v))
	}

	data.Globalbindtype = utils.MapGetString(g, "globalbindtype")
	data.Gotopriorityexpression = utils.MapGetString(g, "gotopriorityexpression")
	data.Priority = utils.MapGetInt64(g, "priority")
	data.Type = utils.MapGetString(g, "type")

	// Read-only (GET-only) attribute.
	data.Bindpolicytype = utils.MapGetInt64(g, "bindpolicytype")
}
