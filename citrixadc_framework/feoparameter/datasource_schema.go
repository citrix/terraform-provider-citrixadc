package feoparameter

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// FeoparameterDataSourceModel is the data-source-specific model, decoupled from
// FeoparameterResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes the resource deliberately omits
// (builtin, feature). Every non-key attribute is Computed.
type FeoparameterDataSourceModel struct {
	Id                 types.String `tfsdk:"id"`
	Cssinlinethressize types.Int64  `tfsdk:"cssinlinethressize"`
	Imginlinethressize types.Int64  `tfsdk:"imginlinethressize"`
	Jpegqualitypercent types.Int64  `tfsdk:"jpegqualitypercent"`
	Jsinlinethressize  types.Int64  `tfsdk:"jsinlinethressize"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/feoparameter.json). Never settable; populated from GET.
	Builtin types.List   `tfsdk:"builtin"`
	Feature types.String `tfsdk:"feature"`
}

func FeoparameterDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"cssinlinethressize": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Threshold value of the file size (in bytes) for converting external CSS files to inline CSS files.",
			},
			"imginlinethressize": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum file size of an image (in bytes), for coverting linked images to inline images.",
			},
			"jpegqualitypercent": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The percentage value of a JPEG image quality to be reduced. Range: 0 - 100",
			},
			"jsinlinethressize": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Threshold value of the file size (in bytes), for converting external JavaScript files to inline JavaScript files.",
			},

			// Read-only (GET-only) attributes surfaced by the data source
			// (intentionally NOT modeled on the resource). All Computed.
			"builtin": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "Indicates that a variable is a built-in (SYSTEM INTERNAL) type.",
			},
			"feature": schema.StringAttribute{
				Computed:    true,
				Description: "The feature to be checked while applying this config.",
			},
		},
	}
}

// feoparameterDataSourceSetAttrFromGet projects a NITRO feoparameter GET
// response onto the data-source model. Attributes are filled from the GET (or
// left Null when the GET omits them) via the shared utils.MapGet* helpers.
func feoparameterDataSourceSetAttrFromGet(ctx context.Context, data *FeoparameterDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In feoparameterDataSourceSetAttrFromGet Function")

	// Read/write attributes as read-back outputs.
	data.Cssinlinethressize = utils.MapGetInt64(g, "cssinlinethressize")
	data.Imginlinethressize = utils.MapGetInt64(g, "imginlinethressize")
	data.Jpegqualitypercent = utils.MapGetInt64(g, "jpegqualitypercent")
	data.Jsinlinethressize = utils.MapGetInt64(g, "jsinlinethressize")

	// Read-only attributes.
	data.Builtin = utils.MapGetStringList(g, "builtin")
	data.Feature = utils.MapGetString(g, "feature")

	// Singleton (unnamed) resource - static ID matching the resource behavior.
	data.Id = types.StringValue("feoparameter-config")
}
