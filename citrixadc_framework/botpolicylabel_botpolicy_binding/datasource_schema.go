package botpolicylabel_botpolicy_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func BotpolicylabelBotpolicyBindingDataSourceSchema() schema.Schema {
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
				Description: "If the current policy evaluates to TRUE, terminate evaluation of policies bound to the current policy label and evaluate the specified policy label.",
			},
			"invoke_labelname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "* If labelType is policylabel, name of the policy label to invoke. \n* If labelType is vserver, name of the virtual server.",
			},
			"labelname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the bot policy label to which to bind the policy.",
			},
			"labeltype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of policy label to invoke. Available settings function as follows:\n* vserver - Invoke an unnamed policy label associated with a virtual server.\n* policylabel - Invoke a user-defined policy label.",
			},
			"policyname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the bot policy.",
			},
			"priority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Specifies the priority of the policy.",
			},
		},
	}
}
