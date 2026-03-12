package iptunnel

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func IptunnelDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"destport": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Specifies UDP destination port for Geneve packets. Default port is 6081.",
			},
			"grepayload": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The payload GRE will carry",
			},
			"ipsecprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of IPSec profile to be associated.",
			},
			"local": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of Citrix ADC owned public IPv4 address, configured on the local Citrix ADC and used to set up the tunnel.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the IP tunnel. Leading character must be a number or letter. Other characters allowed, after the first character, are @ _ - . (period) : (colon) # and space ( ).",
			},
			"ownergroup": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The owner node group in a Cluster for the iptunnel.",
			},
			"protocol": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the protocol to be used on this tunnel.",
			},
			"remote": schema.StringAttribute{
				Required:    true,
				Description: "Public IPv4 address, of the remote device, used to set up the tunnel. For this parameter, you can alternatively specify a network address.",
			},
			"remotesubnetmask": schema.StringAttribute{
				Required:    true,
				Description: "Subnet mask of the remote IP address of the tunnel.",
			},
			"tosinherit": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Default behavior is to copy the ToS field of the internal IP Packet (Payload) to the outer IP packet (Transport packet). But the user can configure a new ToS field using this option.",
			},
			"vlan": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The vlan for mulicast packets",
			},
			"vlantagging": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Option to select Vlan Tagging.",
			},
			"vnid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Virtual network identifier (VNID) is the value that identifies a specific virtual network in the data plane.",
			},
		},
	}
}
