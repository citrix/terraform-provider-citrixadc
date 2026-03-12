package gslbservice_lbmonitor_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/gslb"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// GslbserviceLbmonitorBindingResourceModel describes the resource data model.
type GslbserviceLbmonitorBindingResourceModel struct {
	Id          types.String `tfsdk:"id"`
	MonitorName types.String `tfsdk:"monitor_name"`
	Monstate    types.String `tfsdk:"monstate"`
	Servicename types.String `tfsdk:"servicename"`
	Weight      types.Int64  `tfsdk:"weight"`
}

func (r *GslbserviceLbmonitorBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the gslbservice_lbmonitor_binding resource.",
			},
			"monitor_name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Monitor name.",
			},
			"monstate": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State of the monitor bound to gslb service.",
			},
			"servicename": schema.StringAttribute{
				Required:    true,
				Description: "Name of the GSLB service.",
			},
			"weight": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Weight to assign to the monitor-service binding. A larger number specifies a greater weight. Contributes to the monitoring threshold, which determines the state of the service.",
			},
		},
	}
}

func gslbservice_lbmonitor_bindingGetThePayloadFromtheConfig(ctx context.Context, data *GslbserviceLbmonitorBindingResourceModel) gslb.Gslbservicelbmonitorbinding {
	tflog.Debug(ctx, "In gslbservice_lbmonitor_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	gslbservice_lbmonitor_binding := gslb.Gslbservicelbmonitorbinding{}
	if !data.MonitorName.IsNull() {
		gslbservice_lbmonitor_binding.Monitorname = data.MonitorName.ValueString()
	}
	if !data.Monstate.IsNull() {
		gslbservice_lbmonitor_binding.Monstate = data.Monstate.ValueString()
	}
	if !data.Servicename.IsNull() {
		gslbservice_lbmonitor_binding.Servicename = data.Servicename.ValueString()
	}
	if !data.Weight.IsNull() {
		gslbservice_lbmonitor_binding.Weight = utils.IntPtr(int(data.Weight.ValueInt64()))
	}

	return gslbservice_lbmonitor_binding
}

func gslbservice_lbmonitor_bindingSetAttrFromGet(ctx context.Context, data *GslbserviceLbmonitorBindingResourceModel, getResponseData map[string]interface{}) *GslbserviceLbmonitorBindingResourceModel {
	tflog.Debug(ctx, "In gslbservice_lbmonitor_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["monitor_name"]; ok && val != nil {
		data.MonitorName = types.StringValue(val.(string))
	} else {
		data.MonitorName = types.StringNull()
	}
	if val, ok := getResponseData["monstate"]; ok && val != nil {
		data.Monstate = types.StringValue(val.(string))
	} else {
		data.Monstate = types.StringNull()
	}
	if val, ok := getResponseData["servicename"]; ok && val != nil {
		data.Servicename = types.StringValue(val.(string))
	} else {
		data.Servicename = types.StringNull()
	}
	if val, ok := getResponseData["weight"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Weight = types.Int64Value(intVal)
		}
	} else {
		data.Weight = types.Int64Null()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:base64(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("monitor_name:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.MonitorName.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("servicename:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Servicename.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
