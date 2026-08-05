package route6

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
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
			// SDK v2: Optional + Computed (updateable, no ForceNew)
			"advertise": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Advertise this route.",
			},
			// SDK v2: Optional + Computed (no Default)
			"cost": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Positive integer used by the routing algorithms to determine preference for this route. The lower the cost, the higher the preference.",
			},
			// SDK v2: Optional + Computed (no ForceNew)
			"detail": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "To get a detailed view.",
			},
			// SDK v2: Optional + Computed (no Default)
			"distance": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Administrative distance of this route from the appliance.",
			},
			// SDK v2: Optional + Computed
			"gateway": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The gateway for this route. The value for this parameter is either an IPv6 address or null.",
			},
			// SDK v2: Optional + Computed + ForceNew
			"mgmt": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Route in management plane.",
			},
			// SDK v2: Optional + Computed
			"monitor": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the monitor, of type ND6 or PING, configured on the Citrix ADC to monitor this route.",
			},
			// SDK v2: Optional + Computed (no Default)
			"msr": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Monitor this route with a monitor of type ND6 or PING.",
			},
			// SDK v2: Required + ForceNew
			"network": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "IPv6 network address for which to add a route entry to the routing table of the Citrix ADC.",
			},
			// SDK v2: Optional + Computed (no ForceNew, no Default)
			"ownergroup": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The owner node group in a Cluster for this route6. If owner node group is not specified then the route is treated as Striped route.",
			},
			// SDK v2: Optional + Computed (no ForceNew)
			"routetype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of IPv6 routes to remove from the routing table of the Citrix ADC.",
			},
			// SDK v2: Optional + Computed
			"td": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.",
			},
			// SDK v2: Optional + Computed
			"vlan": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Integer value that uniquely identifies a VLAN through which the Citrix ADC forwards the packets for this route.",
			},
			// SDK v2: Optional + Computed
			"vxlan": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Integer value that uniquely identifies a VXLAN through which the Citrix ADC forwards the packets for this route.",
			},
			// SDK v2: Optional + Computed (no Default)
			"weight": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Positive integer used by the routing algorithms to determine preference for this route over others of equal cost. The lower the weight, the higher the preference.",
			},
		},
	}
}

// route6GetThePayloadFromtheConfig builds the NITRO add payload. It mirrors the
// SDK v2 Create struct: every configured field is sent (unknown/unset values are
// skipped, and the struct's omitempty drops zero values), so add accepts the
// same field set the legacy resource sent.
func route6GetThePayloadFromtheConfig(ctx context.Context, data *Route6ResourceModel) network.Route6 {
	tflog.Debug(ctx, "In route6GetThePayloadFromtheConfig Function")

	route6 := network.Route6{}
	if !data.Advertise.IsNull() && !data.Advertise.IsUnknown() {
		route6.Advertise = data.Advertise.ValueString()
	}
	if !data.Cost.IsNull() && !data.Cost.IsUnknown() {
		route6.Cost = utils.IntPtr(int(data.Cost.ValueInt64()))
	}
	if !data.Detail.IsNull() && !data.Detail.IsUnknown() {
		route6.Detail = data.Detail.ValueBool()
	}
	if !data.Distance.IsNull() && !data.Distance.IsUnknown() {
		route6.Distance = utils.IntPtr(int(data.Distance.ValueInt64()))
	}
	if !data.Gateway.IsNull() && !data.Gateway.IsUnknown() {
		route6.Gateway = data.Gateway.ValueString()
	}
	if !data.Mgmt.IsNull() && !data.Mgmt.IsUnknown() {
		route6.Mgmt = data.Mgmt.ValueBool()
	}
	if !data.Monitor.IsNull() && !data.Monitor.IsUnknown() {
		route6.Monitor = data.Monitor.ValueString()
	}
	if !data.Msr.IsNull() && !data.Msr.IsUnknown() {
		route6.Msr = data.Msr.ValueString()
	}
	if !data.Network.IsNull() && !data.Network.IsUnknown() {
		route6.Network = data.Network.ValueString()
	}
	if !data.Ownergroup.IsNull() && !data.Ownergroup.IsUnknown() {
		route6.Ownergroup = data.Ownergroup.ValueString()
	}
	if !data.Routetype.IsNull() && !data.Routetype.IsUnknown() {
		route6.Routetype = data.Routetype.ValueString()
	}
	if !data.Td.IsNull() && !data.Td.IsUnknown() {
		route6.Td = utils.IntPtr(int(data.Td.ValueInt64()))
	}
	if !data.Vlan.IsNull() && !data.Vlan.IsUnknown() {
		route6.Vlan = utils.IntPtr(int(data.Vlan.ValueInt64()))
	}
	if !data.Vxlan.IsNull() && !data.Vxlan.IsUnknown() {
		route6.Vxlan = utils.IntPtr(int(data.Vxlan.ValueInt64()))
	}
	if !data.Weight.IsNull() && !data.Weight.IsUnknown() {
		route6.Weight = utils.IntPtr(int(data.Weight.ValueInt64()))
	}

	return route6
}

