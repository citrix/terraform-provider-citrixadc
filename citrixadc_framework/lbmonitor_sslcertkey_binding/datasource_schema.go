package lbmonitor_sslcertkey_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func LbmonitorSslcertkeyBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"ca": schema.BoolAttribute{
				Required:    true,
				Description: "The rule for use of CRL corresponding to this CA certificate during client authentication. If crlCheck is set to Mandatory, the system will deny all SSL clients if the CRL is missing, expired - NextUpdate date is in the past, or is incomplete with remote CRL refresh enabled. If crlCheck is set to optional, the system will allow SSL clients in the above error cases.However, in any case if the client certificate is revoked in the CRL, the SSL client will be denied access.",
			},
			"certkeyname": schema.StringAttribute{
				Required:    true,
				Description: "The name of the certificate bound to the monitor.",
			},
			"crlcheck": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The state of the CRL check parameter. (Mandatory/Optional)",
			},
			"monitorname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the monitor.",
			},
			"ocspcheck": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The state of the OCSP check parameter. (Mandatory/Optional)",
			},
		},
	}
}
