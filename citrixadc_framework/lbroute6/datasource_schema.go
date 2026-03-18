package lbroute6

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func Lbroute6DataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"gatewayname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The name of the route.",
			},
			"network": schema.StringAttribute{
				Required:    true,
				Description: "The destination network.",
			},
			"td": schema.Int64Attribute{
				Required:    true,
				Description: "Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.",
			},
		},
	}
}
