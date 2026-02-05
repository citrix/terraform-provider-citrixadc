package authenticationcitrixauthaction

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func AuthenticationcitrixauthactionDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"authentication": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Authentication needs to be disabled for searching user object without performing authentication.",
			},
			"authenticationtype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of the Citrix Authentication implementation. Default implementation uses Citrix Cloud Connector.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the new Citrix Authentication action. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after an action is created.\n\nThe following requirement applies only to the NetScaler CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my action\" or 'my action').",
			},
		},
	}
}
