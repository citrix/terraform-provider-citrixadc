package route

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// RouteDataSourceModel is the data-source-specific model, decoupled from
// RouteResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only metadata attributes that the resource deliberately
// omits (protocol advertisement flags, monitoring probe counters, dynamic-route
// state, ...). Every non-key attribute is Computed; the Framework's
// per-attribute model <-> schema reflection requires this model to have exactly
// the attributes the data-source schema declares.
type RouteDataSourceModel struct {
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
	Netmask                types.String `tfsdk:"netmask"` // Required lookup key
	Network                types.String `tfsdk:"network"` // Required lookup key
	Ownergroup             types.String `tfsdk:"ownergroup"`
	Protocol               types.List   `tfsdk:"protocol"`
	Routetype              types.String `tfsdk:"routetype"`
	Td                     types.Int64  `tfsdk:"td"` // Required lookup key
	Vlan                   types.Int64  `tfsdk:"vlan"`
	Weight                 types.Int64  `tfsdk:"weight"`
	DeleteDefaultRoute     types.Bool   `tfsdk:"delete_default_route"`
	OriginalDefaultGateway types.String `tfsdk:"original_default_gateway"`

	// Read-only (GET-only) metadata from the NITRO doc read-only set
	// (zion73x_readonly/route.json). Never settable; populated from GET.
	Gatewayname       types.String `tfsdk:"gatewayname"`
	Type              types.Bool   `tfsdk:"type"`
	Dynamic           types.Bool   `tfsdk:"dynamic"`
	Static            types.Bool   `tfsdk:"static"`
	Permanent         types.Bool   `tfsdk:"permanent"`
	Direct            types.Bool   `tfsdk:"direct"`
	Nat               types.Bool   `tfsdk:"nat"`
	Lbroute           types.Bool   `tfsdk:"lbroute"`
	Adv               types.Bool   `tfsdk:"adv"`
	Tunnel            types.Bool   `tfsdk:"tunnel"`
	Data              types.Bool   `tfsdk:"data"`
	Data0             types.Bool   `tfsdk:"data0"`
	Flags             types.Bool   `tfsdk:"flags"`
	Routeowners       types.List   `tfsdk:"routeowners"`
	Retain            types.Int64  `tfsdk:"retain"`
	Ospf              types.Bool   `tfsdk:"ospf"`
	Isis              types.Bool   `tfsdk:"isis"`
	Rip               types.Bool   `tfsdk:"rip"`
	Bgp               types.Bool   `tfsdk:"bgp"`
	Dhcp              types.Bool   `tfsdk:"dhcp"`
	Advospf           types.Bool   `tfsdk:"advospf"`
	Advisis           types.Bool   `tfsdk:"advisis"`
	Advrip            types.Bool   `tfsdk:"advrip"`
	Advbgp            types.Bool   `tfsdk:"advbgp"`
	State             types.Int64  `tfsdk:"state"`
	Totalprobes       types.Int64  `tfsdk:"totalprobes"`
	Totalfailedprobes types.Int64  `tfsdk:"totalfailedprobes"`
	Failedprobes      types.Int64  `tfsdk:"failedprobes"`
	Monstatcode       types.Int64  `tfsdk:"monstatcode"`
	Monstatparam1     types.Int64  `tfsdk:"monstatparam1"`
	Monstatparam2     types.Int64  `tfsdk:"monstatparam2"`
	Monstatparam3     types.Int64  `tfsdk:"monstatparam3"`
}

