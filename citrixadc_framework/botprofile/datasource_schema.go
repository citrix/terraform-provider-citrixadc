package botprofile

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

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
		},
	}
}
