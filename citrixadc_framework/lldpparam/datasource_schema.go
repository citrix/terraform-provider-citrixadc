package lldpparam

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func LldpparamDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"holdtimetxmult": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "A multiplier for calculating the duration for which the receiving device stores the LLDP information in its database before discarding or removing it. The duration is calculated as the holdtimeTxMult (Holdtime Multiplier) parameter value multiplied by the timer (Timer) parameter value.",
			},
			"mode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Global mode of Link Layer Discovery Protocol (LLDP) on the Citrix ADC. The resultant LLDP mode of an interface depends on the LLDP mode configured at the global and the interface levels.",
			},
			"timer": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Interval, in seconds, between LLDP packet data units (LLDPDUs).  that the Citrix ADC sends to a directly connected device.",
			},
		},
	}
}
