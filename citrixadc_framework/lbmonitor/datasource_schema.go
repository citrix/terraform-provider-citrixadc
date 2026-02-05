package lbmonitor

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func LbmonitorDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"snmpoid": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "SNMP OID for SNMP monitors.",
			},
			"acctapplicationid": schema.ListAttribute{
				ElementType: types.Int64Type,
				Optional:    true,
				Computed:    true,
				Description: "List of Acct-Application-Id attribute value pairs (AVPs) for the Capabilities-Exchange-Request (CER) message to use for monitoring Diameter servers. A maximum of eight of these AVPs are supported in a monitoring message.",
			},
			"action": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Action to perform when the response to an inline monitor (a monitor of type HTTP-INLINE) indicates that the service is down. A service monitored by an inline monitor is considered DOWN if the response code is not one of the codes that have been specified for the Response Code parameter.\nAvailable settings function as follows:\n* NONE - Do not take any action. However, the show service command and the show lb monitor command indicate the total number of responses that were checked and the number of consecutive error responses received after the last successful probe.\n* LOG - Log the event in NSLOG or SYSLOG.\n* DOWN - Mark the service as being down, and then do not direct any traffic to the service until the configured down time has expired. Persistent connections to the service are terminated as soon as the service is marked as DOWN. Also, log the event in NSLOG or SYSLOG.",
			},
			"alertretries": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of consecutive probe failures after which the appliance generates an SNMP trap called monProbeFailed.",
			},
			"application": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the application used to determine the state of the service. Applicable to monitors of type CITRIX-XML-SERVICE.",
			},
			"attribute": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Attribute to evaluate when the LDAP server responds to the query. Success or failure of the monitoring probe depends on whether the attribute exists in the response. Optional.",
			},
			"authapplicationid": schema.ListAttribute{
				ElementType: types.Int64Type,
				Optional:    true,
				Computed:    true,
				Description: "List of Auth-Application-Id attribute value pairs (AVPs) for the Capabilities-Exchange-Request (CER) message to use for monitoring Diameter servers. A maximum of eight of these AVPs are supported in a monitoring CER message.",
			},
			"basedn": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The base distinguished name of the LDAP service, from where the LDAP server can begin the search for the attributes in the monitoring query. Required for LDAP service monitoring.",
			},
			"binddn": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The distinguished name with which an LDAP monitor can perform the Bind operation on the LDAP server. Optional. Applicable to LDAP monitors.",
			},
			"customheaders": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Custom header string to include in the monitoring probes.",
			},
			"database": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the database to connect to during authentication.",
			},
			"destip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address of the service to which to send probes. If the parameter is set to 0, the IP address of the server to which the monitor is bound is considered the destination IP address.",
			},
			"destport": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "TCP or UDP port to which to send the probe. If the parameter is set to 0, the port number of the service to which the monitor is bound is considered the destination port. For a monitor of type USER, however, the destination port is the port number that is included in the HTTP request sent to the dispatcher. Does not apply to monitors of type PING.",
			},
			"deviation": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time value added to the learned average response time in dynamic response time monitoring (DRTM). When a deviation is specified, the appliance learns the average response time of bound services and adds the deviation to the average. The final value is then continually adjusted to accommodate response time variations over time. Specified in milliseconds, seconds, or minutes.",
			},
			"dispatcherip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address of the dispatcher to which to send the probe.",
			},
			"dispatcherport": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Port number on which the dispatcher listens for the monitoring probe.",
			},
			"domain": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Domain in which the XenDesktop Desktop Delivery Controller (DDC) servers or Web Interface servers are present. Required by CITRIX-XD-DDC and CITRIX-WI-EXTENDED monitors for logging on to the DDC servers and Web Interface servers, respectively.",
			},
			"downtime": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time duration for which to wait before probing a service that has been marked as DOWN. Expressed in milliseconds, seconds, or minutes.",
			},
			"evalrule": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression that evaluates the database server's response to a MYSQL-ECV or MSSQL-ECV monitoring query. Must produce a Boolean result. The result determines the state of the server. If the expression returns TRUE, the probe succeeds.\nFor example, if you want the appliance to evaluate the error message to determine the state of the server, use the rule MYSQL.RES.ROW(10) .TEXT_ELEM(2).EQ(\"MySQL\").",
			},
			"failureretries": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of retries that must fail, out of the number specified for the Retries parameter, for a service to be marked as DOWN. For example, if the Retries parameter is set to 10 and the Failure Retries parameter is set to 6, out of the ten probes sent, at least six probes must fail if the service is to be marked as DOWN. The default value of 0 means that all the retries must fail if the service is to be marked as DOWN.",
			},
			"filename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of a file on the FTP server. The appliance monitors the FTP service by periodically checking the existence of the file on the server. Applicable to FTP-EXTENDED monitors.",
			},
			"filter": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Filter criteria for the LDAP query. Optional.",
			},
			"firmwarerevision": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Firmware-Revision value for the Capabilities-Exchange-Request (CER) message to use for monitoring Diameter servers.",
			},
			"group": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of a newsgroup available on the NNTP service that is to be monitored. The appliance periodically generates an NNTP query for the name of the newsgroup and evaluates the response. If the newsgroup is found on the server, the service is marked as UP. If the newsgroup does not exist or if the search fails, the service is marked as DOWN. Applicable to NNTP monitors.",
			},
			"grpchealthcheck": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Option to enable or disable gRPC health check service.",
			},
			"grpcservicename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Option to specify gRPC service name on which gRPC health check need to be performed",
			},
			"grpcstatuscode": schema.ListAttribute{
				ElementType: types.Int64Type,
				Optional:    true,
				Computed:    true,
				Description: "gRPC status codes for which to mark the service as UP. The default value is 12(health check unimplemented). If the gRPC status code 0 is received from the backend this configuration is ignored.",
			},
			"hostipaddress": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Host-IP-Address value for the Capabilities-Exchange-Request (CER) message to use for monitoring Diameter servers. If Host-IP-Address is not specified, the appliance inserts the mapped IP (MIP) address or subnet IP (SNIP) address from which the CER request (the monitoring probe) is sent.",
			},
			"hostname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Hostname in the FQDN format (Example: porche.cars.org). Applicable to STOREFRONT monitors.",
			},
			"httprequest": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "HTTP request to send to the server (for example, \"HEAD /file.html\").",
			},
			"inbandsecurityid": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Inband-Security-Id for the Capabilities-Exchange-Request (CER) message to use for monitoring Diameter servers.",
			},
			"interval": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time interval between two successive probes. Must be greater than the value of Response Time-out.",
			},
			"ipaddress": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Set of IP addresses expected in the monitoring response from the DNS server, if the record type is A or AAAA. Applicable to DNS monitors.",
			},
			"iptunnel": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Send the monitoring probe to the service through an IP tunnel. A destination IP address must be specified.",
			},
			"kcdaccount": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "KCD Account used by MSSQL monitor",
			},
			"lasversion": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Version number of the Citrix Advanced Access Control Logon Agent. Required by the CITRIX-AAC-LAS monitor.",
			},
			"logonpointname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the logon point that is configured for the Citrix Access Gateway Advanced Access Control software. Required if you want to monitor the associated login page or Logon Agent. Applicable to CITRIX-AAC-LAS and CITRIX-AAC-LOGINPAGE monitors.",
			},
			"lrtm": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Calculate the least response times for bound services. If this parameter is not enabled, the appliance does not learn the response times of the bound services. Also used for LRTM load balancing.",
			},
			"maxforwards": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of hops that the SIP request used for monitoring can traverse to reach the server. Applicable only to monitors of type SIP-UDP.",
			},
			"metric": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Metric name in the metric table, whose setting is changed. A value zero disables the metric and it will not be used for load calculation",
			},
			"metrictable": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Metric table to which to bind metrics.",
			},
			"metricthreshold": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Threshold to be used for that metric.",
			},
			"metricweight": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The weight for the specified service metric with respect to others.",
			},
			"monitorname": schema.StringAttribute{
				Required:    true,
				Description: "Name for the monitor. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.\n\nCLI Users:  If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my monitor\" or 'my monitor').",
			},
			"mqttclientidentifier": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Client id to be used in Connect command",
			},
			"mqttversion": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Version of MQTT protocol used in connect message, default is version 3.1.1 [4]",
			},
			"mssqlprotocolversion": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Version of MSSQL server that is to be monitored.",
			},
			"netprofile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the network profile.",
			},
			"oraclesid": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the service identifier that is used to connect to the Oracle database during authentication.",
			},
			"originhost": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Origin-Host value for the Capabilities-Exchange-Request (CER) message to use for monitoring Diameter servers.",
			},
			"originrealm": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Origin-Realm value for the Capabilities-Exchange-Request (CER) message to use for monitoring Diameter servers.",
			},
			"password": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Password that is required for logging on to the RADIUS, NNTP, FTP, FTP-EXTENDED, MYSQL, MSSQL, POP3, CITRIX-AG, CITRIX-XD-DDC, CITRIX-WI-EXTENDED, CITRIX-XNC-ECV or CITRIX-XDM server. Used in conjunction with the user name specified for the User Name parameter.",
			},
			"productname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Product-Name value for the Capabilities-Exchange-Request (CER) message to use for monitoring Diameter servers.",
			},
			"query": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Domain name to resolve as part of monitoring the DNS service (for example, example.com).",
			},
			"querytype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of DNS record for which to send monitoring queries. Set to Address for querying A records, AAAA for querying AAAA records, and Zone for querying the SOA record.",
			},
			"radaccountsession": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Account Session ID to be used in Account Request Packet. Applicable to monitors of type RADIUS_ACCOUNTING.",
			},
			"radaccounttype": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Account Type to be used in Account Request Packet. Applicable to monitors of type RADIUS_ACCOUNTING.",
			},
			"radapn": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Called Station Id to be used in Account Request Packet. Applicable to monitors of type RADIUS_ACCOUNTING.",
			},
			"radframedip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Source ip with which the packet will go out . Applicable to monitors of type RADIUS_ACCOUNTING.",
			},
			"radkey": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Authentication key (shared secret text string) for RADIUS clients and servers to exchange. Applicable to monitors of type RADIUS and RADIUS_ACCOUNTING.",
			},
			"radmsisdn": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Calling Stations Id to be used in Account Request Packet. Applicable to monitors of type RADIUS_ACCOUNTING.",
			},
			"radnasid": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "NAS-Identifier to send in the Access-Request packet. Applicable to monitors of type RADIUS.",
			},
			"radnasip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Network Access Server (NAS) IP address to use as the source IP address when monitoring a RADIUS server. Applicable to monitors of type RADIUS and RADIUS_ACCOUNTING.",
			},
			"recv": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "String expected from the server for the service to be marked as UP. Applicable to TCP-ECV, HTTP-ECV, and UDP-ECV monitors.",
			},
			"respcode": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Response codes for which to mark the service as UP. For any other response code, the action performed depends on the monitor type. HTTP monitors and RADIUS monitors mark the service as DOWN, while HTTP-INLINE monitors perform the action indicated by the Action parameter.",
			},
			"resptimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Amount of time for which the appliance must wait before it marks a probe as FAILED.  Must be less than the value specified for the Interval parameter.\n\nNote: For UDP-ECV monitors for which a receive string is not configured, response timeout does not apply. For UDP-ECV monitors with no receive string, probe failure is indicated by an ICMP port unreachable error received from the service.",
			},
			"resptimeoutthresh": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Response time threshold, specified as a percentage of the Response Time-out parameter. If the response to a monitor probe has not arrived when the threshold is reached, the appliance generates an SNMP trap called monRespTimeoutAboveThresh. After the response time returns to a value below the threshold, the appliance generates a monRespTimeoutBelowThresh SNMP trap. For the traps to be generated, the \"MONITOR-RTO-THRESHOLD\" alarm must also be enabled.",
			},
			"retries": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of probes to send to establish the state of a service for which a monitoring probe failed.",
			},
			"reverse": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Mark a service as DOWN, instead of UP, when probe criteria are satisfied, and as UP instead of DOWN when probe criteria are not satisfied.",
			},
			"rtsprequest": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RTSP request to send to the server (for example, \"OPTIONS *\").",
			},
			"scriptargs": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "String of arguments for the script. The string is copied verbatim into the request.",
			},
			"scriptname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Path and name of the script to execute. The script must be available on the Citrix ADC, in the /nsconfig/monitors/ directory.",
			},
			"secondarypassword": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Secondary password that users might have to provide to log on to the Access Gateway server. Applicable to CITRIX-AG monitors.",
			},
			"secure": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use a secure SSL connection when monitoring a service. Applicable only to TCP based monitors. The secure option cannot be used with a CITRIX-AG monitor, because a CITRIX-AG monitor uses a secure connection by default.",
			},
			"secureargs": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "List of arguments for the script which should be secure",
			},
			"send": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "String to send to the service. Applicable to TCP-ECV, HTTP-ECV, and UDP-ECV monitors.",
			},
			"servicegroupname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The name of the service group to which the monitor is to be bound.",
			},
			"servicename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The name of the service to which the monitor is bound.",
			},
			"sipmethod": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "SIP method to use for the query. Applicable only to monitors of type SIP-UDP.",
			},
			"sipreguri": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "SIP user to be registered. Applicable only if the monitor is of type SIP-UDP and the SIP Method parameter is set to REGISTER.",
			},
			"sipuri": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "SIP URI string to send to the service (for example, sip:sip.test). Applicable only to monitors of type SIP-UDP.",
			},
			"sitepath": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "URL of the logon page. For monitors of type CITRIX-WEB-INTERFACE, to monitor a dynamic page under the site path, terminate the site path with a slash (/). Applicable to CITRIX-WEB-INTERFACE, CITRIX-WI-EXTENDED and CITRIX-XDM monitors.",
			},
			"snmpcommunity": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Community name for SNMP monitors.",
			},
			"snmpthreshold": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Threshold for SNMP monitors.",
			},
			"snmpversion": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "SNMP version to be used for SNMP monitors.",
			},
			"sqlquery": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "SQL query for a MYSQL-ECV or MSSQL-ECV monitor. Sent to the database server after the server authenticates the connection.",
			},
			"sslprofile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "SSL Profile associated with the monitor",
			},
			"state": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State of the monitor. The DISABLED setting disables not only the monitor being configured, but all monitors of the same type, until the parameter is set to ENABLED. If the monitor is bound to a service, the state of the monitor is not taken into account when the state of the service is determined.",
			},
			"storedb": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Store the database list populated with the responses to monitor probes. Used in database specific load balancing if MSSQL-ECV/MYSQL-ECV  monitor is configured.",
			},
			"storefrontacctservice": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable/Disable probing for Account Service. Applicable only to Store Front monitors. For multi-tenancy configuration users my skip account service",
			},
			"storefrontcheckbackendservices": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This option will enable monitoring of services running on storefront server. Storefront services are monitored by probing to a Windows service that runs on the Storefront server and exposes details of which storefront services are running.",
			},
			"storename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Store Name. For monitors of type STOREFRONT, STORENAME is an optional argument defining storefront service store name. Applicable to STOREFRONT monitors.",
			},
			"successretries": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of consecutive successful probes required to transition a service's state from DOWN to UP.",
			},
			"supportedvendorids": schema.ListAttribute{
				ElementType: types.Int64Type,
				Optional:    true,
				Computed:    true,
				Description: "List of Supported-Vendor-Id attribute value pairs (AVPs) for the Capabilities-Exchange-Request (CER) message to use for monitoring Diameter servers. A maximum eight of these AVPs are supported in a monitoring message.",
			},
			"tos": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Probe the service by encoding the destination IP address in the IP TOS (6) bits.",
			},
			"tosid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The TOS ID of the specified destination IP. Applicable only when the TOS parameter is set.",
			},
			"transparent": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The monitor is bound to a transparent device such as a firewall or router. The state of a transparent device depends on the responsiveness of the services behind it. If a transparent device is being monitored, a destination IP address must be specified. The probe is sent to the specified IP address by using the MAC address of the transparent device.",
			},
			"trofscode": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Code expected when the server is under maintenance",
			},
			"trofsstring": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "String expected from the server for the service to be marked as trofs. Applicable to HTTP-ECV/TCP-ECV monitors.",
			},
			"type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of monitor that you want to create.",
			},
			"units1": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Unit of measurement for the Deviation parameter. Cannot be changed after the monitor is created.",
			},
			"units2": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Unit of measurement for the Down Time parameter. Cannot be changed after the monitor is created.",
			},
			"units3": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "monitor interval units",
			},
			"units4": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "monitor response timeout units",
			},
			"username": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "User name with which to probe the RADIUS, NNTP, FTP, FTP-EXTENDED, MYSQL, MSSQL, POP3, CITRIX-AG, CITRIX-XD-DDC, CITRIX-WI-EXTENDED, CITRIX-XNC or CITRIX-XDM server.",
			},
			"validatecred": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Validate the credentials of the Xen Desktop DDC server user. Applicable to monitors of type CITRIX-XD-DDC.",
			},
			"vendorid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Vendor-Id value for the Capabilities-Exchange-Request (CER) message to use for monitoring Diameter servers.",
			},
			"vendorspecificacctapplicationids": schema.ListAttribute{
				ElementType: types.Int64Type,
				Optional:    true,
				Computed:    true,
				Description: "List of Vendor-Specific-Acct-Application-Id attribute value pairs (AVPs) to use for monitoring Diameter servers. A maximum of eight of these AVPs are supported in a monitoring message. The specified value is combined with the value of vendorSpecificVendorId to obtain the Vendor-Specific-Application-Id AVP in the CER monitoring message.",
			},
			"vendorspecificauthapplicationids": schema.ListAttribute{
				ElementType: types.Int64Type,
				Optional:    true,
				Computed:    true,
				Description: "List of Vendor-Specific-Auth-Application-Id attribute value pairs (AVPs) for the Capabilities-Exchange-Request (CER) message to use for monitoring Diameter servers. A maximum of eight of these AVPs are supported in a monitoring message. The specified value is combined with the value of vendorSpecificVendorId to obtain the Vendor-Specific-Application-Id AVP in the CER monitoring message.",
			},
			"vendorspecificvendorid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Vendor-Id to use in the Vendor-Specific-Application-Id grouped attribute-value pair (AVP) in the monitoring CER message. To specify Auth-Application-Id or Acct-Application-Id in Vendor-Specific-Application-Id, use vendorSpecificAuthApplicationIds or vendorSpecificAcctApplicationIds, respectively. Only one Vendor-Id is supported for all the Vendor-Specific-Application-Id AVPs in a CER monitoring message.",
			},
		},
	}
}
