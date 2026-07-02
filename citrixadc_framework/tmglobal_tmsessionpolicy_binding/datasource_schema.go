package tmglobal_tmsessionpolicy_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func TmglobalTmsessionpolicyBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"feature": schema.StringAttribute{
				Computed:    true,
				Description: "The feature to be checked while applying this config",
			},
			"gotopriorityexpression": schema.StringAttribute{
				Computed:    true,
				Description: "Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.",
			},
			"policyname": schema.StringAttribute{
				Required:    true,
				Description: "The name of the policy.",
			},
			"priority": schema.Int64Attribute{
				Computed:    true,
				Description: "The priority of the policy.",
			},
		},
	}
}
