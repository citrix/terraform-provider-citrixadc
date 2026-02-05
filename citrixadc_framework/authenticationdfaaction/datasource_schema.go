package authenticationdfaaction

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func AuthenticationdfaactionDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"clientid": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "If configured, this string is sent to the DFA server as the X-Citrix-Exchange header value.",
			},
			"defaultauthenticationgroup": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This is the default group that is chosen when the authentication succeeds in addition to extracted groups.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the DFA action.\nMust begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the DFA action is added.",
			},
			"passphrase": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Key shared between the DFA server and the Citrix ADC.\nRequired to allow the Citrix ADC to communicate with the DFA server.",
			},
			"serverurl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "DFA Server URL",
			},
		},
	}
}
