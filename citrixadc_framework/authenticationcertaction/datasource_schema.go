package authenticationcertaction

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func AuthenticationcertactionDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"defaultauthenticationgroup": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This is the default group that is chosen when the authentication succeeds in addition to extracted groups.",
			},
			"groupnamefield": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Client-cert field from which the group is extracted.  Must be set to either \"\"Subject\"\" and \"\"Issuer\"\" (include both sets of double quotation marks).\nFormat: <field>:<subfield>",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the client cert authentication server profile (action).\nMust begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after certifcate action is created.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my authentication action\" or 'my authentication action').",
			},
			"twofactor": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enables or disables two-factor authentication.\nTwo factor authentication is client cert authentication followed by password authentication.",
			},
			"usernamefield": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Client-cert field from which the username is extracted. Must be set to either \"\"Subject\"\" and \"\"Issuer\"\" (include both sets of double quotation marks).\nFormat: <field>:<subfield>.",
			},
		},
	}
}
