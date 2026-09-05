package cacheselector

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// CacheselectorDataSourceModel is the data-source-specific model, decoupled from
// CacheselectorResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes that the resource deliberately omits.
// Every non-key attribute is Computed; the Framework's per-attribute model <->
// schema reflection requires this model to have exactly the attributes the
// data-source schema declares, which is why it cannot reuse the resource model.
type CacheselectorDataSourceModel struct {
	Id           types.String `tfsdk:"id"`
	Selectorname types.String `tfsdk:"selectorname"` // Required lookup key

	// Existing read/write attributes, surfaced here as Computed outputs.
	Rule types.List `tfsdk:"rule"`

	// Read-only (GET-only) metadata from the NITRO doc read-only set
	// (zion73x_readonly/cacheselector.json). Never settable; populated from GET.
	Flags   types.Int64  `tfsdk:"flags"`
	Builtin types.List   `tfsdk:"builtin"`
	Feature types.String `tfsdk:"feature"`
}

func CacheselectorDataSourceSchema() schema.Schema {
	return schema.Schema{
		Description: "Data source to read cache selector configuration.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"rule": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "One or multiple PIXL expressions for evaluating an HTTP request or response.",
			},
			"selectorname": schema.StringAttribute{
				Required:    true,
				Description: "Name for the selector.  Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.",
			},

			// Read-only (GET-only) metadata surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"flags": schema.Int64Attribute{
				Computed:    true,
				Description: "Flags.",
			},
			"builtin": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "Flag to determine whether the cache selector is built-in.",
			},
			"feature": schema.StringAttribute{
				Computed:    true,
				Description: "The feature to be checked while applying this config.",
			},
		},
	}
}

// cacheselectorDataSourceSetAttrFromGet projects a NITRO cacheselector GET
// response onto the data-source model. Because a data source has no plan/apply
// reconciliation, attributes are simply filled from the GET (or left Null when
// the GET omits them). The shared utils.MapGet* helpers implement that projection.
func cacheselectorDataSourceSetAttrFromGet(ctx context.Context, data *CacheselectorDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In cacheselectorDataSourceSetAttrFromGet Function")

	if v, ok := g["selectorname"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Selectorname = types.StringValue(utils.AnyToString(v))
	}

	// Existing read/write attributes as read-back outputs.
	data.Rule = utils.MapGetStringList(g, "rule")

	// Read-only metadata.
	data.Flags = utils.MapGetInt64(g, "flags")
	data.Builtin = utils.MapGetStringList(g, "builtin")
	data.Feature = utils.MapGetString(g, "feature")
}
