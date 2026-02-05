package nsfeature

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func NsfeatureDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nsfeature datasource.",
			},
			"wl": schema.BoolAttribute{
				Computed:    true,
				Description: "Web Logging.",
			},
			"sp": schema.BoolAttribute{
				Computed:    true,
				Description: "Surge Protection.",
			},
			"lb": schema.BoolAttribute{
				Computed:    true,
				Description: "Load Balancing.",
			},
			"cs": schema.BoolAttribute{
				Computed:    true,
				Description: "Content Switching.",
			},
			"cr": schema.BoolAttribute{
				Computed:    true,
				Description: "Cache Redirection.",
			},
			"cmp": schema.BoolAttribute{
				Computed:    true,
				Description: "Compression.",
			},
			"pq": schema.BoolAttribute{
				Computed:    true,
				Description: "Priority Queuing.",
			},
			"ssl": schema.BoolAttribute{
				Computed:    true,
				Description: "SSL Offloading.",
			},
			"gslb": schema.BoolAttribute{
				Computed:    true,
				Description: "Global Server Load Balancing.",
			},
			"hdosp": schema.BoolAttribute{
				Computed:    true,
				Description: "DoS Protection.",
			},
			"cf": schema.BoolAttribute{
				Computed:    true,
				Description: "Content Filtering.",
			},
			"ic": schema.BoolAttribute{
				Computed:    true,
				Description: "Integrated Caching.",
			},
			"sslvpn": schema.BoolAttribute{
				Computed:    true,
				Description: "SSL VPN.",
			},
			"aaa": schema.BoolAttribute{
				Computed:    true,
				Description: "AAA.",
			},
			"ospf": schema.BoolAttribute{
				Computed:    true,
				Description: "OSPF Routing.",
			},
			"rip": schema.BoolAttribute{
				Computed:    true,
				Description: "RIP Routing.",
			},
			"bgp": schema.BoolAttribute{
				Computed:    true,
				Description: "BGP Routing.",
			},
			"rewrite": schema.BoolAttribute{
				Computed:    true,
				Description: "Rewrite.",
			},
			"ipv6pt": schema.BoolAttribute{
				Computed:    true,
				Description: "IPv6 Protocol Translation.",
			},
			"appfw": schema.BoolAttribute{
				Computed:    true,
				Description: "Application Firewall.",
			},
			"responder": schema.BoolAttribute{
				Computed:    true,
				Description: "Responder.",
			},
			"htmlinjection": schema.BoolAttribute{
				Computed:    true,
				Description: "HTML Injection.",
			},
			"push": schema.BoolAttribute{
				Computed:    true,
				Description: "Push.",
			},
			"appflow": schema.BoolAttribute{
				Computed:    true,
				Description: "AppFlow.",
			},
			"cloudbridge": schema.BoolAttribute{
				Computed:    true,
				Description: "CloudBridge.",
			},
			"isis": schema.BoolAttribute{
				Computed:    true,
				Description: "ISIS Routing.",
			},
			"ch": schema.BoolAttribute{
				Computed:    true,
				Description: "Call Home.",
			},
			"appqoe": schema.BoolAttribute{
				Computed:    true,
				Description: "AppQoE.",
			},
			"contentaccelerator": schema.BoolAttribute{
				Computed:    true,
				Description: "Content Accelerator.",
			},
			"rise": schema.BoolAttribute{
				Computed:    true,
				Description: "RISE.",
			},
			"feo": schema.BoolAttribute{
				Computed:    true,
				Description: "Front End Optimization.",
			},
			"lsn": schema.BoolAttribute{
				Computed:    true,
				Description: "Large Scale NAT.",
			},
			"rdpproxy": schema.BoolAttribute{
				Computed:    true,
				Description: "RDP Proxy.",
			},
			"rep": schema.BoolAttribute{
				Computed:    true,
				Description: "Reputation.",
			},
			"urlfiltering": schema.BoolAttribute{
				Computed:    true,
				Description: "URL Filtering.",
			},
			"videooptimization": schema.BoolAttribute{
				Computed:    true,
				Description: "Video Optimization.",
			},
			"forwardproxy": schema.BoolAttribute{
				Computed:    true,
				Description: "Forward Proxy.",
			},
			"sslinterception": schema.BoolAttribute{
				Computed:    true,
				Description: "SSL Interception.",
			},
			"adaptivetcp": schema.BoolAttribute{
				Computed:    true,
				Description: "Adaptive TCP.",
			},
			"cqa": schema.BoolAttribute{
				Computed:    true,
				Description: "Connection Quality Analytics.",
			},
			"ci": schema.BoolAttribute{
				Computed:    true,
				Description: "Content Inspection.",
			},
			"bot": schema.BoolAttribute{
				Computed:    true,
				Description: "Bot Management.",
			},
			"apigateway": schema.BoolAttribute{
				Computed:    true,
				Description: "API Gateway.",
			},
		},
	}
}
