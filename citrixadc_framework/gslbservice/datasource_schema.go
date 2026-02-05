package gslbservice

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func GslbserviceDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"appflowlog": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable logging appflow flow information",
			},
			"cip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "In the request that is forwarded to the GSLB service, insert a header that stores the client's IP address. Client IP header insertion is used in connection-proxy based site persistence.",
			},
			"cipheader": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name for the HTTP header that stores the client's IP address. Used with the Client IP option. If client IP header insertion is enabled on the service and a name is not specified for the header, the Citrix ADC uses the name specified by the cipHeader parameter in the set ns param command or, in the GUI, the Client IP Header parameter in the Configure HTTP Parameters dialog box.",
			},
			"clttimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Idle time, in seconds, after which a client connection is terminated. Applicable if connection proxy based site persistence is used.",
			},
			"cnameentry": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Canonical name of the GSLB service. Used in CNAME-based GSLB.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments that you might want to associate with the GSLB service.",
			},
			"cookietimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Timeout value, in minutes, for the cookie, when cookie based site persistence is enabled.",
			},
			"downstateflush": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Flush all active transactions associated with the GSLB service when its state transitions from UP to DOWN. Do not enable this option for services that must complete their transactions. Applicable if connection proxy based site persistence is used.",
			},
			"hashid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Unique hash identifier for the GSLB service, used by hash based load balancing methods.",
			},
			"healthmonitor": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Monitor the health of the GSLB service.",
			},
			"ip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address for the GSLB service. Should represent a load balancing, content switching, or VPN virtual server on the Citrix ADC, or the IP address of another load balancing device.",
			},
			"ipaddress": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The new IP address of the service.",
			},
			"maxaaausers": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of SSL VPN users that can be logged on concurrently to the VPN virtual server that is represented by this GSLB service. A GSLB service whose user count reaches the maximum is not considered when a GSLB decision is made, until the count drops below the maximum.",
			},
			"maxbandwidth": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Integer specifying the maximum bandwidth allowed for the service. A GSLB service whose bandwidth reaches the maximum is not considered when a GSLB decision is made, until its bandwidth consumption drops below the maximum.",
			},
			"maxclient": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The maximum number of open connections that the service can support at any given time. A GSLB service whose connection count reaches the maximum is not considered when a GSLB decision is made, until the connection count drops below the maximum.",
			},
			"monitor_name_svc": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the monitor to bind to the service.",
			},
			"monthreshold": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Monitoring threshold value for the GSLB service. If the sum of the weights of the monitors that are bound to this GSLB service and are in the UP state is not equal to or greater than this threshold value, the service is marked as DOWN.",
			},
			"naptrdomainttl": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Modify the TTL of the internally created naptr domain",
			},
			"naptrorder": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "An integer specifying the order in which the NAPTR records MUST be processed in order to accurately represent the ordered list of Rules. The ordering is from lowest to highest",
			},
			"naptrpreference": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "An integer specifying the preference of this NAPTR among NAPTR records having same order. lower the number, higher the preference.",
			},
			"naptrreplacement": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The replacement domain name for this NAPTR.",
			},
			"naptrservices": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Service Parameters applicable to this delegation path.",
			},
			"newname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "New name for the GSLB service.",
			},
			"port": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Port on which the load balancing entity represented by this GSLB service listens.",
			},
			"publicip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The public IP address that a NAT device translates to the GSLB service's private IP address. Optional.",
			},
			"publicport": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The public port associated with the GSLB service's public IP address. The port is mapped to the service's private port number. Applicable to the local GSLB service. Optional.",
			},
			"servername": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the server hosting the GSLB service.",
			},
			"servicename": schema.StringAttribute{
				Required:    true,
				Description: "Name for the GSLB service. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Can be changed after the GSLB service is created.\n\nCLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my gslbsvc\" or 'my gslbsvc').",
			},
			"servicetype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of service to create.",
			},
			"sitename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the GSLB site to which the service belongs.",
			},
			"sitepersistence": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use cookie-based site persistence. Applicable only to HTTP and SSL GSLB services.",
			},
			"siteprefix": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The site's prefix string. When the service is bound to a GSLB virtual server, a GSLB site domain is generated internally for each bound service-domain pair by concatenating the site prefix of the service and the name of the domain. If the special string NONE is specified, the site-prefix string is unset. When implementing HTTP redirect site persistence, the Citrix ADC redirects GSLB requests to GSLB services by using their site domains.",
			},
			"state": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable or disable the service.",
			},
			"svrtimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Idle time, in seconds, after which a server connection is terminated. Applicable if connection proxy based site persistence is used.",
			},
			"viewip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address to be used for the given view",
			},
			"viewname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the DNS view of the service. A DNS view is used in global server load balancing (GSLB) to return a predetermined IP address to a specific group of clients, which are identified by using a DNS policy.",
			},
			"weight": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Weight to assign to the monitor-service binding. A larger number specifies a greater weight. Contributes to the monitoring threshold, which determines the state of the service.",
			},
		},
	}
}
