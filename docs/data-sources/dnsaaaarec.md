---
subcategory: "DNS"
---

# Data Source `dnsaaaarec`

The dnsaaaarec data source allows you to retrieve information about DNS AAAA (IPv6 Address) records configured on the Citrix ADC.


## Example usage

```terraform
data "citrixadc_dnsaaaarec" "tf_dnsaaaarec" {
  hostname    = "www.example.com"
  ipv6address = "2001:db8:85a3::8a2e:370:7334"
  type        = "ALL"
}

output "ttl" {
  value = data.citrixadc_dnsaaaarec.tf_dnsaaaarec.ttl
}

output "hostname" {
  value = data.citrixadc_dnsaaaarec.tf_dnsaaaarec.hostname
}

output "ipv6address" {
  value = data.citrixadc_dnsaaaarec.tf_dnsaaaarec.ipv6address
}
```


## Argument Reference

* `hostname` - (Required) Domain name.
* `ipv6address` - (Required) One or more IPv6 addresses to assign to the domain name.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the dnsaaaarec. It is a combination of hostname, ipv6address, and type.
* `ecssubnet` - Subnet for which the cached records need to be removed.
* `nodeid` - Unique number that identifies the cluster node.
* `ttl` - Time to Live (TTL), in seconds, for the record. TTL is the time for which the record must be cached by DNS proxies. The specified TTL is applied to all the resource records that are of the same record type and belong to the specified domain name. For example, if you add an address record, with a TTL of 36000, to the domain name example.com, the TTLs of all the address records of example.com are changed to 36000. If the TTL is not specified, the Citrix ADC uses either the DNS zone's minimum TTL or, if the SOA record is not available on the appliance, the default value of 3600.
