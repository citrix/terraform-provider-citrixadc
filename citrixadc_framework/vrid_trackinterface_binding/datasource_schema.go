package vrid_trackinterface_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func VridTrackinterfaceBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"vrid_id": schema.Int64Attribute{
				Required:    true,
				Description: "Integer that uniquely identifies the VMAC address. The generic VMAC address is in the form of 00:00:5e:00:01:<VRID>. For example, if you add a VRID with a value of 60 and bind it to an interface, the resulting VMAC address is 00:00:5e:00:01:3c, where 3c is the hexadecimal representation of 60.",
			},
			"trackifnum": schema.StringAttribute{
				Required:    true,
				Description: "Interfaces which need to be tracked for this vrID.",
			},
			"flags": schema.Int64Attribute{
				Computed:    true,
				Description: "Flags.",
			},
		},
	}
}
