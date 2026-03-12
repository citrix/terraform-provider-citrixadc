package cacheforwardproxy

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func CacheforwardproxyDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"ipaddress": schema.StringAttribute{
				Required:    true,
				Description: "IP address of the Citrix ADC or a cache server for which the cache acts as a proxy. Requests coming to the Citrix ADC with the configured IP address are forwarded to the particular address, without involving the Integrated Cache in any way.",
			},
			"port": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Port on the Citrix ADC or a server for which the cache acts as a proxy",
			},
		},
	}
}
