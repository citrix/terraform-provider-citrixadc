package tmtrafficaction

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func TmtrafficactionDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"apptimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time interval, in minutes, of user inactivity after which the connection is closed.",
			},
			"forcedtimeout": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Setting to start, stop or reset TM session force timer",
			},
			"forcedtimeoutval": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time interval, in minutes, for which force timer should be set.",
			},
			"formssoaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the configured form-based single sign-on profile.",
			},
			"initiatelogout": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Initiate logout for the traffic management (TM) session if the policy evaluates to true. The session is then terminated after two minutes.",
			},
			"kcdaccount": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Kerberos constrained delegation account name",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the traffic action. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after a traffic action is created.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my action\" or 'my action').",
			},
			"passwdexpression": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "expression that will be evaluated to obtain password for SingleSignOn",
			},
			"persistentcookie": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use persistent cookies for the traffic session. A persistent cookie remains on the user device and is sent with each HTTP request. The cookie becomes stale if the session ends.",
			},
			"samlssoprofile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Profile to be used for doing SAML SSO to remote relying party",
			},
			"sso": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use single sign-on for the resource that the user is accessing now.",
			},
			"userexpression": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "expression that will be evaluated to obtain username for SingleSignOn",
			},
		},
	}
}
