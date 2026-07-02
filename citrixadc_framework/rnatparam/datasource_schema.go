package rnatparam

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func RnatparamDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"srcippersistency": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable source ip persistency, which enables the Citrix ADC to use the RNAT ips using source ip.",
			},
			"tcpproxy": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable TCP proxy, which enables the Citrix ADC to optimize the RNAT TCP traffic by using Layer 4 features.",
			},
		},
	}
}
