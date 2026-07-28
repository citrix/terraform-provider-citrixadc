package sslservicegroup_sslcipher_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func SslservicegroupSslcipherBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"cipheraliasname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The name of the cipher group/alias/name configured for the SSL service group.",
			},
			"ciphername": schema.StringAttribute{
				Required:    true,
				Description: "A cipher-suite can consist of an individual cipher name, the system predefined cipher-alias name, or user defined cipher-group name.",
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
