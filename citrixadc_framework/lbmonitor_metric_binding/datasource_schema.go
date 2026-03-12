package lbmonitor_metric_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func LbmonitorMetricBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"metric": schema.StringAttribute{
				Required:    true,
				Description: "Metric name in the metric table, whose setting is changed. A value zero disables the metric and it will not be used for load calculation",
			},
			"metricthreshold": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Threshold to be used for that metric.",
			},
			"metricweight": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The weight for the specified service metric with respect to others.",
			},
			"monitorname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the monitor.",
			},
		},
	}
}
