package feopolicy

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func FeopolicyDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"action": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The front end optimization action that has to be performed when the rule matches.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "The name of the front end optimization policy.",
			},
			"rule": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The rule associated with the front end optimization policy.",
			},
		},
	}
}
