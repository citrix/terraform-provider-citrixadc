package appfwjsoncontenttype

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AppfwjsoncontenttypeDataSourceModel is the data-source-specific model,
// decoupled from AppfwjsoncontenttypeResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes the resource deliberately omits. Every
// non-key attribute is Computed; the Framework's per-attribute model <-> schema
// reflection requires this model to have exactly the attributes the data-source
// schema declares, which is why it cannot reuse the resource model.
type AppfwjsoncontenttypeDataSourceModel struct {
	Id                   types.String `tfsdk:"id"`
	Isregex              types.String `tfsdk:"isregex"`
	Jsoncontenttypevalue types.String `tfsdk:"jsoncontenttypevalue"` // Required lookup key

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/appfwjsoncontenttype.json). Never settable; populated
	// from GET.
	Builtin types.List   `tfsdk:"builtin"`
	Feature types.String `tfsdk:"feature"`
}

func AppfwjsoncontenttypeDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"isregex": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Is json content type a regular expression?",
			},
			"jsoncontenttypevalue": schema.StringAttribute{
				Required:    true,
				Description: "Content type to be classified as JSON",
			},

			// Read-only (GET-only) attributes surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"builtin": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "Flag to determine if jsoncontenttype is built-in or not.",
			},
			"feature": schema.StringAttribute{
				Computed:    true,
				Description: "The feature to be checked while applying this config.",
			},
		},
	}
}

// appfwjsoncontenttypeDataSourceSetAttrFromGet projects a NITRO
// appfwjsoncontenttype GET response onto the data-source model. Because a data
// source has no plan/apply reconciliation, attributes are simply filled from the
// GET (or left Null when the GET omits them). The shared utils.MapGet* helpers
// implement that projection.
func appfwjsoncontenttypeDataSourceSetAttrFromGet(ctx context.Context, data *AppfwjsoncontenttypeDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In appfwjsoncontenttypeDataSourceSetAttrFromGet Function")

	if v, ok := g["jsoncontenttypevalue"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Jsoncontenttypevalue = types.StringValue(utils.AnyToString(v))
	}

	data.Isregex = utils.MapGetString(g, "isregex")

	// Read-only metadata.
	data.Builtin = utils.MapGetStringList(g, "builtin")
	data.Feature = utils.MapGetString(g, "feature")
}