func RouteDataSourceSchema() schema.Schema {
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
				Description: "Positive integer used by the routing algorithms to determine preference for using this route. The lower the cost, the higher the preference.",
			},
			"cost1": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The cost of a route is used to compare routes of the same type. The route having the lowest cost is the most preferred route. Possible values: 0 through 65535. Default: 0.",
			},
			"detail": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Display a detailed view.",
			},
			"distance": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Administrative distance of this route, which determines the preference of this route over other routes, with same destination, from different routing protocols. A lower value is preferred.",
			},
			"gateway": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address of the gateway for this route. Can be either the IP address of the gateway, or can be null to specify a null interface route.",
			},
			"mgmt": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Route in management plane.",
			},
			"monitor": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the monitor, of type ARP or PING, configured on the Citrix ADC to monitor this route.",
			},
			"msr": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
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
				Optional:    true,
				Computed:    true,
				Description: "The owner node group in a Cluster for this route. If owner node group is not specified then the route is treated as Striped route.",
			},
			"protocol": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Routing protocol used for advertising this route.",
			},
			"routetype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Protocol used by routes that you want to remove from the routing table of the Citrix ADC.",
			},
			"td": schema.Int64Attribute{
				Required:    true,
				Description: "Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.",
			},
			"vlan": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "VLAN as the gateway for this route.",
			},
			"weight": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Positive integer used by the routing algorithms to determine preference for this route over others of equal cost. The lower the weight, the higher the preference.",
			},
			// Convenience attributes carried on the shared resource model; not
			// meaningful for a datasource read but must exist in the schema.
			"delete_default_route": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "If true, delete the default static route (network 0.0.0.0, netmask 0.0.0.0) after adding this route",
			},
			"original_default_gateway": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Stores the gateway of the original default route that was deleted, used to restore it on destroy",
			},

			// Read-only (GET-only) metadata surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"gatewayname": schema.StringAttribute{
				Computed:    true,
				Description: "The name of the gateway for this route. For a route other than a link load balancing (LLB) route, this value is null.",
			},
			"type": schema.BoolAttribute{
				Computed:    true,
				Description: "State of the RNAT.",
			},
			"dynamic": schema.BoolAttribute{
				Computed:    true,
				Description: "State of the route.",
			},
			"static": schema.BoolAttribute{
				Computed:    true,
				Description: "Whether this is a static route.",
			},
			"permanent": schema.BoolAttribute{
				Computed:    true,
				Description: "Whether this is a permanent route.",
			},
			"direct": schema.BoolAttribute{
				Computed:    true,
				Description: "Whether this is a direct route.",
			},
			"nat": schema.BoolAttribute{
				Computed:    true,
				Description: "Whether this is a NAT route.",
			},
			"lbroute": schema.BoolAttribute{
				Computed:    true,
				Description: "Whether this is a link load balancing (LLB) route.",
			},
			"adv": schema.BoolAttribute{
				Computed:    true,
				Description: "Whether this route is advertised.",
			},
			"tunnel": schema.BoolAttribute{
				Computed:    true,
				Description: "Show whether it is a tunnel route or not.",
			},
			"data": schema.BoolAttribute{
				Computed:    true,
				Description: "Internal data of this route.",
			},
			"data0": schema.BoolAttribute{
				Computed:    true,
				Description: "Internal route type is stored, used for get.",
			},
			"flags": schema.BoolAttribute{
				Computed:    true,
				Description: "If this route is dynamic, the name of the routing protocol from which it was learned.",
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
			"ospf": schema.BoolAttribute{
				Computed:    true,
				Description: "OSPF protocol.",
			},
			"isis": schema.BoolAttribute{
				Computed:    true,
				Description: "ISIS protocol.",
			},
			"rip": schema.BoolAttribute{
				Computed:    true,
				Description: "RIP protocol.",
			},
			"bgp": schema.BoolAttribute{
				Computed:    true,
				Description: "BGP protocol.",
			},
			"dhcp": schema.BoolAttribute{
				Computed:    true,
				Description: "DHCP protocol.",
			},
			"advospf": schema.BoolAttribute{
				Computed:    true,
				Description: "Advertised through OSPF protocol.",
			},
			"advisis": schema.BoolAttribute{
				Computed:    true,
				Description: "Advertised through ISIS protocol.",
			},
			"advrip": schema.BoolAttribute{
				Computed:    true,
				Description: "Advertised through RIP protocol.",
			},
			"advbgp": schema.BoolAttribute{
				Computed:    true,
				Description: "Advertised through BGP protocol.",
			},
			"state": schema.Int64Attribute{
				Computed:    true,
				Description: "The state of the static route. Possible values: UP, DOWN.",
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
				Description: "Number of the current failed monitoring probes.",
			},
			"monstatcode": schema.Int64Attribute{
				Computed:    true,
				Description: "The code indicating the monitor response.",
			},
			"monstatparam1": schema.Int64Attribute{
				Computed:    true,
				Description: "First parameter used with the message code.",
			},
			"monstatparam2": schema.Int64Attribute{
				Computed:    true,
				Description: "Second parameter used with the message code.",
			},
			"monstatparam3": schema.Int64Attribute{
				Computed:    true,
				Description: "Third parameter used with the message code.",
			},
		},
	}
}

