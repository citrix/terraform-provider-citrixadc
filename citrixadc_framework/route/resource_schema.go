package route

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// RouteResourceModel describes the resource data model.
type RouteResourceModel struct {
	Id         types.String `tfsdk:"id"`
	Advertise  types.String `tfsdk:"advertise"`
	Cost       types.Int64  `tfsdk:"cost"`
	Cost1      types.Int64  `tfsdk:"cost1"`
	Detail     types.Bool   `tfsdk:"detail"`
	Distance   types.Int64  `tfsdk:"distance"`
	Gateway    types.String `tfsdk:"gateway"`
	Mgmt       types.Bool   `tfsdk:"mgmt"`
	Monitor    types.String `tfsdk:"monitor"`
	Msr        types.String `tfsdk:"msr"`
	Netmask    types.String `tfsdk:"netmask"`
	Network    types.String `tfsdk:"network"`
	Ownergroup types.String `tfsdk:"ownergroup"`
	Protocol   types.List   `tfsdk:"protocol"`
	Routetype  types.String `tfsdk:"routetype"`
	Td         types.Int64  `tfsdk:"td"`
	Vlan       types.Int64  `tfsdk:"vlan"`
	Weight     types.Int64  `tfsdk:"weight"`
}

func (r *RouteResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the route resource.",
			},
			"advertise": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Advertise this route.",
			},
			"cost": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Positive integer used by the routing algorithms to determine preference for using this route. The lower the cost, the higher the preference.",
			},
			"cost1": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The cost of a route is used to compare routes of the same type. The route having the lowest cost is the most preferred route. Possible values: 0 through 65535. Default: 0.",
			},
			"detail": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Display a detailed view.",
			},
			"distance": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(1),
				Description: "Administrative distance of this route, which determines the preference of this route over other routes, with same destination, from different routing protocols. A lower value is preferred.",
			},
			"gateway": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address of the gateway for this route. Can be either the IP address of the gateway, or can be null to specify a null interface route.",
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
				Description: "Name of the monitor, of type ARP or PING, configured on the Citrix ADC to monitor this route.",
			},
			"msr": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Monitor this route using a monitor of type ARP or PING.",
			},
			"netmask": schema.StringAttribute{
				Required:    true,
				Description: "The subnet mask associated with the network address.",
			},
			"network": schema.StringAttribute{
				Required:    true,
				Description: "IPv4 network address for which to add a route entry in the routing table of the Citrix ADC.",
			},
			"ownergroup": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Default:     stringdefault.StaticString("DEFAULT_NG"),
				Description: "The owner node group in a Cluster for this route. If owner node group is not specified then the route is treated as Striped route.",
			},
			"protocol": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Description: "Routing protocol used for advertising this route.",
			},
			"routetype": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Protocol used by routes that you want to remove from the routing table of the Citrix ADC.",
			},
			"td": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.",
			},
			"vlan": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "VLAN as the gateway for this route.",
			},
			"weight": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(1),
				Description: "Positive integer used by the routing algorithms to determine preference for this route over others of equal cost. The lower the weight, the higher the preference.",
			},
		},
	}
}

func routeGetThePayloadFromtheConfig(ctx context.Context, data *RouteResourceModel) network.Route {
	tflog.Debug(ctx, "In routeGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	route := network.Route{}
	if !data.Advertise.IsNull() {
		route.Advertise = data.Advertise.ValueString()
	}
	if !data.Cost.IsNull() {
		route.Cost = utils.IntPtr(int(data.Cost.ValueInt64()))
	}
	if !data.Cost1.IsNull() {
		route.Cost1 = utils.IntPtr(int(data.Cost1.ValueInt64()))
	}
	if !data.Detail.IsNull() {
		route.Detail = data.Detail.ValueBool()
	}
	if !data.Distance.IsNull() {
		route.Distance = utils.IntPtr(int(data.Distance.ValueInt64()))
	}
	if !data.Gateway.IsNull() {
		route.Gateway = data.Gateway.ValueString()
	}
	if !data.Mgmt.IsNull() {
		route.Mgmt = data.Mgmt.ValueBool()
	}
	if !data.Monitor.IsNull() {
		route.Monitor = data.Monitor.ValueString()
	}
	if !data.Msr.IsNull() {
		route.Msr = data.Msr.ValueString()
	}
	if !data.Netmask.IsNull() {
		route.Netmask = data.Netmask.ValueString()
	}
	if !data.Network.IsNull() {
		route.Network = data.Network.ValueString()
	}
	if !data.Ownergroup.IsNull() {
		route.Ownergroup = data.Ownergroup.ValueString()
	}
	if !data.Routetype.IsNull() {
		route.Routetype = data.Routetype.ValueString()
	}
	if !data.Td.IsNull() {
		route.Td = utils.IntPtr(int(data.Td.ValueInt64()))
	}
	if !data.Vlan.IsNull() {
		route.Vlan = utils.IntPtr(int(data.Vlan.ValueInt64()))
	}
	if !data.Weight.IsNull() {
		route.Weight = utils.IntPtr(int(data.Weight.ValueInt64()))
	}

	return route
}

func routeSetAttrFromGet(ctx context.Context, data *RouteResourceModel, getResponseData map[string]interface{}) *RouteResourceModel {
	tflog.Debug(ctx, "In routeSetAttrFromGet Function")

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
	if val, ok := getResponseData["cost1"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Cost1 = types.Int64Value(intVal)
		}
	} else {
		data.Cost1 = types.Int64Null()
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
	if val, ok := getResponseData["netmask"]; ok && val != nil {
		data.Netmask = types.StringValue(val.(string))
	} else {
		data.Netmask = types.StringNull()
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
	if val, ok := getResponseData["weight"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Weight = types.Int64Value(intVal)
		}
	} else {
		data.Weight = types.Int64Null()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated
	data.Id = types.StringValue(fmt.Sprintf("%s,%s,%d", data.Network.ValueString(), data.Netmask.ValueString(), data.Td.ValueInt64()))

	return data
}
