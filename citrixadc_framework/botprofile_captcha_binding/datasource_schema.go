package botprofile_captcha_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func BotprofileCaptchaBindingDataSourceSchema() schema.Schema {
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
			"bot_captcha_action": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "One or more actions to be taken when client fails captcha challenge. Only, log action can be configured with DROP, REDIRECT or RESET action.",
			},
			"bot_captcha_enabled": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable or disable the captcha binding.",
			},
			"bot_captcha_url": schema.StringAttribute{
				Required:    true,
				Description: "URL for which the Captcha action, if configured under IP reputation, TPS or device fingerprint, need to be applied.",
			},
			"captcharesource": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Captcha action binding. For each URL, only one binding is allowed. To update the values of an existing URL binding, user has to first unbind that binding, and then needs to bind the URL again with new values. Maximum 30 bindings can be configured per profile.",
			},
			"graceperiod": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time (in seconds) duration for which no new captcha challenge is sent after current captcha challenge has been answered successfully.",
			},
			"logmessage": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Message to be logged for this binding.",
			},
			"muteperiod": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time (in seconds) duration for which client which failed captcha need to wait until allowed to try again. The requests from this client are silently dropped during the mute period.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the profile. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.), pound (#), space ( ), at (@), equals (=), colon (:), and underscore (_) characters. Cannot be changed after the profile is added.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my profile\" or 'my profile').",
			},
			"requestsizelimit": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Length of body request (in Bytes) up to (equal or less than) which captcha challenge will be provided to client. Above this length threshold the request will be dropped. This is to avoid DOS and DDOS attacks.",
			},
			"retryattempts": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of times client can retry solving the captcha.",
			},
			"waittime": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Wait time in seconds for which ADC needs to wait for the Captcha response. This is to avoid DOS attacks.",
			},
		},
	}
}
