package lbmonitor_metric_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/lb"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// LbmonitorMetricBindingResourceModel describes the resource data model.
type LbmonitorMetricBindingResourceModel struct {
	Id              types.String `tfsdk:"id"`
	Metric          types.String `tfsdk:"metric"`
	Metricthreshold types.Int64  `tfsdk:"metricthreshold"`
	Metricweight    types.Int64  `tfsdk:"metricweight"`
	Monitorname     types.String `tfsdk:"monitorname"`
}

func (r *LbmonitorMetricBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the lbmonitor_metric_binding resource.",
			},
			"metric": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
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

func lbmonitor_metric_bindingGetThePayloadFromtheConfig(ctx context.Context, data *LbmonitorMetricBindingResourceModel) lb.Lbmonitormetricbinding {
	tflog.Debug(ctx, "In lbmonitor_metric_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	lbmonitor_metric_binding := lb.Lbmonitormetricbinding{}
	if !data.Metric.IsNull() {
		lbmonitor_metric_binding.Metric = data.Metric.ValueString()
	}
	if !data.Metricthreshold.IsNull() {
		lbmonitor_metric_binding.Metricthreshold = utils.IntPtr(int(data.Metricthreshold.ValueInt64()))
	}
	if !data.Metricweight.IsNull() {
		lbmonitor_metric_binding.Metricweight = utils.IntPtr(int(data.Metricweight.ValueInt64()))
	}
	if !data.Monitorname.IsNull() {
		lbmonitor_metric_binding.Monitorname = data.Monitorname.ValueString()
	}

	return lbmonitor_metric_binding
}

func lbmonitor_metric_bindingSetAttrFromGet(ctx context.Context, data *LbmonitorMetricBindingResourceModel, getResponseData map[string]interface{}) *LbmonitorMetricBindingResourceModel {
	tflog.Debug(ctx, "In lbmonitor_metric_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["metric"]; ok && val != nil {
		data.Metric = types.StringValue(val.(string))
	} else {
		data.Metric = types.StringNull()
	}
	if val, ok := getResponseData["metricthreshold"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Metricthreshold = types.Int64Value(intVal)
		}
	} else {
		data.Metricthreshold = types.Int64Null()
	}
	if val, ok := getResponseData["metricweight"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Metricweight = types.Int64Value(intVal)
		}
	} else {
		data.Metricweight = types.Int64Null()
	}
	if val, ok := getResponseData["monitorname"]; ok && val != nil {
		data.Monitorname = types.StringValue(val.(string))
	} else {
		data.Monitorname = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("metric:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Metric.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("monitorname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Monitorname.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
