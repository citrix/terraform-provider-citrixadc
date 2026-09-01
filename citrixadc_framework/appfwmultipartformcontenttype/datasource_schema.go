package appfwmultipartformcontenttype

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AppfwmultipartformcontenttypeDataSourceModel is the data-source-specific
// model, decoupled from AppfwmultipartformcontenttypeResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes the resource deliberately omits. Every
// non-key attribute is Computed; the Framework's per-attribute model <-> schema
// reflection requires this model to have exactly the attributes the data-source
// schema declares, which is why it cannot reuse the resource model.
type AppfwmultipartformcontenttypeDataSourceModel struct {
	Id                            types.String `tfsdk:"id"`
	Isregex                       types.String `tfsdk:"isregex"`
	Multipartformcontenttypevalue types.String `tfsdk:"multipartformcontenttypevalue"` // Required lookup key

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/appfwmultipartformcontenttype.json). Never settable;
	// populated from GET.
	Builtin types.List   `tfsdk:"builtin"`
	Feature types.String `tfsdk:"feature"`
}

func AppfwmultipartformcontenttypeDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"isregex": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Is multipart_form content type a regular expression?",
			},
			"multipartformcontenttypevalue": schema.StringAttribute{
				Required:    true,
				Description: "Content type to be classified as multipart form",
			},

			// Read-only (GET-only) attributes surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"builtin": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "Flag to determine if multipart form contenttype is built-in or not.",
			},
			"feature": schema.StringAttribute{
				Computed:    true,
				Description: "The feature to be checked while applying this config.",
			},
		},
	}
}

// appfwmultipartformcontenttypeDataSourceSetAttrFromGet projects a NITRO
// appfwmultipartformcontenttype GET response onto the data-source model. Because
// a data source has no plan/apply reconciliation, attributes are simply filled
// from the GET (or left Null when the GET omits them). The shared utils.MapGet*
// helpers implement that projection.
func appfwmultipartformcontenttypeDataSourceSetAttrFromGet(ctx context.Context, data *AppfwmultipartformcontenttypeDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In appfwmultipartformcontenttypeDataSourceSetAttrFromGet Function")

	if v, ok := g["multipartformcontenttypevalue"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Multipartformcontenttypevalue = types.StringValue(utils.AnyToString(v))
	}

	data.Isregex = utils.MapGetString(g, "isregex")

	// Read-only metadata.
	data.Builtin = utils.MapGetStringList(g, "builtin")
	data.Feature = utils.MapGetString(g, "feature")
}
