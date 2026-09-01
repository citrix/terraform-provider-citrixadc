package appfwgrpccontenttype

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AppfwgrpccontenttypeDataSourceModel is the data-source-specific model,
// decoupled from AppfwgrpccontenttypeResourceModel. A data source is a pure read
// surface, so it exposes the existing lookup/config attributes (as Computed
// outputs) PLUS the read-only (GET-only) attributes the resource omits.
type AppfwgrpccontenttypeDataSourceModel struct {
	Id                   types.String `tfsdk:"id"`
	Grpccontenttypevalue types.String `tfsdk:"grpccontenttypevalue"` // Required lookup key
	Isregex              types.String `tfsdk:"isregex"`

	// Read-only (GET-only) attributes from the NITRO read-only set
	// (zion73x_readonly/appfwgrpccontenttype.json). Populated from GET; never settable.
	Builtin types.List   `tfsdk:"builtin"`
	Feature types.String `tfsdk:"feature"`
}

func AppfwgrpccontenttypeDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"grpccontenttypevalue": schema.StringAttribute{
				Required:    true,
				Description: "Content type to be classified as gRPC",
			},
			"isregex": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Is gRPC content type a regular expression?",
			},

			// Read-only (GET-only) attributes surfaced by the data source.
			"builtin": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "Flag to determine if gRPC content type is built-in or not. Possible values = MODIFIABLE, DELETABLE, IMMUTABLE, PARTITION_ALL.",
			},
			"feature": schema.StringAttribute{
				Computed:    true,
				Description: "The feature to be checked while applying this config.",
			},
		},
	}
}

// appfwgrpccontenttypeDataSourceSetAttrFromGet projects a NITRO
// appfwgrpccontenttype GET response onto the data-source model. Every attribute
// is filled from the GET (or left Null when the GET omits it); id mirrors the
// returned grpccontenttypevalue key.
func appfwgrpccontenttypeDataSourceSetAttrFromGet(ctx context.Context, data *AppfwgrpccontenttypeDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In appfwgrpccontenttypeDataSourceSetAttrFromGet Function")

	if v, ok := g["grpccontenttypevalue"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Grpccontenttypevalue = types.StringValue(utils.AnyToString(v))
	}

	data.Isregex = utils.MapGetString(g, "isregex")

	// Read-only (GET-only) attributes.
	data.Builtin = utils.MapGetStringList(g, "builtin")
	data.Feature = utils.MapGetString(g, "feature")
}
