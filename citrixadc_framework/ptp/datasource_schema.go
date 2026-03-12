package ptp

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func PtpDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"state": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enables or disables Precision Time Protocol (PTP) on the appliance. If you disable PTP, make sure you enable Network Time Protocol (NTP) on the cluster.",
			},
		},
	}
}
