package botprofile_whitelist_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func BotprofileWhitelistBindingDataSourceSchema() schema.Schema {
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
			"bot_whitelist": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Whitelist binding. Maximum 32 bindings can be configured per profile for Whitelist detection.",
			},
			"bot_whitelist_enabled": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enabled or disabled white-list binding.",
			},
			"bot_whitelist_type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of the white-list entry.",
			},
			"bot_whitelist_value": schema.StringAttribute{
				Required:    true,
				Description: "Value of bot white-list entry.",
			},
			"log": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable logging for Whitelist binding.",
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
