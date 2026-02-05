package autoscaleaction

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func AutoscaleactionDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "ActionScale action name.",
			},
			"parameters": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Parameters to use in the action",
			},
			"profilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "AutoScale profile name.",
			},
			"quiettime": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time in seconds no other policy is evaluated or action is taken",
			},
			"type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The type of action.",
			},
			"vmdestroygraceperiod": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time in minutes a VM is kept in inactive state before destroying",
			},
			"vserver": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the vserver on which autoscale action has to be taken.",
			},
		},
	}
}
