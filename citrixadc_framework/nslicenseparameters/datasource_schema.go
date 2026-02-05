package nslicenseparameters

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func NslicenseparametersDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"alert1gracetimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "If ADC remains in grace for the configured hours then first grace alert will be raised",
			},
			"alert2gracetimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "If ADC remains in grace for the configured hours then major grace alert will be raised",
			},
			"heartbeatinterval": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Heartbeat between ADC and Licenseserver is configurable and applicable in case of pooled licensing",
			},
			"inventoryrefreshinterval": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Inventory refresh interval between ADC and Licenseserver is configurable and applicable in case of pooled licensing",
			},
			"licenseexpiryalerttime": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "If ADC license contract expiry date is nearer then GUI/SNMP license expiry alert will be raised",
			},
		},
	}
}
