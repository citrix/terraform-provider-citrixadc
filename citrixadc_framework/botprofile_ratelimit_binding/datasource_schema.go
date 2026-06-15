package botprofile_ratelimit_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func BotprofileRatelimitBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"bot_bind_comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments about this binding.",
			},
			"bot_rate_limit_action": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "One or more actions to be taken when the current rate becomes more than the configured rate. Only LOG action can be combined with DROP, REDIRECT, RESPOND_STATUS_TOO_MANY_REQUESTS or RESET action.",
			},
			"bot_rate_limit_enabled": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable or disable rate-limit binding.",
			},
			"bot_rate_limit_type": schema.StringAttribute{
				Required:    true,
				Description: "Rate-limiting type Following rate-limiting types are allowed:\n*SOURCE_IP - Rate-limiting based on the client IP.\n*SESSION - Rate-limiting based on the configured cookie name.\n*URL - Rate-limiting based on the configured URL.\n*GEOLOCATION - Rate-limiting based on the configured country name.\n*JA3_FINGERPRINT - Rate-limiting based on client SSL JA3 fingerprint.",
			},
			"bot_rate_limit_url": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "URL for the resource based rate-limiting.",
			},
			"bot_ratelimit": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Rate-limit binding. Maximum 30 bindings can be configured per profile for rate-limit detection. For SOURCE_IP type, only one binding can be configured, and for URL type, only one binding is allowed per URL, and for SESSION type, only one binding is allowed for a cookie name. To update the values of an existing binding, user has to first unbind that binding, and then needs to bind again with new values.",
			},
			"condition": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression to be used in a rate-limiting condition. This expression result must be a boolean value.",
			},
			"cookiename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Cookie name which is used to identify the session for session rate-limiting.",
			},
			"countrycode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Country name which is used for geolocation rate-limiting.",
			},
			"limittype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Rate-Limiting traffic Type",
			},
			"logmessage": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Message to be logged for this binding.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the profile. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.), pound (#), space ( ), at (@), equals (=), colon (:), and underscore (_) characters. Cannot be changed after the profile is added.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my profile\" or 'my profile').",
			},
			"rate": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of requests that are allowed in this session in the given period time.",
			},
			"timeslice": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time interval during which requests are tracked to check if they cross the given rate.",
			},
		},
	}
}
