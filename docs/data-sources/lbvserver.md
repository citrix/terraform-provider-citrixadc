---
subcategory: "Load Balancing"
---

# Data Source: citrixadc_lbvserver

The lbvserver data source is used to retrieve information about an existing load balancing virtual server.

## Example usage

```terraform
data "citrixadc_lbvserver" "tf_lbvserver" {
  name = "terraform-lb"
}

output "servicetype" {
  value = data.citrixadc_lbvserver.tf_lbvserver.servicetype
}

output "ipv46" {
  value = data.citrixadc_lbvserver.tf_lbvserver.ipv46
}

output "lbmethod" {
  value = data.citrixadc_lbvserver.tf_lbvserver.lbmethod
}
```

## Argument Reference

* `name` - (Required) Name of the load balancing virtual server.

## Attribute Reference

In addition to the argument above, the following attributes are exported:

* `id` - The ID of the lbvserver resource.
* `servicetype` - Protocol used by the service.
* `ipv46` - IPv4 or IPv6 address assigned to the virtual server.
* `port` - Port number for the virtual server.
* `lbmethod` - Load balancing method used by the virtual server.
* `persistencetype` - Type of persistence for the virtual server.
* `timeout` - Time period for which a persistence session is in effect, in minutes.
* `persistencebackup` - Backup persistence type for the virtual server.
* `backuppersistencetimeout` - Time period for which backup persistence is in effect.
* `persistmask` - Persistence mask for IP based persistence types, for IPv4 virtual servers.
* `v6persistmasklen` - Persistence mask for IP based persistence types, for IPv6 virtual servers.
* `cookiename` - Cookie name for COOKIE persistence type.
* `persistavpno` - Persist AVP number for Diameter Persistency.
* `state` - State of the load balancing virtual server (ENABLED or DISABLED).
* `backupvserver` - Name of the backup virtual server.
* `backuplbmethod` - Backup load balancing method.
* `netmask` - IPv4 subnet mask to apply when the load balancing method is DESTINATIONIPHASH or SOURCEIPHASH.
* `v6netmasklen` - Number of bits to consider in an IPv6 address for DESTINATIONIPHASH or SOURCEIPHASH load balancing methods.
* `hashlength` - Number of bytes to consider for the hash value used in URLHASH and DOMAINHASH load balancing methods.
* `ippattern` - IP address pattern for identifying packets to be accepted by the virtual server.
* `ipmask` - IP mask for the IP Pattern parameter.
* `ipset` - The list of IPv4/IPv6 addresses bound to ipset.
* `range` - Number of IP addresses that the appliance must generate and assign to the virtual server.
* `clttimeout` - Idle time, in seconds, after which a client connection is terminated.
* `somethod` - Type of spillover threshold.
* `sopersistence` - If spillover occurs, maintain source IP address based persistence.
* `sopersistencetimeout` - Timeout for spillover persistence, in minutes.
* `sothreshold` - Threshold at which spillover occurs.
* `sobackupaction` - Action to be performed if spillover is to take effect but no backup chain is usable.
* `healththreshold` - Threshold in percent of active services below which vserver state is made down.
* `rule` - Expression against which traffic is evaluated.
* `listenpolicy` - Expression identifying traffic accepted by the virtual server.
* `listenpriority` - Priority of the listen policy.
* `resrule` - Expression for creating rule based persistence sessions.
* `m` - Redirection mode for load balancing (IP, MAC, IPTUNNEL, or TOS).
* `tosid` - TOS ID of the virtual server.
* `datalength` - Length of the token to be extracted for token method of load balancing.
* `dataoffset` - Offset to be considered when extracting a token from the TCP payload.
* `sessionless` - Perform load balancing on a per-packet basis.
* `trofspersistence` - Honor Trofs persistence.
* `connfailover` - Connection failover mode for the virtual server.
* `redirurl` - URL to redirect traffic if the virtual server becomes unavailable.
* `redirurlflags` - The redirect URL to be unset.
* `cacheable` - Route cacheable requests to a cache redirection virtual server.
* `rtspnat` - Use network address translation (NAT) for RTSP data connections.
* `authn401` - Enable user authentication with HTTP 401 responses.
* `authentication` - Enable or disable user authentication.
* `authenticationhost` - FQDN of the authentication virtual server.
* `authnvsname` - Name of an authentication virtual server for authenticating users.
* `authnprofile` - Name of the authentication profile.
* `push` - Process traffic with the push virtual server.
* `pushvserver` - Name of the load balancing virtual server to receive server push updates.
* `pushlabel` - Expression for extracting a label from the server's response.
* `pushmulticlients` - Allow multiple Web 2.0 connections from the same client.
* `tcpprofilename` - Name of the TCP profile.
* `httpprofilename` - Name of the HTTP profile.
* `dbprofilename` - Name of the DB profile.
* `lbprofilename` - Name of the LB profile.
* `dnsprofilename` - Name of the DNS profile.
* `netprofile` - Name of the network profile.
* `apiprofile` - The API profile name.
* `adfsproxyprofile` - Name of the adfsProxy profile for ADFSPIP protocol.
* `comment` - Comments associated with the virtual server.
* `mssqlserverversion` - Microsoft SQL Server version.
* `mysqlprotocolversion` - MySQL protocol version.
* `mysqlserverversion` - MySQL server version string.
* `mysqlcharacterset` - Character set that the virtual server advertises to clients.
* `mysqlservercapabilities` - Server capabilities that the virtual server advertises to clients.
* `oracleserverversion` - Oracle server version.
* `appflowlog` - Apply AppFlow logging to the virtual server.
* `icmpvsrresponse` - How the Citrix ADC responds to ping requests.
* `disableprimaryondown` - If the primary virtual server goes down, do not allow it to return to primary status until manually enabled.
* `insertvserveripport` - Insert an HTTP header with the virtual server's IP address and port.
* `vipheader` - Name for the inserted header.
* `downstateflush` - Flush all active transactions when state transitions from UP to DOWN.
* `dns64` - Enable or disable DNS64 on lbvserver.
* `dnsoverhttps` - Enable or disable DNS over HTTPS (DoH) processing.
* `bypassaaaa` - Do not send AAAA queries to back end dns server while resolving DNS64 query.
* `recursionavailable` - Set the RA bit in DNS replies.
* `processlocal` - By turning on this option packets destined to a vserver in a cluster will not undergo steering.
* `redirectportrewrite` - Rewrite the port and change the protocol to ensure successful HTTP redirects.
* `redirectfromport` - Port number from which we absorb the traffic for http redirect.
* `httpsredirecturl` - URL to which all HTTP traffic is redirected.
* `retainconnectionsoncluster` - Retain existing connections on a node joining a Cluster system.
* `dbslb` - Enable database specific load balancing for MySQL and MSSQL service types.
* `l2conn` - Use Layer 2 parameters in addition to the 4-tuple to identify a connection.
* `macmoderetainvlan` - Retain vlan information of incoming packet when macmode is enabled.
* `newservicerequest` - Number of requests or percentage of load to increase on a new service.
* `newservicerequestunit` - Units in which to increment load in slow-start mode.
* `newservicerequestincrementinterval` - Interval between successive increments in slow-start mode.
* `minautoscalemembers` - Minimum number of members expected when vserver is used in Autoscale.
* `maxautoscalemembers` - Maximum number of members expected when vserver is used in Autoscale.
* `skippersistency` - Behavior when selected service has reached threshold.
* `td` - Traffic domain ID.
* `order` - Order number to be assigned to the service.
* `orderthreshold` - Threshold of minimum number of services to be UP in an order.
* `toggleorder` - Configure this option to toggle order preference.
* `rhistate` - Route Health Injection (RHI) functionality.
* `probeport` - Port for external health check of the vserver status.
* `probeprotocol` - Protocol for external health check (HTTP or TCP).
* `probesuccessresponsecode` - HTTP code to return in SUCCESS case.
* `tcpprobeport` - Port number for external TCP probe.
* `quicprofilename` - Name of QUIC profile.
* `quicbridgeprofilename` - Name of the QUIC Bridge profile.
* `weight` - Weight to assign to the specified service.
* `servicename` - Service to bind to the virtual server.
* `newname` - New name for the virtual server.
