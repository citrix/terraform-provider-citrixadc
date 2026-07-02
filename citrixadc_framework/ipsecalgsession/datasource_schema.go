package ipsecalgsession

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func IpsecalgsessionDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"destip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Destination IP address.",
			},
			"destip_alg": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Destination IP address.",
			},
			"natip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Natted Source IP address.",
			},
			"natip_alg": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Natted Source IP address.",
			},
			"sourceip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Original Source IP address.",
			},
			"sourceip_alg": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Original Source IP address.",
			},
		},
	}
}
