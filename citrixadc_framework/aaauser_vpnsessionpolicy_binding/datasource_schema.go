package aaauser_vpnsessionpolicy_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func AaauserVpnsessionpolicyBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"gotopriorityexpression": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.",
			},
			"policy": schema.StringAttribute{
				Required:    true,
				Description: "The policy Name.",
			},
			"priority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Integer specifying the priority of the policy.  A lower number indicates a higher priority. Policies are evaluated in the order of their priority numbers. Maximum value for default syntax policies is 2147483647 and for classic policies max priority is 64000.",
			},
			"type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Bindpoint to which the policy is bound.",
			},
			"username": schema.StringAttribute{
				Required:    true,
				Description: "User account to which to bind the policy.",
			},
		},
	}
}
