package route6

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Route6DataSourceModel is the data-source-specific model, decoupled from
// Route6ResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only metadata attributes that the resource deliberately
// omits (protocol-learned flags, monitoring probe counters, dynamic-route
// state, ...). Every non-key attribute is Computed; the Framework's
// per-attribute model <-> schema reflection requires this model to have exactly
// the attributes the data-source schema declares.
type Route6DataSourceModel struct {
	Id         types.String `tfsdk:"id"`
	Advertise  types.String `tfsdk:"advertise"`
	Cost       types.Int64  `tfsdk:"cost"`
	Detail     types.Bool   `tfsdk:"detail"`
	Distance   types.Int64  `tfsdk:"distance"`
	Gateway    types.String `tfsdk:"gateway"`
	Mgmt       types.Bool   `tfsdk:"mgmt"`
	Monitor    types.String `tfsdk:"monitor"`
	Msr        types.String `tfsdk:"msr"`
	Network    types.String `tfsdk:"network"` // Required lookup key
	Ownergroup types.String `tfsdk:"ownergroup"`
	Routetype  types.String `tfsdk:"routetype"`
	Td         types.Int64  `tfsdk:"td"` // Required lookup key
	Vlan       types.Int64  `tfsdk:"vlan"`
	Vxlan      types.Int64  `tfsdk:"vxlan"`
	Weight     types.Int64  `tfsdk:"weight"`

	// Read-only (GET-only) metadata from the NITRO doc read-only set
	// (zion73x_readonly/route6.json). Never settable; populated from GET.
	Gatewayname       types.String `tfsdk:"gatewayname"`
	Type              types.Bool   `tfsdk:"type"`
	Dynamic           types.Bool   `tfsdk:"dynamic"`
	Data              types.Bool   `tfsdk:"data"`
	Flags             types.Bool   `tfsdk:"flags"`
	State             types.Int64  `tfsdk:"state"`
	Totalprobes       types.Int64  `tfsdk:"totalprobes"`
	Totalfailedprobes types.Int64  `tfsdk:"totalfailedprobes"`
	Failedprobes      types.Int64  `tfsdk:"failedprobes"`
	Monstatcode       types.Int64  `tfsdk:"monstatcode"`
	Monstatparam1     types.Int64  `tfsdk:"monstatparam1"`
	Monstatparam2     types.Int64  `tfsdk:"monstatparam2"`
	Monstatparam3     types.Int64  `tfsdk:"monstatparam3"`
	Data1             types.String `tfsdk:"data1"`
	Routeowners       types.List   `tfsdk:"routeowners"`
	Retain            types.Int64  `tfsdk:"retain"`
	Static            types.Bool   `tfsdk:"static"`
	Permanent         types.Bool   `tfsdk:"permanent"`
	Connected         types.Bool   `tfsdk:"connected"`
	Ospfv3            types.Bool   `tfsdk:"ospfv3"`
	Isis              types.Bool   `tfsdk:"isis"`
	Active            types.Bool   `tfsdk:"active"`
	Bgp               types.Bool   `tfsdk:"bgp"`
	Rip               types.Bool   `tfsdk:"rip"`
	Raroute           types.Bool   `tfsdk:"raroute"`
}

func Route6DataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"advertise": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Advertise this route.",
			},
			"cost": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Positive integer used by the routing algorithms to determine preference for this route. The lower the cost, the higher the preference.",
			},
			"detail": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "To get a detailed view.",
			},
			"distance": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Administrative distance of this route from the appliance.",
			},
			"gateway": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The gateway for this route. The value for this parameter is either an IPv6 address or null.",
			},
			"mgmt": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Route in management plane.",
			},
			"monitor": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the monitor, of type ND6 or PING, configured on the Citrix ADC to monitor this route.",
			},
			"msr": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Monitor this route with a monitor of type ND6 or PING.",
			},
			"network": schema.StringAttribute{
				Required:    true,
				Description: "IPv6 network address for which to add a route entry to the routing table of the Citrix ADC.",
			},
			"ownergroup": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The owner node group in a Cluster for this route6. If owner node group is not specified then the route is treated as Striped route.",
			},
			"routetype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of IPv6 routes to remove from the routing table of the Citrix ADC.",
			},
			"td": schema.Int64Attribute{
				Required:    true,
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
				Computed:    true,
				Description: "Positive integer used by the routing algorithms to determine preference for this route over others of equal cost. The lower the weight, the higher the preference.",
			},

			// Read-only (GET-only) metadata surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"gatewayname": schema.StringAttribute{
				Computed:    true,
				Description: "The name of the gateway for this route.",
			},
			"type": schema.BoolAttribute{
				Computed:    true,
				Description: "State of the RNAT.",
			},
			"dynamic": schema.BoolAttribute{
				Computed:    true,
				Description: "Whether this route is dynamically learned or not.",
			},
			"data": schema.BoolAttribute{
				Computed:    true,
				Description: "Internal data of this route.",
			},
			"flags": schema.BoolAttribute{
				Computed:    true,
				Description: "For a dynamic route, the routing protocol from which the route was learned.",
			},
			"state": schema.Int64Attribute{
				Computed:    true,
				Description: "Whether this route is UP or DOWN.",
			},
			"totalprobes": schema.Int64Attribute{
				Computed:    true,
				Description: "The total number of probes sent.",
			},
			"totalfailedprobes": schema.Int64Attribute{
				Computed:    true,
				Description: "The total number of failed probes.",
			},
			"failedprobes": schema.Int64Attribute{
				Computed:    true,
				Description: "Current number of failed monitoring probes.",
			},
			"monstatcode": schema.Int64Attribute{
				Computed:    true,
				Description: "The code indicating the monitor response.",
			},
			"monstatparam1": schema.Int64Attribute{
				Computed:    true,
				Description: "First parameter for use with message code.",
			},
			"monstatparam2": schema.Int64Attribute{
				Computed:    true,
				Description: "Second parameter for use with message code.",
			},
			"monstatparam3": schema.Int64Attribute{
				Computed:    true,
				Description: "Third parameter for use with message code.",
			},
			"data1": schema.StringAttribute{
				Computed:    true,
				Description: "Internal data of this route. Possible values = ENABLED, DISABLED.",
			},
			"routeowners": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "In a cluster, the set of nodes from which this dynamic route has been learnt. Possible values = 0 through 31.",
			},
			"retain": schema.Int64Attribute{
				Computed:    true,
				Description: "Internal retain value of this route.",
			},
			"static": schema.BoolAttribute{
				Computed:    true,
				Description: "Static route.",
			},
			"permanent": schema.BoolAttribute{
				Computed:    true,
				Description: "Permanent Route.",
			},
			"connected": schema.BoolAttribute{
				Computed:    true,
				Description: "Connected Route.",
			},
			"ospfv3": schema.BoolAttribute{
				Computed:    true,
				Description: "For a dynamic route, the routing protocol from which the route was learned.",
			},
			"isis": schema.BoolAttribute{
				Computed:    true,
				Description: "If this route is dynamic then which routing protocol was it learnt from.",
			},
			"active": schema.BoolAttribute{
				Computed:    true,
				Description: "For a dynamic route, the routing protocol from which the route was learned.",
			},
			"bgp": schema.BoolAttribute{
				Computed:    true,
				Description: "For a dynamic route, the routing protocol from which the route was learned.",
			},
			"rip": schema.BoolAttribute{
				Computed:    true,
				Description: "For a dynamic route, the routing protocol from which the route was learned.",
			},
			"raroute": schema.BoolAttribute{
				Computed:    true,
				Description: "For a dynamic route, the routing protocol from which the route was learned.",
			},
		},
	}
}

