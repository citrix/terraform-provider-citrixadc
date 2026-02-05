package appfwsettings

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func AppfwsettingsDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"ceflogging": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable CEF format logs.",
			},
			"centralizedlearning": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Flag used to enable/disable ADM centralized learning",
			},
			"clientiploggingheader": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of an HTTP header that contains the IP address that the client used to connect to the protected web site or service.",
			},
			"cookieflags": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Add the specified flags to AppFW cookies. Available setttings function as follows:\n* None - Do not add flags to AppFW cookies.\n* HTTP Only - Add the HTTP Only flag to AppFW cookies, which prevent scripts from accessing them.\n* Secure - Add Secure flag to AppFW cookies.\n* All - Add both HTTPOnly and Secure flag to AppFW cookies.",
			},
			"cookiepostencryptprefix": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "String that is prepended to all encrypted cookie values.",
			},
			"defaultprofile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Profile to use when a connection does not match any policy. Default setting is APPFW_BYPASS, which sends unmatched connections back to the Citrix ADC without attempting to filter them further.",
			},
			"entitydecoding": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Transform multibyte (double- or half-width) characters to single width characters.",
			},
			"geolocationlogging": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable Geo-Location Logging in CEF format logs.",
			},
			"importsizelimit": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum cumulative size in bytes of all objects imported to Netscaler. The user is not allowed to import an object if the operation exceeds the currently configured limit.",
			},
			"learnratelimit": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of connections per second that the application firewall learning engine examines to generate new relaxations for learning-enabled security checks. The application firewall drops any connections above this limit from the list of connections used by the learning engine.",
			},
			"logmalformedreq": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Log requests that are so malformed that application firewall parsing doesn't occur.",
			},
			"malformedreqaction": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "flag to define action on malformed requests that application firewall cannot parse",
			},
			"proxypassword": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Password with which proxy user logs on.",
			},
			"proxyport": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Proxy Server Port to get updated signatures from AWS.",
			},
			"proxyserver": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Proxy Server IP to get updated signatures from AWS.",
			},
			"proxyusername": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Proxy Username",
			},
			"sessioncookiename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the session cookie that the application firewall uses to track user sessions.\nMust begin with a letter or number, and can consist of from 1 to 31 letters, numbers, and the hyphen (-) and underscore (_) symbols.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my cookie name\" or 'my cookie name').",
			},
			"sessionlifetime": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum amount of time (in seconds) that the application firewall allows a user session to remain active, regardless of user activity. After this time, the user session is terminated. Before continuing to use the protected web site, the user must establish a new session by opening a designated start URL. A value of 0 represents infinite time.",
			},
			"sessionlimit": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of sessions that the application firewall allows to be active, regardless of user activity. After the max_limit reaches, No more user session will be created .",
			},
			"sessiontimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Timeout, in seconds, after which a user session is terminated. Before continuing to use the protected web site, the user must establish a new session by opening a designated start URL.",
			},
			"signatureautoupdate": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Flag used to enable/disable auto update signatures",
			},
			"signatureurl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "URL to download the mapping file from server",
			},
			"undefaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Profile to use when an application firewall policy evaluates to undefined (UNDEF).\nAn UNDEF event indicates an internal error condition. The APPFW_BLOCK built-in profile is the default setting. You can specify a different built-in or user-created profile as the UNDEF profile.",
			},
			"useconfigurablesecretkey": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use configurable secret key in AppFw operations",
			},
		},
	}
}
