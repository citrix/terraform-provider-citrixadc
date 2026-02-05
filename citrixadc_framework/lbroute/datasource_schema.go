package lbroute

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func LbrouteDataSourceSchema() schema.Schema {
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
			"netmask": schema.StringAttribute{
				Required:    true,
				Description: "The netmask to which the route belongs.",
			},
			"network": schema.StringAttribute{
				Required:    true,
				Description: "The IP address of the network to which the route belongs.",
			},
			"td": schema.Int64Attribute{
				Required:    true,
				Description: "Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.",
			},
		},
	}
}
