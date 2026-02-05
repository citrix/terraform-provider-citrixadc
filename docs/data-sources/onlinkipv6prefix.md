---
subcategory: "Network"
---

# Data Source `onlinkipv6prefix`

The onlinkipv6prefix data source allows you to retrieve information about an existing IPv6 prefix configuration for Router Advertisement (RA) messages.

## Example usage

```terraform
data "citrixadc_onlinkipv6prefix" "tf_onlinkipv6prefix" {
  ipv6prefix = "9000::/64"
}

output "ipv6prefix" {
  value = data.citrixadc_onlinkipv6prefix.tf_onlinkipv6prefix.ipv6prefix
}

output "onlinkprefix" {
  value = data.citrixadc_onlinkipv6prefix.tf_onlinkipv6prefix.onlinkprefix
}

output "autonomusprefix" {
  value = data.citrixadc_onlinkipv6prefix.tf_onlinkipv6prefix.autonomusprefix
}
```

## Argument Reference

* `ipv6prefix` - (Required) Onlink prefixes for RA messages.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the onlinkipv6prefix resource. It has the same value as the `ipv6prefix` attribute.
* `autonomusprefix` - RA Prefix Autonomus flag. Possible values: [ YES, NO ]
* `decrementprefixlifetimes` - RA Prefix Autonomus flag.
* `depricateprefix` - Depricate the prefix.
* `onlinkprefix` - RA Prefix onlink flag. Possible values: [ YES, NO ]
* `prefixpreferredlifetime` - Preferred life time of the prefix, in seconds.
* `prefixvalidelifetime` - Valid life time of the prefix, in seconds.
