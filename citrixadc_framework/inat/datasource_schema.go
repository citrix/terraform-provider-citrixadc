package inat

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func InatDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"connfailover": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Synchronize connection information with the secondary appliance in a high availability (HA) pair. That is, synchronize all connection-related information for the INAT session",
			},
			"ftp": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable the FTP protocol on the server for transferring files between the client and the server.",
			},
			"mode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Stateless translation.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the Inbound NAT (INAT) entry. Leading character must be a number or letter. Other characters allowed, after the first character, are @ _ - . (period) : (colon) # and space ( ).",
			},
			"privateip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address of the server to which the packet is sent by the Citrix ADC. Can be an IPv4 or IPv6 address.",
			},
			"proxyip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Unique IP address used as the source IP address in packets sent to the server. Must be a MIP or SNIP address.",
			},
			"publicip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Public IP address of packets received on the Citrix ADC. Can be aNetScaler-owned VIP or VIP6 address.",
			},
			"tcpproxy": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable TCP proxy, which enables the Citrix ADC to optimize the RNAT TCP traffic by using Layer 4 features.",
			},
			"td": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.",
			},
			"tftp": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "To enable/disable TFTP (Default DISABLED).",
			},
			"useproxyport": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable the Citrix ADC to proxy the source port of packets before sending the packets to the server.",
			},
			"usip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable the Citrix ADC to retain the source IP address of packets before sending the packets to the server.",
			},
			"usnip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable the Citrix ADC to use a SNIP address as the source IP address of packets before sending the packets to the server.",
			},
		},
	}
}
