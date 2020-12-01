---
subcategory: "Content Switching"
---

# Resource: csvserver

This resource is used to manage Content switching virtual servers.


## Example usage

```hcl
resource "citrixadc_csvserver" "tf_csvserver" {
  ipv46       = "10.10.10.22"
  name        = "tf_csvserver"
  port        = 80
  servicetype = "HTTP"

}
```


## Argument Reference

* `name` - (Optional) Name for the content switching virtual server.
* `td` - (Optional) Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.
* `servicetype` - (Optional) Protocol used by the virtual server. Possible values: [ HTTP, SSL, TCP, FTP, RTSP, SSL\_TCP, UDP, DNS, SIP\_UDP, SIP\_TCP, SIP\_SSL, ANY, RADIUS, RDP, MYSQL, MSSQL, DIAMETER, SSL\_DIAMETER, DNS\_TCP, ORACLE, SMPP, PROXY, MONGO, MONGO\_TLS ]
* `ipv46` - (Optional) IP address of the content switching virtual server.
* `targettype` - (Optional) Virtual server target type. Possible values: [ GSLB ]
* `dnsrecordtype` - (Optional) Possible values: [ A, AAAA, CNAME, NAPTR ]
* `persistenceid` - (Optional)
* `ippattern` - (Optional) IP address pattern, in dotted decimal notation, for identifying packets to be accepted by the virtual server. The IP Mask parameter specifies which part of the destination IP address is matched against the pattern. Mutually exclusive with the IP Address parameter.

    For example, if the IP pattern assigned to the virtual server is 198.51.100.0 and the IP mask is 255.255.240.0 (a forward mask), the first 20 bits in the destination IP addresses are matched with the first 20 bits in the pattern. The virtual server accepts requests with IP addresses that range from 198.51.96.1 to 198.51.111.254. You can also use a pattern such as 0.0.2.2 and a mask such as 0.0.255.255 (a reverse mask).

    If a destination IP address matches more than one IP pattern, the pattern with the longest match is selected, and the associated virtual server processes the request. For example, if the virtual servers, vs1 and vs2, have the same IP pattern, 0.0.100.128, but different IP masks of 0.0.255.255 and 0.0.224.255, a destination IP address of 198.51.100.128 has the longest match with the IP pattern of vs1. If a destination IP address matches two or more virtual servers to the same extent, the request is processed by the virtual server whose port number matches the port number in the request.
* `ipmask` - (Optional) IP mask, in dotted decimal notation, for the IP Pattern parameter. Can have leading or trailing non-zero octets (for example, 255.255.240.0 or 0.0.255.255). Accordingly, the mask specifies whether the first n bits or the last n bits of the destination IP address in a client request are to be matched with the corresponding bits in the IP pattern. The former is called a forward mask. The latter is called a reverse mask.
* `range` - (Optional) Number of consecutive IP addresses, starting with the address specified by the IP Address parameter, to include in a range of addresses assigned to this virtual server.
* `port` - (Optional) Port number for content switching virtual server.
* `ipset` - (Optional) The list of IPv4/IPv6 addresses bound to ipset would form a part of listening service on the current cs vserver.
* `state` - (Optional) Initial state of the load balancing virtual server. Possible values: [ ENABLED, DISABLED ]
* `stateupdate` - (Optional) Enable state updates for a specific content switching virtual server. By default, the Content Switching virtual server is always UP, regardless of the state of the Load Balancing virtual servers bound to it. This parameter interacts with the global setting as follows:

    Global Level + Vserver Level = Result

    ENABLED + ENABLED = ENABLED

    ENABLED + DISABLED = ENABLED

    DISABLED + ENABLED = ENABLED

    DISABLED + DISABLED = DISABLED

    If you want to enable state updates for only some content switching virtual servers, be sure to disable the state update parameter.
    Possible values: [ ENABLED, DISABLED, UPDATEONBACKENDUPDATE ]
