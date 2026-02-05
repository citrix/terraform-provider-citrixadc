---
subcategory: "Load Balancing"
---

# Data Source `lbmonitor`

The lbmonitor data source allows you to retrieve information about load balancing monitors.


## Example usage

```terraform
data "citrixadc_lbmonitor" "tf_lbmonitor" {
  monitorname = "my_lbmonitor"
}

output "type" {
  value = data.citrixadc_lbmonitor.tf_lbmonitor.type
}

output "interval" {
  value = data.citrixadc_lbmonitor.tf_lbmonitor.interval
}
```


## Argument Reference

* `monitorname` - (Required) Name for the monitor. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `acctapplicationid` - List of Acct-Application-Id attribute value pairs (AVPs) for the Capabilities-Exchange-Request (CER) message to use for monitoring Diameter servers. A maximum of eight of these AVPs are supported in a monitoring message.
* `action` - Action to perform when the response to an inline monitor (a monitor of type HTTP-INLINE) indicates that the service is down. A service monitored by an inline monitor is considered DOWN if the response code is not one of the codes that have been specified for the Response Code parameter.
* `alertretries` - Number of consecutive probe failures after which the appliance generates an SNMP trap called monProbeFailed.
* `application` - Name of the application used to determine the state of the service. Applicable to monitors of type CITRIX-XML-SERVICE.
* `attribute` - Attribute to evaluate when the LDAP server responds to the query. Success or failure of the monitoring probe depends on whether the attribute exists in the response. Optional.
* `authapplicationid` - List of Auth-Application-Id attribute value pairs (AVPs) for the Capabilities-Exchange-Request (CER) message to use for monitoring Diameter servers. A maximum of eight of these AVPs are supported in a monitoring CER message.
* `basedn` - The base distinguished name of the LDAP service, from where the LDAP server can begin the search for the attributes in the monitoring query. Required for LDAP service monitoring.
* `binddn` - The distinguished name with which an LDAP monitor can perform the Bind operation on the LDAP server. Optional. Applicable to LDAP monitors.
* `customheaders` - Custom header string to include in the monitoring probes.
* `database` - Name of the database to connect to during authentication.
* `destip` - IP address of the service to which to send probes. If the parameter is set to 0, the IP address of the server to which the monitor is bound is considered the destination IP address.
* `destport` - TCP or UDP port to which to send the probe. If the parameter is set to 0, the port number of the service to which the monitor is bound is considered the destination port. For a monitor of type USER, however, the destination port is the port number that is included in the HTTP request sent to the dispatcher. Does not apply to monitors of type PING.
* `deviation` - Time value added to the learned average response time in dynamic response time monitoring (DRTM). When a deviation is specified, the appliance learns the average response time of bound services and adds the deviation to the average. The final value is then continually adjusted to accommodate response time variations over time. Specified in milliseconds, seconds, or minutes.
* `dispatcherip` - IP address of the dispatcher to which to send the probe.
* `dispatcherport` - Port number on which the dispatcher listens for the monitoring probe.
* `domain` - Domain in which the XenDesktop Desktop Delivery Controller (DDC) servers or Web Interface servers are present. Required by CITRIX-XD-DDC and CITRIX-WI-EXTENDED monitors for logging on to the DDC servers and Web Interface servers, respectively.
* `downtime` - Time duration for which to wait before probing a service that has been marked as DOWN. Expressed in milliseconds, seconds, or minutes.
* `evalrule` - Expression that evaluates the database server's response to a MYSQL-ECV or MSSQL-ECV monitoring query. Must produce a Boolean result. The result determines the state of the server. If the expression returns TRUE, the probe succeeds.
* `failureretries` - Number of retries that must fail, out of the number specified for the Retries parameter, for a service to be marked as DOWN. For example, if the Retries parameter is set to 10 and the Failure Retries parameter is set to 6, out of the ten probes sent, at least six probes must fail if the service is to be marked as DOWN. The default value of 0 means that all the retries must fail if the service is to be marked as DOWN.
* `filename` - Name of a file on the FTP server. The appliance monitors the FTP service by periodically checking the existence of the file on the server. Applicable to FTP-EXTENDED monitors.
* `filter` - Filter criteria for the LDAP query. Optional.
* `firmwarerevision` - Firmware-Revision value for the Capabilities-Exchange-Request (CER) message to use for monitoring Diameter servers.
* `group` - Name of a newsgroup available on the NNTP service that is to be monitored. The appliance periodically generates an NNTP query for the name of the newsgroup and evaluates the response. If the newsgroup is found on the server, the service is marked as UP. If the newsgroup does not exist or if the search fails, the service is marked as DOWN. Applicable to NNTP monitors.
* `grpchealthcheck` - Option to enable or disable gRPC health check service.
* `grpcservicename` - Option to specify gRPC service name on which gRPC health check need to be performed
* `grpcstatuscode` - gRPC status codes for which to mark the service as UP. The default value is 12(health check unimplemented). If the gRPC status code 0 is received from the backend this configuration is ignored.
* `hostipaddress` - Host-IP-Address value for the Capabilities-Exchange-Request (CER) message to use for monitoring Diameter servers. If Host-IP-Address is not specified, the appliance inserts the mapped IP (MIP) address or subnet IP (SNIP) address from which the CER request (the monitoring probe) is sent.
* `hostname` - Hostname in the FQDN format (Example: porche.cars.org). Applicable to STOREFRONT monitors.
* `httprequest` - HTTP request to send to the server (for example, "HEAD /file.html").
* `inbandsecurityid` - Inband-Security-Id for the Capabilities-Exchange-Request (CER) message to use for monitoring Diameter servers.
* `interval` - Time interval between two successive probes. Must be greater than the value of Response Time-out.
* `ipaddress` - Set of IP addresses expected in the monitoring response from the DNS server, if the record type is A or AAAA. Applicable to DNS monitors.
* `iptunnel` - Send the monitoring probe to the service through an IP tunnel. A destination IP address must be specified.
* `kcdaccount` - KCD Account used by MSSQL monitor
* `lasversion` - Version number of the Citrix Advanced Access Control Logon Agent. Required by the CITRIX-AAC-LAS monitor.
* `logonpointname` - Name of the logon point that is configured for the Citrix Access Gateway Advanced Access Control software. Required if you want to monitor the associated login page or Logon Agent. Applicable to CITRIX-AAC-LAS and CITRIX-AAC-LOGINPAGE monitors.
* `lrtm` - Calculate the least response times for bound services. If this parameter is not enabled, the appliance does not learn the response times of the bound services. Also used for LRTM load balancing.
* `maxforwards` - Maximum number of hops that the SIP request used for monitoring can traverse to reach the server. Applicable only to monitors of type SIP-UDP.
* `metric` - Metric name in the metric table, whose setting is changed. A value zero disables the metric and it will not be used for load calculation
* `metrictable` - Metric table to which to bind metrics.
* `metricthreshold` - Threshold to be used for that metric.
* `metricweight` - The weight for the specified service metric with respect to others.
* `mqttclientidentifier` - Client id to be used in Connect command
* `mqttversion` - Version of MQTT protocol used in connect message, default is version 3.1.1 [4]
* `mssqlprotocolversion` - Version of MSSQL server that is to be monitored.
* `netprofile` - Name of the network profile.
* `oraclesid` - Name of the service identifier that is used to connect to the Oracle database during authentication.
* `originhost` - Origin-Host value for the Capabilities-Exchange-Request (CER) message to use for monitoring Diameter servers.
* `originrealm` - Origin-Realm value for the Capabilities-Exchange-Request (CER) message to use for monitoring Diameter servers.
* `password` - Password that is required for logging on to the RADIUS, NNTP, FTP, FTP-EXTENDED, MYSQL, MSSQL, POP3, CITRIX-AG, CITRIX-XD-DDC, CITRIX-WI-EXTENDED, CITRIX-XNC-ECV or CITRIX-XDM server. Used in conjunction with the user name specified for the User Name parameter.
* `productname` - Product-Name value for the Capabilities-Exchange-Request (CER) message to use for monitoring Diameter servers.
* `query` - Domain name to resolve as part of monitoring the DNS service (for example, example.com).
* `querytype` - Type of DNS record for which to send monitoring queries. Set to Address for querying A records, AAAA for querying AAAA records, and Zone for querying the SOA record.
* `radaccountsession` - Account Session ID to be used in Account Request Packet. Applicable to monitors of type RADIUS_ACCOUNTING.
* `radaccounttype` - Account Type to be used in Account Request Packet. Applicable to monitors of type RADIUS_ACCOUNTING.
* `radapn` - Called Station Id to be used in Account Request Packet. Applicable to monitors of type RADIUS_ACCOUNTING.
* `radframedip` - Source ip with which the packet will go out . Applicable to monitors of type RADIUS_ACCOUNTING.
* `radkey` - Authentication key (shared secret text string) for RADIUS clients and servers to exchange. Applicable to monitors of type RADIUS and RADIUS_ACCOUNTING.
* `radmsisdn` - Calling Stations Id to be used in Account Request Packet. Applicable to monitors of type RADIUS_ACCOUNTING.
* `radnasid` - NAS-Identifier to send in the Access-Request packet. Applicable to monitors of type RADIUS.
* `radnasip` - Network Access Server (NAS) IP address to use as the source IP address when monitoring a RADIUS server. Applicable to monitors of type RADIUS and RADIUS_ACCOUNTING.
* `recv` - String expected from the server for the service to be marked as UP. Applicable to TCP-ECV, HTTP-ECV, and UDP-ECV monitors.
* `respcode` - Response codes for which to mark the service as UP. For any other response code, the action performed depends on the monitor type. HTTP monitors and RADIUS monitors mark the service as DOWN, while HTTP-INLINE monitors perform the action indicated by the Action parameter.
* `resptimeout` - Amount of time for which the appliance must wait before it marks a probe as FAILED. Must be less than the value specified for the Interval parameter.
* `resptimeoutthresh` - Response time threshold, specified as a percentage of the Response Time-out parameter. If the response to a monitor probe has not arrived when the threshold is reached, the appliance generates an SNMP trap called monRespTimeoutAboveThresh.
* `retries` - Maximum number of probes to send to establish the state of a service for which a monitoring probe failed.
* `reverse` - Mark a service as DOWN, instead of UP, when probe criteria are satisfied, and as UP instead of DOWN when probe criteria are not satisfied.
* `rtsprequest` - RTSP request to send to the server (for example, "OPTIONS *").
* `scriptargs` - String of arguments for the script. The string is copied verbatim into the request.
* `scriptname` - Path and name of the script to execute. The script must be available on the Citrix ADC, in the /nsconfig/monitors/ directory.
* `secondarypassword` - Secondary password that users might have to provide to log on to the Access Gateway server. Applicable to CITRIX-AG monitors.
* `secure` - Use a secure SSL connection when monitoring a service. Applicable only to TCP based monitors. The secure option cannot be used with a CITRIX-AG monitor, because a CITRIX-AG monitor uses a secure connection by default.
* `secureargs` - List of arguments for the script which should be secure
* `send` - String to send to the service. Applicable to TCP-ECV, HTTP-ECV, and UDP-ECV monitors.
* `servicegroupname` - The name of the service group to which the monitor is to be bound.
* `servicename` - The name of the service to which the monitor is bound.
* `sipmethod` - SIP method to use for the query. Applicable only to monitors of type SIP-UDP.
* `sipreguri` - SIP user to be registered. Applicable only if the monitor is of type SIP-UDP and the SIP Method parameter is set to REGISTER.
* `sipuri` - SIP URI string to send to the service (for example, sip:sip.test). Applicable only to monitors of type SIP-UDP.
* `sitepath` - URL of the logon page. For monitors of type CITRIX-WEB-INTERFACE, to monitor a dynamic page under the site path, terminate the site path with a slash (/). Applicable to CITRIX-WEB-INTERFACE, CITRIX-WI-EXTENDED and CITRIX-XDM monitors.
* `snmpcommunity` - Community name for SNMP monitors.
* `snmpoid` - SNMP OID for SNMP monitors.
* `snmpthreshold` - Threshold for SNMP monitors.
* `snmpversion` - SNMP version to be used for SNMP monitors.
* `sqlquery` - SQL query for a MYSQL-ECV or MSSQL-ECV monitor. Sent to the database server after the server authenticates the connection.
* `sslprofile` - SSL Profile associated with the monitor
* `state` - State of the monitor. The DISABLED setting disables not only the monitor being configured, but all monitors of the same type, until the parameter is set to ENABLED.
* `storedb` - Store the database list populated with the responses to monitor probes. Used in database specific load balancing if MSSQL-ECV/MYSQL-ECV monitor is configured.
* `storefrontacctservice` - Enable/Disable probing for Account Service. Applicable only to Store Front monitors. For multi-tenancy configuration users my skip account service
* `storefrontcheckbackendservices` - This option will enable monitoring of services running on storefront server. Storefront services are monitored by probing to a Windows service that runs on the Storefront server and exposes details of which storefront services are running.
* `storename` - Store Name. For monitors of type STOREFRONT, STORENAME is an optional argument defining storefront service store name. Applicable to STOREFRONT monitors.
* `successretries` - Number of consecutive successful probes required to transition a service's state from DOWN to UP.
* `supportedvendorids` - List of Supported-Vendor-Id attribute value pairs (AVPs) for the Capabilities-Exchange-Request (CER) message to use for monitoring Diameter servers. A maximum eight of these AVPs are supported in a monitoring message.
* `tos` - Probe the service by encoding the destination IP address in the IP TOS (6) bits.
* `tosid` - The TOS ID of the specified destination IP. Applicable only when the TOS parameter is set.
* `transparent` - The monitor is bound to a transparent device such as a firewall or router. The state of a transparent device depends on the responsiveness of the services behind it.
* `trofscode` - Code expected when the server is under maintenance
* `trofsstring` - String expected from the server for the service to be marked as trofs. Applicable to HTTP-ECV/TCP-ECV monitors.
* `type` - Type of monitor that you want to create.
* `units1` - Unit of measurement for the Deviation parameter. Cannot be changed after the monitor is created.
* `units2` - Unit of measurement for the Down Time parameter. Cannot be changed after the monitor is created.
* `units3` - monitor interval units
* `units4` - monitor response timeout units
* `username` - User name with which to probe the RADIUS, NNTP, FTP, FTP-EXTENDED, MYSQL, MSSQL, POP3, CITRIX-AG, CITRIX-XD-DDC, CITRIX-WI-EXTENDED, CITRIX-XNC or CITRIX-XDM server.
* `validatecred` - Validate the credentials of the Xen Desktop DDC server user. Applicable to monitors of type CITRIX-XD-DDC.
* `vendorid` - Vendor-Id value for the Capabilities-Exchange-Request (CER) message to use for monitoring Diameter servers.
* `vendorspecificacctapplicationids` - List of Vendor-Specific-Acct-Application-Id attribute value pairs (AVPs) to use for monitoring Diameter servers. A maximum of eight of these AVPs are supported in a monitoring message.
* `vendorspecificauthapplicationids` - List of Vendor-Specific-Auth-Application-Id attribute value pairs (AVPs) for the Capabilities-Exchange-Request (CER) message to use for monitoring Diameter servers. A maximum of eight of these AVPs are supported in a monitoring message.
* `vendorspecificvendorid` - Vendor-Id to use in the Vendor-Specific-Application-Id grouped attribute-value pair (AVP) in the monitoring CER message.

## Attribute Reference

* `id` - The id of the lbmonitor. It has the same value as the `monitorname` attribute.
