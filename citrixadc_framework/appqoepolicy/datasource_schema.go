package appqoepolicy

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func AppqoepolicyDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"action": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Configured AppQoE action to trigger",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "0",
			},
			"rule": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression or name of a named expression, against which the request is evaluated. The policy is applied if the rule evaluates to true.",
			},
		},
	}
}
