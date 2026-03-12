package gslbvserver_lbpolicy_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func GslbvserverLbpolicyBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"gotopriorityexpression": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.\n	o	If gotoPriorityExpression is not present or if it is equal to END then the policy bank evaluation ends here\n	o	Else if the gotoPriorityExpression is equal to NEXT then the next policy in the priority order is evaluated.\n	o	Else gotoPriorityExpression is evaluated. The result of gotoPriorityExpression (which has to be a number) is processed as follows:\n		-	An UNDEF event is triggered if\n			.	gotoPriorityExpression cannot be evaluated\n			.	gotoPriorityExpression evaluates to number which is smaller than the maximum priority in the policy bank but is not same as any policy's priority\n			.	gotoPriorityExpression evaluates to a priority that is smaller than the current policy's priority\n		-	If the gotoPriorityExpression evaluates to the priority of the current policy then the next policy in the priority order is evaluated.\n		-	If the gotoPriorityExpression evaluates to the priority of a policy further ahead in the list then that policy will be evaluated next.\n		This field is applicable only to rewrite and responder policies.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the virtual server on which to perform the binding operation.",
			},
			"order": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Order number to be assigned to the service when it is bound to the lb vserver.",
			},
			"policyname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the policy bound to the GSLB vserver.",
			},
			"priority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Priority.",
			},
			"type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The bindpoint to which the policy is bound",
			},
		},
	}
}
