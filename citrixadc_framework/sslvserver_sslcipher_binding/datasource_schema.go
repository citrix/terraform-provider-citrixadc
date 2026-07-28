package sslvserver_sslcipher_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func SslvserverSslcipherBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"cipheraliasname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The name of the cipher group/alias/individual cipheri bindings.",
			},
			"ciphername": schema.StringAttribute{
				Required:    true,
				Description: "Name of the individual cipher, user-defined cipher group, or predefined (built-in) cipher alias.",
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
