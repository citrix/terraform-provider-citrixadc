package gslbsite

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func GslbsiteDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"backupparentlist": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "The list of backup gslb sites configured in preferred order. Need to be parent gsb sites.",
			},
			"clip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Cluster IP address. Specify this parameter to connect to the remote cluster site for GSLB auto-sync. Note: The cluster IP address is defined when creating the cluster.",
			},
			"metricexchange": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Exchange metrics with other sites. Metrics are exchanged by using Metric Exchange Protocol (MEP). The appliances in the GSLB setup exchange health information once every second.\n\nIf you disable metrics exchange, you can use only static load balancing methods (such as round robin, static proximity, or the hash-based methods), and if you disable metrics exchange when a dynamic load balancing method (such as least connection) is in operation, the appliance falls back to round robin. Also, if you disable metrics exchange, you must use a monitor to determine the state of GSLB services. Otherwise, the service is marked as DOWN.",
			},
			"naptrreplacementsuffix": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The naptr replacement suffix configured here will be used to construct the naptr replacement field in NAPTR record.",
			},
			"newname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "New name for the GSLB site.",
			},
			"nwmetricexchange": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Exchange, with other GSLB sites, network metrics such as round-trip time (RTT), learned from communications with various local DNS (LDNS) servers used by clients. RTT information is used in the dynamic RTT load balancing method, and is exchanged every 5 seconds.",
			},
			"parentsite": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Parent site of the GSLB site, in a parent-child topology.",
			},
			"publicclip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address to be used to globally access the remote cluster when it is deployed behind a NAT. It can be same as the normal cluster IP address.",
			},
			"publicip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Public IP address for the local site. Required only if the appliance is deployed in a private address space and the site has a public IP address hosted on an external firewall or a NAT device.",
			},
			"sessionexchange": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Exchange persistent session entries with other GSLB sites every five seconds.",
			},
			"siteipaddress": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address for the GSLB site. The GSLB site uses this IP address to communicate with other GSLB sites. For a local site, use any IP address that is owned by the appliance (for example, a SNIP or MIP address, or the IP address of the ADNS service).",
			},
			"sitename": schema.StringAttribute{
				Required:    true,
				Description: "Name for the GSLB site. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the virtual server is created.\n\nCLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my gslbsite\" or 'my gslbsite').",
			},
			"sitepassword": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Password to be used for mep communication between gslb site nodes.",
			},
			"sitetype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of site to create. If the type is not specified, the appliance automatically detects and sets the type on the basis of the IP address being assigned to the site. If the specified site IP address is owned by the appliance (for example, a MIP address or SNIP address), the site is a local site. Otherwise, it is a remote site.",
			},
			"triggermonitor": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify the conditions under which the GSLB service must be monitored by a monitor, if one is bound. Available settings function as follows:\n* ALWAYS - Monitor the GSLB service at all times.\n* MEPDOWN - Monitor the GSLB service only when the exchange of metrics through the Metrics Exchange Protocol (MEP) is disabled.\nMEPDOWN_SVCDOWN - Monitor the service in either of the following situations:\n* The exchange of metrics through MEP is disabled.\n* The exchange of metrics through MEP is enabled but the status of the service, learned through metrics exchange, is DOWN.",
			},
		},
	}
}