// routeDataSourceSetAttrFromGet projects a NITRO route GET response onto the
// data-source model. Because a data source has no plan/apply reconciliation,
// attributes are simply filled from the GET (or left Null when the GET omits
// them). The shared utils.MapGet* helpers implement that projection.
func routeDataSourceSetAttrFromGet(ctx context.Context, data *RouteDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In routeDataSourceSetAttrFromGet Function")

	// Read/write attributes as read-back outputs.
	data.Advertise = utils.MapGetString(g, "advertise")
	data.Cost = utils.MapGetInt64(g, "cost")
	data.Cost1 = utils.MapGetInt64(g, "cost1")
	data.Detail = utils.MapGetBool(g, "detail")
	data.Distance = utils.MapGetInt64(g, "distance")
	data.Gateway = utils.MapGetString(g, "gateway")
	data.Mgmt = utils.MapGetBool(g, "mgmt")
	data.Monitor = utils.MapGetString(g, "monitor")
	data.Msr = utils.MapGetString(g, "msr")
	data.Netmask = utils.MapGetString(g, "netmask")
	data.Network = utils.MapGetString(g, "network")
	data.Ownergroup = utils.MapGetString(g, "ownergroup")
	data.Protocol = utils.MapGetStringList(g, "protocol")
	data.Routetype = utils.MapGetString(g, "routetype")
	// td is a config-supplied key; NITRO omits it for the default traffic
	// domain (0), so preserve the configured value instead of nulling it.
	if tdv, tdok := g["td"]; tdok && tdv != nil {
		if iv, err := utils.ConvertToInt64(tdv); err == nil {
			data.Td = types.Int64Value(iv)
		}
	}
	data.Vlan = utils.MapGetInt64(g, "vlan")
	data.Weight = utils.MapGetInt64(g, "weight")

	// Convenience attributes managed by Create/Delete; never returned by GET.
	data.DeleteDefaultRoute = types.BoolNull()
	data.OriginalDefaultGateway = types.StringNull()

	// Read-only metadata.
	data.Gatewayname = utils.MapGetString(g, "gatewayname")
	data.Type = utils.MapGetBool(g, "type")
	data.Dynamic = utils.MapGetBool(g, "dynamic")
	data.Static = utils.MapGetBool(g, "Static")
	data.Permanent = utils.MapGetBool(g, "permanent")
	data.Direct = utils.MapGetBool(g, "direct")
	data.Nat = utils.MapGetBool(g, "nat")
	data.Lbroute = utils.MapGetBool(g, "lbroute")
	data.Adv = utils.MapGetBool(g, "adv")
	data.Tunnel = utils.MapGetBool(g, "tunnel")
	data.Data = utils.MapGetBool(g, "data")
	data.Data0 = utils.MapGetBool(g, "data0")
	data.Flags = utils.MapGetBool(g, "flags")
	data.Routeowners = utils.MapGetStringList(g, "routeowners")
	data.Retain = utils.MapGetInt64(g, "retain")
	data.Ospf = utils.MapGetBool(g, "ospf")
	data.Isis = utils.MapGetBool(g, "isis")
	data.Rip = utils.MapGetBool(g, "rip")
	data.Bgp = utils.MapGetBool(g, "bgp")
	data.Dhcp = utils.MapGetBool(g, "dhcp")
	data.Advospf = utils.MapGetBool(g, "advospf")
	data.Advisis = utils.MapGetBool(g, "advisis")
	data.Advrip = utils.MapGetBool(g, "advrip")
	data.Advbgp = utils.MapGetBool(g, "advbgp")
	data.State = utils.MapGetInt64(g, "state")
	data.Totalprobes = utils.MapGetInt64(g, "totalprobes")
	data.Totalfailedprobes = utils.MapGetInt64(g, "totalfailedprobes")
	data.Failedprobes = utils.MapGetInt64(g, "failedprobes")
	data.Monstatcode = utils.MapGetInt64(g, "monstatcode")
	data.Monstatparam1 = utils.MapGetInt64(g, "monstatparam1")
	data.Monstatparam2 = utils.MapGetInt64(g, "monstatparam2")
	data.Monstatparam3 = utils.MapGetInt64(g, "monstatparam3")

	// Preserve the SDK v2 ID scheme: network__netmask__gateway
	data.Id = types.StringValue(data.Network.ValueString() + "__" + data.Netmask.ValueString() + "__" + data.Gateway.ValueString())
}
