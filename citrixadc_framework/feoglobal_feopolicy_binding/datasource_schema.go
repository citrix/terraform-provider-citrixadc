package feoglobal_feopolicy_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func FeoglobalFeopolicyBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"globalbindtype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "0",
			},
			"gotopriorityexpression": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.",
			},
			"policyname": schema.StringAttribute{
				Required:    true,
				Description: "The name of the globally bound front end optimization policy.",
			},
			"priority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The priority assigned to the policy binding.",
			},
			"type": schema.StringAttribute{
				Required:    true,
				Description: "Bindpoint to which the policy is bound.",
			},
		},
	}
}
