package route

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// RouteResourceModel describes the resource data model.
type RouteResourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Advertise              types.String `tfsdk:"advertise"`
	Cost                   types.Int64  `tfsdk:"cost"`
	Cost1                  types.Int64  `tfsdk:"cost1"`
	Detail                 types.Bool   `tfsdk:"detail"`
	Distance               types.Int64  `tfsdk:"distance"`
	Gateway                types.String `tfsdk:"gateway"`
	Mgmt                   types.Bool   `tfsdk:"mgmt"`
	Monitor                types.String `tfsdk:"monitor"`
	Msr                    types.String `tfsdk:"msr"`
	Netmask                types.String `tfsdk:"netmask"`
	Network                types.String `tfsdk:"network"`
	Ownergroup             types.String `tfsdk:"ownergroup"`
	Protocol               types.List   `tfsdk:"protocol"`
	Routetype              types.String `tfsdk:"routetype"`
	Td                     types.Int64  `tfsdk:"td"`
	Vlan                   types.Int64  `tfsdk:"vlan"`
	Weight                 types.Int64  `tfsdk:"weight"`
	DeleteDefaultRoute     types.Bool   `tfsdk:"delete_default_route"`
	OriginalDefaultGateway types.String `tfsdk:"original_default_gateway"`
}

func (r *RouteResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the route resource.",
			},
			// Updateable in SDK v2 (not ForceNew).
			"advertise": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Advertise this route.",
			},
			// SDK v2: Optional+Computed, updateable (no ForceNew).
			"cost": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Positive integer used by the routing algorithms to determine preference for using this route. The lower the cost, the higher the preference.",
			},
			"cost1": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(0),
				Description: "The cost of a route is used to compare routes of the same type. The route having the lowest cost is the most preferred route. Possible values: 0 through 65535. Default: 0.",
			},
			// SDK v2: Optional+Computed, updateable (no ForceNew).
			"detail": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Display a detailed view.",
			},
			// SDK v2: Optional+Computed, no Default (value read from ADC).
			"distance": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(1),
				Description: "Administrative distance of this route, which determines the preference of this route over other routes, with same destination, from different routing protocols. A lower value is preferred.",
			},
			// SDK v2: Required + ForceNew.
			"gateway": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "IP address of the gateway for this route. Can be either the IP address of the gateway, or can be null to specify a null interface route.",
			},
			// SDK v2: Optional+Computed + ForceNew.
			"mgmt": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Route in management plane.",
			},
			"monitor": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the monitor, of type ARP or PING, configured on the Citrix ADC to monitor this route.",
			},
			// SDK v2: Optional+Computed, no Default (value read from ADC).
			"msr": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Monitor this route using a monitor of type ARP or PING.",
			},
			// SDK v2: Required + ForceNew.
			"netmask": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The subnet mask associated with the network address.",
			},
			// SDK v2: Required + ForceNew.
			"network": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "IPv4 network address for which to add a route entry in the routing table of the Citrix ADC.",
			},
			// SDK v2: Optional+Computed, updateable (no ForceNew, no Default).
			"ownergroup": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The owner node group in a Cluster for this route. If owner node group is not specified then the route is treated as Striped route.",
			},
			// SDK v2: Optional+Computed, updateable.
			"protocol": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Routing protocol used for advertising this route.",
			},
			// SDK v2: Optional+Computed, updateable (no ForceNew).
			"routetype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Protocol used by routes that you want to remove from the routing table of the Citrix ADC.",
			},
			// SDK v2: Optional+Computed, updateable.
			"td": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.",
			},
			// SDK v2: Optional+Computed, updateable (no ForceNew).
			"vlan": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "VLAN as the gateway for this route.",
			},
			// SDK v2: Optional+Computed, no Default (value read from ADC).
			"weight": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(1),
				Description: "Positive integer used by the routing algorithms to determine preference for this route over others of equal cost. The lower the weight, the higher the preference.",
			},
			// Convenience attribute preserved from SDK v2 (Optional, Default false, ForceNew).
			"delete_default_route": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				Default:  booldefault.StaticBool(false),
				PlanModifiers: []planmodifier.Bool{
					// GH #1436: avoid spurious destroy+recreate on upgrade.
					boolplanmodifier.UseStateForUnknown(),
					boolplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "If true, delete the default static route (network 0.0.0.0, netmask 0.0.0.0) after adding this route",
			},
			// Convenience attribute preserved from SDK v2 (Computed, ForceNew).
			"original_default_gateway": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Stores the gateway of the original default route that was deleted, used to restore it on destroy",
			},
		},
	}
}

