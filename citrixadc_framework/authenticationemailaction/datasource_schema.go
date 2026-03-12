package authenticationemailaction

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func AuthenticationemailactionDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"content": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Content to be delivered to the user. \"$code\" string within the content will be replaced with the actual one-time-code to be sent.",
			},
			"defaultauthenticationgroup": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This is the group that is added to user sessions that match current IdP policy. It can be used in policies to identify relying party trust.",
			},
			"emailaddress": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "An optional expression that yields user's email. When not configured, user's default mail address would be used. When configured, result of this expression is used as destination email address.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the new email action. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after an action is created.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my action\" or 'my action').",
			},
			"password": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Password/Clientsecret to use when authenticating to the server.",
			},
			"serverurl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Address of the server that delivers the message. It is fully qualified fqdn such as http(s):// or smtp(s):// for http and smtp protocols respectively. For SMTP, the port number is mandatory like smtps://smtp.example.com:25.",
			},
			"timeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time after which the code expires.",
			},
			"type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of the email action. Default type is SMTP.",
			},
			"username": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Username/Clientid/EmailID to be used to authenticate to the server.",
			},
		},
	}
}
