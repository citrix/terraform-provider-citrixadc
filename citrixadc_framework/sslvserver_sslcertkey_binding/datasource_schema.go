package sslvserver_sslcertkey_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func SslvserverSslcertkeyBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"ca": schema.BoolAttribute{
				Required:    true,
				Description: "CA certificate.",
			},
			"certkeyname": schema.StringAttribute{
				Required:    true,
				Description: "The name of the certificate key pair binding.",
			},
			"crlcheck": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The state of the CRL check parameter. (Mandatory/Optional)",
			},
			"ocspcheck": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The state of the OCSP check parameter. (Mandatory/Optional)",
			},
			"skipcaname": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The flag is used to indicate whether this particular CA certificate's CA_Name needs to be sent to the SSL client while requesting for client certificate in a SSL handshake",
			},
			"snicert": schema.BoolAttribute{
				Required:    true,
				Description: "The name of the CertKey. Use this option to bind Certkey(s) which will be used in SNI processing.",
			},
			"vservername": schema.StringAttribute{
				Required:    true,
				Description: "Name of the SSL virtual server.",
			},
		},
	}
}
