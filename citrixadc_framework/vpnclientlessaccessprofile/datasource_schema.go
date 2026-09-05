package vpnclientlessaccessprofile

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// VpnclientlessaccessprofileDataSourceModel is the data-source-specific model,
// decoupled from VpnclientlessaccessprofileResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes the resource deliberately omits. Every
// non-key attribute is Computed; the Framework's per-attribute model <-> schema
// reflection requires this model to have exactly the attributes the data-source
// schema declares, which is why it cannot reuse the resource model.
type VpnclientlessaccessprofileDataSourceModel struct {
	Id          types.String `tfsdk:"id"`
	Profilename types.String `tfsdk:"profilename"` // Required lookup key

	// Read/write attributes, surfaced here as Computed outputs.
	Clientconsumedcookies          types.String `tfsdk:"clientconsumedcookies"`
	Javascriptrewritepolicylabel   types.String `tfsdk:"javascriptrewritepolicylabel"`
	Regexforfindingcustomurls      types.String `tfsdk:"regexforfindingcustomurls"`
	Regexforfindingurlincss        types.String `tfsdk:"regexforfindingurlincss"`
	Regexforfindingurlinjavascript types.String `tfsdk:"regexforfindingurlinjavascript"`
	Regexforfindingurlinxcomponent types.String `tfsdk:"regexforfindingurlinxcomponent"`
	Regexforfindingurlinxml        types.String `tfsdk:"regexforfindingurlinxml"`
	Reqhdrrewritepolicylabel       types.String `tfsdk:"reqhdrrewritepolicylabel"`
	Requirepersistentcookie        types.String `tfsdk:"requirepersistentcookie"`
	Reshdrrewritepolicylabel       types.String `tfsdk:"reshdrrewritepolicylabel"`
	Urlrewritepolicylabel          types.String `tfsdk:"urlrewritepolicylabel"`

	// Read-only (GET-only) metadata from the NITRO doc read-only set
	// (zion73x_readonly/vpnclientlessaccessprofile.json). Never settable;
	// populated from GET.
	Cssrewritepolicylabel        types.String `tfsdk:"cssrewritepolicylabel"`
	Xmlrewritepolicylabel        types.String `tfsdk:"xmlrewritepolicylabel"`
	Xcomponentrewritepolicylabel types.String `tfsdk:"xcomponentrewritepolicylabel"`
	Isdefault                    types.Bool   `tfsdk:"isdefault"`
	Description                  types.String `tfsdk:"description"`
	Builtin                      types.List   `tfsdk:"builtin"`
	Feature                      types.String `tfsdk:"feature"`
}

func VpnclientlessaccessprofileDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"clientconsumedcookies": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify the name of the pattern set containing the names of the cookies, which are allowed between the client and the server. If a pattern set is not specified, Citrix Gateway does not allow any cookies between the client and the server. A cookie that is not specified in the pattern set is handled by Citrix Gateway on behalf of the client.",
			},
			"javascriptrewritepolicylabel": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the configured JavaScript rewrite policy label.  If you do not specify a policy label name, then JAVA scripts are not rewritten.",
			},
			"profilename": schema.StringAttribute{
				Required:    true,
				Description: "Name for the Citrix Gateway clientless access profile. Must begin with an ASCII alphabetic or underscore (_) character, and must consist only of ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the profile is created.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my profile\" or 'my profile').",
			},
			"regexforfindingcustomurls": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the pattern set that contains the regular expressions, which match the URLs in the custom content type other than HTML, CSS, XML, XCOMP, and JavaScript. The custom content type should be included in the patset ns_cvpn_custom_content_types.",
			},
			"regexforfindingurlincss": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the pattern set that contains the regular expressions, which match the URL in the CSS.",
			},
			"regexforfindingurlinjavascript": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the pattern set that contains the regular expressions, which match the URL in Java script.",
			},
			"regexforfindingurlinxcomponent": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the pattern set that contains the regular expressions, which match the URL in X Component.",
			},
			"regexforfindingurlinxml": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the pattern set that contains the regular expressions, which match the URL in XML.",
			},
			"reqhdrrewritepolicylabel": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the configured Request rewrite policy label.  If you do not specify a policy label name, then requests are not rewritten.",
			},
			"requirepersistentcookie": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify whether a persistent session cookie is set and accepted for clientless access. If this parameter is set to ON, COM objects, such as MSOffice, which are invoked by the browser can access the files using clientless access. Use caution because the persistent cookie is stored on the disk.",
			},
			"reshdrrewritepolicylabel": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the configured Response rewrite policy label.",
			},
			"urlrewritepolicylabel": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the configured URL rewrite policy label. If you do not specify a policy label name, then URLs are not rewritten.",
			},

			// Read-only (GET-only) metadata surfaced by the data source (these are
			// intentionally NOT modeled on the resource). All Computed.
			"cssrewritepolicylabel": schema.StringAttribute{
				Computed:    true,
				Description: "The configured CSS rewrite policylabel.",
			},
			"xmlrewritepolicylabel": schema.StringAttribute{
				Computed:    true,
				Description: "The configured XML rewrite policylabel.",
			},
			"xcomponentrewritepolicylabel": schema.StringAttribute{
				Computed:    true,
				Description: "The configured X-Component rewrite policylabel.",
			},
			"isdefault": schema.BoolAttribute{
				Computed:    true,
				Description: "A value of true is returned if it is a default vpnclientlessrwprofile.",
			},
			"description": schema.StringAttribute{
				Computed:    true,
				Description: "Description of the clientless access profile.",
			},
			"builtin": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "Flag to determine if vpn clientless rewrite profile is built-in or not. Possible values = MODIFIABLE, DELETABLE, IMMUTABLE, PARTITION_ALL.",
			},
			"feature": schema.StringAttribute{
				Computed:    true,
				Description: "The feature to be checked while applying this config.",
			},
		},
	}
}

// vpnclientlessaccessprofileDataSourceSetAttrFromGet projects a NITRO
// vpnclientlessaccessprofile GET response onto the data-source model. Because a
// data source has no plan/apply reconciliation, attributes are simply filled
// from the GET (or left Null when the GET omits them). The shared utils.MapGet*
// helpers implement that projection.
func vpnclientlessaccessprofileDataSourceSetAttrFromGet(ctx context.Context, data *VpnclientlessaccessprofileDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In vpnclientlessaccessprofileDataSourceSetAttrFromGet Function")

	if v, ok := g["profilename"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Profilename = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Clientconsumedcookies = utils.MapGetString(g, "clientconsumedcookies")
	data.Javascriptrewritepolicylabel = utils.MapGetString(g, "javascriptrewritepolicylabel")
	data.Regexforfindingcustomurls = utils.MapGetString(g, "regexforfindingcustomurls")
	data.Regexforfindingurlincss = utils.MapGetString(g, "regexforfindingurlincss")
	data.Regexforfindingurlinjavascript = utils.MapGetString(g, "regexforfindingurlinjavascript")
	data.Regexforfindingurlinxcomponent = utils.MapGetString(g, "regexforfindingurlinxcomponent")
	data.Regexforfindingurlinxml = utils.MapGetString(g, "regexforfindingurlinxml")
	data.Reqhdrrewritepolicylabel = utils.MapGetString(g, "reqhdrrewritepolicylabel")
	data.Requirepersistentcookie = utils.MapGetString(g, "requirepersistentcookie")
	data.Reshdrrewritepolicylabel = utils.MapGetString(g, "reshdrrewritepolicylabel")
	data.Urlrewritepolicylabel = utils.MapGetString(g, "urlrewritepolicylabel")

	// Read-only metadata.
	data.Cssrewritepolicylabel = utils.MapGetString(g, "cssrewritepolicylabel")
	data.Xmlrewritepolicylabel = utils.MapGetString(g, "xmlrewritepolicylabel")
	data.Xcomponentrewritepolicylabel = utils.MapGetString(g, "xcomponentrewritepolicylabel")
	data.Isdefault = utils.MapGetBool(g, "isdefault")
	data.Description = utils.MapGetString(g, "description")
	data.Builtin = utils.MapGetStringList(g, "builtin")
	data.Feature = utils.MapGetString(g, "feature")
}
