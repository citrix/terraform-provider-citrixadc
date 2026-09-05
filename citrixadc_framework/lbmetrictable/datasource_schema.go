package lbmetrictable

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// LbmetrictableDataSourceModel is the data-source-specific model, decoupled from
// LbmetrictableResourceModel. A data source is a pure read surface (Read only),
// so it exposes the FULL GET projection: the lookup key AND the read-only
// attributes the resource deliberately omits (metrictype, builtin, feature).
type LbmetrictableDataSourceModel struct {
	Id          types.String `tfsdk:"id"`
	Metrictable types.String `tfsdk:"metrictable"` // Required lookup key

	// Read-only (GET-only) metadata from the NITRO read-only set
	// (zion73x_readonly/lbmetrictable.json). Never settable; populated from GET.
	Metrictype types.String `tfsdk:"metrictype"`
	Builtin    types.List   `tfsdk:"builtin"`
	Feature    types.String `tfsdk:"feature"`
}

func LbmetrictableDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"metrictable": schema.StringAttribute{
				Required:    true,
				Description: "Name for the metric table. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.\n\nCLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my metrictable\" or 'my metrictable').",
			},

			// Read-only (GET-only) metadata surfaced by the data source
			// (intentionally NOT modeled on the resource). All Computed.
			"metrictype": schema.StringAttribute{
				Computed:    true,
				Description: "Indication if it is a configured or internal metric table. Possible values: [ INTERNAL, CONFIGURED ].",
			},
			"builtin": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "Flag to determine whether the metric table is built-in. Possible values: [ MODIFIABLE, DELETABLE, IMMUTABLE, PARTITION_ALL ]. A list of strings.",
			},
			"feature": schema.StringAttribute{
				Computed:    true,
				Description: "The feature to be checked while applying this configuration.",
			},
		},
	}
}

// lbmetrictableDataSourceSetAttrFromGet projects a NITRO lbmetrictable GET
// response onto the data-source model. Attributes are simply filled from the GET
// (or left Null when the GET omits them) via the shared utils.MapGet* helpers.
func lbmetrictableDataSourceSetAttrFromGet(ctx context.Context, data *LbmetrictableDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In lbmetrictableDataSourceSetAttrFromGet Function")

	if v, ok := g["metrictable"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Metrictable = types.StringValue(utils.AnyToString(v))
	} else {
		data.Id = data.Metrictable
	}

	// Read-only metadata.
	data.Metrictype = utils.MapGetString(g, "metrictype")
	data.Builtin = utils.MapGetStringList(g, "builtin")
	data.Feature = utils.MapGetString(g, "feature")
}
