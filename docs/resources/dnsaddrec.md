---
subcategory: "DNS"
---

# Resource: dnsaddrec

The dnsaddrec resource is used to create DNS addrec.


## Example usage

```hcl
resource "citrixadc_dnsaddrec" "dnsaddrec" {
  hostname  = "a.root-servers.net"
  ipaddress = "65.200.211.129"
  ttl       = 3600
}
```


## Argument Reference

* `hostname` - (Required) Domain name.
* `ipaddress` - (Required) One or more IPv4 addresses to assign to the domain name.
* `ecssubnet` - (Optional) Subnet for which the cached address records need to be removed.
* `nodeid` - (Optional) Unique number that identifies the cluster node.
* `ttl` - (Optional) Time to Live (TTL), in seconds, for the record. TTL is the time for which the record must be cached by DNS proxies. The specified TTL is applied to all the resource records that are of the same record type and belong to the specified domain name. For example, if you add an address record, with a TTL of 36000, to the domain name example.com, the TTLs of all the address records of example.com are changed to 36000. If the TTL is not specified, the Citrix ADC uses either the DNS zone's minimum TTL or, if the SOA record is not available on the appliance, the default value of 3600.
* `type` - (Optional) The address record type. The type can take 3 values: ADNS -  If this is specified, all of the authoritative address records will be displayed. PROXY - If this is specified, all of the proxy address records will be displayed. ALL  -  If this is specified, all of the address records will be displayed.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the dnsaddrec It has the same value as the `domain` attribute.


## Import

A dnsaddrec can be imported using its name, e.g.

```shell
terraform import citrixadc_dnsaddrec.dnsaddrec ab.root-servers.net
```
