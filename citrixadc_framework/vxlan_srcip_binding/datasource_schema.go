package vxlan_srcip_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func VxlanSrcipBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"vxlanid": schema.Int64Attribute{
				Required:    true,
				Description: "A positive integer, which is also called VXLAN Network Identifier (VNI), that uniquely identifies a VXLAN.",
			},
			"srcip": schema.StringAttribute{
				Required:    true,
				Description: "The source IP address to use in outgoing vxlan packets.",
			},
		},
	}
}
