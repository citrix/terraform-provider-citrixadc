package sslcertificatechain

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func SslcertificatechainDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"certkeyname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the certificate-key pair.",
			},
		},
	}
}
