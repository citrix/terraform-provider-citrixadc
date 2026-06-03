package videooptimizationparameter

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func VideooptimizationparameterDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"quicpacingrate": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "QUIC Video Pacing Rate (Kbps).",
			},
			"randomsamplingpercentage": schema.Float64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Random Sampling Percentage.",
			},
		},
	}
}