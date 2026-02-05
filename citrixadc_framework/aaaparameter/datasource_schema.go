package aaaparameter

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func AaaparameterDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"aaadloglevel": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "AAAD log level, which specifies the types of AAAD events to log in nsvpn.log.\nAvailable values function as follows:\n* EMERGENCY - Events that indicate an immediate crisis on the server.\n* ALERT - Events that might require action.\n* CRITICAL - Events that indicate an imminent server crisis.\n* ERROR - Events that indicate some type of error.\n* WARNING - Events that require action in the near future.\n* NOTICE - Events that the administrator should know about.\n* INFORMATIONAL - All but low-level events.\n* DEBUG - All events, in extreme detail.",
			},
			"aaadnatip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Source IP address to use for traffic that is sent to the authentication server.",
			},
			"aaasessionloglevel": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Audit log level, which specifies the types of events to log for cli executed commands.\nAvailable values function as follows:\n* EMERGENCY - Events that indicate an immediate crisis on the server.\n* ALERT - Events that might require action.\n* CRITICAL - Events that indicate an imminent server crisis.\n* ERROR - Events that indicate some type of error.\n* WARNING - Events that require action in the near future.\n* NOTICE - Events that the administrator should know about.\n* INFORMATIONAL - All but low-level events.\n* DEBUG - All events, in extreme detail.",
			},
			"apitokencache": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Option to enable/disable API cache feature.",
			},
			"defaultauthtype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The default authentication server type.",
			},
			"defaultcspheader": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Parameter to enable/disable default CSP header",
			},
			"dynaddr": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Set by the DHCP client when the IP address was fetched dynamically.",
			},
			"enableenhancedauthfeedback": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enhanced auth feedback provides more information to the end user about the reason for an authentication failure.  The default value is set to NO.",
			},
			"enablesessionstickiness": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enables/Disables stickiness to authentication servers",
			},
			"enablestaticpagecaching": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The default state of VPN Static Page caching. Static Page caching is enabled by default.",
			},
			"enhancedepa": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Parameter to enable/disable EPA v2 functionality",
			},
			"failedlogintimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of minutes an account will be locked if user exceeds maximum permissible attempts",
			},
			"ftmode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "First time user mode determines which configuration options are shown by default when logging in to the GUI. This setting is controlled by the GUI.",
			},
			"httponlycookie": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Parameter to set/reset HttpOnly Flag for NSC_AAAC/NSC_TMAS cookies in nfactor",
			},
			"loginencryption": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Parameter to encrypt login information for nFactor flow",
			},
			"maxaaausers": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of concurrent users allowed to log on to VPN simultaneously.",
			},
			"maxkbquestions": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "This will set maximum number of Questions to be asked for KB Validation. Default value is 2, Max Value is 6",
			},
			"maxloginattempts": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum Number of login Attempts",
			},
			"maxsamldeflatesize": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "This will set the maximum deflate size in case of SAML Redirect binding.",
			},
			"persistentloginattempts": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Persistent storage of unsuccessful user login attempts",
			},
			"pwdexpirynotificationdays": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "This will set the threshold time in days for password expiry notification. Default value is 0, which means no notification is sent",
			},
			"samesite": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "SameSite attribute value for Cookies generated in AAATM context. This attribute value will be appended only for the cookies which are specified in the builtin patset ns_cookies_samesite",
			},
			"securityinsights": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "On enabling this option, the Citrix ADC will send the security insight records to the configured collectors when request comes to Authentication endpoint.\n* If cs vserver is frontend with Authentication vserver as target for cs action, then record is sent using Authentication vserver name.\n* If vpn/lb/cs vserver are configured with Authentication ON, then then record is sent using vpn/lb/cs vserver name accordingly.\n* If authentication vserver is frontend, then record is sent using Authentication vserver name.",
			},
			"tokenintrospectioninterval": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Frequency at which a token must be verified at the Authorization Server (AS) despite being found in cache.",
			},
			"wafprotection": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Entities for which WAF Protection need to be applied.\nAvailable settings function as follows:\n* DEFAULT - No Endpoint WAF protection.\n* AUTH - Endpoints used for Authentication applicable for both AAATM, IDP, GATEWAY use cases.\n* VPN - Endpoints used for Gateway use cases.\n* PORTAL - Endpoints related to web portal.\n* DISABLED - No Endpoint WAF protection.\nCurrently supported only in default partition",
			},
		},
	}
}
