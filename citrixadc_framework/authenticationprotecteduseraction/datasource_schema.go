package authenticationprotecteduseraction

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func AuthenticationprotecteduseractionDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"maxconcurrentusers": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Max number of concurrent users allowed.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the action to configure.",
			},
			"realmstr": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Kerberos Realm.",
			},
		},
	}
}
