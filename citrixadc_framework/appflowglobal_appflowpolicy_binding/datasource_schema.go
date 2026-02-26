package appflowglobal_appflowpolicy_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func AppflowglobalAppflowpolicyBindingDataSourceSchema() schema.Schema {
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
			"invoke": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Invoke policies bound to a virtual server or a user-defined policy label. After the invoked policies are evaluated, the flow returns to the policy with the next priority.",
			},
			"labelname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the label to invoke if the current policy evaluates to TRUE.",
			},
			"labeltype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of policy label to invoke. Specify vserver for a policy label associated with a virtual server, or policylabel for a user-defined policy label.",
			},
			"policyname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the AppFlow policy.",
			},
			"priority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Specifies the priority of the policy.",
			},
			"type": schema.StringAttribute{
				Required:    true,
				Description: "Global bind point for which to show detailed information about the policies bound to the bind point.",
			},
		},
	}
}
