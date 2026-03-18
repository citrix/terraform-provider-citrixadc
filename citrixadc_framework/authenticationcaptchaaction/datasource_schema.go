package authenticationcaptchaaction

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func AuthenticationcaptchaactionDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"defaultauthenticationgroup": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This is the group that is added to user sessions that match current policy.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the new captcha action. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after an action is created.\n\nThe following requirement applies only to the NetScaler CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my action\" or 'my action').",
			},
			"scorethreshold": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "This is the score threshold value for recaptcha v3.",
			},
			"secretkey": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Secret of gateway as established at the captcha source.",
			},
			"serverurl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This is the endpoint at which captcha response is validated.",
			},
			"sitekey": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Sitekey to identify gateway fqdn while loading captcha.",
			},
		},
	}
}
