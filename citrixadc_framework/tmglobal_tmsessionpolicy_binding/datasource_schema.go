package tmglobal_tmsessionpolicy_binding

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// TmglobalTmsessionpolicyBindingDataSourceModel is the data-source-specific
// model, decoupled from TmglobalTmsessionpolicyBindingResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attribute the resource deliberately omits
// (bindpolicytype). The Framework's per-attribute model <-> schema reflection
// requires this model to have exactly the attributes the data-source schema
// declares.
type TmglobalTmsessionpolicyBindingDataSourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Feature                types.String `tfsdk:"feature"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Policyname             types.String `tfsdk:"policyname"`
	Priority               types.Int64  `tfsdk:"priority"`

	// Read-only (GET-only) attribute from the NITRO doc read-only set
	// (zion73x_readonly/tmglobal_tmsessionpolicy_binding.json). Never settable;
	// populated from GET.
	Bindpolicytype types.Int64 `tfsdk:"bindpolicytype"`
}

func TmglobalTmsessionpolicyBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"feature": schema.StringAttribute{
				Computed:    true,
				Description: "The feature to be checked while applying this config",
			},
			"gotopriorityexpression": schema.StringAttribute{
				Computed:    true,
				Description: "Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.",
			},
			"policyname": schema.StringAttribute{
				Required:    true,
				Description: "The name of the policy.",
			},
			"priority": schema.Int64Attribute{
				Computed:    true,
				Description: "The priority of the policy.",
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

// tmglobal_tmsessionpolicy_bindingDataSourceSetAttrFromGet projects a NITRO
// tmglobal_tmsessionpolicy_binding GET response onto the data-source model. A
// data source has no plan/apply reconciliation, so attributes are simply filled
// from the GET (or left Null when the GET omits them) via the shared
// utils.MapGet* helpers, and the ID is set.
func tmglobal_tmsessionpolicy_bindingDataSourceSetAttrFromGet(ctx context.Context, data *TmglobalTmsessionpolicyBindingDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In tmglobal_tmsessionpolicy_bindingDataSourceSetAttrFromGet Function")

	if v, ok := g["policyname"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Policyname = types.StringValue(utils.AnyToString(v))
	}

	data.Feature = utils.MapGetString(g, "feature")
	data.Gotopriorityexpression = utils.MapGetString(g, "gotopriorityexpression")
	data.Priority = utils.MapGetInt64(g, "priority")

	// Read-only (GET-only) attribute.
	data.Bindpolicytype = utils.MapGetInt64(g, "bindpolicytype")
}