// route6DataSourceSetAttrFromGet projects a NITRO route6 GET response onto the
// data-source model. Because a data source has no plan/apply reconciliation,
// attributes are simply filled from the GET (or left Null when the GET omits
// them). The shared utils.MapGet* helpers implement that projection.
func route6DataSourceSetAttrFromGet(ctx context.Context, data *Route6DataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In route6DataSourceSetAttrFromGet Function")

	// Read/write attributes as read-back outputs.
	data.Advertise = utils.MapGetString(g, "advertise")
	data.Cost = utils.MapGetInt64(g, "cost")
	data.Detail = utils.MapGetBool(g, "detail")
	data.Distance = utils.MapGetInt64(g, "distance")
	data.Gateway = utils.MapGetString(g, "gateway")
	data.Mgmt = utils.MapGetBool(g, "mgmt")
	data.Monitor = utils.MapGetString(g, "monitor")
	data.Msr = utils.MapGetString(g, "msr")
	data.Network = utils.MapGetString(g, "network")
	data.Ownergroup = utils.MapGetString(g, "ownergroup")
	data.Routetype = utils.MapGetString(g, "routetype")
	// td is a config-supplied key; NITRO omits it for the default traffic
	// domain (0), so preserve the configured value instead of nulling it.
	if tdv, tdok := g["td"]; tdok && tdv != nil {
		if iv, err := utils.ConvertToInt64(tdv); err == nil {
			data.Td = types.Int64Value(iv)
		}
	}
	data.Vlan = utils.MapGetInt64(g, "vlan")
	data.Vxlan = utils.MapGetInt64(g, "vxlan")
	data.Weight = utils.MapGetInt64(g, "weight")

	// Read-only metadata.
	data.Gatewayname = utils.MapGetString(g, "gatewayname")
	data.Type = utils.MapGetBool(g, "type")
	data.Dynamic = utils.MapGetBool(g, "dynamic")
	data.Data = utils.MapGetBool(g, "data")
	data.Flags = utils.MapGetBool(g, "flags")
	data.State = utils.MapGetInt64(g, "state")
	data.Totalprobes = utils.MapGetInt64(g, "totalprobes")
	data.Totalfailedprobes = utils.MapGetInt64(g, "totalfailedprobes")
	data.Failedprobes = utils.MapGetInt64(g, "failedprobes")
	data.Monstatcode = utils.MapGetInt64(g, "monstatcode")
	data.Monstatparam1 = utils.MapGetInt64(g, "monstatparam1")
	data.Monstatparam2 = utils.MapGetInt64(g, "monstatparam2")
	data.Monstatparam3 = utils.MapGetInt64(g, "monstatparam3")
	data.Data1 = utils.MapGetString(g, "data1")
	data.Routeowners = utils.MapGetStringList(g, "routeowners")
	data.Retain = utils.MapGetInt64(g, "retain")
	data.Static = utils.MapGetBool(g, "Static")
	data.Permanent = utils.MapGetBool(g, "permanent")
	data.Connected = utils.MapGetBool(g, "connected")
	data.Ospfv3 = utils.MapGetBool(g, "ospfv3")
	data.Isis = utils.MapGetBool(g, "isis")
	data.Active = utils.MapGetBool(g, "active")
	data.Bgp = utils.MapGetBool(g, "bgp")
	data.Rip = utils.MapGetBool(g, "rip")
	data.Raroute = utils.MapGetBool(g, "raroute")

	// SDK v2 ID scheme: plain network value (single_unique).
	data.Id = types.StringValue(data.Network.ValueString())
}
