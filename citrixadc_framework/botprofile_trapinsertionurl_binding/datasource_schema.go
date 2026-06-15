package botprofile_trapinsertionurl_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func BotprofileTrapinsertionurlBindingDataSourceSchema() schema.Schema {
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
			"bot_trap_url": schema.StringAttribute{
				Required:    true,
				Description: "Request URL regex pattern for which Trap URL is inserted.",
			},
			"bot_trap_url_insertion_enabled": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable or disable the request URL pattern.",
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
			"trapinsertionurl": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Bind the trap URL for the configured request URLs. Maximum 30 bindings can be configured per profile.",
			},
		},
	}
}
