---
subcategory: "DNS"
---

# Resource: dnsaaaarec

The dnsaaaarec resource is used to create DNS aaaarec.


## Example usage

```hcl
resource "citrixadc_dnsaaaarec" "dnsaaaarec" {
	hostname = "www.adfihrwpi.com"
    ipv6address = "2001:db8:85a3::8a2e:370:7334"
    ttl = 3600
}

```


## Argument Reference

* `hostname` - (Required) Domain name.
* `ipv6address` - (Required) One or more IPv6 addresses to assign to the domain name.
* `ecssubnet` - (Optional) Subnet for which the cached records need to be removed.
* `nodeid` - (Optional) Unique number that identifies the cluster node.
* `ttl` - (Optional) Time to Live (TTL), in seconds, for the record. TTL is the time for which the record must be cached by DNS proxies. The specified TTL is applied to all the resource records that are of the same record type and belong to the specified domain name. For example, if you add an address record, with a TTL of 36000, to the domain name example.com, the TTLs of all the address records of example.com are changed to 36000. If the TTL is not specified, the Citrix ADC uses either the DNS zone's minimum TTL or, if the SOA record is not available on the appliance, the default value of 3600.
* `type` - (Optional) Type of records to display. Available settings function as follows: * ADNS - Display all authoritative address records. * PROXY - Display all proxy address records. * ALL - Display all address records.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the dnsaaaarec. It has the same value as the `hostname` attribute.


## Import

A dnsaaaarec can be imported using its name, e.g.

```shell
terraform import citrixadc_dnsaaaarec.dnsaaaarec www.adfihrwpi.com
```
