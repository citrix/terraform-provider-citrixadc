package sslservice_sslcipher_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func SslserviceSslcipherBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"cipheraliasname": schema.StringAttribute{
				Computed:    true,
				Description: "The cipher group/alias/individual cipher configuration.",
			},
			"cipherdefaulton": schema.Int64Attribute{
				Computed:    true,
				Description: "Flag indicating whether the bound cipher was the DEFAULT cipher, bound at boot time, or any other cipher from the CLI",
			},
			"ciphername": schema.StringAttribute{
				Required:    true,
				Description: "Name of the individual cipher, user-defined cipher group, or predefined (built-in) cipher alias.",
			},
			"description": schema.StringAttribute{
				Computed:    true,
				Description: "The cipher suite description.",
			},
			"servicename": schema.StringAttribute{
				Required:    true,
				Description: "Name of the SSL service for which to set advanced configuration.",
			},
		},
	}
}
