package appqoepolicy

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AppqoepolicyDataSourceModel is the data-source-specific model, decoupled from
// AppqoepolicyResourceModel. A data source is a pure read surface, so it can
// expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes the resource deliberately omits.
type AppqoepolicyDataSourceModel struct {
	Id     types.String `tfsdk:"id"`
	Action types.String `tfsdk:"action"`
	Name   types.String `tfsdk:"name"`
	Rule   types.String `tfsdk:"rule"`

	// Read-only (GET-only) attribute from the NITRO doc read-only set
	// (zion73x_readonly/appqoepolicy.json). Never settable; populated from GET.
	Hits types.Int64 `tfsdk:"hits"`
}

func AppqoepolicyDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"action": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Configured AppQoE action to trigger",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "0",
			},
			"rule": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression or name of a named expression, against which the request is evaluated. The policy is applied if the rule evaluates to true.",
			},

			// Read-only (GET-only) attribute surfaced by the data source.
			"hits": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of hits.",
			},
		},
	}
}

// appqoepolicyDataSourceSetAttrFromGet projects a NITRO appqoepolicy GET
// response onto the data-source model. Attributes are simply filled from the GET
// (or left Null when the GET omits them) via the shared utils.MapGet* helpers.
func appqoepolicyDataSourceSetAttrFromGet(ctx context.Context, data *AppqoepolicyDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In appqoepolicyDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Action = utils.MapGetString(g, "action")
	data.Rule = utils.MapGetString(g, "rule")

	// Read-only (GET-only) attribute.
	data.Hits = utils.MapGetInt64(g, "hits")
}
