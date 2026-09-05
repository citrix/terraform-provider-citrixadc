package cmpaction

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// CmpactionDataSourceModel is the data-source-specific model, decoupled from
// CmpactionResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes that the resource deliberately omits
// (builtin, feature, isdefault). Every non-key attribute is Computed; the
// Framework's per-attribute model <-> schema reflection requires this model to
// have exactly the attributes the data-source schema declares, which is why it
// cannot reuse the resource model.
type CmpactionDataSourceModel struct {
	Id              types.String `tfsdk:"id"`
	Addvaryheader   types.String `tfsdk:"addvaryheader"`
	Cmptype         types.String `tfsdk:"cmptype"`
	Deltatype       types.String `tfsdk:"deltatype"`
	Name            types.String `tfsdk:"name"`
	Newname         types.String `tfsdk:"newname"`
	Varyheadervalue types.String `tfsdk:"varyheadervalue"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/cmpaction.json). Never settable; populated from GET.
	Builtin   types.List   `tfsdk:"builtin"`
	Feature   types.String `tfsdk:"feature"`
	Isdefault types.Bool   `tfsdk:"isdefault"`
}

func CmpactionDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"addvaryheader": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Control insertion of the Vary header in HTTP responses compressed by Citrix ADC. Intermediate caches store different versions of the response for different values of the headers present in the Vary response header.",
			},
			"cmptype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of compression performed by this action.\nAvailable settings function as follows:\n* COMPRESS - Apply GZIP or DEFLATE compression to the response, depending on the request header. Prefer GZIP.\n* GZIP - Apply GZIP compression.\n* DEFLATE - Apply DEFLATE compression.\n* NOCOMPRESS - Do not compress the response if the request matches a policy that uses this action.",
			},
			"deltatype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The type of delta action (if delta type compression action is defined).",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the compression action. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Can be changed after the action is added.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my cmp action\" or 'my cmp action').",
			},
			"newname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "New name for the compression action. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at\n(@), equals (=), and hyphen (-) characters.\nChoose a name that can be correlated with the function that the action performs.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my cmp action\" or 'my cmp action').",
			},
			"varyheadervalue": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The value of the HTTP Vary header for compressed responses.",
			},

			// Read-only (GET-only) attributes surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"builtin": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "Flag to determine whether compression is default or not.",
			},
			"feature": schema.StringAttribute{
				Computed:    true,
				Description: "The feature to be checked while applying this config.",
			},
			"isdefault": schema.BoolAttribute{
				Computed:    true,
				Description: "A value of true is returned if it is a default policy.",
			},
		},
	}
}

// cmpactionDataSourceSetAttrFromGet projects a NITRO cmpaction GET response onto
// the data-source model. Because a data source has no plan/apply reconciliation,
// attributes are simply filled from the GET (or left Null when the GET omits
// them). The shared utils.MapGet* helpers implement that projection.
func cmpactionDataSourceSetAttrFromGet(ctx context.Context, data *CmpactionDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In cmpactionDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Addvaryheader = utils.MapGetString(g, "addvaryheader")
	data.Cmptype = utils.MapGetString(g, "cmptype")
	data.Deltatype = utils.MapGetString(g, "deltatype")
	data.Varyheadervalue = utils.MapGetString(g, "varyheadervalue")

	// newname is a rename-only input the GET never returns -> Null.
	data.Newname = types.StringNull()

	// Read-only attributes.
	data.Builtin = utils.MapGetStringList(g, "builtin")
	data.Feature = utils.MapGetString(g, "feature")
	data.Isdefault = utils.MapGetBool(g, "isdefault")
}
