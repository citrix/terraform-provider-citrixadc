package lbvserver_appqoepolicy_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func LbvserverAppqoepolicyBindingDataSourceSchema() schema.Schema {
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
			"invoke": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Invoke policies bound to a virtual server or policy label.",
			},
			"labelname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the virtual server or user-defined policy label to invoke if the policy evaluates to TRUE.",
			},
			"labeltype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of policy label to invoke. Applicable only to rewrite, videooptimization and cache policies. Available settings function as follows:\n* reqvserver - Evaluate the request against the request-based policies bound to the specified virtual server.\n* resvserver - Evaluate the response against the response-based policies bound to the specified virtual server.\n* policylabel - invoke the request or response against the specified user-defined policy label.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the virtual server. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Can be changed after the virtual server is created.\n\nCLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my vserver\" or 'my vserver').",
			},
			"order": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Integer specifying the order of the service. A larger number specifies a lower order. Defines the order of the service relative to the other services in the load balancing vserver's bindings. Determines the priority given to the service among all the services bound.",
			},
			"policyname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the policy bound to the LB vserver.",
			},
			"priority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Priority.",
			},
		},
	}
}
