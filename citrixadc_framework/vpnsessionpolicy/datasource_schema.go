package vpnsessionpolicy

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// VpnsessionpolicyDataSourceModel is the data-source-specific model, decoupled
// from VpnsessionpolicyResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only metadata attributes that the resource deliberately
// omits (builtin, feature, expressiontype, hits, ...). Every non-key attribute
// is Computed.
type VpnsessionpolicyDataSourceModel struct {
	Id     types.String `tfsdk:"id"`
	Action types.String `tfsdk:"action"`
	Name   types.String `tfsdk:"name"` // Required lookup key
	Rule   types.String `tfsdk:"rule"`

	// Read-only (GET-only) metadata from the NITRO doc read-only set
	// (zion73x_readonly/vpnsessionpolicy.json). Never settable; populated from GET.
	Builtin        types.List   `tfsdk:"builtin"`
	Feature        types.String `tfsdk:"feature"`
	Expressiontype types.String `tfsdk:"expressiontype"`
	Hits           types.Int64  `tfsdk:"hits"`
}

func VpnsessionpolicyDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"action": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Action to be applied by the new session policy if the rule criteria are met.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the new session policy that is applied after the user logs on to Citrix Gateway.",
			},
			"rule": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression, or name of a named expression, specifying the traffic that matches the policy.\n\nThe following requirements apply only to the Citrix ADC CLI:\n* If the expression includes one or more spaces, enclose the entire expression in double quotation marks.\n* If the expression itself includes double quotation marks, escape the quotations by using the \\ character.\n* Alternatively, you can use single quotation marks to enclose the rule, in which case you do not have to escape the double quotation marks.",
			},

			// Read-only (GET-only) metadata surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"builtin": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "Indicates that a variable is a built-in (SYSTEM INTERNAL) type.",
			},
			"feature": schema.StringAttribute{
				Computed:    true,
				Description: "The feature to be checked while applying this config.",
			},
			"expressiontype": schema.StringAttribute{
				Computed:    true,
				Description: "Type of policy (Classic/Advanced).",
			},
			"hits": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of hits.",
			},
		},
	}
}

// vpnsessionpolicyDataSourceSetAttrFromGet projects a NITRO vpnsessionpolicy GET
// response onto the data-source model. Because a data source has no plan/apply
// reconciliation, attributes are simply filled from the GET (or left Null when
// the GET omits them). The shared utils.MapGet* helpers implement that projection.
func vpnsessionpolicyDataSourceSetAttrFromGet(ctx context.Context, data *VpnsessionpolicyDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In vpnsessionpolicyDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Action = utils.MapGetString(g, "action")
	data.Rule = utils.MapGetString(g, "rule")

	// Read-only metadata.
	data.Builtin = utils.MapGetStringList(g, "builtin")
	data.Feature = utils.MapGetString(g, "feature")
	data.Expressiontype = utils.MapGetString(g, "expressiontype")
	data.Hits = utils.MapGetInt64(g, "hits")
}