// route6GetThePayloadForUpdate builds the NITRO update payload. The route6
// `update` endpoint is PUT /route6 (unnamed) and accepts only the identifying
// keys plus the mutable properties (network, gateway, vlan, vxlan, td, weight,
// distance, cost, advertise, msr, monitor). ownergroup/mgmt/routetype/detail are
// NOT part of the update payload, so they are excluded here (ownergroup/mgmt are
// route-identity attributes; mgmt is also RequiresReplace).
func route6GetThePayloadForUpdate(ctx context.Context, data *Route6ResourceModel) network.Route6 {
	tflog.Debug(ctx, "In route6GetThePayloadForUpdate Function")

	route6 := network.Route6{}
	// Identifying keys — always send network; send the rest when configured so
	// NITRO can locate the exact route to modify.
	if !data.Network.IsNull() && !data.Network.IsUnknown() {
		route6.Network = data.Network.ValueString()
	}
	if !data.Gateway.IsNull() && !data.Gateway.IsUnknown() {
		route6.Gateway = data.Gateway.ValueString()
	}
	if !data.Vlan.IsNull() && !data.Vlan.IsUnknown() {
		route6.Vlan = utils.IntPtr(int(data.Vlan.ValueInt64()))
	}
	if !data.Vxlan.IsNull() && !data.Vxlan.IsUnknown() {
		route6.Vxlan = utils.IntPtr(int(data.Vxlan.ValueInt64()))
	}
	if !data.Td.IsNull() && !data.Td.IsUnknown() {
		route6.Td = utils.IntPtr(int(data.Td.ValueInt64()))
	}
	// Mutable properties
	if !data.Weight.IsNull() && !data.Weight.IsUnknown() {
		route6.Weight = utils.IntPtr(int(data.Weight.ValueInt64()))
	}
	if !data.Distance.IsNull() && !data.Distance.IsUnknown() {
		route6.Distance = utils.IntPtr(int(data.Distance.ValueInt64()))
	}
	if !data.Cost.IsNull() && !data.Cost.IsUnknown() {
		route6.Cost = utils.IntPtr(int(data.Cost.ValueInt64()))
	}
	if !data.Advertise.IsNull() && !data.Advertise.IsUnknown() {
		route6.Advertise = data.Advertise.ValueString()
	}
	if !data.Msr.IsNull() && !data.Msr.IsUnknown() {
		route6.Msr = data.Msr.ValueString()
	}
	if !data.Monitor.IsNull() && !data.Monitor.IsUnknown() {
		route6.Monitor = data.Monitor.ValueString()
	}

	return route6
}

