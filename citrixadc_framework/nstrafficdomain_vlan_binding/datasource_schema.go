package nstrafficdomain_vlan_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func NstrafficdomainVlanBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"td": schema.Int64Attribute{
				Required:    true,
				Description: "Integer value that uniquely identifies a traffic domain.",
			},
			"vlan": schema.Int64Attribute{
				Required:    true,
				Description: "ID of the VLAN to bind to this traffic domain. More than one VLAN can be bound to a traffic domain, but the same VLAN cannot be a part of multiple traffic domains.",
			},
		},
	}
}
