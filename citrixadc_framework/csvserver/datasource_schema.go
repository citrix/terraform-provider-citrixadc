package csvserver

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func CsvserverDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"apiprofile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The API profile where one or more API specs are bounded to.",
			},
			"appflowlog": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable logging appflow flow information",
			},
			"authentication": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Authenticate users who request a connection to the content switching virtual server.",
			},
			"authenticationhost": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "FQDN of the authentication virtual server. The service type of the virtual server should be either HTTP or SSL.",
			},
			"authn401": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable HTTP 401-response based authentication.",
			},
			"authnprofile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the authentication profile to be used when authentication is turned on.",
			},
			"authnvsname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of authentication virtual server that authenticates the incoming user requests to this content switching virtual server.",
			},
			"backupip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "0",
			},
			"backuppersistencetimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time period for which backup persistence is in effect.",
			},
			"backupvserver": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the backup virtual server that you are configuring. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Can be changed after the backup virtual server is created. You can assign a different backup virtual server or rename the existing virtual server.\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks.",
			},
			"cacheable": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use this option to specify whether a virtual server, used for load balancing or content switching, routes requests to the cache redirection virtual server before sending it to the configured servers.",
			},
			"casesensitive": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Consider case in URLs (for policies that use URLs instead of RULES). For example, with the ON setting, the URLs /a/1.html and /A/1.HTML are treated differently and can have different targets (set by content switching policies). With the OFF setting, /a/1.html and /A/1.HTML are switched to the same target.",
			},
			"clttimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Idle time, in seconds, after which the client connection is terminated. The default values are:\n180 seconds for HTTP/SSL-based services.\n9000 seconds for other TCP-based services.\n120 seconds for DNS-based services.\n120 seconds for other UDP-based services.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Information about this virtual server.",
			},
			"cookiedomain": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "0",
			},
			"cookiename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use this parameter to  specify the cookie name for COOKIE peristence type. It specifies the name of cookie with a maximum of 32 characters. If not specified, cookie name is internally generated.",
			},
			"cookietimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "0",
			},
			"dbprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the DB profile.",
			},
			"disableprimaryondown": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Continue forwarding the traffic to backup virtual server even after the primary server comes UP from the DOWN state.",
			},
			"dnsoverhttps": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This option is used to enable/disable DNS over HTTPS (DoH) processing.",
			},
			"dnsprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the DNS profile to be associated with the VServer. DNS profile properties will applied to the transactions processed by a VServer. This parameter is valid only for DNS and DNS-TCP VServers.",
			},
			"dnsrecordtype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "0",
			},
			"domainname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Domain name for which to change the time to live (TTL) and/or backup service IP address.",
			},
			"downstateflush": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Flush all active transactions associated with a virtual server whose state transitions from UP to DOWN. Do not enable this option for applications that must complete their transactions.",
			},
			"dtls": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This option starts/stops the dtls service on the vserver",
			},
			"httpprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the HTTP profile containing HTTP configuration settings for the virtual server. The service type of the virtual server should be either HTTP or SSL.",
			},
			"httpsredirecturl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "URL to which all HTTP traffic received on the port specified in the -redirectFromPort parameter is redirected.",
			},
			"icmpvsrresponse": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Can be active or passive",
			},
			"insertvserveripport": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Insert the virtual server's VIP address and port number in the request header. Available values function as follows:\n        VIPADDR - Header contains the vserver's IP address and port number without any translation.\n        OFF     - The virtual IP and port header insertion option is disabled.\n        V6TOV4MAPPING - Header contains the mapped IPv4 address corresponding to the IPv6 address of the vserver and the port number. An IPv6 address can be mapped to a user-specified IPv4 address using the set ns ip6 command.",
			},
			"ipmask": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP mask, in dotted decimal notation, for the IP Pattern parameter. Can have leading or trailing non-zero octets (for example, 255.255.240.0 or 0.0.255.255). Accordingly, the mask specifies whether the first n bits or the last n bits of the destination IP address in a client request are to be matched with the corresponding bits in the IP pattern. The former is called a forward mask. The latter is called a reverse mask.",
			},
			"ippattern": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address pattern, in dotted decimal notation, for identifying packets to be accepted by the virtual server. The IP Mask parameter specifies which part of the destination IP address is matched against the pattern. Mutually exclusive with the IP Address parameter.\nFor example, if the IP pattern assigned to the virtual server is 198.51.100.0 and the IP mask is 255.255.240.0 (a forward mask), the first 20 bits in the destination IP addresses are matched with the first 20 bits in the pattern. The virtual server accepts requests with IP addresses that range from 198.51.96.1 to 198.51.111.254. You can also use a pattern such as 0.0.2.2 and a mask such as 0.0.255.255 (a reverse mask).\nIf a destination IP address matches more than one IP pattern, the pattern with the longest match is selected, and the associated virtual server processes the request. For example, if the virtual servers, vs1 and vs2, have the same IP pattern, 0.0.100.128, but different IP masks of 0.0.255.255 and 0.0.224.255, a destination IP address of 198.51.100.128 has the longest match with the IP pattern of vs1. If a destination IP address matches two or more virtual servers to the same extent, the request is processed by the virtual server whose port number matches the port number in the request.",
			},
			"ipset": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The list of IPv4/IPv6 addresses bound to ipset would form a part of listening service on the current cs vserver",
			},
			"ipv46": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address of the content switching virtual server.",
			},
			"l2conn": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use L2 Parameters to identify a connection",
			},
			"listenpolicy": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "String specifying the listen policy for the content switching virtual server. Can be either the name of an existing expression or an in-line expression.",
			},
			"listenpriority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Integer specifying the priority of the listen policy. A higher number specifies a lower priority. If a request matches the listen policies of more than one virtual server the virtual server whose listen policy has the highest priority (the lowest priority number) accepts the request.",
			},
			"mssqlserverversion": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The version of the MSSQL server",
			},
			"mysqlcharacterset": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The character set returned by the mysql vserver.",
			},
			"mysqlprotocolversion": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The protocol version returned by the mysql vserver.",
			},
			"mysqlservercapabilities": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The server capabilities returned by the mysql vserver.",
			},
			"mysqlserverversion": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The server version string returned by the mysql vserver.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the content switching virtual server. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters.\nCannot be changed after the CS virtual server is created.\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, my server or my server).",
			},
			"netprofile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The name of the network profile.",
			},
			"newname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "New name for the virtual server. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters.\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my name\" or 'my name').",
			},
			"oracleserverversion": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Oracle server version",
			},
			"persistencebackup": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Backup persistence type for the virtual server. Becomes operational if the primary persistence mechanism fails.",
			},
			"persistenceid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "0",
			},
			"persistencetype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of persistence for the virtual server. Available settings function as follows:\n* SOURCEIP - Connections from the same client IP address belong to the same persistence session.\n* COOKIEINSERT - Connections that have the same HTTP Cookie, inserted by a Set-Cookie directive from a server, belong to the same persistence session.\n* SSLSESSION - Connections that have the same SSL Session ID belong to the same persistence session.",
			},
			"persistmask": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Persistence mask for IP based persistence types, for IPv4 virtual servers.",
			},
			"port": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Port number for content switching virtual server.",
			},
			"precedence": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of precedence to use for both RULE-based and URL-based policies on the content switching virtual server. With the default (RULE) setting, incoming requests are evaluated against the rule-based content switching policies. If none of the rules match, the URL in the request is evaluated against the URL-based content switching policies.",
			},
			"probeport": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Citrix ADC provides support for external health check of the vserver status. Select port for HTTP/TCP monitring",
			},
			"probeprotocol": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Citrix ADC provides support for external health check of the vserver status. Select HTTP or TCP probes for healthcheck",
			},
			"probesuccessresponsecode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "HTTP code to return in SUCCESS case.",
			},
			"push": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Process traffic with the push virtual server that is bound to this content switching virtual server (specified by the Push VServer parameter). The service type of the push virtual server should be either HTTP or SSL.",
			},
			"pushlabel": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression for extracting the label from the response received from server. This string can be either an existing rule name or an inline expression. The service type of the virtual server should be either HTTP or SSL.",
			},
			"pushmulticlients": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Allow multiple Web 2.0 connections from the same client to connect to the virtual server and expect updates.",
			},
			"pushvserver": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the load balancing virtual server, of type PUSH or SSL_PUSH, to which the server pushes updates received on the client-facing load balancing virtual server.",
			},
			"quicprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of QUIC profile which will be attached to the Content Switching VServer.",
			},
			"range": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of consecutive IP addresses, starting with the address specified by the IP Address parameter, to include in a range of addresses assigned to this virtual server.",
			},
			"redirectfromport": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Port number for the virtual server, from which we absorb the traffic for http redirect",
			},
			"redirectportrewrite": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State of port rewrite while performing HTTP redirect.",
			},
			"redirecturl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "URL to which traffic is redirected if the virtual server becomes unavailable. The service type of the virtual server should be either HTTP or SSL.\nCaution: Make sure that the domain in the URL does not match the domain specified for a content switching policy. If it does, requests are continuously redirected to the unavailable virtual server.",
			},
			"rhistate": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "A host route is injected according to the setting on the virtual servers\n            * If set to PASSIVE on all the virtual servers that share the IP address, the appliance always injects the hostroute.\n            * If set to ACTIVE on all the virtual servers that share the IP address, the appliance injects even if one virtual server is UP.\n            * If set to ACTIVE on some virtual servers and PASSIVE on the others, the appliance, injects even if one virtual server set to ACTIVE is UP.",
			},
			"rtspnat": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable network address translation (NAT) for real-time streaming protocol (RTSP) connections.",
			},
			"servicetype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Protocol used by the virtual server.",
			},
			"sitedomainttl": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "0",
			},
			"sobackupaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Action to be performed if spillover is to take effect, but no backup chain to spillover is usable or exists",
			},
			"somethod": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of spillover used to divert traffic to the backup virtual server when the primary virtual server reaches the spillover threshold. Connection spillover is based on the number of connections. Bandwidth spillover is based on the total Kbps of incoming and outgoing traffic.",
			},
			"sopersistence": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Maintain source-IP based persistence on primary and backup virtual servers.",
			},
			"sopersistencetimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time-out value, in minutes, for spillover persistence.",
			},
			"sothreshold": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Depending on the spillover method, the maximum number of connections or the maximum total bandwidth (Kbps) that a virtual server can handle before spillover occurs.",
			},
			"state": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Initial state of the load balancing virtual server.",
			},
			"stateupdate": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable state updates for a specific content switching virtual server. By default, the Content Switching virtual server is always UP, regardless of the state of the Load Balancing virtual servers bound to it. This parameter interacts with the global setting as follows:\nGlobal Level | Vserver Level | Result\nENABLED      ENABLED        ENABLED\nENABLED      DISABLED       ENABLED\nDISABLED     ENABLED        ENABLED\nDISABLED     DISABLED       DISABLED\nIf you want to enable state updates for only some content switching virtual servers, be sure to disable the state update parameter.",
			},
			"targettype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Virtual server target type.",
			},
			"tcpprobeport": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Port number for external TCP probe. NetScaler provides support for external TCP health check of the vserver status over the selected port. This option is only supported for vservers assigned with an IPAddress or ipset.",
			},
			"tcpprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the TCP profile containing TCP configuration settings for the virtual server.",
			},
			"td": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.",
			},
			"timeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time period for which a persistence session is in effect.",
			},
			"ttl": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "0",
			},
			"v6persistmasklen": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Persistence mask for IP based persistence types, for IPv6 virtual servers.",
			},
			"vipheader": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of virtual server IP and port header, for use with the VServer IP Port Insertion parameter.",
			},
		},
	}
}
