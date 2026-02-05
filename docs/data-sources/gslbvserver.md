---
subcategory: "GSLB"
---

# Data Source `gslbvserver`

The gslbvserver data source allows you to retrieve information about GSLB virtual servers.


## Example usage

```terraform
data "citrixadc_gslbvserver" "tf_gslbvserver" {
  name = "GSLB-East-Coast-Vserver"
}

output "servicetype" {
  value = data.citrixadc_gslbvserver.tf_gslbvserver.servicetype
}

output "dnsrecordtype" {
  value = data.citrixadc_gslbvserver.tf_gslbvserver.dnsrecordtype
}

output "lbmethod" {
  value = data.citrixadc_gslbvserver.tf_gslbvserver.lbmethod
}
```


## Argument Reference

* `name` - (Required) Name for the GSLB virtual server. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Can be changed after the virtual server is created.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the gslbvserver. It has the same value as the `name` attribute.
* `appflowlog` - Enable logging appflow flow information.
* `backupip` - The IP address of the backup service for the specified domain name. Used when all the services bound to the domain are down, or when the backup chain of virtual servers is down.
* `backuplbmethod` - Backup load balancing method. Becomes operational if the primary load balancing method fails or cannot be used. Valid only if the primary method is based on either round-trip time (RTT) or static proximity.
* `backupsessiontimeout` - A non zero value enables the feature whose minimum value is 2 minutes. The feature can be disabled by setting the value to zero. The created session is in effect for a specific client per domain.
* `backupvserver` - Name of the backup GSLB virtual server to which the appliance should to forward requests if the status of the primary GSLB virtual server is down or exceeds its spillover threshold.
* `comment` - Any comments that you might want to associate with the GSLB virtual server.
* `considereffectivestate` - If the primary state of all bound GSLB services is DOWN, consider the effective states of all the GSLB services, obtained through the Metrics Exchange Protocol (MEP), when determining the state of the GSLB virtual server. To consider the effective state, set the parameter to STATE_ONLY. To disregard the effective state, set the parameter to NONE.
* `cookie_domain` - The cookie domain for the GSLB site. Used when inserting the GSLB site cookie in the HTTP response.
* `cookietimeout` - Timeout, in minutes, for the GSLB site cookie.
* `disableprimaryondown` - Continue to direct traffic to the backup chain even after the primary GSLB virtual server returns to the UP state. Used when spillover is configured for the virtual server.
* `dnsrecordtype` - DNS record type to associate with the GSLB virtual server's domain name.
* `domainname` - Domain name for which to change the time to live (TTL) and/or backup service IP address.
* `dynamicweight` - Specify if the appliance should consider the service count, service weights, or ignore both when using weight-based load balancing methods. The state of the number of services bound to the virtual server help the appliance to select the service.
* `ecs` - If enabled, respond with EDNS Client Subnet (ECS) option in the response for a DNS query with ECS. The ECS address will be used for persistence and spillover persistence (if enabled) instead of the LDNS address. Persistence mask is ignored if ECS is enabled.
* `ecsaddrvalidation` - Validate if ECS address is a private or unroutable address and in such cases, use the LDNS IP.
* `edr` - Send clients an empty DNS response when the GSLB virtual server is DOWN.
* `iptype` - The IP type for this GSLB vserver.
* `lbmethod` - Load balancing method for the GSLB virtual server.
* `mir` - Include multiple IP addresses in the DNS responses sent to clients.
* `netmask` - IPv4 network mask for use in the SOURCEIPHASH load balancing method.
* `newname` - New name for the GSLB virtual server.
* `order` - Order number to be assigned to the service when it is bound to the lb vserver.
* `orderthreshold` - This option is used to to specify the threshold of minimum number of services to be UP in an order, for it to be considered in Lb decision.
* `persistenceid` - The persistence ID for the GSLB virtual server. The ID is a positive integer that enables GSLB sites to identify the GSLB virtual server, and is required if source IP address based or spill over based persistence is enabled on the virtual server.
* `persistencetype` - Use source IP address based persistence for the virtual server. After the load balancing method selects a service for the first packet, the IP address received in response to the DNS query is used for subsequent requests from the same client.
* `persistmask` - The optional IPv4 network mask applied to IPv4 addresses to establish source IP address based persistence.
* `rule` - Expression, or name of a named expression, against which traffic is evaluated. This field is applicable only if gslb method or gslb backup method are set to API.
* `servicegroupname` - The GSLB service group name bound to the selected GSLB virtual server.
* `servicename` - Name of the GSLB service for which to change the weight.
* `servicetype` - Protocol used by services bound to the virtual server.
* `sitedomainttl` - TTL, in seconds, for all internally created site domains (created when a site prefix is configured on a GSLB service) that are associated with this virtual server.
* `sobackupaction` - Action to be performed if spillover is to take effect, but no backup chain to spillover is usable or exists.
* `somethod` - Type of threshold that, when exceeded, triggers spillover. Available settings function as follows: CONNECTION - Spillover occurs when the number of client connections exceeds the threshold. DYNAMICCONNECTION - Spillover occurs when the number of client connections at the GSLB virtual server exceeds the sum of the maximum client (Max Clients) settings for bound GSLB services. BANDWIDTH - Spillover occurs when the bandwidth consumed by the GSLB virtual server's incoming and outgoing traffic exceeds the threshold. HEALTH - Spillover occurs when the percentage of weights of the GSLB services that are UP drops below the threshold. NONE - Spillover does not occur.
* `sopersistence` - If spillover occurs, maintain source IP address based persistence for both primary and backup GSLB virtual servers.
* `sopersistencetimeout` - Timeout for spillover persistence, in minutes.
* `sothreshold` - Threshold at which spillover occurs. Specify an integer for the CONNECTION spillover method, a bandwidth value in kilobits per second for the BANDWIDTH method (do not enter the units), or a percentage for the HEALTH method (do not enter the percentage symbol).
* `state` - State of the GSLB virtual server.
* `timeout` - Idle time, in minutes, after which a persistence entry is cleared.
* `toggleorder` - Configure this option to toggle order preference.
* `tolerance` - Tolerance in milliseconds. Tolerance value is used in deciding which sites in a GSLB configuration must be considered for implementing the RTT load balancing method. The sites having the RTT value less than or equal to the sum of the lowest RTT and tolerance value are considered.
* `ttl` - Time to live (TTL) for the domain.
* `v6netmasklen` - Number of bits to consider, in an IPv6 source IP address, for creating the hash that is required by the SOURCEIPHASH load balancing method.
* `v6persistmasklen` - Number of bits to consider in an IPv6 source IP address when creating source IP address based persistence sessions.
* `weight` - Weight for the service.
