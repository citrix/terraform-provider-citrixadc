package vxlanvlanmap

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func VxlanvlanmapDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the mapping table.",
			},
		},
	}
}
