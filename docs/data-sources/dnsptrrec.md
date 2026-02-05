---
subcategory: "DNS"
---

# Data Source `dnsptrrec`

The dnsptrrec data source allows you to retrieve information about DNS PTR (Pointer) records configured on the Citrix ADC.


## Example usage

```terraform
data "citrixadc_dnsptrrec" "tf_dnsptrrec" {
  reversedomain = "0.2.0.192.in-addr.arpa"
}

output "reversedomain" {
  value = data.citrixadc_dnsptrrec.tf_dnsptrrec.reversedomain
}

output "domain" {
  value = data.citrixadc_dnsptrrec.tf_dnsptrrec.domain
}

output "ttl" {
  value = data.citrixadc_dnsptrrec.tf_dnsptrrec.ttl
}
```


## Argument Reference

* `reversedomain` - (Required) Reversed domain name representation of the IPv4 or IPv6 address for which to retrieve the PTR record. Use the "in-addr.arpa." suffix for IPv4 addresses and the "ip6.arpa." suffix for IPv6 addresses.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the dnsptrrec. It is a unique identifier for the record.
* `domain` - Domain name for which the reverse mapping is configured.
* `ecssubnet` - Subnet for which the cached PTR record need to be removed.
* `nodeid` - Unique number that identifies the cluster node.
* `ttl` - Time to Live (TTL), in seconds, for the record. TTL is the time for which the record must be cached by DNS proxies. The specified TTL is applied to all the resource records that are of the same record type and belong to the specified domain name. For example, if you add an address record, with a TTL of 36000, to the domain name example.com, the TTLs of all the address records of example.com are changed to 36000. If the TTL is not specified, the Citrix ADC uses either the DNS zone's minimum TTL or, if the SOA record is not available on the appliance, the default value of 3600.
