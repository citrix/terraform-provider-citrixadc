package sslservicegroup_sslciphersuite_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func SslservicegroupSslciphersuiteBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"ciphername": schema.StringAttribute{
				Required:    true,
				Description: "The name of the cipher group/alias/name configured for the SSL service group.",
			},
			"description": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The description of the cipher.",
			},
			"servicegroupname": schema.StringAttribute{
				Required:    true,
				Description: "The name of the SSL service to which the SSL policy needs to be bound.",
			},
		},
	}
}
