package botprofile_tps_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func BotprofileTpsBindingDataSourceSchema() schema.Schema {
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
			"bot_tps": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "TPS binding. For each type only binding can be configured. To  update the values of an existing binding, user has to first unbind that binding, and then needs to bind again with new values.",
			},
			"bot_tps_action": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "One to more actions to be taken if bot is detected based on this TPS binding. Only LOG action can be combined with DROP, RESET, REDIRECT, or MITIGIATION action.",
			},
			"bot_tps_enabled": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enabled or disabled TPS binding.",
			},
			"bot_tps_type": schema.StringAttribute{
				Required:    true,
				Description: "Type of TPS binding.",
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
			"percentage": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum percentage increase in the requests from (or to) a IP, Geolocation, URL or Host in 30 minutes interval.",
			},
			"threshold": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of requests that are allowed from (or to) a IP, Geolocation, URL or Host in 1 second time interval.",
			},
		},
	}
}
