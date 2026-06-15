package bridgegroup_vlan_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func BridgegroupVlanBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"bridgegroup_id": schema.Int64Attribute{
				Required:    true,
				Description: "The integer that uniquely identifies the bridge group.",
			},
			"vlan": schema.Int64Attribute{
				Required:    true,
				Description: "Names of all member VLANs.",
			},
		},
	}
}
