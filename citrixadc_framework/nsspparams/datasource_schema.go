package nsspparams

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// NsspparamsDataSourceModel is the data-source-specific model, decoupled from
// NsspparamsResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes that the resource deliberately omits.
// Every non-key attribute is Computed; the Framework's per-attribute model
// <-> schema reflection requires this model to have exactly the attributes the
// data-source schema declares, which is why it cannot reuse the resource model.
type NsspparamsDataSourceModel struct {
	Id            types.String `tfsdk:"id"`
	Basethreshold types.Int64  `tfsdk:"basethreshold"`
	Throttle      types.String `tfsdk:"throttle"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/nsspparams.json). Never settable; populated from GET.
	Table0  types.List   `tfsdk:"table0"`
	Builtin types.List   `tfsdk:"builtin"`
	Feature types.String `tfsdk:"feature"`
}

func NsspparamsDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"basethreshold": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of server connections that can be opened before surge protection is activated.",
			},
			"throttle": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Rate at which the system opens connections to the server.",
			},

			// Read-only (GET-only) attributes surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"table0": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "Table.",
			},
			"builtin": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "Flag to determine if sp param is built-in or not.",
			},
			"feature": schema.StringAttribute{
				Computed:    true,
				Description: "The feature to be checked while applying this config.",
			},
		},
	}
}

// nsspparamsDataSourceSetAttrFromGet projects a NITRO nsspparams GET response
// onto the data-source model. Because a data source has no plan/apply
// reconciliation, attributes are simply filled from the GET (or left Null when
// the GET omits them). The shared utils.MapGet* helpers implement that
// projection.
func nsspparamsDataSourceSetAttrFromGet(ctx context.Context, data *NsspparamsDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In nsspparamsDataSourceSetAttrFromGet Function")

	// nsspparams is a singleton with no lookup key; use a static ID.
	data.Id = types.StringValue("nsspparams-config")

	// Read/write attributes as read-back outputs.
	data.Basethreshold = utils.MapGetInt64(g, "basethreshold")
	data.Throttle = utils.MapGetString(g, "throttle")

	// Read-only attributes.
	data.Table0 = utils.MapGetStringList(g, "table0")
	data.Builtin = utils.MapGetStringList(g, "builtin")
	data.Feature = utils.MapGetString(g, "feature")
}
