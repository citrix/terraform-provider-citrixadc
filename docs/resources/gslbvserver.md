---
subcategory: "GSLB"
---

# Resource: glsbvserver

This resource is used to manage Global Service Load Balancing vserver.


## Example usage

```hcl
resource "citrixadc_gslbvserver" "tf_gslbvserver" {
  name        = "GSLB_East_Coast_Vserver"
  servicetype = "HTTP"
  state       = "ENABLED"
  edr         = "ENABLED"
  mir         = "DISABLED"
  lbmethod    = "ROUNDROBIN"
}
```


## Argument Reference

* `name` - (Required) Name for the GSLB virtual server.
* `servicetype` - (Required) Protocol used by services bound to the virtual server. Possible values: [ HTTP, FTP, TCP, UDP, SSL, SSL\_BRIDGE, SSL\_TCP, NNTP, ANY, SIP\_UDP, SIP\_TCP, SIP\_SSL, RADIUS, RDP, RTSP, MYSQL, MSSQL, ORACLE ]
* `iptype` - (Optional) The IP type for this GSLB vserver. Possible values: [ IPV4, IPV6 ]
* `dnsrecordtype` - (Optional) DNS record type to associate with the GSLB virtual server's domain name. Possible values: [ A, AAAA, CNAME, NAPTR ]
* `lbmethod` - (Optional) Load balancing method for the GSLB virtual server. Possible values: [ ROUNDROBIN, LEASTCONNECTION, LEASTRESPONSETIME, SOURCEIPHASH, LEASTBANDWIDTH, LEASTPACKETS, STATICPROXIMITY, RTT, CUSTOMLOAD, API ]
* `backupsessiontimeout` - (Optional) A non zero value enables the feature whose minimum value is 2 minutes. The feature can be disabled by setting the value to zero. The created session is in effect for a specific client per domain.
* `backuplbmethod` - (Optional) Backup load balancing method. Becomes operational if the primary load balancing method fails or cannot be used. Valid only if the primary method is based on either round-trip time (RTT) or static proximity. Possible values: [ ROUNDROBIN, LEASTCONNECTION, LEASTRESPONSETIME, SOURCEIPHASH, LEASTBANDWIDTH, LEASTPACKETS, STATICPROXIMITY, RTT, CUSTOMLOAD, API ]
* `netmask` - (Optional) IPv4 network mask for use in the SOURCEIPHASH load balancing method.
* `v6netmasklen` - (Optional) Number of bits to consider, in an IPv6 source IP address, for creating the hash that is required by the SOURCEIPHASH load balancing method.
* `tolerance` - (Optional) Site selection tolerance, in milliseconds, for implementing the RTT load balancing method. If a site's RTT deviates from the lowest RTT by more than the specified tolerance, the site is not considered when the Citrix ADC makes a GSLB decision. The appliance implements the round robin method of global server load balancing between sites whose RTT values are within the specified tolerance. If the tolerance is 0 (zero), the appliance always sends clients the IP address of the site with the lowest RTT.
* `persistencetype` - (Optional) Use source IP address based persistence for the virtual server. After the load balancing method selects a service for the first packet, the IP address received in response to the DNS query is used for subsequent requests from the same client. Possible values: [ SOURCEIP, NONE  ]
* `persistenceid` - (Optional) The persistence ID for the GSLB virtual server. The ID is a positive integer that enables GSLB sites to identify the GSLB virtual server, and is required if source IP address based or spill over based persistence is enabled on the virtual server.
* `persistmask` - (Optional) The optional IPv4 network mask applied to IPv4 addresses to establish source IP address based persistence.
* `v6persistmasklen` - (Optional) Number of bits to consider in an IPv6 source IP address when creating source IP address based persistence sessions.
* `timeout` - (Optional) Idle time, in minutes, after which a persistence entry is cleared.
* `edr` - (Optional) Send clients an empty DNS response when the GSLB virtual server is DOWN. Possible values: [ ENABLED, DISABLED ]
* `ecs` - (Optional) If enabled, respond with EDNS Client Subnet (ECS) option in the response for a DNS query with ECS. The ECS address will be used for persistence and spillover persistence (if enabled) instead of the LDNS address. Persistence mask is ignored if ECS is enabled. Possible values: [ ENABLED, DISABLED ]
* `ecsaddrvalidation` - (Optional) Validate if ECS address is a private or unroutable address and in such cases, use the LDNS IP. Possible values: [ ENABLED, DISABLED ]
* `mir` - (Optional) Include multiple IP addresses in the DNS responses sent to clients. Possible values: [ ENABLED, DISABLED ]
* `disableprimaryondown` - (Optional) Continue to direct traffic to the backup chain even after the primary GSLB virtual server returns to the UP state. Used when spillover is configured for the virtual server. Possible values: [ ENABLED, DISABLED ]
* `dynamicweight` - (Optional) Specify if the appliance should consider the service count, service weights, or ignore both when using weight-based load balancing methods. The state of the number of services bound to the virtual server help the appliance to select the service. Possible values : [ SERVICECOUNT, SERVICEWEIGHT, DISABLED ]
* `state` - (Optional) State of the GSLB virtual server. Possible values: [ ENABLED, DISABLED ]
* `considereffectivestate` - (Optional) If the primary state of all bound GSLB services is DOWN, consider the effective states of all the GSLB services, obtained through the Metrics Exchange Protocol (MEP), when determining the state of the GSLB virtual server. To consider the effective state, set the parameter to STATE\_ONLY. To disregard the effective state, set the parameter to NONE.

    The effective state of a GSLB service is the ability of the corresponding virtual server to serve traffic. The effective state of the load balancing virtual server, which is transferred to the GSLB service, is UP even if only one virtual server in the backup chain of virtual servers is in the UP state.

    Possible values: [ NONE, STATE_ONLY ]

