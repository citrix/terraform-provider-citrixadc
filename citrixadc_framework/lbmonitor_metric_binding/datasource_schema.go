package lbmonitor_metric_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// LbmonitorMetricBindingDataSourceModel is the data-source-specific model,
// decoupled from LbmonitorMetricBindingResourceModel. A data source is a pure
// read surface, so it can expose the FULL GET projection: the configurable
// attributes (as Computed outputs) AND the read-only metadata attributes the
// resource deliberately omits.
type LbmonitorMetricBindingDataSourceModel struct {
	Id              types.String `tfsdk:"id"`
	Metric          types.String `tfsdk:"metric"` // Required lookup key
	Metricthreshold types.Int64  `tfsdk:"metricthreshold"`
	Metricweight    types.Int64  `tfsdk:"metricweight"`
	Monitorname     types.String `tfsdk:"monitorname"` // Required lookup key

	// Read-only (GET-only) metadata from the NITRO doc read-only set
	// (zion73x_readonly/lbmonitor_metric_binding.json).
	MetricUnit  types.String `tfsdk:"metric_unit"`
	Metrictable types.String `tfsdk:"metrictable"`
}

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

			// Read-only (GET-only) metadata surfaced by the data source. All Computed.
			"metric_unit": schema.StringAttribute{
				Computed:    true,
				Description: "Giving the unit of the metric. Possible values = Bytes/s, ms, pkts/s, users",
			},
			"metrictable": schema.StringAttribute{
				Computed:    true,
				Description: "Metric table to which to bind metrics.",
			},
		},
	}
}

// lbmonitor_metric_bindingDataSourceSetAttrFromGet projects a NITRO
// lbmonitor_metric_binding GET response onto the data-source model.
func lbmonitor_metric_bindingDataSourceSetAttrFromGet(ctx context.Context, data *LbmonitorMetricBindingDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In lbmonitor_metric_bindingDataSourceSetAttrFromGet Function")

	// Preserve the config-provided lookup keys when the GET omits them.
	metric := data.Metric
	monitorname := data.Monitorname

	if v, ok := g["metric"]; ok && v != nil {
		data.Metric = types.StringValue(utils.AnyToString(v))
	} else {
		data.Metric = metric
	}
	if v, ok := g["monitorname"]; ok && v != nil {
		data.Monitorname = types.StringValue(utils.AnyToString(v))
	} else {
		data.Monitorname = monitorname
	}

	data.Metricthreshold = utils.MapGetInt64(g, "metricthreshold")
	data.Metricweight = utils.MapGetInt64(g, "metricweight")

	// Read-only (GET-only) metadata.
	data.MetricUnit = utils.MapGetString(g, "metric_unit")
	data.Metrictable = utils.MapGetString(g, "metrictable")

	// Set the composite ID (metric:<v>,monitorname:<v>).
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("metric:%s", utils.UrlEncode(data.Metric.ValueString())))
	idParts = append(idParts, fmt.Sprintf("monitorname:%s", utils.UrlEncode(data.Monitorname.ValueString())))
	data.Id = types.StringValue(strings.Join(idParts, ","))
}
