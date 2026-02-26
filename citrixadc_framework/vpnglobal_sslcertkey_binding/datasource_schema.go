package vpnglobal_sslcertkey_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func VpnglobalSslcertkeyBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"cacert": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The name of the CA certificate binding.",
			},
			"certkeyname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "SSL certkey to use in signing tokens. Only RSA cert key is allowed",
			},
			"crlcheck": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The state of the CRL check parameter (Mandatory/Optional).",
			},
			"gotopriorityexpression": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Applicable only to advance vpn session policy. An expression or other value specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.",
			},
			"ocspcheck": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The state of the OCSP check parameter (Mandatory/Optional).",
			},
			"userdataencryptionkey": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Certificate to be used for encrypting user data like KB Question and Answers, Alternate Email Address, etc.",
			},
		},
	}
}
