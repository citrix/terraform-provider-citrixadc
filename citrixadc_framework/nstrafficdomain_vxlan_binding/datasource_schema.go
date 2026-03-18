package nstrafficdomain_vxlan_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func NstrafficdomainVxlanBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"td": schema.Int64Attribute{
				Required:    true,
				Description: "Integer value that uniquely identifies a traffic domain.",
			},
			"vxlan": schema.Int64Attribute{
				Required:    true,
				Description: "ID of the VXLAN to bind to this traffic domain. More than one VXLAN can be bound to a traffic domain, but the same VXLAN cannot be a part of multiple traffic domains.",
			},
		},
	}
}
