package servicegroup_lbmonitor_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/basic"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// ServicegroupLbmonitorBindingResourceModel describes the resource data model.
type ServicegroupLbmonitorBindingResourceModel struct {
	Id               types.String `tfsdk:"id"`
	Customserverid   types.String `tfsdk:"customserverid"`
	Dbsttl           types.Int64  `tfsdk:"dbsttl"`
	Hashid           types.Int64  `tfsdk:"hashid"`
	MonitorName      types.String `tfsdk:"monitor_name"`
	Monstate         types.String `tfsdk:"monstate"`
	Nameserver       types.String `tfsdk:"nameserver"`
	Order            types.Int64  `tfsdk:"order"`
	Passive          types.Bool   `tfsdk:"passive"`
	Port             types.Int64  `tfsdk:"port"`
	Serverid         types.Int64  `tfsdk:"serverid"`
	Servicegroupname types.String `tfsdk:"servicegroupname"`
	State            types.String `tfsdk:"state"`
	Weight           types.Int64  `tfsdk:"weight"`
}

func (r *ServicegroupLbmonitorBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the servicegroup_lbmonitor_binding resource.",
			},
			"customserverid": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Unique service identifier. Used when the persistency type for the virtual server is set to Custom Server ID.",
			},
			"dbsttl": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Specify the TTL for DNS record for domain based service.The default value of ttl is 0 which indicates to use the TTL received in DNS response for monitors",
			},
			"hashid": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Unique numerical identifier used by hash based load balancing methods to identify a service.",
			},
			"monitor_name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Monitor name.",
			},
			"monstate": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Monitor state.",
			},
			"nameserver": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Specify the nameserver to which the query for bound domain needs to be sent. If not specified, use the global nameserver",
			},
			"order": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Order number to be assigned to the servicegroup member",
			},
			"passive": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Indicates if load monitor is passive. A passive load monitor does not remove service from LB decision when threshold is breached.",
			},
			"port": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Port number of the service. Each service must have a unique port number.",
			},
			"serverid": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "The  identifier for the service. This is used when the persistency type is set to Custom Server ID.",
			},
			"servicegroupname": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the service group.",
			},
			"state": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Initial state of the service after binding.",
			},
			"weight": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Weight to assign to the servers in the service group. Specifies the capacity of the servers relative to the other servers in the load balancing configuration. The higher the weight, the higher the percentage of requests sent to the service.",
			},
		},
	}
}

func servicegroup_lbmonitor_bindingGetThePayloadFromthePlan(ctx context.Context, data *ServicegroupLbmonitorBindingResourceModel) basic.Servicegrouplbmonitorbinding {
	tflog.Debug(ctx, "In servicegroup_lbmonitor_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	servicegroup_lbmonitor_binding := basic.Servicegrouplbmonitorbinding{}
	if !data.Customserverid.IsNull() && !data.Customserverid.IsUnknown() {
		servicegroup_lbmonitor_binding.Customserverid = data.Customserverid.ValueString()
	}
	if !data.Dbsttl.IsNull() && !data.Dbsttl.IsUnknown() {
		servicegroup_lbmonitor_binding.Dbsttl = utils.IntPtr(int(data.Dbsttl.ValueInt64()))
	}
	if !data.Hashid.IsNull() && !data.Hashid.IsUnknown() {
		servicegroup_lbmonitor_binding.Hashid = utils.IntPtr(int(data.Hashid.ValueInt64()))
	}
	if !data.MonitorName.IsNull() && !data.MonitorName.IsUnknown() {
		servicegroup_lbmonitor_binding.Monitorname = data.MonitorName.ValueString()
	}
	if !data.Monstate.IsNull() && !data.Monstate.IsUnknown() {
		servicegroup_lbmonitor_binding.Monstate = data.Monstate.ValueString()
	}
	if !data.Nameserver.IsNull() && !data.Nameserver.IsUnknown() {
		servicegroup_lbmonitor_binding.Nameserver = data.Nameserver.ValueString()
	}
	if !data.Order.IsNull() && !data.Order.IsUnknown() {
		servicegroup_lbmonitor_binding.Order = utils.IntPtr(int(data.Order.ValueInt64()))
	}
	if !data.Passive.IsNull() && !data.Passive.IsUnknown() {
		servicegroup_lbmonitor_binding.Passive = data.Passive.ValueBool()
	}
	if !data.Port.IsNull() && !data.Port.IsUnknown() {
		servicegroup_lbmonitor_binding.Port = utils.IntPtr(int(data.Port.ValueInt64()))
	}
	if !data.Serverid.IsNull() && !data.Serverid.IsUnknown() {
		servicegroup_lbmonitor_binding.Serverid = utils.IntPtr(int(data.Serverid.ValueInt64()))
	}
	if !data.Servicegroupname.IsNull() && !data.Servicegroupname.IsUnknown() {
		servicegroup_lbmonitor_binding.Servicegroupname = data.Servicegroupname.ValueString()
	}
	if !data.State.IsNull() && !data.State.IsUnknown() {
		servicegroup_lbmonitor_binding.State = data.State.ValueString()
	}
	if !data.Weight.IsNull() && !data.Weight.IsUnknown() {
		servicegroup_lbmonitor_binding.Weight = utils.IntPtr(int(data.Weight.ValueInt64()))
	}

	return servicegroup_lbmonitor_binding
}

func servicegroup_lbmonitor_bindingSetAttrFromGet(ctx context.Context, data *ServicegroupLbmonitorBindingResourceModel, getResponseData map[string]interface{}) *ServicegroupLbmonitorBindingResourceModel {
	tflog.Debug(ctx, "In servicegroup_lbmonitor_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["customserverid"]; ok && val != nil {
		data.Customserverid = types.StringValue(val.(string))
	} else {
		data.Customserverid = types.StringNull()
	}
	if val, ok := getResponseData["dbsttl"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Dbsttl = types.Int64Value(intVal)
		}
	} else {
		data.Dbsttl = types.Int64Null()
	}
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
	if val, ok := getResponseData["nameserver"]; ok && val != nil {
		data.Nameserver = types.StringValue(val.(string))
	} else {
		data.Nameserver = types.StringNull()
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
	if val, ok := getResponseData["serverid"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Serverid = types.Int64Value(intVal)
		}
	} else {
		data.Serverid = types.Int64Null()
	}
	if val, ok := getResponseData["servicegroupname"]; ok && val != nil {
		data.Servicegroupname = types.StringValue(val.(string))
	} else {
		data.Servicegroupname = types.StringNull()
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
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("monitor_name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.MonitorName.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("port:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Port.ValueInt64()))))
	idParts = append(idParts, fmt.Sprintf("servicegroupname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Servicegroupname.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
