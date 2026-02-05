package gslbservicegroup

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func GslbservicegroupDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"appflowlog": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable logging of AppFlow information for the specified GSLB service group.",
			},
			"autodelayedtrofs": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Indicates graceful movement of the service to TROFS. System will wait for monitor response time out before moving to TROFS",
			},
			"autoscale": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Auto scale option for a GSLB servicegroup",
			},
			"cip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Insert the Client IP header in requests forwarded to the GSLB service.",
			},
			"cipheader": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the HTTP header whose value must be set to the IP address of the client. Used with the Client IP parameter. If client IP insertion is enabled, and the client IP header is not specified, the value of Client IP Header parameter or the value set by the set ns config command is used as client's IP header name.",
			},
			"clttimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time, in seconds, after which to terminate an idle client connection.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any information about the GSLB service group.",
			},
			"delay": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The time allowed (in seconds) for a graceful shutdown. During this period, new connections or requests will continue to be sent to this service for clients who already have a persistent session on the system. Connections or requests from fresh or new clients who do not yet have a persistence sessions on the system will not be sent to the service. Instead, they will be load balanced among other available services. After the delay time expires, no new requests or connections will be sent to the service.",
			},
			"downstateflush": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Flush all active transactions associated with all the services in the GSLB service group whose state transitions from UP to DOWN. Do not enable this option for applications that must complete their transactions.",
			},
			"dup_weight": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "weight of the monitor that is bound to GSLB servicegroup.",
			},
			"graceful": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Wait for all existing connections to the service to terminate before shutting down the service.",
			},
			"hashid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The hash identifier for the service. This must be unique for each service. This parameter is used by hash based load balancing methods.",
			},
			"healthmonitor": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Monitor the health of this GSLB service.Available settings function are as follows:\nYES - Send probes to check the health of the GSLB service.\nNO - Do not send probes to check the health of the GSLB service. With the NO option, the appliance shows the service as UP at all times.",
			},
			"includemembers": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Display the members of the listed GSLB service groups in addition to their settings. Can be specified when no service group name is provided in the command. In that case, the details displayed for each service group are identical to the details displayed when a service group name is provided, except that bound monitors are not displayed.",
			},
			"maxbandwidth": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum bandwidth, in Kbps, allocated for all the services in the GSLB service group.",
			},
			"maxclient": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of simultaneous open connections for the GSLB service group.",
			},
			"monitor_name_svc": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the monitor bound to the GSLB service group. Used to assign a weight to the monitor.",
			},
			"monthreshold": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Minimum sum of weights of the monitors that are bound to this GSLB service. Used to determine whether to mark a GSLB service as UP or DOWN.",
			},
			"newname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "New name for the GSLB service group.",
			},
			"order": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Order number to be assigned to the gslb servicegroup member",
			},
			"port": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Server port number.",
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
				Description: "Name of the server to which to bind the service group.",
			},
			"servicegroupname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the GSLB service group. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Can be changed after the name is created.",
			},
			"servicetype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Protocol used to exchange data with the GSLB service.",
			},
			"sitename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the GSLB site to which the service group belongs.",
			},
			"sitepersistence": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use cookie-based site persistence. Applicable only to HTTP and SSL non-autoscale enabled GSLB servicegroups.",
			},
			"siteprefix": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The site's prefix string. When the GSLB service group is bound to a GSLB virtual server, a GSLB site domain is generated internally for each bound serviceitem-domain pair by concatenating the site prefix of the service item and the name of the domain. If the special string NONE is specified, the site-prefix string is unset. When implementing HTTP redirect site persistence, the Citrix ADC redirects GSLB requests to GSLB services by using their site domains.",
			},
			"state": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Initial state of the GSLB service group.",
			},
			"svrtimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time, in seconds, after which to terminate an idle server connection.",
			},
			"weight": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Weight to assign to the servers in the service group. Specifies the capacity of the servers relative to the other servers in the load balancing configuration. The higher the weight, the higher the percentage of requests sent to the service.",
			},
		},
	}
}