// route6SetAttrFromGet maps a NITRO GET response record onto the model. It is
// shared by the resource Read/readback and the datasource Read.
//   - `network` is the ID key: preserve the already-known (configured) value to
//     avoid inconsistent-result on any server-side normalization; adopt the GET
//     value only when it is null/unknown (import).
//   - For every other field, adopt the GET value when present; when NITRO omits
//     it, only null the field if it is still Unknown (a not-yet-resolved Computed
//     value). Never clobber a known/configured value that NITRO omitted on default
//     (omit-on-default trap).
func route6SetAttrFromGet(ctx context.Context, data *Route6ResourceModel, getResponseData map[string]interface{}) *Route6ResourceModel {
	tflog.Debug(ctx, "In route6SetAttrFromGet Function")

	if val, ok := getResponseData["advertise"]; ok && val != nil {
		data.Advertise = types.StringValue(val.(string))
	} else if data.Advertise.IsUnknown() {
		data.Advertise = types.StringNull()
	}
	if val, ok := getResponseData["cost"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Cost = types.Int64Value(intVal)
		} else if data.Cost.IsUnknown() {
			data.Cost = types.Int64Null()
		}
	} else if data.Cost.IsUnknown() {
		data.Cost = types.Int64Null()
	}
	if val, ok := getResponseData["detail"]; ok && val != nil {
		data.Detail = types.BoolValue(route6ToBool(val))
	} else if data.Detail.IsUnknown() {
		data.Detail = types.BoolNull()
	}
	if val, ok := getResponseData["distance"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Distance = types.Int64Value(intVal)
		} else if data.Distance.IsUnknown() {
			data.Distance = types.Int64Null()
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
		data.Mgmt = types.BoolValue(route6ToBool(val))
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
	// network is the ID key — preserve the configured value; adopt GET only on import.
	if data.Network.IsNull() || data.Network.IsUnknown() {
		if val, ok := getResponseData["network"]; ok && val != nil {
			data.Network = types.StringValue(val.(string))
		}
	}
	if val, ok := getResponseData["ownergroup"]; ok && val != nil {
		data.Ownergroup = types.StringValue(val.(string))
	} else if data.Ownergroup.IsUnknown() {
		data.Ownergroup = types.StringNull()
	}
	if val, ok := getResponseData["routetype"]; ok && val != nil {
		data.Routetype = types.StringValue(val.(string))
	} else if data.Routetype.IsUnknown() {
		data.Routetype = types.StringNull()
	}
	if val, ok := getResponseData["td"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Td = types.Int64Value(intVal)
		} else if data.Td.IsUnknown() {
			data.Td = types.Int64Null()
		}
	} else if data.Td.IsUnknown() {
		data.Td = types.Int64Null()
	}
	if val, ok := getResponseData["vlan"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Vlan = types.Int64Value(intVal)
		} else if data.Vlan.IsUnknown() {
			data.Vlan = types.Int64Null()
		}
	} else if data.Vlan.IsUnknown() {
		data.Vlan = types.Int64Null()
	}
	if val, ok := getResponseData["vxlan"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Vxlan = types.Int64Value(intVal)
		} else if data.Vxlan.IsUnknown() {
			data.Vxlan = types.Int64Null()
		}
	} else if data.Vxlan.IsUnknown() {
		data.Vxlan = types.Int64Null()
	}
	if val, ok := getResponseData["weight"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Weight = types.Int64Value(intVal)
		} else if data.Weight.IsUnknown() {
			data.Weight = types.Int64Null()
		}
	} else if data.Weight.IsUnknown() {
		data.Weight = types.Int64Null()
	}

	// SDK v2 ID scheme: d.SetId(network) — plain network value (single_unique).
	data.Id = types.StringValue(data.Network.ValueString())

	return data
}

// route6ToBool coerces a NITRO GET value (bool or the usual string encodings)
// into a Go bool.
func route6ToBool(val interface{}) bool {
	switch v := val.(type) {
	case bool:
		return v
	case string:
		switch v {
		case "true", "True", "TRUE", "yes", "YES", "ENABLED", "enabled":
			return true
		default:
			return false
		}
	default:
		return false
	}
}
