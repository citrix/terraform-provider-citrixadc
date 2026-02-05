---
subcategory: "DNS"
---

# Data Source `dnsaddrec`

The dnsaddrec data source allows you to retrieve information about DNS A (Address) records configured on the Citrix ADC.


## Example usage

```terraform
data "citrixadc_dnsaddrec" "tf_dnsaddrec" {
  hostname  = "example.com"
  ipaddress = "192.168.1.10"
}

output "hostname" {
  value = data.citrixadc_dnsaddrec.tf_dnsaddrec.hostname
}

output "ipaddress" {
  value = data.citrixadc_dnsaddrec.tf_dnsaddrec.ipaddress
}

output "ttl" {
  value = data.citrixadc_dnsaddrec.tf_dnsaddrec.ttl
}
```


## Argument Reference

* `hostname` - (Required) Domain name for which to retrieve the address record.
* `ipaddress` - (Required) IPv4 address to assign to the domain name.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the dnsaddrec. It is a unique identifier for the record.
* `ecssubnet` - Subnet for which the cached address records need to be removed.
* `nodeid` - Unique number that identifies the cluster node.
* `ttl` - Time to Live (TTL), in seconds, for the record. TTL is the time for which the record must be cached by DNS proxies. The specified TTL is applied to all the resource records that are of the same record type and belong to the specified domain name. For example, if you add an address record, with a TTL of 36000, to the domain name example.com, the TTLs of all the address records of example.com are changed to 36000. If the TTL is not specified, the Citrix ADC uses either the DNS zone's minimum TTL or, if the SOA record is not available on the appliance, the default value of 3600.
