package authenticationloginschema

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func AuthenticationloginschemaDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"authenticationschema": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the file for reading authentication schema to be sent for Login Page UI. This file should contain xml definition of elements as per Citrix Forms Authentication Protocol to be able to render login form. If administrator does not want to prompt users for additional credentials but continue with previously obtained credentials, then \"noschema\" can be given as argument. Please note that this applies only to loginSchemas that are used with user-defined factors, and not the vserver factor.",
			},
			"authenticationstrength": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Weight of the current authentication",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the new login schema. Login schema defines the way login form is rendered. It provides a way to customize the fields that are shown to the user. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after an action is created.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my action\" or 'my action').",
			},
			"passwdexpression": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression for password extraction during login. This can be any relevant advanced policy expression.",
			},
			"passwordcredentialindex": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The index at which user entered password should be stored in session.",
			},
			"ssocredentials": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This option indicates whether current factor credentials are the default SSO (SingleSignOn) credentials.",
			},
			"usercredentialindex": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The index at which user entered username should be stored in session.",
			},
			"userexpression": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression for username extraction during login. This can be any relevant advanced policy expression.",
			},
		},
	}
}
