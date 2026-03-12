package vxlan_nsip_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func VxlanNsipBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"vxlanid": schema.Int64Attribute{
				Required:    true,
				Description: "A positive integer, which is also called VXLAN Network Identifier (VNI), that uniquely identifies a VXLAN.",
			},
			"ipaddress": schema.StringAttribute{
				Required:    true,
				Description: "The IP address assigned to the VXLAN.",
			},
			"netmask": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Subnet mask for the network address defined for this VXLAN.",
			},
		},
	}
}
