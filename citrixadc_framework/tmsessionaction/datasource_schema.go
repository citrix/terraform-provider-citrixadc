package tmsessionaction

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func TmsessionactionDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"defaultauthorizationaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Allow or deny access to content for which there is no specific authorization policy.",
			},
			"homepage": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Web address of the home page that a user is displayed when authentication vserver is bookmarked and used to login.",
			},
			"httponlycookie": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Allow only an HTTP session cookie, in which case the cookie cannot be accessed by scripts.",
			},
			"kcdaccount": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Kerberos constrained delegation account name",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the session action. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after a session action is created.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my action\" or 'my action').",
			},
			"persistentcookie": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable or disable persistent SSO cookies for the traffic management (TM) session. A persistent cookie remains on the user device and is sent with each HTTP request. The cookie becomes stale if the session ends. This setting is overwritten if a traffic action sets persistent cookie to OFF.\nNote: If persistent cookie is enabled, make sure you set the persistent cookie validity.",
			},
			"persistentcookievalidity": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Integer specifying the number of minutes for which the persistent cookie remains valid. Can be set only if the persistent cookie setting is enabled.",
			},
			"sesstimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Session timeout, in minutes. If there is no traffic during the timeout period, the user is disconnected and must reauthenticate to access intranet resources.",
			},
			"sso": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use single sign-on (SSO) to log users on to all web applications automatically after they authenticate, or pass users to the web application logon page to authenticate to each application individually. Note that this configuration does not honor the following authentication types for security reason. BASIC, DIGEST, and NTLM (without Negotiate NTLM2 Key or Negotiate Sign Flag). Use TM TrafficAction to configure SSO for these authentication types.",
			},
			"ssocredential": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use the primary or secondary authentication credentials for single sign-on (SSO).",
			},
			"ssodomain": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Domain to use for single sign-on (SSO).",
			},
		},
	}
}
