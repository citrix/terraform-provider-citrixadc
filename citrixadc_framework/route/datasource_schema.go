package route

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

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
		},
	}
}
