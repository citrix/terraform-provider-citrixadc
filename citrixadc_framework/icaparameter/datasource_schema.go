package icaparameter

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func IcaparameterDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"dfpersistence": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable/Disable DF Persistence",
			},
			"edtlosstolerant": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable/Disable EDT Loss Tolerant feature",
			},
			"edtpmtuddf": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable/Disable DF enforcement for EDT PMTUD Control Blocks",
			},
			"edtpmtuddftimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "DF enforcement timeout for EDTPMTUDDF",
			},
			"edtpmtudrediscovery": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable/Disable EDT PMTUD Rediscovery",
			},
			"enablesronhafailover": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable/Disable Session Reliability on HA failover. The default value is No",
			},
			"hdxinsightnonnsap": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable/Disable HDXInsight for Non NSAP ICA Sessions. The default value is Yes",
			},
			"l7latencyfrequency": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify the time interval/period for which L7 Client Latency value is to be calculated. By default, L7 Client Latency is calculated for every packet. The default value is 0",
			},
		},
	}
}