* `cacheable` - (Optional) Use this option to specify whether a virtual server, used for load balancing or content switching, routes requests to the cache redirection virtual server before sending it to the configured servers. Possible values: [ YES, NO ]
* `redirecturl` - (Optional) URL to which traffic is redirected if the virtual server becomes unavailable. The service type of the virtual server should be either HTTP or SSL.

    ~> Make sure that the domain in the URL does not match the domain specified for a content switching policy. If it does, requests are continuously redirected to the unavailable virtual server.
* `clttimeout` - (Optional) Idle time, in seconds, after which the client connection is terminated.
* `precedence` - (Optional) Type of precedence to use for both RULE-based and URL-based policies on the content switching virtual server. With the default (RULE) setting, incoming requests are evaluated against the rule-based content switching policies. If none of the rules match, the URL in the request is evaluated against the URL-based content switching policies. Possible values: [ RULE, URL ]
* `casesensitive` - (Optional) Consider case in URLs (for policies that use URLs instead of RULES). For example, with the ON setting, the URLs /a/1.html and /A/1.HTML are treated differently and can have different targets (set by content switching policies). With the OFF setting, /a/1.html and /A/1.HTML are switched to the same target. Possible values: [ ON, OFF ]
* `somethod` - (Optional) Type of spillover used to divert traffic to the backup virtual server when the primary virtual server reaches the spillover threshold. Connection spillover is based on the number of connections. Bandwidth spillover is based on the total Kbps of incoming and outgoing traffic. Possible values: [ CONNECTION, DYNAMICCONNECTION, BANDWIDTH, HEALTH, NONE ]
* `sopersistence` - (Optional) Maintain source-IP based persistence on primary and backup virtual servers. Possible values: [ ENABLED, DISABLED ]
* `sopersistencetimeout` - (Optional) Time-out value, in minutes, for spillover persistence.
* `sothreshold` - (Optional) Depending on the spillover method, the maximum number of connections or the maximum total bandwidth (Kbps) that a virtual server can handle before spillover occurs.
* `sobackupaction` - (Optional) Action to be performed if spillover is to take effect, but no backup chain to spillover is usable or exists. Possible values: [ DROP, ACCEPT, REDIRECT ]
* `redirectportrewrite` - (Optional) State of port rewrite while performing HTTP redirect.Possible values: [ ENABLED, DISABLED ]
* `downstateflush` - (Optional) Flush all active transactions associated with a virtual server whose state transitions from UP to DOWN. Do not enable this option for applications that must complete their transactions. Possible values: [ ENABLED, DISABLED ]
* `backupvserver` - (Optional) Name of the backup virtual server that you are configuring.
* `disableprimaryondown` - (Optional) Continue forwarding the traffic to backup virtual server even after the primary server comes UP from the DOWN state. Possible values: [ ENABLED, DISABLED ]
* `insertvserveripport` - (Optional) Insert the virtual server's VIP address and port number in the request header. Available values function as follows:

    VIPADDR - Header contains the vserver's IP address and port number without any translation.

    OFF - The virtual IP and port header insertion option is disabled.

    V6TOV4MAPPING - Header contains the mapped IPv4 address corresponding to the IPv6 address of the vserver and the port number. An IPv6 address can be mapped to a user-specified IPv4 address using the set ns ip6 command.

    Possible values: [ OFF, VIPADDR, V6TOV4MAPPING ]
