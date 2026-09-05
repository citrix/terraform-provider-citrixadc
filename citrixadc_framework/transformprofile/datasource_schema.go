package transformprofile

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// TransformprofileDataSourceModel is the data-source-specific model, decoupled
// from TransformprofileResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes that the resource deliberately omits
// (regexforfindingurl*, additional*headerslist, ...). Every non-key attribute is
// Computed.
type TransformprofileDataSourceModel struct {
	Id                        types.String `tfsdk:"id"`
	Comment                   types.String `tfsdk:"comment"`
	Name                      types.String `tfsdk:"name"` // Required lookup key
	Onlytransformabsurlinbody types.String `tfsdk:"onlytransformabsurlinbody"`
	Type                      types.String `tfsdk:"type"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/transformprofile.json). Never settable; populated from GET.
	Regexforfindingurlinjavascript types.String `tfsdk:"regexforfindingurlinjavascript"`
	Regexforfindingurlincss        types.String `tfsdk:"regexforfindingurlincss"`
	Regexforfindingurlinxcomponent types.String `tfsdk:"regexforfindingurlinxcomponent"`
	Regexforfindingurlinxml        types.String `tfsdk:"regexforfindingurlinxml"`
	Additionalreqheaderslist       types.String `tfsdk:"additionalreqheaderslist"`
	Additionalrespheaderslist      types.String `tfsdk:"additionalrespheaderslist"`
}

func TransformprofileDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments to preserve information about this URL Transformation profile.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the URL transformation profile. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the URL transformation profile is added.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, my transform profile or my transform profile).",
			},
			"onlytransformabsurlinbody": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "In the HTTP body, transform only absolute URLs. Relative URLs are ignored.",
			},
			"type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of transformation. Always URL for URL Transformation profiles.",
			},

			// Read-only (GET-only) attributes surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"regexforfindingurlinjavascript": schema.StringAttribute{
				Computed:    true,
				Description: "Patclass having regexes to find the URLs in JavaScript.",
			},
			"regexforfindingurlincss": schema.StringAttribute{
				Computed:    true,
				Description: "Patclass having regexes to find the URLs in CSS.",
			},
			"regexforfindingurlinxcomponent": schema.StringAttribute{
				Computed:    true,
				Description: "Patclass having regexes to find the URLs in X-Component.",
			},
			"regexforfindingurlinxml": schema.StringAttribute{
				Computed:    true,
				Description: "Patclass having regexes to find the URLs in XML.",
			},
			"additionalreqheaderslist": schema.StringAttribute{
				Computed:    true,
				Description: "Patclass having a list of additional request header names that should transformed.",
			},
			"additionalrespheaderslist": schema.StringAttribute{
				Computed:    true,
				Description: "Patclass having a list of additional response header names that should transformed.",
			},
		},
	}
}

// transformprofileDataSourceSetAttrFromGet projects a NITRO transformprofile GET
// response onto the data-source model. Because a data source has no plan/apply
// reconciliation, attributes are simply filled from the GET (or left Null when
// the GET omits them). The shared utils.MapGet* helpers implement that projection.
func transformprofileDataSourceSetAttrFromGet(ctx context.Context, data *TransformprofileDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In transformprofileDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Comment = utils.MapGetString(g, "comment")
	data.Onlytransformabsurlinbody = utils.MapGetString(g, "onlytransformabsurlinbody")
	data.Type = utils.MapGetString(g, "type")

	// Read-only attributes.
	data.Regexforfindingurlinjavascript = utils.MapGetString(g, "regexforfindingurlinjavascript")
	data.Regexforfindingurlincss = utils.MapGetString(g, "regexforfindingurlincss")
	data.Regexforfindingurlinxcomponent = utils.MapGetString(g, "regexforfindingurlinxcomponent")
	data.Regexforfindingurlinxml = utils.MapGetString(g, "regexforfindingurlinxml")
	data.Additionalreqheaderslist = utils.MapGetString(g, "additionalreqheaderslist")
	data.Additionalrespheaderslist = utils.MapGetString(g, "additionalrespheaderslist")
}
