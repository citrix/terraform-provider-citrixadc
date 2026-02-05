package service

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func ServiceDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"internal": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Display only dynamically learned services.",
			},
			"accessdown": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use Layer 2 mode to bridge the packets sent to this service if it is marked as DOWN. If the service is DOWN, and this parameter is disabled, the packets are dropped.",
			},
			"all": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Display both user-configured and dynamically learned services.",
			},
			"appflowlog": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable logging of AppFlow information.",
			},
			"cacheable": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use the transparent cache redirection virtual server to forward requests to the cache server.\nNote: Do not specify this parameter if you set the Cache Type parameter.",
			},
			"cachetype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Cache type supported by the cache server.",
			},
			"cip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Before forwarding a request to the service, insert an HTTP header with the client's IPv4 or IPv6 address as its value. Used if the server needs the client's IP address for security, accounting, or other purposes, and setting the Use Source IP parameter is not a viable option.",
			},
			"cipheader": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name for the HTTP header whose value must be set to the IP address of the client. Used with the Client IP parameter. If you set the Client IP parameter, and you do not specify a name for the header, the appliance uses the header name specified for the global Client IP Header parameter (the cipHeader parameter in the set ns param CLI command or the Client IP Header parameter in the Configure HTTP Parameters dialog box at System > Settings > Change HTTP parameters). If the global Client IP Header parameter is not specified, the appliance inserts a header with the name \"client-ip.\"",
			},
			"cka": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable client keep-alive for the service.",
			},
			"cleartextport": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Port to which clear text data must be sent after the appliance decrypts incoming SSL traffic. Applicable to transparent SSL services.",
			},
			"clttimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time, in seconds, after which to terminate an idle client connection.",
			},
			"cmp": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable compression for the service.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any information about the service.",
			},
			"contentinspectionprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the ContentInspection profile that contains IPS/IDS communication related setting for the service",
			},
			"customserverid": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Unique identifier for the service. Used when the persistency type for the virtual server is set to Custom Server ID.",
			},
			"delay": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time, in seconds, allocated to the NetScaler for a graceful shutdown of the service. During this period, new requests are sent to the service only for clients who already have persistent sessions on the appliance. Requests from new clients are load balanced among other available services. After the delay time expires, no requests are sent to the service, and the service is marked as unavailable (OUT OF SERVICE).",
			},
			"dnsprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the DNS profile to be associated with the service. DNS profile properties will applied to the transactions processed by a service. This parameter is valid only for ADNS, ADNS-TCP and ADNS-DOT services.",
			},
			"downstateflush": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Flush all active transactions associated with a service whose state transitions from UP to DOWN. Do not enable this option for applications that must complete their transactions.",
			},
			"graceful": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Shut down gracefully, not accepting any new connections, and disabling the service when all of its connections are closed.",
			},
			"hashid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "A numerical identifier that can be used by hash based load balancing methods. Must be unique for each service.",
			},
			"healthmonitor": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Monitor the health of this service. Available settings function as follows:\nYES - Send probes to check the health of the service.\nNO - Do not send probes to check the health of the service. With the NO option, the appliance shows the service as UP at all times.",
			},
			"httpprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the HTTP profile that contains HTTP configuration settings for the service.",
			},
			"ip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP to assign to the service.",
			},
			"ipaddress": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The new IP address of the service.",
			},
			"maxbandwidth": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum bandwidth, in Kbps, allocated to the service.",
			},
			"maxclient": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of simultaneous open connections to the service.",
			},
			"maxreq": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of requests that can be sent on a persistent connection to the service.\nNote: Connection requests beyond this value are rejected.",
			},
			"monconnectionclose": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Close monitoring connections by sending the service a connection termination message with the specified bit set.",
			},
			"monitor_name_svc": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the monitor bound to the specified service.",
			},
			"monthreshold": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Minimum sum of weights of the monitors that are bound to this service. Used to determine whether to mark a service as UP or DOWN.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the service. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the service has been created.",
			},
			"netprofile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Network profile to use for the service.",
			},
			"newname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "New name for the service. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.",
			},
			"pathmonitor": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Path monitoring for clustering",
			},
			"pathmonitorindv": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Individual Path monitoring decisions",
			},
			"port": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Port number of the service.",
			},
			"processlocal": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "By turning on this option packets destined to a service in a cluster will not under go any steering. Turn this option for single packet request response mode or when the upstream device is performing a proper RSS for connection based distribution.",
			},
			"quicprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of QUIC profile which will be attached to the service.",
			},
			"rtspsessionidremap": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable RTSP session ID mapping for the service.",
			},
			"serverid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The  identifier for the service. This is used when the persistency type is set to Custom Server ID.",
			},
			"servername": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the server that hosts the service.",
			},
			"servicetype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Protocol in which data is exchanged with the service.",
			},
			"sp": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable surge protection for the service.",
			},
			"state": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Initial state of the service.",
			},
			"svrtimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time, in seconds, after which to terminate an idle server connection.",
			},
			"tcpb": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable TCP buffering for the service.",
			},
			"tcpprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the TCP profile that contains TCP configuration settings for the service.",
			},
			"td": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.",
			},
			"useproxyport": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use the proxy port as the source port when initiating connections with the server. With the NO setting, the client-side connection port is used as the source port for the server-side connection.\nNote: This parameter is available only when the Use Source IP (USIP) parameter is set to YES.",
			},
			"usip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use the client's IP address as the source IP address when initiating a connection to the server. When creating a service, if you do not set this parameter, the service inherits the global Use Source IP setting (available in the enable ns mode and disable ns mode CLI commands, or in the System > Settings > Configure modes > Configure Modes dialog box). However, you can override this setting after you create the service.",
			},
			"weight": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Weight to assign to the monitor-service binding. When a monitor is UP, the weight assigned to its binding with the service determines how much the monitor contributes toward keeping the health of the service above the value configured for the Monitor Threshold parameter.",
			},
		},
	}
}
