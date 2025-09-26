---
subcategory: "Load Balancing"
---

# Resource: lbmonitor

The lbmonitor resource is used to create load balancing monitors.


## Example usage

```hcl
resource "citrixadc_lbmonitor" "tf_lbmonitor" {
  monitorname = "tf_lbmonitor"
  type = "HTTP"
}
```


## Argument Reference

* `monitorname` - (Optional) Name for the monitor. 
* `type` - (Optional) Type of monitor that you want to create. Possible values: [ PING, TCP, HTTP, TCP-ECV, HTTP-ECV, UDP-ECV, DNS, FTP, LDNS-PING, LDNS-TCP, LDNS-DNS, RADIUS, USER, HTTP-INLINE, SIP-UDP, SIP-TCP, LOAD, FTP-EXTENDED, SMTP, SNMP, NNTP, MYSQL, MYSQL-ECV, MSSQL-ECV, ORACLE-ECV, LDAP, POP3, CITRIX-XML-SERVICE, CITRIX-WEB-INTERFACE, DNS-TCP, RTSP, ARP, CITRIX-AG, CITRIX-AAC-LOGINPAGE, CITRIX-AAC-LAS, CITRIX-XD-DDC, ND6, CITRIX-WI-EXTENDED, DIAMETER, RADIUS_ACCOUNTING, STOREFRONT, APPC, SMPP, CITRIX-XNC-ECV, CITRIX-XDM, CITRIX-STA-SERVICE, CITRIX-STA-SERVICE-NHOP ]
* `action` - (Optional) Action to perform when the response to an inline monitor (a monitor of type HTTP-INLINE) indicates that the service is down. A service monitored by an inline monitor is considered DOWN if the response code is not one of the codes that have been specified for the Response Code parameter. Available settings function as follows: NONE - Do not take any action. However, the show service command and the show lb monitor command indicate the total number of responses that were checked and the number of consecutive error responses received after the last successful probe. LOG - Log the event in NSLOG or SYSLOG. DOWN - Mark the service as being down, and then do not direct any traffic to the service until the configured down time has expired. Persistent connections to the service are terminated as soon as the service is marked as DOWN. Also, log the event in NSLOG or SYSLOG. Possible values: [ NONE, LOG, DOWN ]
* `respcode` - (Optional) Response codes for which to mark the service as UP. For any other response code, the action performed depends on the monitor type. HTTP monitors and RADIUS monitors mark the service as DOWN, while HTTP-INLINE monitors perform the action indicated by the Action parameter.
* `httprequest` - (Optional) HTTP request to send to the server (for example, "HEAD /file.html").
* `rtsprequest` - (Optional) RTSP request to send to the server (for example, "OPTIONS \*").
* `customheaders` - (Optional) Custom header string to include in the monitoring probes.
* `maxforwards` - (Optional) 
* `sipmethod` - (Optional) SIP method to use for the query. Applicable only to monitors of type SIP-UDP. Possible values: [ OPTIONS, INVITE, REGISTER ]
* `sipuri` - (Optional) SIP URI string to send to the service (for example, sip:sip.test). Applicable only to monitors of type SIP-UDP.
* `sipreguri` - (Optional) SIP user to be registered. Applicable only if the monitor is of type SIP-UDP and the SIP Method parameter is set to REGISTER.
* `send` - (Optional) String to send to the service. Applicable to TCP-ECV, HTTP-ECV, and UDP-ECV monitors.
* `recv` - (Optional) String expected from the server for the service to be marked as UP. Applicable to TCP-ECV, HTTP-ECV, and UDP-ECV monitors.
* `query` - (Optional) Domain name to resolve as part of monitoring the DNS service (for example, example.com).
* `querytype` - (Optional) Type of DNS record for which to send monitoring queries. Set to Address for querying A records, AAAA for querying AAAA records, and Zone for querying the SOA record. Possible values: [ Address, Zone, AAAA ]
* `scriptname` - (Optional) Path and name of the script to execute. The script must be available on the Citrix ADC, in the /nsconfig/monitors/ directory.
* `scriptargs` - (Optional) String of arguments for the script. The string is copied verbatim into the request.
* `dispatcherip` - (Optional) IP address of the dispatcher to which to send the probe.
* `dispatcherport` - (Optional) Port number on which the dispatcher listens for the monitoring probe.
* `username` - (Optional) User name with which to probe the RADIUS, NNTP, FTP, FTP-EXTENDED, MYSQL, MSSQL, POP3, CITRIX-AG, CITRIX-XD-DDC, CITRIX-WI-EXTENDED, CITRIX-XNC or CITRIX-XDM server.
* `password` - (Optional) Password that is required for logging on to the RADIUS, NNTP, FTP, FTP-EXTENDED, MYSQL, MSSQL, POP3, CITRIX-AG, CITRIX-XD-DDC, CITRIX-WI-EXTENDED, CITRIX-XNC-ECV or CITRIX-XDM server. Used in conjunction with the user name specified for the User Name parameter.
* `secondarypassword` - (Optional) Secondary password that users might have to provide to log on to the Access Gateway server. Applicable to CITRIX-AG monitors.
* `logonpointname` - (Optional) Name of the logon point that is configured for the Citrix Access Gateway Advanced Access Control software. Required if you want to monitor the associated login page or Logon Agent. Applicable to CITRIX-AAC-LAS and CITRIX-AAC-LOGINPAGE monitors.
* `lasversion` - (Optional) Version number of the Citrix Advanced Access Control Logon Agent. Required by the CITRIX-AAC-LAS monitor.
* `radkey` - (Optional) Authentication key (shared secret text string) for RADIUS clients and servers to exchange. Applicable to monitors of type RADIUS and RADIUS_ACCOUNTING.
* `radnasid` - (Optional) NAS-Identifier to send in the Access-Request packet. Applicable to monitors of type RADIUS.
* `radnasip` - (Optional) Network Access Server (NAS) IP address to use as the source IP address when monitoring a RADIUS server. Applicable to monitors of type RADIUS and RADIUS_ACCOUNTING.
* `radaccounttype` - (Optional) Account Type to be used in Account Request Packet. Applicable to monitors of type RADIUS_ACCOUNTING.
* `radframedip` - (Optional) Source ip with which the packet will go out . Applicable to monitors of type RADIUS_ACCOUNTING.
* `radapn` - (Optional) Called Station Id to be used in Account Request Packet. Applicable to monitors of type RADIUS_ACCOUNTING.
* `radmsisdn` - (Optional) Calling Stations Id to be used in Account Request Packet. Applicable to monitors of type RADIUS_ACCOUNTING.
* `radaccountsession` - (Optional) Account Session ID to be used in Account Request Packet. Applicable to monitors of type RADIUS_ACCOUNTING.
* `lrtm` - (Optional) Calculate the least response times for bound services. If this parameter is not enabled, the appliance does not learn the response times of the bound services. Also used for LRTM load balancing. Possible values: [ ENABLED, DISABLED ]
* `deviation` - (Optional) Time value added to the learned average response time in dynamic response time monitoring (DRTM). When a deviation is specified, the appliance learns the average response time of bound services and adds the deviation to the average. The final value is then continually adjusted to accommodate response time variations over time. Specified in milliseconds, seconds, or minutes.
* `units1` - (Optional) Unit of measurement for the Deviation parameter. Cannot be changed after the monitor is created. Possible values: [ SEC, MSEC, MIN ]
* `interval` - (Optional) Time interval between two successive probes. Must be greater than the value of Response Time-out.
* `units3` - (Optional) monitor interval units. Possible values: [ SEC, MSEC, MIN ]
* `resptimeout` - (Optional) Amount of time for which the appliance must wait before it marks a probe as FAILED.  Must be less than the value specified for the Interval parameter. Note: For UDP-ECV monitors for which a receive string is not configured, response timeout does not apply. For UDP-ECV monitors with no receive string, probe failure is indicated by an ICMP port unreachable error received from the service.
* `units4` - (Optional) monitor response timeout units. Possible values: [ SEC, MSEC, MIN ]
* `resptimeoutthresh` - (Optional) Response time threshold, specified as a percentage of the Response Time-out parameter. If the response to a monitor probe has not arrived when the threshold is reached, the appliance generates an SNMP trap called monRespTimeoutAboveThresh. After the response time returns to a value below the threshold, the appliance generates a monRespTimeoutBelowThresh SNMP trap. For the traps to be generated, the "MONITOR-RTO-THRESHOLD" alarm must also be enabled.
* `retries` - (Optional) 
* `failureretries` - (Optional) Number of retries that must fail, out of the number specified for the Retries parameter, for a service to be marked as DOWN. For example, if the Retries parameter is set to 10 and the Failure Retries parameter is set to 6, out of the ten probes sent, at least six probes must fail if the service is to be marked as DOWN. The default value of 0 means that all the retries must fail if the service is to be marked as DOWN.
* `alertretries` - (Optional) Number of consecutive probe failures after which the appliance generates an SNMP trap called monProbeFailed.
* `successretries` - (Optional) Number of consecutive successful probes required to transition a service's state from DOWN to UP.
* `downtime` - (Optional) Time duration for which to wait before probing a service that has been marked as DOWN. Expressed in milliseconds, seconds, or minutes.
* `units2` - (Optional) Unit of measurement for the Down Time parameter. Cannot be changed after the monitor is created. Possible values: [ SEC, MSEC, MIN ]
* `destip` - (Optional) IP address of the service to which to send probes. If the parameter is set to 0, the IP address of the server to which the monitor is bound is considered the destination IP address.
* `destport` - (Optional) TCP or UDP port to which to send the probe. If the parameter is set to 0, the port number of the service to which the monitor is bound is considered the destination port. For a monitor of type USER, however, the destination port is the port number that is included in the HTTP request sent to the dispatcher. Does not apply to monitors of type PING.
* `state` - (Optional) State of the monitor. The DISABLED setting disables not only the monitor being configured, but all monitors of the same type, until the parameter is set to ENABLED. If the monitor is bound to a service, the state of the monitor is not taken into account when the state of the service is determined. Possible values: [ ENABLED, DISABLED ]
* `reverse` - (Optional) Mark a service as DOWN, instead of UP, when probe criteria are satisfied, and as UP instead of DOWN when probe criteria are not satisfied. Possible values: [ YES, NO ]
* `transparent` - (Optional) The monitor is bound to a transparent device such as a firewall or router. The state of a transparent device depends on the responsiveness of the services behind it. If a transparent device is being monitored, a destination IP address must be specified. The probe is sent to the specified IP address by using the MAC address of the transparent device. Possible values: [ YES, NO ]
* `iptunnel` - (Optional) Send the monitoring probe to the service through an IP tunnel. A destination IP address must be specified. Possible values: [ YES, NO ]
* `tos` - (Optional) Probe the service by encoding the destination IP address in the IP TOS (6) bits. Possible values: [ YES, NO ]
* `tosid` - (Optional) The TOS ID of the specified destination IP. Applicable only when the TOS parameter is set.
* `secure` - (Optional) Use a secure SSL connection when monitoring a service. Applicable only to TCP based monitors. The secure option cannot be used with a CITRIX-AG monitor, because a CITRIX-AG monitor uses a secure connection by default. Possible values: [ YES, NO ]
* `validatecred` - (Optional) Validate the credentials of the Xen Desktop DDC server user. Applicable to monitors of type CITRIX-XD-DDC. Possible values: [ YES, NO ]
* `domain` - (Optional) Domain in which the XenDesktop Desktop Delivery Controller (DDC) servers or Web Interface servers are present. Required by CITRIX-XD-DDC and CITRIX-WI-EXTENDED monitors for logging on to the DDC servers and Web Interface servers, respectively.
* `group` - (Optional) Name of a newsgroup available on the NNTP service that is to be monitored. The appliance periodically generates an NNTP query for the name of the newsgroup and evaluates the response. If the newsgroup is found on the server, the service is marked as UP. If the newsgroup does not exist or if the search fails, the service is marked as DOWN. Applicable to NNTP monitors.
* `filename` - (Optional) Name of a file on the FTP server. The appliance monitors the FTP service by periodically checking the existence of the file on the server. Applicable to FTP-EXTENDED monitors.
* `basedn` - (Optional) The base distinguished name of the LDAP service, from where the LDAP server can begin the search for the attributes in the monitoring query. Required for LDAP service monitoring.
* `binddn` - (Optional) The distinguished name with which an LDAP monitor can perform the Bind operation on the LDAP server. Optional. Applicable to LDAP monitors.
* `filter` - (Optional) Filter criteria for the LDAP query. Optional.
* `attribute` - (Optional) Attribute to evaluate when the LDAP server responds to the query. Success or failure of the monitoring probe depends on whether the attribute exists in the response. Optional.
* `database` - (Optional) Name of the database to connect to during authentication.
* `oraclesid` - (Optional) Name of the service identifier that is used to connect to the Oracle database during authentication.
* `sqlquery` - (Optional) SQL query for a MYSQL-ECV or MSSQL-ECV monitor. Sent to the database server after the server authenticates the connection.
* `evalrule` - (Optional) Expression that evaluates the database server's response to a MYSQL-ECV or MSSQL-ECV monitoring query. Must produce a Boolean result. The result determines the state of the server. If the expression returns TRUE, the probe succeeds. For example, if you want the appliance to evaluate the error message to determine the state of the server, use the rule MYSQL.RES.ROW(10) .TEXT_ELEM(2).EQ("MySQL").
* `mssqlprotocolversion` - (Optional) Version of MSSQL server that is to be monitored. Possible values: [ 70, 2000, 2000SP1, 2005, 2008, 2008R2, 2012, 2014 ]
* `snmpoid` - (Optional) SNMP OID for SNMP monitors.
* `snmpcommunity` - (Optional) Community name for SNMP monitors.
* `snmpthreshold` - (Optional) Threshold for SNMP monitors.
* `snmpversion` - (Optional) SNMP version to be used for SNMP monitors. Possible values: [ V1, V2 ]
* `metrictable` - (Optional) Metric table to which to bind metrics.
* `application` - (Optional) Name of the application used to determine the state of the service. Applicable to monitors of type CITRIX-XML-SERVICE.
* `sitepath` - (Optional) URL of the logon page. For monitors of type CITRIX-WEB-INTERFACE, to monitor a dynamic page under the site path, terminate the site path with a slash (/). Applicable to CITRIX-WEB-INTERFACE, CITRIX-WI-EXTENDED and CITRIX-XDM monitors.
* `storename` - (Optional) Store Name. For monitors of type STOREFRONT, STORENAME is an optional argument defining storefront service store name. Applicable to STOREFRONT monitors.
* `storefrontacctservice` - (Optional) Enable/Disable probing for Account Service. Applicable only to Store Front monitors. For multi-tenancy configuration users my skip account service. Possible values: [ YES, NO ]
* `hostname` - (Optional) Hostname in the FQDN format (Example: porche.cars.org). Applicable to STOREFRONT monitors.
* `netprofile` - (Optional) Name of the network profile.
* `originhost` - (Optional) Origin-Host value for the Capabilities-Exchange-Request (CER) message to use for monitoring Diameter servers.
* `originrealm` - (Optional) Origin-Realm value for the Capabilities-Exchange-Request (CER) message to use for monitoring Diameter servers.
* `hostipaddress` - (Optional) Host-IP-Address value for the Capabilities-Exchange-Request (CER) message to use for monitoring Diameter servers. If Host-IP-Address is not specified, the appliance inserts the mapped IP (MIP) address or subnet IP (SNIP) address from which the CER request (the monitoring probe) is sent.
* `vendorid` - (Optional) Vendor-Id value for the Capabilities-Exchange-Request (CER) message to use for monitoring Diameter servers.
* `productname` - (Optional) Product-Name value for the Capabilities-Exchange-Request (CER) message to use for monitoring Diameter servers.
* `firmwarerevision` - (Optional) Firmware-Revision value for the Capabilities-Exchange-Request (CER) message to use for monitoring Diameter servers.
* `inbandsecurityid` - (Optional) Inband-Security-Id for the Capabilities-Exchange-Request (CER) message to use for monitoring Diameter servers. Possible values: [ NO_INBAND_SECURITY, TLS ]
* `vendorspecificvendorid` - (Optional) Vendor-Id to use in the Vendor-Specific-Application-Id grouped attribute-value pair (AVP) in the monitoring CER message. To specify Auth-Application-Id or Acct-Application-Id in Vendor-Specific-Application-Id, use vendorSpecificAuthApplicationIds or vendorSpecificAcctApplicationIds, respectively. Only one Vendor-Id is supported for all the Vendor-Specific-Application-Id AVPs in a CER monitoring message.
* `kcdaccount` - (Optional) KCD Account used by MSSQL monitor.
* `storedb` - (Optional) Store the database list populated with the responses to monitor probes. Used in database specific load balancing if MSSQL-ECV/MYSQL-ECV  monitor is configured. Possible values: [ ENABLED, DISABLED ]
* `storefrontcheckbackendservices` - (Optional) This option will enable monitoring of services running on storefront server. Storefront services are monitored by probing to a Windows service that runs on the Storefront server and exposes details of which storefront services are running. Possible values: [ YES, NO ]
* `trofscode` - (Optional) Code expected when the server is under maintenance.
* `trofsstring` - (Optional) String expected from the server for the service to be marked as trofs. Applicable to HTTP-ECV/TCP-ECV monitors.
* `sslprofile` - (Optional) SSL Profile associated with the monitor.
* `metric` - (Optional) Metric name in the metric table, whose setting is changed. A value zero disables the metric and it will not be used for load calculation.
* `metricthreshold` - (Optional) Threshold to be used for that metric.
* `metricweight` - (Optional) The weight for the specified service metric with respect to others.
* `servicename` - (Optional) The name of the service to which the monitor is bound.
* `servicegroupname` - (Optional) The name of the service group to which the monitor is to be bound.
* `secureargs` - (Optional) List of arguments for the script which should be secure.
* `authapplicationid` - (Optional) List of Auth-Application-Id attribute value pairs (AVPs) for the Capabilities-Exchange-Request (CER) message to use for monitoring Diameter servers. A maximum of eight of these AVPs are supported in a monitoring CER message.
* `acctapplicationid` - (Optional) List of Acct-Application-Id attribute value pairs (AVPs) for the Capabilities-Exchange-Request (CER) message to use for monitoring Diameter servers. A maximum of eight of these AVPs are supported in a monitoring message.
* `supportedvendorids` - (Optional) List of Supported-Vendor-Id attribute value pairs (AVPs) for the Capabilities-Exchange-Request (CER) message to use for monitoring Diameter servers. A maximum eight of these AVPs are supported in a monitoring message.
* `vendorspecificauthapplicationids` - (Optional) List of Vendor-Specific-Auth-Application-Id attribute value pairs (AVPs) for the Capabilities-Exchange-Request (CER) message to use for monitoring Diameter servers. A maximum of eight of these AVPs are supported in a monitoring message. The specified value is combined with the value of vendorSpecificVendorId to obtain the Vendor-Specific-Application-Id AVP in the CER monitoring message.
* `vendorspecificacctapplicationids` - (Optional) List of Vendor-Specific-Acct-Application-Id attribute value pairs (AVPs) to use for monitoring Diameter servers. A maximum of eight of these AVPs are supported in a monitoring message. The specified value is combined with the value of vendorSpecificVendorId to obtain the Vendor-Specific-Application-Id AVP in the CER monitoring message.
* `mqttclientidentifier` - (Optional) Client id to be used in Connect command.
* `mqttversion` - (Optional) Version of MQTT protocol used in connect message, default is version 3.1.1 [4].
* `grpchealthcheck` - (Optional) Option to enable or disable gRPC health check service.
* `grpcstatuscode` - (Optional) gRPC status codes for which to mark the service as UP. The default value is 12(health check unimplemented). If the gRPC status code 0 is received from the backend this configuration is ignored.
* `grpcservicename` - (Optional) Option to specify gRPC service name on which gRPC health check need to be performed.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the lbmonitor. It has the same value as the `name` attribute.


## Import

A lbmonitor can be imported using its name, e.g.

```shell
terraform import citrixadc_lbmonitor.tf_lbmonitor tf_lbmonitor
```
