package arp

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func ArpDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"all": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Remove all ARP entries from the ARP table of the Citrix ADC.",
			},
			"ifnum": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Interface through which the network device is accessible. Specify the interface in (slot/port) notation. For example, 1/3.",
			},
			"ipaddress": schema.StringAttribute{
				Required:    true,
				Description: "IP address of the network device that you want to add to the ARP table.",
			},
			"mac": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "MAC address of the network device.",
			},
			"nodeid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Unique number that identifies the cluster node.",
			},
			"ownernode": schema.Int64Attribute{
				Required:    true,
				Description: "The owner node for the Arp entry.",
			},
			"td": schema.Int64Attribute{
				Required:    true,
				Description: "Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.",
			},
			"vlan": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The VLAN ID through which packets are to be sent after matching the ARP entry. This is a numeric value.",
			},
			"vtep": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address of the VXLAN tunnel endpoint (VTEP) through which the IP address of this ARP entry is reachable.",
			},
			"vxlan": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "ID of the VXLAN on which the IP address of this ARP entry is reachable.",
			},
		},
	}
}
