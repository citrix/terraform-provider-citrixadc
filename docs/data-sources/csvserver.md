---
subcategory: "CS"
---

# Data Source `csvserver`

The csvserver data source allows you to retrieve information about Content Switching virtual servers.


## Example usage

```terraform
data "citrixadc_csvserver" "tf_csvserver" {
  name = "terraform-cs"
}

output "ipv46" {
  value = data.citrixadc_csvserver.tf_csvserver.ipv46
}

output "servicetype" {
  value = data.citrixadc_csvserver.tf_csvserver.servicetype
}

output "port" {
  value = data.citrixadc_csvserver.tf_csvserver.port
}
```


## Argument Reference

* `name` - (Required) Name for the content switching virtual server. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the csvserver. It has the same value as the `name` attribute.
* `apiprofile` - The API profile where one or more API specs are bounded to.
* `appflowlog` - Enable logging appflow flow information.
* `authentication` - Authenticate users who request a connection to the content switching virtual server.
* `authenticationhost` - FQDN of the authentication virtual server. The service type of the virtual server should be either HTTP or SSL.
* `authn401` - Enable HTTP 401-response based authentication.
* `authnprofile` - Name of the authentication profile to be used when authentication is turned on.
* `authnvsname` - Name of authentication virtual server that authenticates the incoming user requests to this content switching virtual server.
* `backupip` - Backup IP address.
* `backuppersistencetimeout` - Time period for which backup persistence is in effect.
* `backupvserver` - Name of the backup virtual server that you are configuring.
* `cacheable` - Use this option to specify whether a virtual server, used for load balancing or content switching, routes requests to the cache redirection virtual server before sending it to the configured servers.
* `casesensitive` - Consider case in URLs (for policies that use URLs instead of RULES).
* `clttimeout` - Idle time, in seconds, after which the client connection is terminated.
* `comment` - Information about this virtual server.
* `cookiedomain` - Cookie domain.
* `cookiename` - Use this parameter to specify the cookie name for COOKIE peristence type.
* `cookietimeout` - Cookie timeout.
* `dbprofilename` - Name of the DB profile.
* `disableprimaryondown` - Continue forwarding the traffic to backup virtual server even after the primary server comes UP from the DOWN state.
* `dnsoverhttps` - This option is used to enable/disable DNS over HTTPS (DoH) processing.
* `dnsprofilename` - Name of the DNS profile to be associated with the VServer.
* `dnsrecordtype` - DNS record type.
* `domainname` - Domain name for which to change the time to live (TTL) and/or backup service IP address.
* `downstateflush` - Flush all active transactions associated with a virtual server whose state transitions from UP to DOWN.
* `dtls` - This option starts/stops the dtls service on the vserver.
* `httpprofilename` - Name of the HTTP profile containing HTTP configuration settings for the virtual server.
* `httpsredirecturl` - URL to which all HTTP traffic received on the port specified in the -redirectFromPort parameter is redirected.
* `icmpvsrresponse` - Can be active or passive.
* `insertvserveripport` - Insert the virtual server's VIP address and port number in the request header.
* `ipmask` - IP mask, in dotted decimal notation, for the IP Pattern parameter.
* `ippattern` - IP address pattern, in dotted decimal notation, for identifying packets to be accepted by the virtual server.
* `ipset` - The list of IPv4/IPv6 addresses bound to ipset would form a part of listening service on the current cs vserver.
* `ipv46` - IP address of the content switching virtual server.
* `l2conn` - Use L2 Parameters to identify a connection.
* `listenpolicy` - String specifying the listen policy for the content switching virtual server.
* `listenpriority` - Integer specifying the priority of the listen policy.
* `mssqlserverversion` - The version of the MSSQL server.
* `mysqlcharacterset` - The character set returned by the mysql vserver.
* `mysqlprotocolversion` - The protocol version returned by the mysql vserver.
* `mysqlservercapabilities` - The server capabilities returned by the mysql vserver.
* `mysqlserverversion` - The server version string returned by the mysql vserver.
* `netprofile` - The name of the network profile.
* `newname` - New name for the virtual server.
* `oracleserverversion` - Oracle server version.
* `persistencebackup` - Backup persistence type for the virtual server.
* `persistenceid` - Persistence ID.
* `persistencetype` - Type of persistence for the virtual server.
* `persistmask` - Persistence mask for IP based persistence types, for IPv4 virtual servers.
* `port` - Port number for content switching virtual server.
* `precedence` - Type of precedence to use for both RULE-based and URL-based policies on the content switching virtual server.
* `probeport` - Citrix ADC provides support for external health check of the vserver status. Select port for HTTP/TCP monitoring.
* `probeprotocol` - Citrix ADC provides support for external health check of the vserver status. Select HTTP or TCP probes for healthcheck.
* `probesuccessresponsecode` - HTTP code to return in SUCCESS case.
* `push` - Process traffic with the push virtual server that is bound to this content switching virtual server.
* `pushlabel` - Expression for extracting the label from the response received from server.
* `pushmulticlients` - Allow multiple Web 2.0 connections from the same client to connect to the virtual server and expect updates.
* `pushvserver` - Name of the load balancing virtual server, of type PUSH or SSL_PUSH, to which the server pushes updates received on the client-facing load balancing virtual server.
* `quicprofilename` - Name of QUIC profile which will be attached to the Content Switching VServer.
* `range` - Number of consecutive IP addresses, starting with the address specified by the IP Address parameter, to include in a range of addresses assigned to this virtual server.
* `redirectfromport` - Port number for the virtual server, from which we absorb the traffic for http redirect.
* `redirectportrewrite` - State of port rewrite while performing HTTP redirect.
* `redirecturl` - URL to which traffic is redirected if the virtual server becomes unavailable.
* `rhistate` - A host route is injected according to the setting on the virtual servers.
* `rtspnat` - Enable network address translation (NAT) for real-time streaming protocol (RTSP) connections.
* `servicetype` - Protocol used by the virtual server.
* `sitedomainttl` - Site domain TTL.
* `sobackupaction` - Action to be performed if spillover is to take effect, but no backup chain to spillover is usable or exists.
* `somethod` - Type of spillover used to divert traffic to the backup virtual server when the primary virtual server reaches the spillover threshold.
* `sopersistence` - Maintain source-IP based persistence on primary and backup virtual servers.
* `sopersistencetimeout` - Time-out value, in minutes, for spillover persistence.
* `sothreshold` - Depending on the spillover method, the maximum number of connections or the maximum total bandwidth (Kbps) that a virtual server can handle before spillover occurs.
* `state` - Initial state of the load balancing virtual server.
* `stateupdate` - Enable state updates for a specific content switching virtual server.
* `targettype` - Virtual server target type.
* `tcpprobeport` - Port number for external TCP probe.
* `tcpprofilename` - Name of the TCP profile containing TCP configuration settings for the virtual server.
* `td` - Integer value that uniquely identifies the traffic domain in which you want to configure the entity.
* `timeout` - Time period for which a persistence session is in effect.
* `ttl` - TTL value.
* `v6persistmasklen` - Persistence mask for IP based persistence types, for IPv6 virtual servers.
* `vipheader` - Name of virtual server IP and port header, for use with the VServer IP Port Insertion parameter.
