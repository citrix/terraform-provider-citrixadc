package inatparam

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func InatparamDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"nat46fragheader": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "When disabled, translator will not insert IPv6 fragmentation header for non fragmented IPv4 packets",
			},
			"nat46ignoretos": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Ignore TOS.",
			},
			"nat46v6mtu": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "MTU setting for the IPv6 side. If the incoming IPv4 packet greater than this, either fragment or send icmp need fragmentation error.",
			},
			"nat46v6prefix": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The prefix used for translating packets received from private IPv6 servers into IPv4 packets. This prefix has a length of 96 bits (128-32 = 96). The IPv6 servers embed the destination IP address of the IPv4 servers or hosts in the last 32 bits of the destination IP address field of the IPv6 packets. The first 96 bits of the destination IP address field are set as the IPv6 NAT prefix. IPv6 packets addressed to this prefix have to be routed to the Citrix ADC to ensure that the IPv6-IPv4 translation is done by the appliance.",
			},
			"nat46zerochecksum": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Calculate checksum for UDP packets with zero checksum",
			},
			"td": schema.Int64Attribute{
				Required:    true,
				Description: "Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.",
			},
		},
	}
}
