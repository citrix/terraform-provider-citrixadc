---
subcategory: "NS"
---

# Data Source: nssourceroutecachetable

The nssourceroutecachetable data source retrieves the read-only attributes of the NITRO `nssourceroutecachetable` (Source IP MAC cache table) object via its keyless `get (all)` endpoint.

~> **NOTE** `get (all)` returns a table (list) of cache entries; this data source exposes the first entry. If the cache table is empty the read returns an error.


## Example usage

```terraform
data "citrixadc_nssourceroutecachetable" "example" {
}

output "sourceip" {
  value = data.citrixadc_nssourceroutecachetable.example.sourceip
}
```


## Argument Reference

This data source takes no arguments.


## Attribute Reference

The following attributes are available:

* `id` - The id of the nssourceroutecachetable data source. It is a synthetic value (`nssourceroutecachetable-config`).
* `sourceip` - Source ip of the connection.
* `sourcemac` - Source MAC address of an incoming IPv6 packet.
* `vlan` - ID of the VLAN.
* `interface` - ID of an interface (NITRO key `Interface`).
* `nextgenapiresource` - Read-only attribute (NITRO key `_nextgenapiresource`).
* `count` - Count parameter (NITRO key `__count`).
