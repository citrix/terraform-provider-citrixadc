package botsettings

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func BotsettingsDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"defaultnonintrusiveprofile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Profile to use when the feature is not enabled but feature is licensed. NonIntrusive checks will be disabled and IPRep cronjob(24 Hours) will be removed if this is set to BOT_BYPASS.",
			},
			"defaultprofile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Profile to use when a connection does not match any policy. Default setting is \" \", which sends unmatched connections back to the Citrix ADC without attempting to filter them further.",
			},
			"dfprequestlimit": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of requests to allow without bot session cookie if device fingerprint is enabled",
			},
			"javascriptname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the JavaScript that the Bot Management feature  uses in response.\nMust begin with a letter or number, and can consist of from 1 to 31 letters, numbers, and the hyphen (-) and underscore (_) symbols.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my cookie name\" or 'my cookie name').",
			},
			"proxypassword": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Password with which user logs on.",
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
				Description: "Name of the SessionCookie that the Bot Management feature uses for tracking.\nMust begin with a letter or number, and can consist of from 1 to 31 letters, numbers, and the hyphen (-) and underscore (_) symbols.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my cookie name\" or 'my cookie name').",
			},
			"sessiontimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Timeout, in seconds, after which a user session is terminated.",
			},
			"signatureautoupdate": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Flag used to enable/disable bot auto update signatures",
			},
			"signatureurl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "URL to download the bot signature mapping file from server",
			},
			"trapurlautogenerate": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable/disable trap URL auto generation. When enabled, trap URL is updated within the configured interval.",
			},
			"trapurlinterval": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time in seconds after which trap URL is updated.",
			},
			"trapurllength": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Length of the auto-generated trap URL.",
			},
		},
	}
}
