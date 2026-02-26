package systemgroup_nspartition_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func SystemgroupNspartitionBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"groupname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the system group.",
			},
			"partitionname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the Partition to bind to the system group.",
			},
		},
	}
}
