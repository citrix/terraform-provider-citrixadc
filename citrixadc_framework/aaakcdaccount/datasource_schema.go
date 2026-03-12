package aaakcdaccount

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func AaakcdaccountDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"cacert": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "CA Cert for UserCert or when doing PKINIT backchannel.",
			},
			"delegateduser": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Username that can perform kerberos constrained delegation.",
			},
			"enterpriserealm": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enterprise Realm of the user. This should be given only in certain KDC deployments where KDC expects Enterprise username instead of Principal Name",
			},
			"kcdaccount": schema.StringAttribute{
				Required:    true,
				Description: "The name of the KCD account.",
			},
			"kcdpassword": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Password for Delegated User.",
			},
			"keytab": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The path to the keytab file. If specified other parameters in this command need not be given",
			},
			"realmstr": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Kerberos Realm.",
			},
			"servicespn": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Service SPN. When specified, this will be used to fetch kerberos tickets. If not specified, Citrix ADC will construct SPN using service fqdn",
			},
			"usercert": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "SSL Cert (including private key) for Delegated User.",
			},
			"userrealm": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Realm of the user",
			},
		},
	}
}
