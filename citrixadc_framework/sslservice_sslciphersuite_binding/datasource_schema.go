package sslservice_sslciphersuite_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func SslserviceSslciphersuiteBindingDataSourceSchema() schema.Schema {
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
			"servicename": schema.StringAttribute{
				Required:    true,
				Description: "Name of the SSL service for which to set advanced configuration.",
			},
		},
	}
}
