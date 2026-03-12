package sslcertkey_sslocspresponder_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func SslcertkeySslocspresponderBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"ca": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The certificate-key pair being unbound is a Certificate Authority (CA) certificate. If you choose this option, the certificate-key pair is unbound from the list of CA certificates that were bound to the specified SSL virtual server or SSL service.",
			},
			"certkey": schema.StringAttribute{
				Required:    true,
				Description: "Name of the certificate-key pair.",
			},
			"ocspresponder": schema.StringAttribute{
				Required:    true,
				Description: "OCSP responders bound to this certkey",
			},
			"priority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "ocsp priority",
			},
		},
	}
}
