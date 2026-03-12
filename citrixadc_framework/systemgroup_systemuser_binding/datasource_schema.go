package systemgroup_systemuser_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func SystemgroupSystemuserBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"groupname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the system group.",
			},
			"username": schema.StringAttribute{
				Required:    true,
				Description: "The system user.",
			},
		},
	}
}
