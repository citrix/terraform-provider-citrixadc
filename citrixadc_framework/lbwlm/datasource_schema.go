package lbwlm

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func LbwlmDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"ipaddress": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The IP address of the WLM.",
			},
			"katimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The idle time period after which Citrix ADC would probe the WLM. The value ranges from 1 to 1440 minutes.",
			},
			"lbuid": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The LBUID for the Load Balancer to communicate to the Work Load Manager.",
			},
			"port": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The port of the WLM.",
			},
			"wlmname": schema.StringAttribute{
				Required:    true,
				Description: "The name of the Work Load Manager.",
			},
		},
	}
}
