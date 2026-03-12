package vxlanvlanmap_vxlan_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func VxlanvlanmapVxlanBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the mapping table.",
			},
			"vlan": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "The vlan id or the range of vlan ids in the on-premise network.",
			},
			"vxlan": schema.Int64Attribute{
				Required:    true,
				Description: "The VXLAN assigned to the vlan inside the cloud.",
			},
		},
	}
}
