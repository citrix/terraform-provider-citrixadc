package nslicenseproxyserver

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func NslicenseproxyserverDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"port": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "License proxy server port.",
			},
			"serverip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address of the License proxy server. Either serverip or servername must be specified.",
			},
			"servername": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Fully qualified domain name of the License proxy server. Either serverip or servername must be specified.",
			},
		},
	}
}
