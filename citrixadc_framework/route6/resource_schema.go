package route6

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// Route6ResourceModel describes the resource data model.
type Route6ResourceModel struct {
	Id         types.String `tfsdk:"id"`
	Advertise  types.String `tfsdk:"advertise"`
	Cost       types.Int64  `tfsdk:"cost"`
	Detail     types.Bool   `tfsdk:"detail"`
	Distance   types.Int64  `tfsdk:"distance"`
	Gateway    types.String `tfsdk:"gateway"`
	Mgmt       types.Bool   `tfsdk:"mgmt"`
	Monitor    types.String `tfsdk:"monitor"`
	Msr        types.String `tfsdk:"msr"`
	Network    types.String `tfsdk:"network"`
	Ownergroup types.String `tfsdk:"ownergroup"`
	Routetype  types.String `tfsdk:"routetype"`
	Td         types.Int64  `tfsdk:"td"`
	Vlan       types.Int64  `tfsdk:"vlan"`
	Vxlan      types.Int64  `tfsdk:"vxlan"`
	Weight     types.Int64  `tfsdk:"weight"`
}

func (r *Route6Resource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the route6 resource.",
			},
			"advertise": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Advertise this route.",
			},
			"cost": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(1),
				Description: "Positive integer used by the routing algorithms to determine preference for this route. The lower the cost, the higher the preference.",
			},
			"detail": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "To get a detailed view.",
			},
			"distance": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(1),
				Description: "Administrative distance of this route from the appliance.",
			},
			"gateway": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The gateway for this route. The value for this parameter is either an IPv6 address or null.",
			},
			"mgmt": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Route in management plane.",
			},
			"monitor": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the monitor, of type ND6 or PING, configured on the Citrix ADC to monitor this route.",
			},
			"msr": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Monitor this route with a monitor of type ND6 or PING.",
			},
			"network": schema.StringAttribute{
				Required:    true,
				Description: "IPv6 network address for which to add a route entry to the routing table of the Citrix ADC.",
			},
			"ownergroup": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Default:     stringdefault.StaticString("DEFAULT_NG"),
				Description: "The owner node group in a Cluster for this route6. If owner node group is not specified then the route is treated as Striped route.",
			},
			"routetype": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Type of IPv6 routes to remove from the routing table of the Citrix ADC.",
			},
			"td": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.",
			},
			"vlan": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Integer value that uniquely identifies a VLAN through which the Citrix ADC forwards the packets for this route.",
			},
			"vxlan": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Integer value that uniquely identifies a VXLAN through which the Citrix ADC forwards the packets for this route.",
			},
			"weight": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(1),
				Description: "Positive integer used by the routing algorithms to determine preference for this route over others of equal cost. The lower the weight, the higher the preference.",
			},
		},
	}
}

func route6GetThePayloadFromtheConfig(ctx context.Context, data *Route6ResourceModel) network.Route6 {
	tflog.Debug(ctx, "In route6GetThePayloadFromtheConfig Function")

	// Create API request body from the model
	route6 := network.Route6{}
	if !data.Advertise.IsNull() {
		route6.Advertise = data.Advertise.ValueString()
	}
	if !data.Cost.IsNull() {
		route6.Cost = utils.IntPtr(int(data.Cost.ValueInt64()))
	}
	if !data.Detail.IsNull() {
		route6.Detail = data.Detail.ValueBool()
	}
	if !data.Distance.IsNull() {
		route6.Distance = utils.IntPtr(int(data.Distance.ValueInt64()))
	}
	if !data.Gateway.IsNull() {
		route6.Gateway = data.Gateway.ValueString()
	}
	if !data.Mgmt.IsNull() {
		route6.Mgmt = data.Mgmt.ValueBool()
	}
	if !data.Monitor.IsNull() {
		route6.Monitor = data.Monitor.ValueString()
	}
	if !data.Msr.IsNull() {
		route6.Msr = data.Msr.ValueString()
	}
	if !data.Network.IsNull() {
		route6.Network = data.Network.ValueString()
	}
	if !data.Ownergroup.IsNull() {
		route6.Ownergroup = data.Ownergroup.ValueString()
	}
	if !data.Routetype.IsNull() {
		route6.Routetype = data.Routetype.ValueString()
	}
	if !data.Td.IsNull() {
		route6.Td = utils.IntPtr(int(data.Td.ValueInt64()))
	}
	if !data.Vlan.IsNull() {
		route6.Vlan = utils.IntPtr(int(data.Vlan.ValueInt64()))
	}
	if !data.Vxlan.IsNull() {
		route6.Vxlan = utils.IntPtr(int(data.Vxlan.ValueInt64()))
	}
	if !data.Weight.IsNull() {
		route6.Weight = utils.IntPtr(int(data.Weight.ValueInt64()))
	}

	return route6
}

func route6SetAttrFromGet(ctx context.Context, data *Route6ResourceModel, getResponseData map[string]interface{}) *Route6ResourceModel {
	tflog.Debug(ctx, "In route6SetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["advertise"]; ok && val != nil {
		data.Advertise = types.StringValue(val.(string))
	} else {
		data.Advertise = types.StringNull()
	}
	if val, ok := getResponseData["cost"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Cost = types.Int64Value(intVal)
		}
	} else {
		data.Cost = types.Int64Null()
	}
	if val, ok := getResponseData["detail"]; ok && val != nil {
		data.Detail = types.BoolValue(val.(bool))
	} else {
		data.Detail = types.BoolNull()
	}
	if val, ok := getResponseData["distance"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Distance = types.Int64Value(intVal)
		}
	} else {
		data.Distance = types.Int64Null()
	}
	if val, ok := getResponseData["gateway"]; ok && val != nil {
		data.Gateway = types.StringValue(val.(string))
	} else {
		data.Gateway = types.StringNull()
	}
	if val, ok := getResponseData["mgmt"]; ok && val != nil {
		data.Mgmt = types.BoolValue(val.(bool))
	} else {
		data.Mgmt = types.BoolNull()
	}
	if val, ok := getResponseData["monitor"]; ok && val != nil {
		data.Monitor = types.StringValue(val.(string))
	} else {
		data.Monitor = types.StringNull()
	}
	if val, ok := getResponseData["msr"]; ok && val != nil {
		data.Msr = types.StringValue(val.(string))
	} else {
		data.Msr = types.StringNull()
	}
	if val, ok := getResponseData["network"]; ok && val != nil {
		data.Network = types.StringValue(val.(string))
	} else {
		data.Network = types.StringNull()
	}
	if val, ok := getResponseData["ownergroup"]; ok && val != nil {
		data.Ownergroup = types.StringValue(val.(string))
	} else {
		data.Ownergroup = types.StringNull()
	}
	if val, ok := getResponseData["routetype"]; ok && val != nil {
		data.Routetype = types.StringValue(val.(string))
	} else {
		data.Routetype = types.StringNull()
	}
	if val, ok := getResponseData["td"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Td = types.Int64Value(intVal)
		}
	} else {
		data.Td = types.Int64Null()
	}
	if val, ok := getResponseData["vlan"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Vlan = types.Int64Value(intVal)
		}
	} else {
		data.Vlan = types.Int64Null()
	}
	if val, ok := getResponseData["vxlan"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Vxlan = types.Int64Value(intVal)
		}
	} else {
		data.Vxlan = types.Int64Null()
	}
	if val, ok := getResponseData["weight"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Weight = types.Int64Value(intVal)
		}
	} else {
		data.Weight = types.Int64Null()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated
	data.Id = types.StringValue(fmt.Sprintf("%s,%d", data.Network.ValueString(), data.Td.ValueInt64()))

	return data
}
