package vlan_nsip_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func VlanNsipBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"vlanid": schema.Int64Attribute{
				Required:    true,
				Description: "Specifies the virtual LAN ID.",
			},
			"ipaddress": schema.StringAttribute{
				Required:    true,
				Description: "The IP address assigned to the VLAN.",
			},
			"netmask": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Subnet mask for the network address defined for this VLAN.",
			},
			"ownergroup": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The owner node group in a Cluster for this vlan.",
			},
			"td": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.",
			},
		},
	}
}
