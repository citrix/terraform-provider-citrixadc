package systemuser_nspartition_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func SystemuserNspartitionBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"partitionname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the Partition to bind to the system user.",
			},
			"username": schema.StringAttribute{
				Required:    true,
				Description: "Name of the system-user entry to which to bind the command policy.",
			},
		},
	}
}
