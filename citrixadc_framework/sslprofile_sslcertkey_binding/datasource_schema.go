package sslprofile_sslcertkey_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func SslprofileSslcertkeyBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"cipherpriority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Priority of the cipher binding",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the SSL profile.",
			},
			"sslicacertkey": schema.StringAttribute{
				Required:    true,
				Description: "The certkey (CA certificate + private key) to be used for SSL interception.",
			},
		},
	}
}
