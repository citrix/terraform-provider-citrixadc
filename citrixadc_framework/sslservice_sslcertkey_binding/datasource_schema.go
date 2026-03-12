package sslservice_sslcertkey_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func SslserviceSslcertkeyBindingDataSourceSchema() schema.Schema {
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
				Description: "The certificate key pair binding.",
			},
			"crlcheck": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The state of the CRL check parameter. (Mandatory/Optional)",
			},
			"ocspcheck": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Rule to use for the OCSP responder associated with the CA certificate during client authentication. If MANDATORY is specified, deny all SSL clients if the OCSP check fails because of connectivity issues with the remote OCSP server, or any other reason that prevents the OCSP check. With the OPTIONAL setting, allow SSL clients even if the OCSP check fails except when the client certificate is revoked.",
			},
			"servicename": schema.StringAttribute{
				Required:    true,
				Description: "Name of the SSL service for which to set advanced configuration.",
			},
			"skipcaname": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The flag is used to indicate whether this particular CA certificate's CA_Name needs to be sent to the SSL client while requesting      for client certificate in a SSL handshake",
			},
			"snicert": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The name of the CertKey. Use this option to bind Certkey(s) which will be used in SNI processing.",
			},
		},
	}
}
