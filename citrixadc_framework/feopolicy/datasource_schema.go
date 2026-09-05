package feopolicy

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// FeopolicyDataSourceModel is the data-source-specific model, decoupled from
// FeopolicyResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes the resource deliberately omits
// (builtin, feature, hits, undefhits). Every non-key attribute is Computed.
type FeopolicyDataSourceModel struct {
	Id     types.String `tfsdk:"id"`
	Action types.String `tfsdk:"action"`
	Name   types.String `tfsdk:"name"` // Required lookup key
	Rule   types.String `tfsdk:"rule"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/feopolicy.json). Never settable; populated from GET.
	Builtin   types.List   `tfsdk:"builtin"`
	Feature   types.String `tfsdk:"feature"`
	Hits      types.Int64  `tfsdk:"hits"`
	Undefhits types.Int64  `tfsdk:"undefhits"`
}

func FeopolicyDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"action": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The front end optimization action that has to be performed when the rule matches.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "The name of the front end optimization policy.",
			},
			"rule": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The rule associated with the front end optimization policy.",
			},

			// Read-only (GET-only) attributes surfaced by the data source
			// (intentionally NOT modeled on the resource). All Computed.
			"builtin": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "Flag to determine if the front end optimization policy is built-in or not.",
			},
			"feature": schema.StringAttribute{
				Computed:    true,
				Description: "The feature to be checked while applying this config.",
			},
			"hits": schema.Int64Attribute{
				Computed:    true,
				Description: "Total number of hits.",
			},
			"undefhits": schema.Int64Attribute{
				Computed:    true,
				Description: "Total number of undefined policy hits.",
			},
		},
	}
}

// feopolicyDataSourceSetAttrFromGet projects a NITRO feopolicy GET response onto
// the data-source model. Attributes are filled from the GET (or left Null when
// the GET omits them) via the shared utils.MapGet* helpers.
func feopolicyDataSourceSetAttrFromGet(ctx context.Context, data *FeopolicyDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In feopolicyDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Action = utils.MapGetString(g, "action")
	data.Rule = utils.MapGetString(g, "rule")

	// Read-only attributes.
	data.Builtin = utils.MapGetStringList(g, "builtin")
	data.Feature = utils.MapGetString(g, "feature")
	data.Hits = utils.MapGetInt64(g, "hits")
	data.Undefhits = utils.MapGetInt64(g, "undefhits")
}