// routeGetThePayloadFromthePlan builds the NITRO add payload from the plan.
// Only configured (known, non-null) values are included so that unset
// Optional+Computed attributes are not sent with zero values.
func routeGetThePayloadFromthePlan(ctx context.Context, data *RouteResourceModel) network.Route {
	tflog.Debug(ctx, "In routeGetThePayloadFromthePlan Function")

	route := network.Route{}
	if !data.Advertise.IsNull() && !data.Advertise.IsUnknown() {
		route.Advertise = data.Advertise.ValueString()
	}
	if !data.Cost.IsNull() && !data.Cost.IsUnknown() {
		route.Cost = utils.IntPtr(int(data.Cost.ValueInt64()))
	}
	if !data.Cost1.IsNull() && !data.Cost1.IsUnknown() {
		route.Cost1 = utils.IntPtr(int(data.Cost1.ValueInt64()))
	}
	if !data.Detail.IsNull() && !data.Detail.IsUnknown() {
		route.Detail = data.Detail.ValueBool()
	}
	if !data.Distance.IsNull() && !data.Distance.IsUnknown() {
		route.Distance = utils.IntPtr(int(data.Distance.ValueInt64()))
	}
	if !data.Gateway.IsNull() && !data.Gateway.IsUnknown() {
		route.Gateway = data.Gateway.ValueString()
	}
	if !data.Mgmt.IsNull() && !data.Mgmt.IsUnknown() {
		route.Mgmt = data.Mgmt.ValueBool()
	}
	if !data.Monitor.IsNull() && !data.Monitor.IsUnknown() {
		route.Monitor = data.Monitor.ValueString()
	}
	if !data.Msr.IsNull() && !data.Msr.IsUnknown() {
		route.Msr = data.Msr.ValueString()
	}
	if !data.Netmask.IsNull() && !data.Netmask.IsUnknown() {
		route.Netmask = data.Netmask.ValueString()
	}
	if !data.Network.IsNull() && !data.Network.IsUnknown() {
		route.Network = data.Network.ValueString()
	}
	if !data.Ownergroup.IsNull() && !data.Ownergroup.IsUnknown() {
		route.Ownergroup = data.Ownergroup.ValueString()
	}
	if !data.Protocol.IsNull() && !data.Protocol.IsUnknown() {
		var protocolList []string
		data.Protocol.ElementsAs(ctx, &protocolList, false)
		route.Protocol = protocolList
	}
	if !data.Routetype.IsNull() && !data.Routetype.IsUnknown() {
		route.Routetype = data.Routetype.ValueString()
	}
	if !data.Td.IsNull() && !data.Td.IsUnknown() {
		route.Td = utils.IntPtr(int(data.Td.ValueInt64()))
	}
	if !data.Vlan.IsNull() && !data.Vlan.IsUnknown() {
		route.Vlan = utils.IntPtr(int(data.Vlan.ValueInt64()))
	}
	if !data.Weight.IsNull() && !data.Weight.IsUnknown() {
		route.Weight = utils.IntPtr(int(data.Weight.ValueInt64()))
	}

	return route
}

