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
			"password": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Password for logging on to the database. Must be the same as the password specified in the database.",
			},
			"password_wo": schema.StringAttribute{
				Optional:    true,
				Description: "Password for logging on to the database. Must be the same as the password specified in the database.",
			},
			"password_wo_version": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Increment this version to signal a password_wo update.",
			},
			"username": schema.StringAttribute{
				Required:    true,
				Description: "Name of the database user. Must be the same as the user name specified in the database.",
			},
		},
	}
}
