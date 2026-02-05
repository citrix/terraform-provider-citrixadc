package dbuser

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func DbuserDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"loggedin": schema.BoolAttribute{
				Required:    true,
				Description: "Display the names of all database users currently logged on to the Citrix ADC.",
			},
			"password": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Password for logging on to the database. Must be the same as the password specified in the database.",
			},
			"username": schema.StringAttribute{
				Required:    true,
				Description: "Name of the database user. Must be the same as the user name specified in the database.",
			},
		},
	}
}
