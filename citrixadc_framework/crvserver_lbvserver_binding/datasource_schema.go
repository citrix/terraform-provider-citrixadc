package crvserver_lbvserver_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func CrvserverLbvserverBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"lbvserver": schema.StringAttribute{
				Required:    true,
				Description: "The Default target server name.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the cache redirection virtual server to which to bind the cache redirection policy.",
			},
		},
	}
}
