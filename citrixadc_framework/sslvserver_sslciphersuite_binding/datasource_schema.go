package sslvserver_sslciphersuite_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func SslvserverSslciphersuiteBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"ciphername": schema.StringAttribute{
				Required:    true,
				Description: "The cipher group/alias/individual cipher configuration",
			},
			"description": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The cipher suite description.",
			},
			"vservername": schema.StringAttribute{
				Required:    true,
				Description: "Name of the SSL virtual server.",
			},
		},
	}
}
