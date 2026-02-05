package dnsaction64

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func Dnsaction64DataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"actionname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the dns64 action.",
			},
			"excluderule": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The expression to select the criteria for eliminating the corresponding ipv6 addresses from the response.",
			},
			"mappedrule": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The expression to select the criteria for ipv4 addresses to be used for synthesis.\n                      Only if the mappedrule is evaluated to true the corresponding ipv4 address is used for synthesis using respective prefix,\n                      otherwise the A RR is discarded",
			},
			"prefix": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The dns64 prefix to be used if the after evaluating the rules",
			},
		},
	}
}
