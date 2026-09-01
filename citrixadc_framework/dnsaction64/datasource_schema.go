package dnsaction64

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Dnsaction64DataSourceModel is the data-source-specific model, decoupled from
// Dnsaction64ResourceModel. A data source is a pure read surface, so it can
// expose the full GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes the resource deliberately omits.
type Dnsaction64DataSourceModel struct {
	Id          types.String `tfsdk:"id"`
	Actionname  types.String `tfsdk:"actionname"`
	Excluderule types.String `tfsdk:"excluderule"`
	Mappedrule  types.String `tfsdk:"mappedrule"`
	Prefix      types.String `tfsdk:"prefix"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/dnsaction64.json). Never settable; populated from GET.
	Builtin types.List   `tfsdk:"builtin"`
	Feature types.String `tfsdk:"feature"`
}

func Dnsaction64DataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"actionname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the dns64 action.",
			},
			"excluderule": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The expression to select the criteria for eliminating the corresponding ipv6 addresses from the response.",
			},
			"mappedrule": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The expression to select the criteria for ipv4 addresses to be used for synthesis.\n                      Only if the mappedrule is evaluated to true the corresponding ipv4 address is used for synthesis using respective prefix,\n                      otherwise the A RR is discarded",
			},
			"prefix": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The dns64 prefix to be used if the after evaluating the rules",
			},

			// Read-only (GET-only) attributes surfaced by the data source
			// (intentionally NOT modeled on the resource). All Computed.
			"builtin": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "Flag to determine whether dna64action is default or not. Possible values = MODIFIABLE, DELETABLE, IMMUTABLE, PARTITION_ALL.",
			},
			"feature": schema.StringAttribute{
				Computed:    true,
				Description: "The feature to be checked while applying this config.",
			},
		},
	}
}

// dnsaction64DataSourceSetAttrFromGet projects a NITRO dnsaction64 GET response
// onto the data-source model. Attributes are simply filled from the GET (or left
// Null when the GET omits them) via the shared utils.MapGet* helpers.
func dnsaction64DataSourceSetAttrFromGet(ctx context.Context, data *Dnsaction64DataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In dnsaction64DataSourceSetAttrFromGet Function")

	if v, ok := g["actionname"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Actionname = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Excluderule = utils.MapGetString(g, "excluderule")
	data.Mappedrule = utils.MapGetString(g, "mappedrule")
	data.Prefix = utils.MapGetString(g, "prefix")

	// Read-only attributes.
	data.Builtin = utils.MapGetStringList(g, "builtin")
	data.Feature = utils.MapGetString(g, "feature")
}
