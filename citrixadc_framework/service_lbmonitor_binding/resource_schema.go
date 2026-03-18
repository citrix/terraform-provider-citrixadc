package service_lbmonitor_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/basic"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// ServiceLbmonitorBindingResourceModel describes the resource data model.
type ServiceLbmonitorBindingResourceModel struct {
	Id          types.String `tfsdk:"id"`
	MonitorName types.String `tfsdk:"monitor_name"`
	Monstate    types.String `tfsdk:"monstate"`
	Name        types.String `tfsdk:"name"`
	Passive     types.Bool   `tfsdk:"passive"`
	Weight      types.Int64  `tfsdk:"weight"`
}

func (r *ServiceLbmonitorBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the service_lbmonitor_binding resource.",
			},
			"monitor_name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The monitor Names.",
			},
			"monstate": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The configured state (enable/disable) of the monitor on this server.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the service to which to bind a monitor.",
			},
			"passive": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Indicates if load monitor is passive. A passive load monitor does not remove service from LB decision when threshold is breached.",
			},
			"weight": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Weight to assign to the monitor-service binding. When a monitor is UP, the weight assigned to its binding with the service determines how much the monitor contributes toward keeping the health of the service above the value configured for the Monitor Threshold parameter.",
			},
		},
	}
}

func service_lbmonitor_bindingGetThePayloadFromtheConfig(ctx context.Context, data *ServiceLbmonitorBindingResourceModel) basic.Servicelbmonitorbinding {
	tflog.Debug(ctx, "In service_lbmonitor_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	service_lbmonitor_binding := basic.Servicelbmonitorbinding{}
	if !data.MonitorName.IsNull() {
		service_lbmonitor_binding.Monitorname = data.MonitorName.ValueString()
	}
	if !data.Monstate.IsNull() {
		service_lbmonitor_binding.Monstate = data.Monstate.ValueString()
	}
	if !data.Name.IsNull() {
		service_lbmonitor_binding.Name = data.Name.ValueString()
	}
	if !data.Passive.IsNull() {
		service_lbmonitor_binding.Passive = data.Passive.ValueBool()
	}
	if !data.Weight.IsNull() {
		service_lbmonitor_binding.Weight = utils.IntPtr(int(data.Weight.ValueInt64()))
	}

	return service_lbmonitor_binding
}

func service_lbmonitor_bindingSetAttrFromGet(ctx context.Context, data *ServiceLbmonitorBindingResourceModel, getResponseData map[string]interface{}) *ServiceLbmonitorBindingResourceModel {
	tflog.Debug(ctx, "In service_lbmonitor_bindingSetAttrFromGet Function")

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
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["passive"]; ok && val != nil {
		data.Passive = types.BoolValue(val.(bool))
	} else {
		data.Passive = types.BoolNull()
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
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Name.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
