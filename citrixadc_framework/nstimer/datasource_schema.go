package nstimer

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func NstimerDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Comments associated with this timer.",
			},
			"interval": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The frequency at which the policies bound to this timer are invoked. The minimum value is 20 msec. The maximum value is 20940 in seconds and 349 in minutes",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Timer name.",
			},
			"newname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The new name of the timer.",
			},
			"unit": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Timer interval unit",
			},
		},
	}
}
