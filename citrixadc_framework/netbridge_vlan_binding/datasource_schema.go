package netbridge_vlan_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func NetbridgeVlanBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "The name of the network bridge.",
			},
			"vlan": schema.Int64Attribute{
				Required:    true,
				Description: "The VLAN that is extended by this network bridge.",
			},
		},
	}
}