* `vipheader` - (Optional) Name of virtual server IP and port header, for use with the VServer IP Port Insertion parameter.
* `rtspnat` - (Optional) Enable network address translation (NAT) for real-time streaming protocol (RTSP) connections. Possible values: [ ON, OFF ]
* `authenticationhost` - (Optional) FQDN of the authentication virtual server. The service type of the virtual server should be either HTTP or SSL.
* `authentication`  - (Optional) Authenticate users who request a connection to the content switching virtual server. Possible values: [ ON, OFF ]
* `listenpolicy` - (Optional) String specifying the listen policy for the content switching virtual server. Can be either the name of an existing expression or an in-line expression.
* `listenpriority` - (Optional)     Integer specifying the priority of the listen policy. A higher number specifies a lower priority. If a request matches the listen policies of more than one virtual server the virtual server whose listen policy has the highest priority (the lowest priority number) accepts the request.
* `authn401` - (Optional) Enable HTTP 401-response based authentication. Possible values: [ ON, OFF ]
* `authnvsname` - (Optional) Name of authentication virtual server that authenticates the incoming user requests to this content switching virtual server.
* `push` - (Optional) Process traffic with the push virtual server that is bound to this content switching virtual server (specified by the Push VServer parameter). The service type of the push virtual server should be either HTTP or SSL. Possible values: [ ENABLED, DISABLED ]
* `pushvserver` - (Optional) Name of the load balancing virtual server, of type PUSH or SSL_PUSH, to which the server pushes updates received on the client-facing load balancing virtual server.
* `pushlabel` - (Optional) Expression for extracting the label from the response received from server. This string can be either an existing rule name or an inline expression. The service type of the virtual server should be either HTTP or SSL.
* `pushmulticlients` - (Optional) Allow multiple Web 2.0 connections from the same client to connect to the virtual server and expect updates. Possible values: [ YES, NO ]
* `httpprofilename` - (Optional) Name of the HTTP profile containing HTTP configuration settings for the virtual server. The service type of the virtual server should be either HTTP or SSL.
* `dbprofilename` - (Optional) Name of the DB profile.
* `oracleserverversion` - (Optional) Oracle server version. Possible values: [ 10G, 11G ]
* `comment` - (Optional) Information about this virtual server.
* `mssqlserverversion` - (Optional) The version of the MSSQL server. Possible values: [ 70, 2000, 2000SP1, 2005, 2008, 2008R2, 2012, 2014 ]
* `l2conn` - (Optional) Use L2 Parameters to identify a connection. Possible values: [ ON, OFF ]
* `mysqlprotocolversion` - (Optional) The protocol version returned by the mysql vserver.
* `mysqlserverversion` - (Optional) The server version string returned by the mysql vserver.
* `mysqlcharacterset` - (Optional) The character set returned by the mysql vserver.
* `mysqlservercapabilities` - (Optional) The server capabilities returned by the mysql vserver.
* `appflowlog` - (Optional) Enable logging appflow flow information. Possible values: [ ENABLED, DISABLED ]
* `netprofile` - (Optional) The name of the network profile.
* `icmpvsrresponse` - (Optional) Can be active or passive. Possible values: [ PASSIVE, ACTIVE ]
* `rhistate` - (Optional) A host route is injected according to the setting on the virtual servers

    If set to PASSIVE on all the virtual servers that share the IP address, the appliance always injects the hostroute.

    If set to ACTIVE on all the virtual servers that share the IP address, the appliance injects even if one virtual server is UP.

    If set to ACTIVE on some virtual servers and PASSIVE on the others, the appliance, injects even if one virtual server set to ACTIVE is UP.

    Default value: PASSIVE

    Possible values: [ PASSIVE, ACTIVE ]
* `authnprofile` - (Optional) Name of the authentication profile to be used when authentication is turned on.
* `dnsprofilename` - (Optional) Name of the DNS profile to be associated with the VServer. DNS profile properties will applied to the transactions processed by a VServer. This parameter is valid only for DNS and DNS-TCP VServers.
* `domainname` - (Optional) Domain name for which to change the time to live (TTL) and/or backup service IP address.
* `ttl` - (Optional)
* `backupip` - (Optional)
* `cookiedomain` - (Optional)
* `cookietimeout` - (Optional)
* `sitedomainttl` - (Optional)


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the csvserver. It has the same value as the `name` attribute.


## Import

An instance of the resource can be imported using its name, e.g.

```shell
terraform import citrixadc_csvserver.tf_csvserver tf_csvserver
```