* `comment` - (Optional) Any comments that you might want to associate with the GSLB virtual server.
* `somethod` - (Optional) Type of threshold that, when exceeded, triggers spillover. Possible values: [ CONNECTION, DYNAMICCONNECTION, BANDWIDTH, HEALTH, NONE ]
* `sopersistence` - (Optional) If spillover occurs, maintain source IP address based persistence for both primary and backup GSLB virtual servers.Possible values: [ ENABLED, DISABLED ]
* `sopersistencetimeout` - (Optional) Timeout for spillover persistence, in minutes.
* `sothreshold` - (Optional) Threshold at which spillover occurs. Specify an integer for the CONNECTION spillover method, a bandwidth value in kilobits per second for the BANDWIDTH method (do not enter the units), or a percentage for the HEALTH method (do not enter the percentage symbol).
* `sobackupaction` - (Optional) Action to be performed if spillover is to take effect, but no backup chain to spillover is usable or exists. Possible values: [ DROP, ACCEPT, REDIRECT ]
* `appflowlog` - (Optional) Enable logging appflow flow information. Possible values: [ ENABLED, DISABLED ]
* `backupvserver` - (Optional) Name of the backup GSLB virtual server to which the appliance should to forward requests if the status of the primary GSLB virtual server is down or exceeds its spillover threshold.
* `servicename` - (Optional) Name of the GSLB service for which to change the weight.
* `weight` - (Optional) Weight to assign to the GSLB service.
* `domainname` - (Optional) Domain name for which to change the time to live (TTL) and/or backup service IP address.
* `ttl` - (Optional) Time to live (TTL) for the domain.
* `backupip` - (Optional) The IP address of the backup service for the specified domain name. Used when all the services bound to the domain are down, or when the backup chain of virtual servers is down.
* `cookiedomain` - (Optional) The cookie domain for the GSLB site. Used when inserting the GSLB site cookie in the HTTP response.
* `cookietimeout` - (Optional) Timeout, in minutes, for the GSLB site cookie.
* `sitedomainttl` - (Optional) TTL, in seconds, for all internally created site domains (created when a site prefix is configured on a GSLB service) that are associated with this virtual server.
* `orderthreshold` - (Optional) This option is used to to specify the threshold of minimum number of services to be UP in an order, for it to be considered in Lb decision.
* `rule` - (Optional) Expression, or name of a named expression, against which traffic is evaluated. This field is applicable only if gslb method or gslb backup method are set to API. The following requirements apply only to the Citrix ADC CLI: * If the expression includes one or more spaces, enclose the entire expression in double quotation marks. * If the expression itself includes double quotation marks, escape the quotations by using the \ character. * Alternatively, you can use single quotation marks to enclose the rule, in which case you do not have to escape the double quotation marks.
* `toggleorder` - (Optional) Configure this option to toggle order preference
* `domain` - (Optional) A set of domain binding blocks. Documented below. (deprecates soon)
* `service` - (Optional) A set of GSLB service biding blocks. Documented below. (deprecates soon)

!>
[**DEPRECATED**] Please use [`gslbvserver_domain_binding`](https://registry.terraform.io/providers/citrix/citrixadc/latest/docs/resources/gslbvserver_domain_binding) to bind `domain` to `gslbvserver` instead of this resource. Support for binding `domain` to `gslbvserver` in this resource will be deprecated soon.

!>
[**DEPRECATED**] Please use [`gslbvserver_gslbservice_binding`](https://registry.terraform.io/providers/citrix/citrixadc/latest/docs/resources/gslbvserver_gslbservice_binding) to bind `gslbservice` to `gslbvserver` instead of this resource. Support for binding `gslbservice` to `gslbvserver` in this resource will be deprecated soon.

A domain binding supports the following:

* `backupipflag` - (Optional) The IP address of the backup service for the specified domain name. Used when all the services bound to the domain are down, or when the backup chain of virtual servers is down.
* `cookietimeout` - (Optional) Timeout, in minutes, for the GSLB site cookie.
* `backupip` - (Optional) The IP address of the backup service for the specified domain name. Used when all the services bound to the domain are down, or when the backup chain of virtual servers is down.
* `ttl` - (Optional) Time to live (TTL) for the domain.
* `domainname` - (Optional) Domain name for which to change the time to live (TTL) and/or backup service IP address.
* `sitedomainttl` - (Optional) TTL, in seconds, for all internally created site domains (created when a site prefix is configured on a GSLB service) that are associated with this virtual server.
* `cookiedomain` - (Optional) The cookie domain for the GSLB site. Used when inserting the GSLB site cookie in the HTTP response.
* `cookiedomainflag` - (Optional) The cookie domain for the GSLB site. Used when inserting the GSLB site cookie in the HTTP response.
* `name` - (Optional) Name of the virtual server on which to perform the binding operation.

A GSLB service binding supports the following:

* `weight` - (Optional) Weight to assign to the GSLB service.
* `servicename` - (Optional) Name of the GSLB service for which to change the weight.
* `domainname` - (Optional) Domain name for which to change the time to live (TTL) and/or backup service IP address.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the gslbvserver. It has the same value as the `name` attribute.


## Import

An instance of the resource can be imported using its name, e.g.

```shell
terraform import citrixadc_gslbvserver.tf_gslbvserver GSLB_East_Coast_Vserver
```
