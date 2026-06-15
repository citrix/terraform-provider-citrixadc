package lbmetrictable_metric_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func LbmetrictableMetricBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"snmpoid": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "New SNMP OID of the metric.",
			},
			"metric": schema.StringAttribute{
				Required:    true,
				Description: "Name of the metric for which to change the SNMP OID.",
			},
			"metrictable": schema.StringAttribute{
				Required:    true,
				Description: "Name of the metric table.",
			},
		},
	}
}
