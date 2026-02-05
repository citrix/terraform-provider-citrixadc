package autoscalepolicy

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func AutoscalepolicyDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"action": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The autoscale profile associated with the policy.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Comments associated with this autoscale policy.",
			},
			"logaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The log action associated with the autoscale policy",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "The name of the autoscale policy.",
			},
			"newname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The new name of the autoscale policy.",
			},
			"rule": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The rule associated with the policy.",
			},
		},
	}
}
