---
subcategory: "DNS"
---

# Resource: dnsptrrec

The dnsptrrec resource is used to create DNS ptrRec.


## Example usage

```hcl
resource "citrixadc_dnsptrrec" "tf_dnsptrrec" {
	reversedomain = "0.2.0.192.in-addr.arpa"
	domain        = "example.com"
	ttl           = 3600
  }
```


## Argument Reference

* `reversedomain` - (Required) Reversed domain name representation of the IPv4 or IPv6 address for which to create the PTR record. Use the "in-addr.arpa." suffix for IPv4 addresses and the "ip6.arpa." suffix for IPv6 addresses.
* `domain` - (Required) Domain name for which to configure reverse mapping.
* `ecssubnet` - (Optional) Subnet for which the cached PTR record need to be removed.
* `nodeid` - (Optional) Unique number that identifies the cluster node.
* `ttl` - (Optional) Time to Live (TTL), in seconds, for the record. TTL is the time for which the record must be cached by DNS proxies. The specified TTL is applied to all the resource records that are of the same record type and belong to the specified domain name. For example, if you add an address record, with a TTL of 36000, to the domain name example.com, the TTLs of all the address records of example.com are changed to 36000. If the TTL is not specified, the Citrix ADC uses either the DNS zone's minimum TTL or, if the SOA record is not available on the appliance, the default value of 3600.
* `type` - (Optional) Type of records to display. Available settings function as follows: * ADNS - Display all authoritative address records. * PROXY - Display all proxy address records. * ALL - Display all address records.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the dnsptrrec. It has the same value as the `reversedomain` attribute.


## Import

A dnsptrrec can be imported using its name, e.g.

```shell
terraform import citrixadc_dnsptrrec.tf_dnsptrrec 0.2.0.192.in-addr.arpa
```
