package vrid_channel_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func VridChannelBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"vrid_id": schema.Int64Attribute{
				Required:    true,
				Description: "Integer that uniquely identifies the VMAC address. The generic VMAC address is in the form of 00:00:5e:00:01:<VRID>. For example, if you add a VRID with a value of 60 and bind it to an interface, the resulting VMAC address is 00:00:5e:00:01:3c, where 3c is the hexadecimal representation of 60.",
			},
			"ifnum": schema.StringAttribute{
				Required:    true,
				Description: "Interfaces to bind to the VMAC, specified in (slot/port) notation (for example, 1/2).Use spaces to separate multiple entries.",
			},
			"flags": schema.Int64Attribute{
				Computed:    true,
				Description: "Flags.",
			},
			"vlan": schema.Int64Attribute{
				Computed:    true,
				Description: "The VLAN in which this VRID resides.",
			},
		},
	}
}
