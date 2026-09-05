package cloudtrafficroutes

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func CloudtrafficroutesDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the traffic cloud route.",
			},
			"targetvpcnetwork": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Target VPC network name.",
			},
			"destrange": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Destination IP range in CIDR format.",
			},
			"nexthopip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Next hop IP address.",
			},
			"ownernode": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "cluster owner node id for the nexthopipaddress.",
			},
		},
	}
}
