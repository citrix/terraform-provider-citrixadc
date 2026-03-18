package gslbservicegroup_lbmonitor_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/gslb"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// GslbservicegroupLbmonitorBindingResourceModel describes the resource data model.
type GslbservicegroupLbmonitorBindingResourceModel struct {
	Id               types.String `tfsdk:"id"`
	Hashid           types.Int64  `tfsdk:"hashid"`
	MonitorName      types.String `tfsdk:"monitor_name"`
	Monstate         types.String `tfsdk:"monstate"`
	Order            types.Int64  `tfsdk:"order"`
	Passive          types.Bool   `tfsdk:"passive"`
	Port             types.Int64  `tfsdk:"port"`
	Publicip         types.String `tfsdk:"publicip"`
	Publicport       types.Int64  `tfsdk:"publicport"`
	Servicegroupname types.String `tfsdk:"servicegroupname"`
	Siteprefix       types.String `tfsdk:"siteprefix"`
	State            types.String `tfsdk:"state"`
	Weight           types.Int64  `tfsdk:"weight"`
}

func (r *GslbservicegroupLbmonitorBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the gslbservicegroup_lbmonitor_binding resource.",
			},
			"hashid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Unique numerical identifier used by hash based load balancing methods to identify a service.",
			},
			"monitor_name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Monitor name.",
			},
			"monstate": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Monitor state.",
			},
			"order": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Order number to be assigned to the gslb servicegroup member",
			},
			"passive": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Indicates if load monitor is passive. A passive load monitor does not remove service from LB decision when threshold is breached.",
			},
			"port": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Port number of the GSLB service. Each service must have a unique port number.",
			},
			"publicip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The public IP address that a NAT device translates to the GSLB service's private IP address. Optional.",
			},
			"publicport": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The public port associated with the GSLB service's public IP address. The port is mapped to the service's private port number. Applicable to the local GSLB service. Optional.",
			},
			"servicegroupname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the GSLB service group.",
			},
			"siteprefix": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The site's prefix string. When the GSLB service group is bound to a GSLB virtual server, a GSLB site domain is generated internally for each bound serviceitem-domain pair by concatenating the site prefix of the service item and the name of the domain. If the special string NONE is specified, the site-prefix string is unset. When implementing HTTP redirect site persistence, the Citrix ADC redirects GSLB requests to GSLB services by using their site domains.",
			},
			"state": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Initial state of the service after binding.",
			},
			"weight": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Weight to assign to the servers in the service group. Specifies the capacity of the servers relative to the other servers in the load balancing configuration. The higher the weight, the higher the percentage of requests sent to the service.",
			},
		},
	}
}

func gslbservicegroup_lbmonitor_bindingGetThePayloadFromtheConfig(ctx context.Context, data *GslbservicegroupLbmonitorBindingResourceModel) gslb.Gslbservicegrouplbmonitorbinding {
	tflog.Debug(ctx, "In gslbservicegroup_lbmonitor_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	gslbservicegroup_lbmonitor_binding := gslb.Gslbservicegrouplbmonitorbinding{}
	if !data.Hashid.IsNull() {
		gslbservicegroup_lbmonitor_binding.Hashid = utils.IntPtr(int(data.Hashid.ValueInt64()))
	}
	if !data.MonitorName.IsNull() {
		gslbservicegroup_lbmonitor_binding.Monitorname = data.MonitorName.ValueString()
	}
	if !data.Monstate.IsNull() {
		gslbservicegroup_lbmonitor_binding.Monstate = data.Monstate.ValueString()
	}
	if !data.Order.IsNull() {
		gslbservicegroup_lbmonitor_binding.Order = utils.IntPtr(int(data.Order.ValueInt64()))
	}
	if !data.Passive.IsNull() {
		gslbservicegroup_lbmonitor_binding.Passive = data.Passive.ValueBool()
	}
	if !data.Port.IsNull() {
		gslbservicegroup_lbmonitor_binding.Port = utils.IntPtr(int(data.Port.ValueInt64()))
	}
	if !data.Publicip.IsNull() {
		gslbservicegroup_lbmonitor_binding.Publicip = data.Publicip.ValueString()
	}
	if !data.Publicport.IsNull() {
		gslbservicegroup_lbmonitor_binding.Publicport = utils.IntPtr(int(data.Publicport.ValueInt64()))
	}
	if !data.Servicegroupname.IsNull() {
		gslbservicegroup_lbmonitor_binding.Servicegroupname = data.Servicegroupname.ValueString()
	}
	if !data.Siteprefix.IsNull() {
		gslbservicegroup_lbmonitor_binding.Siteprefix = data.Siteprefix.ValueString()
	}
	if !data.State.IsNull() {
		gslbservicegroup_lbmonitor_binding.State = data.State.ValueString()
	}
	if !data.Weight.IsNull() {
		gslbservicegroup_lbmonitor_binding.Weight = utils.IntPtr(int(data.Weight.ValueInt64()))
	}

	return gslbservicegroup_lbmonitor_binding
}

func gslbservicegroup_lbmonitor_bindingSetAttrFromGet(ctx context.Context, data *GslbservicegroupLbmonitorBindingResourceModel, getResponseData map[string]interface{}) *GslbservicegroupLbmonitorBindingResourceModel {
	tflog.Debug(ctx, "In gslbservicegroup_lbmonitor_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["hashid"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Hashid = types.Int64Value(intVal)
		}
	} else {
		data.Hashid = types.Int64Null()
	}
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
	if val, ok := getResponseData["order"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Order = types.Int64Value(intVal)
		}
	} else {
		data.Order = types.Int64Null()
	}
	if val, ok := getResponseData["passive"]; ok && val != nil {
		data.Passive = types.BoolValue(val.(bool))
	} else {
		data.Passive = types.BoolNull()
	}
	if val, ok := getResponseData["port"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Port = types.Int64Value(intVal)
		}
	} else {
		data.Port = types.Int64Null()
	}
	if val, ok := getResponseData["publicip"]; ok && val != nil {
		data.Publicip = types.StringValue(val.(string))
	} else {
		data.Publicip = types.StringNull()
	}
	if val, ok := getResponseData["publicport"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Publicport = types.Int64Value(intVal)
		}
	} else {
		data.Publicport = types.Int64Null()
	}
	if val, ok := getResponseData["servicegroupname"]; ok && val != nil {
		data.Servicegroupname = types.StringValue(val.(string))
	} else {
		data.Servicegroupname = types.StringNull()
	}
	if val, ok := getResponseData["siteprefix"]; ok && val != nil {
		data.Siteprefix = types.StringValue(val.(string))
	} else {
		data.Siteprefix = types.StringNull()
	}
	if val, ok := getResponseData["state"]; ok && val != nil {
		data.State = types.StringValue(val.(string))
	} else {
		data.State = types.StringNull()
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
	idParts = append(idParts, fmt.Sprintf("port:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Port.ValueInt64()))))
	idParts = append(idParts, fmt.Sprintf("servicegroupname:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Servicegroupname.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
