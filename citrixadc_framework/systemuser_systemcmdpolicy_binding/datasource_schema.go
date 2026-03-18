package systemuser_systemcmdpolicy_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func SystemuserSystemcmdpolicyBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"policyname": schema.StringAttribute{
				Required:    true,
				Description: "The name of command policy.",
			},
			"priority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The priority of the policy.",
			},
			"username": schema.StringAttribute{
				Required:    true,
				Description: "Name of the system-user entry to which to bind the command policy.",
			},
		},
	}
}
