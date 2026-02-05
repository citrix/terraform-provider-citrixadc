package ipv6

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func Ipv6DataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"dodad": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable the Citrix ADC to do Duplicate Address\nDetection (DAD) for all the Citrix ADC owned IPv6 addresses regardless of whether they are obtained through stateless auto configuration, DHCPv6, or manual configuration.",
			},
			"natprefix": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Prefix used for translating packets from private IPv6 servers to IPv4 packets. This prefix has a length of 96 bits (128-32 = 96). The IPv6 servers embed the destination IP address of the IPv4 servers or hosts in the last 32 bits of the destination IP address field of the IPv6 packets. The first 96 bits of the destination IP address field are set as the IPv6 NAT prefix. IPv6 packets addressed to this prefix have to be routed to the Citrix ADC to ensure that the IPv6-IPv4 translation is done by the appliance.",
			},
			"ndbasereachtime": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Base reachable time of the Neighbor Discovery (ND6) protocol. The time, in milliseconds, that the Citrix ADC assumes an adjacent device is reachable after receiving a reachability confirmation.",
			},
			"ndretransmissiontime": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Retransmission time of the Neighbor Discovery (ND6) protocol. The time, in milliseconds, between retransmitted Neighbor Solicitation (NS) messages, to an adjacent device.",
			},
			"ralearning": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable the Citrix ADC to learn about various routes from Router Advertisement (RA) and Router Solicitation (RS) messages sent by the routers.",
			},
			"routerredirection": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable the Citrix ADC to do Router Redirection.",
			},
			"td": schema.Int64Attribute{
				Required:    true,
				Description: "Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.",
			},
			"usipnatprefix": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IPV6 NATPREFIX used in NAT46 scenario when USIP is turned on",
			},
		},
	}
}