// routeSetAttrFromGet copies the GET response into the resource state model.
// When a field is absent from the GET response (NITRO omits some fields when
// they hold their default value), the previously-known value is preserved and
// only an as-yet-unknown value is nulled. This avoids the omit-on-default trap
// where a configured 0/false/"" value would be clobbered on Read.
func routeSetAttrFromGet(ctx context.Context, data *RouteResourceModel, getResponseData map[string]interface{}) *RouteResourceModel {
	tflog.Debug(ctx, "In routeSetAttrFromGet Function")

	if val, ok := getResponseData["advertise"]; ok && val != nil {
		data.Advertise = types.StringValue(val.(string))
	} else if data.Advertise.IsUnknown() {
		data.Advertise = types.StringNull()
	}
	if val, ok := getResponseData["cost"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Cost = types.Int64Value(intVal)
		}
	} else if data.Cost.IsUnknown() {
		data.Cost = types.Int64Null()
	}
	if val, ok := getResponseData["cost1"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Cost1 = types.Int64Value(intVal)
		}
	} else if data.Cost1.IsUnknown() {
		data.Cost1 = types.Int64Null()
	}
	if val, ok := getResponseData["detail"]; ok && val != nil {
		data.Detail = types.BoolValue(val.(bool))
	} else if data.Detail.IsUnknown() {
		data.Detail = types.BoolNull()
	}
	if val, ok := getResponseData["distance"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Distance = types.Int64Value(intVal)
		}
	} else if data.Distance.IsUnknown() {
		data.Distance = types.Int64Null()
	}
	if val, ok := getResponseData["gateway"]; ok && val != nil {
		data.Gateway = types.StringValue(val.(string))
	} else if data.Gateway.IsUnknown() {
		data.Gateway = types.StringNull()
	}
	if val, ok := getResponseData["mgmt"]; ok && val != nil {
		data.Mgmt = types.BoolValue(val.(bool))
	} else if data.Mgmt.IsUnknown() {
		data.Mgmt = types.BoolNull()
	}
	if val, ok := getResponseData["monitor"]; ok && val != nil {
		data.Monitor = types.StringValue(val.(string))
	} else if data.Monitor.IsUnknown() {
		data.Monitor = types.StringNull()
	}
	if val, ok := getResponseData["msr"]; ok && val != nil {
		data.Msr = types.StringValue(val.(string))
	} else if data.Msr.IsUnknown() {
		data.Msr = types.StringNull()
	}
	if val, ok := getResponseData["netmask"]; ok && val != nil {
		data.Netmask = types.StringValue(val.(string))
	} else if data.Netmask.IsUnknown() {
		data.Netmask = types.StringNull()
	}
	if val, ok := getResponseData["network"]; ok && val != nil {
		data.Network = types.StringValue(val.(string))
	} else if data.Network.IsUnknown() {
		data.Network = types.StringNull()
	}
	if val, ok := getResponseData["ownergroup"]; ok && val != nil {
		data.Ownergroup = types.StringValue(val.(string))
	} else if data.Ownergroup.IsUnknown() {
		data.Ownergroup = types.StringNull()
	}
	if val, ok := getResponseData["protocol"]; ok && val != nil {
		if sliceVal, ok := val.([]interface{}); ok {
			stringList := utils.ToStringList(sliceVal)
			listValue, _ := types.ListValueFrom(ctx, types.StringType, stringList)
			data.Protocol = listValue
		} else if data.Protocol.IsUnknown() {
			data.Protocol = types.ListNull(types.StringType)
		}
	} else if data.Protocol.IsUnknown() {
		data.Protocol = types.ListNull(types.StringType)
	}
	if val, ok := getResponseData["routetype"]; ok && val != nil {
		data.Routetype = types.StringValue(val.(string))
	} else if data.Routetype.IsUnknown() {
		data.Routetype = types.StringNull()
	}
	if val, ok := getResponseData["td"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Td = types.Int64Value(intVal)
		}
	} else if data.Td.IsUnknown() {
		data.Td = types.Int64Null()
	}
	if val, ok := getResponseData["vlan"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Vlan = types.Int64Value(intVal)
		}
	} else if data.Vlan.IsUnknown() {
		data.Vlan = types.Int64Null()
	}
	if val, ok := getResponseData["weight"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Weight = types.Int64Value(intVal)
		}
	} else if data.Weight.IsUnknown() {
		data.Weight = types.Int64Null()
	}

	// Note: id, delete_default_route and original_default_gateway are not part
	// of the NITRO GET response. The ID is set once in Create; the two
	// convenience attributes are managed by Create/Delete and preserved here.

	return data
}

// routeSetAttrFromGetForDatasource faithfully copies every field of the GET
// response into the model for the datasource read path (which has no prior
// state to preserve) and sets the ID using the SDK v2 scheme.
func routeSetAttrFromGetForDatasource(ctx context.Context, data *RouteResourceModel, getResponseData map[string]interface{}) *RouteResourceModel {
	tflog.Debug(ctx, "In routeSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["advertise"]; ok && val != nil {
		data.Advertise = types.StringValue(val.(string))
	} else {
		data.Advertise = types.StringNull()
	}
	if val, ok := getResponseData["cost"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Cost = types.Int64Value(intVal)
		} else {
			data.Cost = types.Int64Null()
		}
	} else {
		data.Cost = types.Int64Null()
	}
	if val, ok := getResponseData["cost1"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Cost1 = types.Int64Value(intVal)
		} else {
			data.Cost1 = types.Int64Null()
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
		} else {
			data.Distance = types.Int64Null()
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
	if val, ok := getResponseData["protocol"]; ok && val != nil {
		if sliceVal, ok := val.([]interface{}); ok {
			stringList := utils.ToStringList(sliceVal)
			listValue, _ := types.ListValueFrom(ctx, types.StringType, stringList)
			data.Protocol = listValue
		} else {
			data.Protocol = types.ListNull(types.StringType)
		}
	} else {
		data.Protocol = types.ListNull(types.StringType)
	}
	if val, ok := getResponseData["routetype"]; ok && val != nil {
		data.Routetype = types.StringValue(val.(string))
	} else {
		data.Routetype = types.StringNull()
	}
	if val, ok := getResponseData["td"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Td = types.Int64Value(intVal)
		} else {
			data.Td = types.Int64Null()
		}
	} else {
		data.Td = types.Int64Null()
	}
	if val, ok := getResponseData["vlan"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Vlan = types.Int64Value(intVal)
		} else {
			data.Vlan = types.Int64Null()
		}
	} else {
		data.Vlan = types.Int64Null()
	}
	if val, ok := getResponseData["weight"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Weight = types.Int64Value(intVal)
		} else {
			data.Weight = types.Int64Null()
		}
	} else {
		data.Weight = types.Int64Null()
	}

	// Preserve the SDK v2 ID scheme: network__netmask__gateway
	data.Id = types.StringValue(data.Network.ValueString() + "__" + data.Netmask.ValueString() + "__" + data.Gateway.ValueString())

	return data
}
