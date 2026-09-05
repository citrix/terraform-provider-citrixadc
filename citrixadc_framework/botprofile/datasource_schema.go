package botprofile

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// BotprofileDataSourceModel is the data-source-specific model, decoupled from
// BotprofileResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only metadata attributes that the resource deliberately
// omits (builtin, feature). Every non-key attribute is Computed; the Framework's
// per-attribute model <-> schema reflection requires this model to have exactly
// the attributes the data-source schema declares, which is why it cannot reuse
// the resource model.
type BotprofileDataSourceModel struct {
	Id                                     types.String `tfsdk:"id"`
	Addcookieflags                         types.String `tfsdk:"addcookieflags"`
	BotEnableBlackList                     types.String `tfsdk:"bot_enable_black_list"`
	BotEnableIpReputation                  types.String `tfsdk:"bot_enable_ip_reputation"`
	BotEnableRateLimit                     types.String `tfsdk:"bot_enable_rate_limit"`
	BotEnableTps                           types.String `tfsdk:"bot_enable_tps"`
	BotEnableWhiteList                     types.String `tfsdk:"bot_enable_white_list"`
	Clientipexpression                     types.String `tfsdk:"clientipexpression"`
	Comment                                types.String `tfsdk:"comment"`
	Devicefingerprint                      types.String `tfsdk:"devicefingerprint"`
	Devicefingerprintaction                types.List   `tfsdk:"devicefingerprintaction"`
	Devicefingerprintmobile                types.List   `tfsdk:"devicefingerprintmobile"`
	Dfprequestlimit                        types.Int64  `tfsdk:"dfprequestlimit"`
	Errorurl                               types.String `tfsdk:"errorurl"`
	Headlessbrowserdetection               types.String `tfsdk:"headlessbrowserdetection"`
	Kmdetection                            types.String `tfsdk:"kmdetection"`
	Kmeventspostbodylimit                  types.Int64  `tfsdk:"kmeventspostbodylimit"`
	Kmjavascriptname                       types.String `tfsdk:"kmjavascriptname"`
	Name                                   types.String `tfsdk:"name"` // Required lookup key
	Sessioncookiename                      types.String `tfsdk:"sessioncookiename"`
	Sessiontimeout                         types.Int64  `tfsdk:"sessiontimeout"`
	Signature                              types.String `tfsdk:"signature"`
	Signaturemultipleuseragentheaderaction types.List   `tfsdk:"signaturemultipleuseragentheaderaction"`
	Signaturenouseragentheaderaction       types.List   `tfsdk:"signaturenouseragentheaderaction"`
	Spoofedreqaction                       types.List   `tfsdk:"spoofedreqaction"`
	Trap                                   types.String `tfsdk:"trap"`
	Trapaction                             types.List   `tfsdk:"trapaction"`
	Trapurl                                types.String `tfsdk:"trapurl"`
	Verboseloglevel                        types.String `tfsdk:"verboseloglevel"`

	// Read-only (GET-only) metadata from the NITRO doc read-only set
	// (zion73x_readonly/botprofile.json). Never settable; populated from GET.
	Builtin types.List   `tfsdk:"builtin"`
	Feature types.String `tfsdk:"feature"`
}

func BotprofileDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"addcookieflags": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Add the specified flags to bot session cookies. Available settings function as follows:\n* None - Do not add flags to cookies.\n* HTTP Only - Add the HTTP Only flag to cookies, which prevents scripts from accessing cookies.\n* Secure - Add Secure flag to cookies.\n* All - Add both HTTPOnly and Secure flags to cookies.",
			},
			"bot_enable_black_list": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable black-list bot detection.",
			},
			"bot_enable_ip_reputation": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable IP-reputation bot detection.",
			},
			"bot_enable_rate_limit": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable rate-limit bot detection.",
			},
			"bot_enable_tps": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable TPS.",
			},
			"bot_enable_white_list": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable white-list bot detection.",
			},
			"clientipexpression": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression to get the client IP.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments about the purpose of profile, or other useful information about the profile.",
			},
			"devicefingerprint": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable device-fingerprint bot detection",
			},
			"devicefingerprintaction": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Action to be taken for device-fingerprint based bot detection.",
			},
			"devicefingerprintmobile": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Enabling bot device fingerprint protection for mobile clients",
			},
			"dfprequestlimit": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of requests to allow without bot session cookie if device fingerprint is enabled",
			},
			"errorurl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "URL that Bot protection uses as the Error URL.",
			},
			"headlessbrowserdetection": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable Headless Browser detection.",
			},
			"kmdetection": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable keyboard-mouse based bot detection.",
			},
			"kmeventspostbodylimit": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Size of the KM data send by the browser, needs to be processed on ADC",
			},
			"kmjavascriptname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the JavaScript file that the Bot Management feature will insert in the response for keyboard-mouse based detection.\nMust begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my javascript file name\" or 'my javascript file name').",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the profile. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.), pound (#), space ( ), at (@), equals (=), colon (:), and underscore (_) characters. Cannot be changed after the profile is added.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my profile\" or 'my profile').",
			},
			"sessioncookiename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the SessionCookie that the Bot Management feature uses for tracking.\nMust begin with a letter or number, and can consist of from 1 to 31 letters, numbers, and the hyphen (-) and underscore (_) symbols.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my cookie name\" or 'my cookie name').",
			},
			"sessiontimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Timeout, in seconds, after which a user session is terminated.",
			},
			"signature": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of object containing bot static signature details.",
			},
			"signaturemultipleuseragentheaderaction": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Actions to be taken if multiple User-Agent headers are seen in a request (Applicable if Signature check is enabled). Log action should be combined with other actions",
			},
			"signaturenouseragentheaderaction": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Actions to be taken if no User-Agent header in the request (Applicable if Signature check is enabled).",
			},
			"spoofedreqaction": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Actions to be taken on a spoofed request (A request spoofing good bot user agent string).",
			},
			"trap": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable trap bot detection.",
			},
			"trapaction": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Action to be taken for bot trap based bot detection.",
			},
			"trapurl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "URL that Bot protection uses as the Trap URL.",
			},
			"verboseloglevel": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Bot verbose Logging. Based on the log level, ADC will log additional information whenever client is detected as a bot.",
			},

			// Read-only (GET-only) metadata surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"builtin": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "Flag to determine if bot profile is built-in or not. Possible values = MODIFIABLE, DELETABLE, IMMUTABLE, PARTITION_ALL.",
			},
			"feature": schema.StringAttribute{
				Computed:    true,
				Description: "The feature to be checked while applying this config.",
			},
		},
	}
}

// botprofileDataSourceSetAttrFromGet projects a NITRO botprofile GET response
// onto the data-source model. Because a data source has no plan/apply
// reconciliation, attributes are simply filled from the GET (or left Null when
// the GET omits them). The shared utils.MapGet* helpers implement that
// projection.
func botprofileDataSourceSetAttrFromGet(ctx context.Context, data *BotprofileDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In botprofileDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Addcookieflags = utils.MapGetString(g, "addcookieflags")
	data.BotEnableBlackList = utils.MapGetString(g, "bot_enable_black_list")
	data.BotEnableIpReputation = utils.MapGetString(g, "bot_enable_ip_reputation")
	data.BotEnableRateLimit = utils.MapGetString(g, "bot_enable_rate_limit")
	data.BotEnableTps = utils.MapGetString(g, "bot_enable_tps")
	data.BotEnableWhiteList = utils.MapGetString(g, "bot_enable_white_list")
	data.Clientipexpression = utils.MapGetString(g, "clientipexpression")
	data.Comment = utils.MapGetString(g, "comment")
	data.Devicefingerprint = utils.MapGetString(g, "devicefingerprint")
	data.Devicefingerprintaction = utils.MapGetStringList(g, "devicefingerprintaction")
	data.Devicefingerprintmobile = utils.MapGetStringList(g, "devicefingerprintmobile")
	data.Dfprequestlimit = utils.MapGetInt64(g, "dfprequestlimit")
	data.Errorurl = utils.MapGetString(g, "errorurl")
	data.Headlessbrowserdetection = utils.MapGetString(g, "headlessbrowserdetection")
	data.Kmdetection = utils.MapGetString(g, "kmdetection")
	data.Kmeventspostbodylimit = utils.MapGetInt64(g, "kmeventspostbodylimit")
	data.Kmjavascriptname = utils.MapGetString(g, "kmjavascriptname")
	data.Sessioncookiename = utils.MapGetString(g, "sessioncookiename")
	data.Sessiontimeout = utils.MapGetInt64(g, "sessiontimeout")
	data.Signature = utils.MapGetString(g, "signature")
	data.Signaturemultipleuseragentheaderaction = utils.MapGetStringList(g, "signaturemultipleuseragentheaderaction")
	data.Signaturenouseragentheaderaction = utils.MapGetStringList(g, "signaturenouseragentheaderaction")
	data.Spoofedreqaction = utils.MapGetStringList(g, "spoofedreqaction")
	data.Trap = utils.MapGetString(g, "trap")
	data.Trapaction = utils.MapGetStringList(g, "trapaction")
	data.Trapurl = utils.MapGetString(g, "trapurl")
	data.Verboseloglevel = utils.MapGetString(g, "verboseloglevel")

	// Read-only metadata.
	data.Builtin = utils.MapGetStringList(g, "builtin")
	data.Feature = utils.MapGetString(g, "feature")
}
