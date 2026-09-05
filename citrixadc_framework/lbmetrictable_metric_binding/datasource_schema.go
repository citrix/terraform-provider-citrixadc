package lbmetrictable_metric_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// LbmetrictableMetricBindingDataSourceModel is the data-source-specific model,
// decoupled from LbmetrictableMetricBindingResourceModel. A data source is a pure
// read surface, so it can expose the FULL GET projection: the configurable
// attributes (as Computed outputs) AND the read-only metadata attributes the
// resource deliberately omits.
type LbmetrictableMetricBindingDataSourceModel struct {
	Id          types.String `tfsdk:"id"`
	Snmpoid     types.String `tfsdk:"snmpoid"`
	Metric      types.String `tfsdk:"metric"`      // Required lookup key
	Metrictable types.String `tfsdk:"metrictable"` // Required lookup key

	// Read-only (GET-only) metadata from the NITRO doc read-only set
	// (zion73x_readonly/lbmetrictable_metric_binding.json).
	Metrictype types.String `tfsdk:"metrictype"`
}

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

			// Read-only (GET-only) metadata surfaced by the data source. All Computed.
			"metrictype": schema.StringAttribute{
				Computed:    true,
				Description: "Indication if it is a configured or internal. Possible values = INTERNAL, CONFIGURED",
			},
		},
	}
}

// lbmetrictable_metric_bindingDataSourceSetAttrFromGet projects a NITRO
// lbmetrictable_metric_binding GET response onto the data-source model.
func lbmetrictable_metric_bindingDataSourceSetAttrFromGet(ctx context.Context, data *LbmetrictableMetricBindingDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In lbmetrictable_metric_bindingDataSourceSetAttrFromGet Function")

	// Preserve the config-provided lookup keys when the GET omits them.
	metric := data.Metric
	metrictable := data.Metrictable

	if v, ok := g["metric"]; ok && v != nil {
		data.Metric = types.StringValue(utils.AnyToString(v))
	} else {
		data.Metric = metric
	}
	if v, ok := g["metrictable"]; ok && v != nil {
		data.Metrictable = types.StringValue(utils.AnyToString(v))
	} else {
		data.Metrictable = metrictable
	}

	// The NITRO GET response echoes the SNMP OID under the "Snmpoid" key.
	data.Snmpoid = utils.MapGetString(g, "Snmpoid")

	// Read-only (GET-only) metadata.
	data.Metrictype = utils.MapGetString(g, "metrictype")

	// Set the composite ID (metric:<v>,metrictable:<v>).
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("metric:%s", utils.UrlEncode(data.Metric.ValueString())))
	idParts = append(idParts, fmt.Sprintf("metrictable:%s", utils.UrlEncode(data.Metrictable.ValueString())))
	data.Id = types.StringValue(strings.Join(idParts, ","))
}
