/*
* Copyright (c) 2021 Citrix Systems, Inc.
*
*   Licensed under the Apache License, Version 2.0 (the "License");
*   you may not use this file except in compliance with the License.
*   You may obtain a copy of the License at
*
*       http://www.apache.org/licenses/LICENSE-2.0
*
*  Unless required by applicable law or agreed to in writing, software
*   distributed under the License is distributed on an "AS IS" BASIS,
*   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
*   See the License for the specific language governing permissions and
*   limitations under the License.
*/

package lb

/**
* Configuration for monitor resource.
*/
type Lbmonitor struct {
	/**
	* Name for the monitor. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.
		CLI Users:  If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my monitor" or 'my monitor').
	*/
	Monitorname string `json:"monitorname,omitempty"`
	/**
	* Type of monitor that you want to create.
	*/
	Type string `json:"type,omitempty"`
	/**
	* Action to perform when the response to an inline monitor (a monitor of type HTTP-INLINE) indicates that the service is down. A service monitored by an inline monitor is considered DOWN if the response code is not one of the codes that have been specified for the Response Code parameter.
		Available settings function as follows:
		* NONE - Do not take any action. However, the show service command and the show lb monitor command indicate the total number of responses that were checked and the number of consecutive error responses received after the last successful probe.
		* LOG - Log the event in NSLOG or SYSLOG.
		* DOWN - Mark the service as being down, and then do not direct any traffic to the service until the configured down time has expired. Persistent connections to the service are terminated as soon as the service is marked as DOWN. Also, log the event in NSLOG or SYSLOG.
	*/
	Action string `json:"action,omitempty"`
	/**
	* Response codes for which to mark the service as UP. For any other response code, the action performed depends on the monitor type. HTTP monitors and RADIUS monitors mark the service as DOWN, while HTTP-INLINE monitors perform the action indicated by the Action parameter.
	*/
	Respcode []string `json:"respcode,omitempty"`
	/**
	* HTTP request to send to the server (for example, "HEAD /file.html").
	*/
	Httprequest string `json:"httprequest,omitempty"`
	/**
	* RTSP request to send to the server (for example, "OPTIONS *").
	*/
	Rtsprequest string `json:"rtsprequest,omitempty"`
	/**
	* Custom header string to include in the monitoring probes.
	*/
	Customheaders string `json:"customheaders,omitempty"`
	/**
	* Maximum number of hops that the SIP request used for monitoring can traverse to reach the server. Applicable only to monitors of type SIP-UDP.
	*/
	Maxforwards *int `json:"maxforwards,omitempty"`
	/**
	* SIP method to use for the query. Applicable only to monitors of type SIP-UDP.
	*/
	Sipmethod string `json:"sipmethod,omitempty"`
	/**
	* SIP URI string to send to the service (for example, sip:sip.test). Applicable only to monitors of type SIP-UDP.
	*/
	Sipuri string `json:"sipuri,omitempty"`
	/**
	* SIP user to be registered. Applicable only if the monitor is of type SIP-UDP and the SIP Method parameter is set to REGISTER.
	*/
	Sipreguri string `json:"sipreguri,omitempty"`
	/**
	* String to send to the service. Applicable to TCP-ECV, HTTP-ECV, and UDP-ECV monitors.
	*/
	Send string `json:"send,omitempty"`
	/**
	* String expected from the server for the service to be marked as UP. Applicable to TCP-ECV, HTTP-ECV, and UDP-ECV monitors.
	*/
	Recv string `json:"recv,omitempty"`
	/**
	* Domain name to resolve as part of monitoring the DNS service (for example, example.com).
	*/
	Query string `json:"query,omitempty"`
	/**
	* Type of DNS record for which to send monitoring queries. Set to Address for querying A records, AAAA for querying AAAA records, and Zone for querying the SOA record.
	*/
	Querytype string `json:"querytype,omitempty"`
	/**
	* Path and name of the script to execute. The script must be available on the Citrix ADC, in the /nsconfig/monitors/ directory.
	*/
	Scriptname string `json:"scriptname,omitempty"`
	/**
	* String of arguments for the script. The string is copied verbatim into the request.
	*/
	Scriptargs string `json:"scriptargs,omitempty"`
	/**
	* List of arguments for the script which should be secure
	*/
	Secureargs string `json:"secureargs,omitempty"`
	/**
	* IP address of the dispatcher to which to send the probe.
	*/
	Dispatcherip string `json:"dispatcherip,omitempty"`
	/**
	* Port number on which the dispatcher listens for the monitoring probe.
	*/
	Dispatcherport *int `json:"dispatcherport,omitempty"`
	/**
	* User name with which to probe the RADIUS, NNTP, FTP, FTP-EXTENDED, MYSQL, MSSQL, POP3, CITRIX-AG, CITRIX-XD-DDC, CITRIX-WI-EXTENDED, CITRIX-XNC or CITRIX-XDM server.
	*/
	Username string `json:"username,omitempty"`
	/**
	* Password that is required for logging on to the RADIUS, NNTP, FTP, FTP-EXTENDED, MYSQL, MSSQL, POP3, CITRIX-AG, CITRIX-XD-DDC, CITRIX-WI-EXTENDED, CITRIX-XNC-ECV or CITRIX-XDM server. Used in conjunction with the user name specified for the User Name parameter.
	*/
	Password string `json:"password,omitempty"`
	/**
	* Secondary password that users might have to provide to log on to the Access Gateway server. Applicable to CITRIX-AG monitors.
	*/
	Secondarypassword string `json:"secondarypassword,omitempty"`
	/**
	* Name of the logon point that is configured for the Citrix Access Gateway Advanced Access Control software. Required if you want to monitor the associated login page or Logon Agent. Applicable to CITRIX-AAC-LAS and CITRIX-AAC-LOGINPAGE monitors.
	*/
	Logonpointname string `json:"logonpointname,omitempty"`
	/**
	* Version number of the Citrix Advanced Access Control Logon Agent. Required by the CITRIX-AAC-LAS monitor.
	*/
	Lasversion string `json:"lasversion,omitempty"`
	/**
	* Authentication key (shared secret text string) for RADIUS clients and servers to exchange. Applicable to monitors of type RADIUS and RADIUS_ACCOUNTING.
	*/
	Radkey string `json:"radkey,omitempty"`
	/**
	* NAS-Identifier to send in the Access-Request packet. Applicable to monitors of type RADIUS.
	*/
	Radnasid string `json:"radnasid,omitempty"`
	/**
	* Network Access Server (NAS) IP address to use as the source IP address when monitoring a RADIUS server. Applicable to monitors of type RADIUS and RADIUS_ACCOUNTING.
	*/
	Radnasip string `json:"radnasip,omitempty"`
	/**
	* Account Type to be used in Account Request Packet. Applicable to monitors of type RADIUS_ACCOUNTING.
	*/
	Radaccounttype *int `json:"radaccounttype,omitempty"`
	/**
	* Source ip with which the packet will go out . Applicable to monitors of type RADIUS_ACCOUNTING.
	*/
	Radframedip string `json:"radframedip,omitempty"`
	/**
	* Called Station Id to be used in Account Request Packet. Applicable to monitors of type RADIUS_ACCOUNTING.
	*/
	Radapn string `json:"radapn,omitempty"`
	/**
	* Calling Stations Id to be used in Account Request Packet. Applicable to monitors of type RADIUS_ACCOUNTING.
	*/
	Radmsisdn string `json:"radmsisdn,omitempty"`
	/**
	* Account Session ID to be used in Account Request Packet. Applicable to monitors of type RADIUS_ACCOUNTING.
	*/
	Radaccountsession string `json:"radaccountsession,omitempty"`
	/**
	* Calculate the least response times for bound services. If this parameter is not enabled, the appliance does not learn the response times of the bound services. Also used for LRTM load balancing.
	*/
	Lrtm string `json:"lrtm,omitempty"`
	/**
	* Time value added to the learned average response time in dynamic response time monitoring (DRTM). When a deviation is specified, the appliance learns the average response time of bound services and adds the deviation to the average. The final value is then continually adjusted to accommodate response time variations over time. Specified in milliseconds, seconds, or minutes.
	*/
	Deviation *int `json:"deviation"` // Zero is a valid value
	/**
	* Unit of measurement for the Deviation parameter. Cannot be changed after the monitor is created.
	*/
	Units1 string `json:"units1,omitempty"`
	/**
	* Time interval between two successive probes. Must be greater than the value of Response Time-out.
	*/
	Interval *int `json:"interval,omitempty"`
	/**
	* monitor interval units
	*/
	Units3 string `json:"units3,omitempty"`
	/**
	* Amount of time for which the appliance must wait before it marks a probe as FAILED.  Must be less than the value specified for the Interval parameter.
		Note: For UDP-ECV monitors for which a receive string is not configured, response timeout does not apply. For UDP-ECV monitors with no receive string, probe failure is indicated by an ICMP port unreachable error received from the service.
	*/
	Resptimeout *int `json:"resptimeout,omitempty"`
	/**
	* monitor response timeout units
	*/
	Units4 string `json:"units4,omitempty"`
	/**
	* Response time threshold, specified as a percentage of the Response Time-out parameter. If the response to a monitor probe has not arrived when the threshold is reached, the appliance generates an SNMP trap called monRespTimeoutAboveThresh. After the response time returns to a value below the threshold, the appliance generates a monRespTimeoutBelowThresh SNMP trap. For the traps to be generated, the "MONITOR-RTO-THRESHOLD" alarm must also be enabled.
	*/
	Resptimeoutthresh *int `json:"resptimeoutthresh,omitempty"`
	/**
	* Maximum number of probes to send to establish the state of a service for which a monitoring probe failed.
	*/
	Retries *int `json:"retries,omitempty"`
	/**
	* Number of retries that must fail, out of the number specified for the Retries parameter, for a service to be marked as DOWN. For example, if the Retries parameter is set to 10 and the Failure Retries parameter is set to 6, out of the ten probes sent, at least six probes must fail if the service is to be marked as DOWN. The default value of 0 means that all the retries must fail if the service is to be marked as DOWN.
	*/
	Failureretries *int `json:"failureretries,omitempty"`
	/**
	* Number of consecutive probe failures after which the appliance generates an SNMP trap called monProbeFailed.
	*/
	Alertretries *int `json:"alertretries,omitempty"`
	/**
	* Number of consecutive successful probes required to transition a service's state from DOWN to UP.
	*/
	Successretries *int `json:"successretries,omitempty"`
	/**
	* Time duration for which to wait before probing a service that has been marked as DOWN. Expressed in milliseconds, seconds, or minutes.
	*/
	Downtime *int `json:"downtime,omitempty"`
	/**
	* Unit of measurement for the Down Time parameter. Cannot be changed after the monitor is created.
	*/
	Units2 string `json:"units2,omitempty"`
	/**
	* IP address of the service to which to send probes. If the parameter is set to 0, the IP address of the server to which the monitor is bound is considered the destination IP address.
	*/
	Destip string `json:"destip,omitempty"`
	/**
	* TCP or UDP port to which to send the probe. If the parameter is set to 0, the port number of the service to which the monitor is bound is considered the destination port. For a monitor of type USER, however, the destination port is the port number that is included in the HTTP request sent to the dispatcher. Does not apply to monitors of type PING.
	*/
	Destport *int `json:"destport,omitempty"`
	/**
	* State of the monitor. The DISABLED setting disables not only the monitor being configured, but all monitors of the same type, until the parameter is set to ENABLED. If the monitor is bound to a service, the state of the monitor is not taken into account when the state of the service is determined.
	*/
	State string `json:"state,omitempty"`
	/**
	* Mark a service as DOWN, instead of UP, when probe criteria are satisfied, and as UP instead of DOWN when probe criteria are not satisfied.
	*/
	Reverse string `json:"reverse,omitempty"`
	/**
	* The monitor is bound to a transparent device such as a firewall or router. The state of a transparent device depends on the responsiveness of the services behind it. If a transparent device is being monitored, a destination IP address must be specified. The probe is sent to the specified IP address by using the MAC address of the transparent device.
	*/
	Transparent string `json:"transparent,omitempty"`
	/**
	* Send the monitoring probe to the service through an IP tunnel. A destination IP address must be specified.
	*/
	Iptunnel string `json:"iptunnel,omitempty"`
	/**
	* Probe the service by encoding the destination IP address in the IP TOS (6) bits.
	*/
	Tos string `json:"tos,omitempty"`
	/**
	* The TOS ID of the specified destination IP. Applicable only when the TOS parameter is set.
	*/
	Tosid *int `json:"tosid,omitempty"`
	/**
	* Use a secure SSL connection when monitoring a service. Applicable only to TCP based monitors. The secure option cannot be used with a CITRIX-AG monitor, because a CITRIX-AG monitor uses a secure connection by default.
	*/
	Secure string `json:"secure,omitempty"`
	/**
	* Validate the credentials of the Xen Desktop DDC server user. Applicable to monitors of type CITRIX-XD-DDC.
	*/
	Validatecred string `json:"validatecred,omitempty"`
	/**
	* Domain in which the XenDesktop Desktop Delivery Controller (DDC) servers or Web Interface servers are present. Required by CITRIX-XD-DDC and CITRIX-WI-EXTENDED monitors for logging on to the DDC servers and Web Interface servers, respectively.
	*/
	Domain string `json:"domain,omitempty"`
	/**
	* Set of IP addresses expected in the monitoring response from the DNS server, if the record type is A or AAAA. Applicable to DNS monitors.
	*/
	Ipaddress []string `json:"ipaddress,omitempty"`
	/**
	* Name of a newsgroup available on the NNTP service that is to be monitored. The appliance periodically generates an NNTP query for the name of the newsgroup and evaluates the response. If the newsgroup is found on the server, the service is marked as UP. If the newsgroup does not exist or if the search fails, the service is marked as DOWN. Applicable to NNTP monitors.
	*/
	Group string `json:"group,omitempty"`
	/**
	* Name of a file on the FTP server. The appliance monitors the FTP service by periodically checking the existence of the file on the server. Applicable to FTP-EXTENDED monitors.
	*/
	Filename string `json:"filename,omitempty"`
	/**
	* The base distinguished name of the LDAP service, from where the LDAP server can begin the search for the attributes in the monitoring query. Required for LDAP service monitoring.
	*/
	Basedn string `json:"basedn,omitempty"`
	/**
	* The distinguished name with which an LDAP monitor can perform the Bind operation on the LDAP server. Optional. Applicable to LDAP monitors.
	*/
	Binddn string `json:"binddn,omitempty"`
	/**
	* Filter criteria for the LDAP query. Optional.
	*/
	Filter string `json:"filter,omitempty"`
	/**
	* Attribute to evaluate when the LDAP server responds to the query. Success or failure of the monitoring probe depends on whether the attribute exists in the response. Optional.
	*/
	Attribute string `json:"attribute,omitempty"`
	/**
	* Name of the database to connect to during authentication.
	*/
	Database string `json:"database,omitempty"`
	/**
	* Name of the service identifier that is used to connect to the Oracle database during authentication.
	*/
	Oraclesid string `json:"oraclesid,omitempty"`
	/**
	* SQL query for a MYSQL-ECV or MSSQL-ECV monitor. Sent to the database server after the server authenticates the connection.
	*/
	Sqlquery string `json:"sqlquery,omitempty"`
	/**
	* Expression that evaluates the database server's response to a MYSQL-ECV or MSSQL-ECV monitoring query. Must produce a Boolean result. The result determines the state of the server. If the expression returns TRUE, the probe succeeds.
		For example, if you want the appliance to evaluate the error message to determine the state of the server, use the rule MYSQL.RES.ROW(10) .TEXT_ELEM(2).EQ("MySQL").
	*/
	Evalrule string `json:"evalrule,omitempty"`
	/**
	* Version of MSSQL server that is to be monitored.
	*/
	Mssqlprotocolversion string `json:"mssqlprotocolversion,omitempty"`
	/**
	* SNMP OID for SNMP monitors.
	*/
	Snmpoid string `json:"Snmpoid,omitempty"`
	/**
	* Community name for SNMP monitors.
	*/
	Snmpcommunity string `json:"snmpcommunity,omitempty"`
	/**
	* Threshold for SNMP monitors.
	*/
	Snmpthreshold string `json:"snmpthreshold,omitempty"`
	/**
	* SNMP version to be used for SNMP monitors.
	*/
	Snmpversion string `json:"snmpversion,omitempty"`
	/**
	* Metric table to which to bind metrics.
	*/
	Metrictable string `json:"metrictable,omitempty"`
	/**
	* Name of the application used to determine the state of the service. Applicable to monitors of type CITRIX-XML-SERVICE.
	*/
	Application string `json:"application,omitempty"`
	/**
	* URL of the logon page. For monitors of type CITRIX-WEB-INTERFACE, to monitor a dynamic page under the site path, terminate the site path with a slash (/). Applicable to CITRIX-WEB-INTERFACE, CITRIX-WI-EXTENDED and CITRIX-XDM monitors.
	*/
	Sitepath string `json:"sitepath,omitempty"`
	/**
	* Store Name. For monitors of type STOREFRONT, STORENAME is an optional argument defining storefront service store name. Applicable to STOREFRONT monitors.
	*/
	Storename string `json:"storename,omitempty"`
	/**
	* Enable/Disable probing for Account Service. Applicable only to Store Front monitors. For multi-tenancy configuration users my skip account service
	*/
	Storefrontacctservice string `json:"storefrontacctservice,omitempty"`
	/**
	* Hostname in the FQDN format (Example: porche.cars.org). Applicable to STOREFRONT monitors.
	*/
	Hostname string `json:"hostname,omitempty"`
	/**
	* Name of the network profile.
	*/
	Netprofile string `json:"netprofile,omitempty"`
	/**
	* Origin-Host value for the Capabilities-Exchange-Request (CER) message to use for monitoring Diameter servers.
	*/
	Originhost string `json:"originhost,omitempty"`
	/**
	* Origin-Realm value for the Capabilities-Exchange-Request (CER) message to use for monitoring Diameter servers.
	*/
	Originrealm string `json:"originrealm,omitempty"`
	/**
	* Host-IP-Address value for the Capabilities-Exchange-Request (CER) message to use for monitoring Diameter servers. If Host-IP-Address is not specified, the appliance inserts the mapped IP (MIP) address or subnet IP (SNIP) address from which the CER request (the monitoring probe) is sent.
	*/
	Hostipaddress string `json:"hostipaddress,omitempty"`
	/**
	* Vendor-Id value for the Capabilities-Exchange-Request (CER) message to use for monitoring Diameter servers.
	*/
	Vendorid *int `json:"vendorid,omitempty"`
	/**
	* Product-Name value for the Capabilities-Exchange-Request (CER) message to use for monitoring Diameter servers.
	*/
	Productname string `json:"productname,omitempty"`
	/**
	* Firmware-Revision value for the Capabilities-Exchange-Request (CER) message to use for monitoring Diameter servers.
	*/
	Firmwarerevision *int `json:"firmwarerevision,omitempty"`
	/**
	* List of Auth-Application-Id attribute value pairs (AVPs) for the Capabilities-Exchange-Request (CER) message to use for monitoring Diameter servers. A maximum of eight of these AVPs are supported in a monitoring CER message.
	*/
	Authapplicationid []int `json:"authapplicationid,omitempty"`
	/**
	* List of Acct-Application-Id attribute value pairs (AVPs) for the Capabilities-Exchange-Request (CER) message to use for monitoring Diameter servers. A maximum of eight of these AVPs are supported in a monitoring message.
	*/
	Acctapplicationid []int `json:"acctapplicationid,omitempty"`
	/**
	* Inband-Security-Id for the Capabilities-Exchange-Request (CER) message to use for monitoring Diameter servers.
	*/
	Inbandsecurityid string `json:"inbandsecurityid,omitempty"`
	/**
	* List of Supported-Vendor-Id attribute value pairs (AVPs) for the Capabilities-Exchange-Request (CER) message to use for monitoring Diameter servers. A maximum eight of these AVPs are supported in a monitoring message.
	*/
	Supportedvendorids []int `json:"supportedvendorids,omitempty"`
	/**
	* Vendor-Id to use in the Vendor-Specific-Application-Id grouped attribute-value pair (AVP) in the monitoring CER message. To specify Auth-Application-Id or Acct-Application-Id in Vendor-Specific-Application-Id, use vendorSpecificAuthApplicationIds or vendorSpecificAcctApplicationIds, respectively. Only one Vendor-Id is supported for all the Vendor-Specific-Application-Id AVPs in a CER monitoring message.
	*/
	Vendorspecificvendorid *int `json:"vendorspecificvendorid,omitempty"`
	/**
	* List of Vendor-Specific-Auth-Application-Id attribute value pairs (AVPs) for the Capabilities-Exchange-Request (CER) message to use for monitoring Diameter servers. A maximum of eight of these AVPs are supported in a monitoring message. The specified value is combined with the value of vendorSpecificVendorId to obtain the Vendor-Specific-Application-Id AVP in the CER monitoring message.
	*/
	Vendorspecificauthapplicationids []int `json:"vendorspecificauthapplicationids,omitempty"`
	/**
	* List of Vendor-Specific-Acct-Application-Id attribute value pairs (AVPs) to use for monitoring Diameter servers. A maximum of eight of these AVPs are supported in a monitoring message. The specified value is combined with the value of vendorSpecificVendorId to obtain the Vendor-Specific-Application-Id AVP in the CER monitoring message.
	*/
	Vendorspecificacctapplicationids []int `json:"vendorspecificacctapplicationids,omitempty"`
	/**
	* KCD Account used by MSSQL monitor
	*/
	Kcdaccount string `json:"kcdaccount,omitempty"`
	/**
	* Store the database list populated with the responses to monitor probes. Used in database specific load balancing if MSSQL-ECV/MYSQL-ECV  monitor is configured.
	*/
	Storedb string `json:"storedb,omitempty"`
	/**
	* This option will enable monitoring of services running on storefront server. Storefront services are monitored by probing to a Windows service that runs on the Storefront server and exposes details of which storefront services are running.
	*/
	Storefrontcheckbackendservices string `json:"storefrontcheckbackendservices,omitempty"`
	/**
	* Code expected when the server is under maintenance
	*/
	Trofscode *int `json:"trofscode,omitempty"`
	/**
	* String expected from the server for the service to be marked as trofs. Applicable to HTTP-ECV/TCP-ECV monitors.
	*/
	Trofsstring string `json:"trofsstring,omitempty"`
	/**
	* SSL Profile associated with the monitor
	*/
	Sslprofile string `json:"sslprofile,omitempty"`
	/**
	* Client id to be used in Connect command
	*/
	Mqttclientidentifier string `json:"mqttclientidentifier,omitempty"`
	/**
	* Version of MQTT protocol used in connect message, default is version 3.1.1 [4]
	*/
	Mqttversion *int `json:"mqttversion,omitempty"`
	/**
	* Option to enable or disable gRPC health check service.
	*/
	Grpchealthcheck string `json:"grpchealthcheck,omitempty"`
	/**
	* gRPC status codes for which to mark the service as UP. The default value is 12(health check unimplemented). If the gRPC status code 0 is received from the backend this configuration is ignored.
	*/
	Grpcstatuscode []int `json:"grpcstatuscode,omitempty"`
	/**
	* Option to specify gRPC service name on which gRPC health check need to be performed
	*/
	Grpcservicename string `json:"grpcservicename,omitempty"`
	/**
	* Metric name in the metric table, whose setting is changed. A value zero disables the metric and it will not be used for load calculation
	*/
	Metric string `json:"metric,omitempty"`
	/**
	* Threshold to be used for that metric.
	*/
	Metricthreshold *int `json:"metricthreshold,omitempty"`
	/**
	* The weight for the specified service metric with respect to others.
	*/
	Metricweight *int `json:"metricweight,omitempty"`
	/**
	* The name of the service to which the monitor is bound.
	*/
	Servicename string `json:"servicename,omitempty"`
	/**
	* The name of the service group to which the monitor is to be bound.
	*/
	Servicegroupname string `json:"servicegroupname,omitempty"`

	//------- Read only Parameter ---------;

	Lrtmconf string `json:"lrtmconf,omitempty"`
	Lrtmconfstr string `json:"lrtmconfstr,omitempty"`
	Dynamicresponsetimeout string `json:"dynamicresponsetimeout,omitempty"`
	Dynamicinterval string `json:"dynamicinterval,omitempty"`
	Multimetrictable string `json:"multimetrictable,omitempty"`
	Dupstate string `json:"dup_state,omitempty"`
	Dupweight string `json:"dup_weight,omitempty"`
	Weight string `json:"weight,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
