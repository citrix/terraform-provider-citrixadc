package botprofile_kmdetectionexpr_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func BotprofileKmdetectionexprBindingDataSourceSchema() schema.Schema {
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
			"bot_km_detection_enabled": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable or disable the keyboard-mouse based binding.",
			},
			"bot_km_expression_name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the keyboard-mouse expression object.",
			},
			"bot_km_expression_value": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "JavaScript file for keyboard-mouse detection, would be inserted if the result of the expression is true.",
			},
			"kmdetectionexpr": schema.BoolAttribute{
				Required:    true,
				Description: "Keyboard-mouse based detection binding. For each name, only one binding is allowed. To update the values of an existing binding, user has to first unbind that binding, then needs to bind again with new vlaues. Maximum 30 bindings can be configured per profile.",
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
