---
subcategory: "Network"
---

# Data Source: nd6ravariables_onlinkipv6prefix_binding

The nd6ravariables_onlinkipv6prefix_binding data source allows you to retrieve information about a binding between nd6ravariables and an onlink IPv6 prefix.


## Example Usage

```terraform
data "citrixadc_nd6ravariables_onlinkipv6prefix_binding" "tf_binding" {
  vlan       = 1
  ipv6prefix = "2003::/64"
}

output "binding_id" {
  value = data.citrixadc_nd6ravariables_onlinkipv6prefix_binding.tf_binding.id
}

output "ipv6prefix" {
  value = data.citrixadc_nd6ravariables_onlinkipv6prefix_binding.tf_binding.ipv6prefix
}
```


## Argument Reference

* `vlan` - (Required) The VLAN number. Minimum value = 1, Maximum value = 4094.
* `ipv6prefix` - (Required) Onlink prefixes for RA messages.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nd6ravariables_onlinkipv6prefix_binding. It is a system-generated identifier.
