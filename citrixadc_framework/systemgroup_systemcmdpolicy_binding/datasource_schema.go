package systemgroup_systemcmdpolicy_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func SystemgroupSystemcmdpolicyBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"groupname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the system group.",
			},
			"policyname": schema.StringAttribute{
				Required:    true,
				Description: "The name of command policy.",
			},
			"priority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The priority of the command policy.",
			},
		},
	}
}
