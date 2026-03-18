package route6

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

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
		},
	}
}
