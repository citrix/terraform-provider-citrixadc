package cloudroutes

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func CloudroutesDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the route.",
			},
			"routesvpcnetwork": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "client vpc network name.",
			},
			"vipsubnet": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "vip subnet in CIDR format.",
			},
			"vipvpcnetwork": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "vip vpc network name.",
			},
			"clientipaddress": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IPv4 or IPv6 address attached to the nic interface towards vpc mentiond in vpcnetwork.",
			},
		},
	}
}
