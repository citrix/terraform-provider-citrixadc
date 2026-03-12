package nslicense

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func NslicenseDataSourceSchema() schema.Schema {
	return schema.Schema{
		Description: "Data source for querying Citrix ADC license information",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nslicense datasource (always 'nslicense')",
			},
			"wl": schema.BoolAttribute{
				Computed:    true,
				Description: "Web Logging feature is licensed",
			},
			"sp": schema.BoolAttribute{
				Computed:    true,
				Description: "Surge Protection feature is licensed",
			},
			"lb": schema.BoolAttribute{
				Computed:    true,
				Description: "Load Balancing feature is licensed",
			},
			"cs": schema.BoolAttribute{
				Computed:    true,
				Description: "Content Switching feature is licensed",
			},
			"cr": schema.BoolAttribute{
				Computed:    true,
				Description: "Cache Redirection feature is licensed",
			},
			"cmp": schema.BoolAttribute{
				Computed:    true,
				Description: "Compression feature is licensed",
			},
			"delta": schema.BoolAttribute{
				Computed:    true,
				Description: "Delta Compression feature is licensed",
			},
			"ssl": schema.BoolAttribute{
				Computed:    true,
				Description: "SSL Offloading feature is licensed",
			},
			"gslb": schema.BoolAttribute{
				Computed:    true,
				Description: "Global Server Load Balancing feature is licensed",
			},
			"gslbp": schema.BoolAttribute{
				Computed:    true,
				Description: "GSLB Proximity feature is licensed",
			},
			"routing": schema.BoolAttribute{
				Computed:    true,
				Description: "Routing feature is licensed",
			},
			"cf": schema.BoolAttribute{
				Computed:    true,
				Description: "Content Filtering feature is licensed",
			},
			"contentaccelerator": schema.BoolAttribute{
				Computed:    true,
				Description: "Content Accelerator feature is licensed",
			},
			"ic": schema.BoolAttribute{
				Computed:    true,
				Description: "Integrated Caching feature is licensed",
			},
			"sslvpn": schema.BoolAttribute{
				Computed:    true,
				Description: "SSL VPN feature is licensed",
			},
			"f_sslvpn_users": schema.StringAttribute{
				Computed:    true,
				Description: "Number of SSL VPN users licensed",
			},
			"f_ica_users": schema.StringAttribute{
				Computed:    true,
				Description: "Number of ICA users licensed",
			},
			"aaa": schema.BoolAttribute{
				Computed:    true,
				Description: "AAA (Authentication, Authorization, Accounting) feature is licensed",
			},
			"ospf": schema.BoolAttribute{
				Computed:    true,
				Description: "OSPF routing feature is licensed",
			},
			"rip": schema.BoolAttribute{
				Computed:    true,
				Description: "RIP routing feature is licensed",
			},
			"bgp": schema.BoolAttribute{
				Computed:    true,
				Description: "BGP routing feature is licensed",
			},
			"rewrite": schema.BoolAttribute{
				Computed:    true,
				Description: "Rewrite feature is licensed",
			},
			"ipv6pt": schema.BoolAttribute{
				Computed:    true,
				Description: "IPv6 Protocol Translation feature is licensed",
			},
			"appfw": schema.BoolAttribute{
				Computed:    true,
				Description: "Application Firewall feature is licensed",
			},
			"responder": schema.BoolAttribute{
				Computed:    true,
				Description: "Responder feature is licensed",
			},
			"agee": schema.BoolAttribute{
				Computed:    true,
				Description: "AGEE feature is licensed",
			},
			"nsxn": schema.BoolAttribute{
				Computed:    true,
				Description: "NetScaler XN feature is licensed",
			},
			"modelid": schema.StringAttribute{
				Computed:    true,
				Description: "Model ID of the appliance",
			},
			"push": schema.BoolAttribute{
				Computed:    true,
				Description: "Push feature is licensed",
			},
			"appflow": schema.BoolAttribute{
				Computed:    true,
				Description: "AppFlow feature is licensed",
			},
			"cloudbridge": schema.BoolAttribute{
				Computed:    true,
				Description: "CloudBridge feature is licensed",
			},
			"cloudbridgeappliance": schema.BoolAttribute{
				Computed:    true,
				Description: "CloudBridge Appliance feature is licensed",
			},
			"cloudextenderappliance": schema.BoolAttribute{
				Computed:    true,
				Description: "CloudExtender Appliance feature is licensed",
			},
			"isis": schema.BoolAttribute{
				Computed:    true,
				Description: "ISIS routing feature is licensed",
			},
			"cluster": schema.BoolAttribute{
				Computed:    true,
				Description: "Cluster feature is licensed",
			},
			"ch": schema.BoolAttribute{
				Computed:    true,
				Description: "Call Home feature is licensed",
			},
			"appqoe": schema.BoolAttribute{
				Computed:    true,
				Description: "AppQoE feature is licensed",
			},
			"appflowica": schema.BoolAttribute{
				Computed:    true,
				Description: "AppFlow for ICA feature is licensed",
			},
			"isstandardlic": schema.BoolAttribute{
				Computed:    true,
				Description: "Standard license is applied",
			},
			"isenterpriselic": schema.BoolAttribute{
				Computed:    true,
				Description: "Enterprise license is applied",
			},
			"isplatinumlic": schema.BoolAttribute{
				Computed:    true,
				Description: "Platinum license is applied",
			},
			"issgwylic": schema.BoolAttribute{
				Computed:    true,
				Description: "Secure Gateway license is applied",
			},
			"isswglic": schema.BoolAttribute{
				Computed:    true,
				Description: "SWG license is applied",
			},
			"feo": schema.BoolAttribute{
				Computed:    true,
				Description: "Front End Optimization feature is licensed",
			},
			"lsn": schema.BoolAttribute{
				Computed:    true,
				Description: "Large Scale NAT feature is licensed",
			},
			"licensingmode": schema.StringAttribute{
				Computed:    true,
				Description: "Licensing mode (e.g., EXPRESS, POOLED)",
			},
			"rdpproxy": schema.BoolAttribute{
				Computed:    true,
				Description: "RDP Proxy feature is licensed",
			},
			"rep": schema.BoolAttribute{
				Computed:    true,
				Description: "Reputation feature is licensed",
			},
			"urlfiltering": schema.BoolAttribute{
				Computed:    true,
				Description: "URL Filtering feature is licensed",
			},
			"videooptimization": schema.BoolAttribute{
				Computed:    true,
				Description: "Video Optimization feature is licensed",
			},
			"forwardproxy": schema.BoolAttribute{
				Computed:    true,
				Description: "Forward Proxy feature is licensed",
			},
			"sslinterception": schema.BoolAttribute{
				Computed:    true,
				Description: "SSL Interception feature is licensed",
			},
			"remotecontentinspection": schema.BoolAttribute{
				Computed:    true,
				Description: "Remote Content Inspection feature is licensed",
			},
			"adaptivetcp": schema.BoolAttribute{
				Computed:    true,
				Description: "Adaptive TCP feature is licensed",
			},
			"cqa": schema.BoolAttribute{
				Computed:    true,
				Description: "CQA feature is licensed",
			},
			"bot": schema.BoolAttribute{
				Computed:    true,
				Description: "Bot Management feature is licensed",
			},
			"apigateway": schema.BoolAttribute{
				Computed:    true,
				Description: "API Gateway feature is licensed",
			},
		},
	}
}
