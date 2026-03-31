package lbmetrictable_metric_binding

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

// LbmetrictableMetricBindingResourceModel describes the resource data model.
type LbmetrictableMetricBindingResourceModel struct {
	Id          types.String `tfsdk:"id"`
	Snmpoid     types.String `tfsdk:"snmpoid"`
	Metric      types.String `tfsdk:"metric"`
	Metrictable types.String `tfsdk:"metrictable"`
}

func (r *LbmetrictableMetricBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the lbmetrictable_metric_binding resource.",
			},
			"snmpoid": schema.StringAttribute{
				Required:    true,
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

func lbmetrictable_metric_bindingGetThePayloadFromtheConfig(ctx context.Context, data *LbmetrictableMetricBindingResourceModel) lb.Lbmetrictablemetricbinding {
	tflog.Debug(ctx, "In lbmetrictable_metric_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	lbmetrictable_metric_binding := lb.Lbmetrictablemetricbinding{}
	if !data.Snmpoid.IsNull() {
		lbmetrictable_metric_binding.Snmpoid = data.Snmpoid.ValueString()
	}
	if !data.Metric.IsNull() {
		lbmetrictable_metric_binding.Metric = data.Metric.ValueString()
	}
	if !data.Metrictable.IsNull() {
		lbmetrictable_metric_binding.Metrictable = data.Metrictable.ValueString()
	}

	return lbmetrictable_metric_binding
}

func lbmetrictable_metric_bindingSetAttrFromGet(ctx context.Context, data *LbmetrictableMetricBindingResourceModel, getResponseData map[string]interface{}) *LbmetrictableMetricBindingResourceModel {
	tflog.Debug(ctx, "In lbmetrictable_metric_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["Snmpoid"]; ok && val != nil {
		data.Snmpoid = types.StringValue(val.(string))
	} else {
		data.Snmpoid = types.StringNull()
	}
	if val, ok := getResponseData["metric"]; ok && val != nil {
		data.Metric = types.StringValue(val.(string))
	} else {
		data.Metric = types.StringNull()
	}
	if val, ok := getResponseData["metrictable"]; ok && val != nil {
		data.Metrictable = types.StringValue(val.(string))
	} else {
		data.Metrictable = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("metric:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Metric.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("metrictable:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Metrictable.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
