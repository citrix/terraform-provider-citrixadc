package appfwurlencodedformcontenttype

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AppfwurlencodedformcontenttypeDataSourceModel is the data-source-specific
// model, decoupled from AppfwurlencodedformcontenttypeResourceModel. A data
// source is a pure read surface (Read only; no plan/apply lifecycle), so it can
// expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes the resource deliberately omits. The
// Framework's per-attribute model <-> schema reflection requires this model to
// have exactly the attributes the data-source schema declares.
type AppfwurlencodedformcontenttypeDataSourceModel struct {
	Id                             types.String `tfsdk:"id"`
	Isregex                        types.String `tfsdk:"isregex"`
	Urlencodedformcontenttypevalue types.String `tfsdk:"urlencodedformcontenttypevalue"` // Required lookup key

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/appfwurlencodedformcontenttype.json). Never settable;
	// populated from GET.
	Builtin types.List   `tfsdk:"builtin"`
	Feature types.String `tfsdk:"feature"`
}

func AppfwurlencodedformcontenttypeDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"isregex": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Is urlencoded form content type a regular expression?",
			},
			"urlencodedformcontenttypevalue": schema.StringAttribute{
				Required:    true,
				Description: "Content type to be classified as urlencoded form",
			},

			// Read-only (GET-only) attributes surfaced by the data source.
			"builtin": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "Flag to determine if urlencoded form contenttype is built-in or not.",
			},
			"feature": schema.StringAttribute{
				Computed:    true,
				Description: "The feature to be checked while applying this config.",
			},
		},
	}
}

// appfwurlencodedformcontenttypeDataSourceSetAttrFromGet projects a NITRO
// appfwurlencodedformcontenttype GET response onto the data-source model.
// Because a data source has no plan/apply reconciliation, attributes are simply
// filled from the GET (or left Null when the GET omits them). The shared
// utils.MapGet* helpers implement that projection.
func appfwurlencodedformcontenttypeDataSourceSetAttrFromGet(ctx context.Context, data *AppfwurlencodedformcontenttypeDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In appfwurlencodedformcontenttypeDataSourceSetAttrFromGet Function")

	if v, ok := g["urlencodedformcontenttypevalue"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Urlencodedformcontenttypevalue = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attribute as a read-back output (Null when the GET omits it).
	data.Isregex = utils.MapGetString(g, "isregex")

	// Read-only (GET-only) attributes.
	data.Builtin = utils.MapGetStringList(g, "builtin")
	data.Feature = utils.MapGetString(g, "feature")
}
