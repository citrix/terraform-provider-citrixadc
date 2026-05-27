package authenticationsmartaccessprofile

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func AuthenticationsmartaccessprofileDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Optional comment for the profile.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the Smartaccess profile",
			},
			"tags": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The tag that is associated with Smartaccess profile.",
			},
		},
	}
}
