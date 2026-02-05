package sslhsmkey

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func SslhsmkeyDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"hsmkeyname": schema.StringAttribute{
				Required:    true,
				Description: "0",
			},
			"hsmtype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of HSM.",
			},
			"key": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the key. optionally, for Thales, path to the HSM key file; /var/opt/nfast/kmdata/local/ is the default path. Applies when HSMTYPE is THALES or KEYVAULT.",
			},
			"keystore": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of keystore object representing HSM where key is stored. For example, name of keyvault object or azurekeyvault authentication object. Applies only to KEYVAULT type HSM.",
			},
			"password": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Password for a partition. Applies only to SafeNet HSM.",
			},
			"serialnum": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Serial number of the partition on which the key is present. Applies only to SafeNet HSM.",
			},
		},
	}
}
