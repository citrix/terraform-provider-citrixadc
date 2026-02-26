package vlan_interface_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func VlanInterfaceBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"vlanid": schema.Int64Attribute{
				Required:    true,
				Description: "Specifies the virtual LAN ID.",
			},
			"ifnum": schema.StringAttribute{
				Required:    true,
				Description: "The interface to be bound to the VLAN, specified in slot/port notation (for example, 1/3).",
			},
			"ownergroup": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The owner node group in a Cluster for this vlan.",
			},
			"tagged": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Make the interface an 802.1q tagged interface. Packets sent on this interface on this VLAN have an additional 4-byte 802.1q tag, which identifies the VLAN. To use 802.1q tagging, you must also configure the switch connected to the appliance's interfaces.",
			},
		},
	}
}
