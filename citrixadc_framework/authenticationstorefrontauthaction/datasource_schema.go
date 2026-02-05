package authenticationstorefrontauthaction

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func AuthenticationstorefrontauthactionDataSourceSchema() schema.Schema {
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
			"domain": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Domain of the server that is used for authentication. If users enter name without domain, this parameter is added to username in the authentication request to server.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the Storefront Authentication action.\nMust begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the profile is created.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my authentication action\" or 'my authentication action').",
			},
			"serverurl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "URL of the Storefront server. This is the FQDN of the Storefront server. example: https://storefront.com/.  Authentication endpoints are learned dynamically by Gateway.",
			},
		},
	}
}
