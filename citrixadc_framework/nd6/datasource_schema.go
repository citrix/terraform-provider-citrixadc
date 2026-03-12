package nd6

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func Nd6DataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"ifnum": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Interface through which the adjacent network device is available, specified in slot/port notation (for example, 1/3). Use spaces to separate multiple entries.",
			},
			"mac": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "MAC address of the adjacent network device.",
			},
			"neighbor": schema.StringAttribute{
				Required:    true,
				Description: "Link-local IPv6 address of the adjacent network device to add to the ND6 table.",
			},
			"nodeid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Unique number that identifies the cluster node.",
			},
			"td": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.",
			},
			"vlan": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Integer value that uniquely identifies the VLAN on which the adjacent network device exists.",
			},
			"vtep": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address of the VXLAN tunnel endpoint (VTEP) through which the IPv6 address of this ND6 entry is reachable.",
			},
			"vxlan": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "ID of the VXLAN on which the IPv6 address of this ND6 entry is reachable.",
			},
		},
	}
}
