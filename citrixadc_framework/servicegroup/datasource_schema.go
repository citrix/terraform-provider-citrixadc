package servicegroup

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func ServicegroupDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"appflowlog": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable logging of AppFlow information for the specified service group.",
			},
			"autodelayedtrofs": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Indicates graceful movement of IP-Port binding/s to TROFS when IP addresses are removed from DNS response. System will wait for monitor response timeout period before moving to TROFS .",
			},
			"autodisabledelay": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The time allowed (in seconds) for a graceful shutdown. During this period, new connections or requests will continue to be sent to this service for clients who already have a persistent session on the system. Connections or requests from fresh or new clients who do not yet have a persistence sessions on the system will not be sent to the service. Instead, they will be load balanced among other available services. After the delay time expires, no new requests or connections will be sent to the service.",
			},
			"autodisablegraceful": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Indicates graceful shutdown of the service. System will wait for all outstanding connections to this service to be closed before disabling the service.",
			},
			"autoscale": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Auto scale option for a servicegroup",
			},
			"bootstrap": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Flag to check if kafka broker servicegroup is of type bootstrap or not.",
			},
			"cacheable": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use the transparent cache redirection virtual server to forward the request to the cache server.\nNote: Do not set this parameter if you set the Cache Type.",
			},
			"cachetype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Cache type supported by the cache server.",
			},
			"cip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Insert the Client IP header in requests forwarded to the service.",
			},
			"cipheader": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the HTTP header whose value must be set to the IP address of the client. Used with the Client IP parameter. If client IP insertion is enabled, and the client IP header is not specified, the value of Client IP Header parameter or the value set by the set ns config command is used as client's IP header name.",
			},
			"cka": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable client keep-alive for the service group.",
			},
			"clttimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time, in seconds, after which to terminate an idle client connection.",
			},
			"cmp": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable compression for the specified service.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any information about the service group.",
			},
			"customserverid": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The identifier for this IP:Port pair. Used when the persistency type is set to Custom Server ID.",
			},
			"dbsttl": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify the TTL for DNS record for domain based service.The default value of ttl is 0 which indicates to use the TTL received in DNS response for monitors",
			},
			"delay": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time, in seconds, allocated for a shutdown of the services in the service group. During this period, new requests are sent to the service only for clients who already have persistent sessions on the appliance. Requests from new clients are load balanced among other available services. After the delay time expires, no requests are sent to the service, and the service is marked as unavailable (OUT OF SERVICE).",
			},
			"downstateflush": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Flush all active transactions associated with all the services in the service group whose state transitions from UP to DOWN. Do not enable this option for applications that must complete their transactions.",
			},
			"dup_weight": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "weight of the monitor that is bound to servicegroup.",
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
				Description: "Monitor the health of this service.  Available settings function as follows:\nYES - Send probes to check the health of the service.\nNO - Do not send probes to check the health of the service. With the NO option, the appliance shows the service as UP at all times.",
			},
			"httpprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the HTTP profile that contains HTTP configuration settings for the service group.",
			},
			"includemembers": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Display the members of the listed service groups in addition to their settings. Can be specified when no service group name is provided in the command. In that case, the details displayed for each service group are identical to the details displayed when a service group name is provided, except that bound monitors are not displayed.",
			},
			"maxbandwidth": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum bandwidth, in Kbps, allocated for all the services in the service group.",
			},
			"maxclient": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of simultaneous open connections for the service group.",
			},
			"maxreq": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of requests that can be sent on a persistent connection to the service group.\nNote: Connection requests beyond this value are rejected.",
			},
			"memberport": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "member port",
			},
			"monconnectionclose": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Close monitoring connections by sending the service a connection termination message with the specified bit set.",
			},
			"monitor_name_svc": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the monitor bound to the service group. Used to assign a weight to the monitor.",
			},
			"monthreshold": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Minimum sum of weights of the monitors that are bound to this service. Used to determine whether to mark a service as UP or DOWN.",
			},
			"nameserver": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify the nameserver to which the query for bound domain needs to be sent. If not specified, use the global nameserver",
			},
			"netprofile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Network profile for the service group.",
			},
			"newname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "New name for the service group.",
			},
			"order": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Order number to be assigned to the servicegroup member",
			},
			"pathmonitor": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Path monitoring for clustering",
			},
			"pathmonitorindv": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Individual Path monitoring decisions.",
			},
			"port": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Server port number.",
			},
			"quicprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of QUIC profile which will be attached to the service group.",
			},
			"rtspsessionidremap": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable RTSP session ID mapping for the service group.",
			},
			"serverid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The  identifier for the service. This is used when the persistency type is set to Custom Server ID.",
			},
			"servername": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the server to which to bind the service group.",
			},
			"servicegroupname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the service group. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Can be changed after the name is created.",
			},
			"servicetype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Protocol used to exchange data with the service.",
			},
			"sp": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable surge protection for the service group.",
			},
			"state": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Initial state of the service group.",
			},
			"svrtimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time, in seconds, after which to terminate an idle server connection.",
			},
			"tcpb": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable TCP buffering for the service group.",
			},
			"tcpprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the TCP profile that contains TCP configuration settings for the service group.",
			},
			"td": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.",
			},
			"topicname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the Kafka topic.",
			},
			"useproxyport": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use the proxy port as the source port when initiating connections with the server. With the NO setting, the client-side connection port is used as the source port for the server-side connection.\nNote: This parameter is available only when the Use Source IP (USIP) parameter is set to YES.",
			},
			"usip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use client's IP address as the source IP address when initiating connection to the server. With the NO setting, which is the default, a mapped IP (MIP) address or subnet IP (SNIP) address is used as the source IP address to initiate server side connections.",
			},
			"weight": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Weight to assign to the servers in the service group. Specifies the capacity of the servers relative to the other servers in the load balancing configuration. The higher the weight, the higher the percentage of requests sent to the service.",
			},
		},
	}
}
