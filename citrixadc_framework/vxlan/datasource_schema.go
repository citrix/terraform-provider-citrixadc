package vxlan

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func VxlanDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"dynamicrouting": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable dynamic routing on this VXLAN.",
			},
			"vxlanid": schema.Int64Attribute{
				Required:    true,
				Description: "A positive integer, which is also called VXLAN Network Identifier (VNI), that uniquely identifies a VXLAN.",
			},
			"innervlantagging": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specifies whether Citrix ADC should generate VXLAN packets with inner VLAN tag.",
			},
			"ipv6dynamicrouting": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable all IPv6 dynamic routing protocols on this VXLAN. Note: For the ENABLED setting to work, you must configure IPv6 dynamic routing protocols from the VTYSH command line.",
			},
			"port": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Specifies UDP destination port for VXLAN packets.",
			},
			"protocol": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "VXLAN-GPE next protocol. RESERVED, IPv4, IPv6, ETHERNET, NSH",
			},
			"type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "VXLAN encapsulation type. VXLAN, VXLANGPE",
			},
			"vlan": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "ID of VLANs whose traffic is allowed over this VXLAN. If you do not specify any VLAN IDs, the Citrix ADC allows traffic of all VLANs that are not part of any other VXLANs.",
			},
		},
	}
}
