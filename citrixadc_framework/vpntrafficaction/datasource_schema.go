package vpntrafficaction

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func VpntrafficactionDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"apptimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum amount of time, in minutes, a user can stay logged on to the web application.",
			},
			"formssoaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the form-based single sign-on profile. Form-based single sign-on allows users to log on one time to all protected applications in your network, instead of requiring them to log on separately to access each one.",
			},
			"fta": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify file type association, which is a list of file extensions that users are allowed to open.",
			},
			"hdx": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Provide hdx proxy to the ICA traffic",
			},
			"kcdaccount": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Kerberos constrained delegation account name",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the traffic action. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after a traffic action is created.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my action\" or 'my action').",
			},
			"passwdexpression": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "expression that will be evaluated to obtain password for SingleSignOn",
			},
			"proxy": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address and Port of the proxy server to be used for HTTP access for this request.",
			},
			"qual": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Protocol, either HTTP or TCP, to be used with the action.",
			},
			"samlssoprofile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Profile to be used for doing SAML SSO to remote relying party",
			},
			"sso": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Provide single sign-on to the web application.\n	    NOTE : Authentication mechanisms like Basic-authentication  require the user credentials to be sent in plaintext which is not secure if the server is running on HTTP (instead of HTTPS).",
			},
			"userexpression": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "expression that will be evaluated to obtain username for SingleSignOn",
			},
			"wanscaler": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use the Repeater Plug-in to optimize network traffic.",
			},
		},
	}
}
