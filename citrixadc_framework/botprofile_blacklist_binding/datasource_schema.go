package botprofile_blacklist_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func BotprofileBlacklistBindingDataSourceSchema() schema.Schema {
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
			"bot_blacklist": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Blacklist binding. Maximum 32 bindings can be configured per profile for Blacklist detection.",
			},
			"bot_blacklist_action": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "One or more actions to be taken if  bot is detected based on this Blacklist binding. Only LOG action can be combined with DROP or RESET action.",
			},
			"bot_blacklist_enabled": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enabled or disbaled black-list binding.",
			},
			"bot_blacklist_type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of the black-list entry.",
			},
			"bot_blacklist_value": schema.StringAttribute{
				Required:    true,
				Description: "Value of the bot black-list entry.",
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
		},
	}
}
