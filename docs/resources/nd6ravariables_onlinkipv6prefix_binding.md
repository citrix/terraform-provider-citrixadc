---
subcategory: "Network"
---

# Resource: nd6ravariables_onlinkipv6prefix_binding

The nd6ravariables_onlinkipv6prefix_binding resource is used to create nd6ravariables_onlinkipv6prefix_binding.


## Example usage

```hcl
resource "citrixadc_nd6ravariables_onlinkipv6prefix_binding" "tf_nd6ravariables_onlinkipv6prefix_binding" {
  vlan      = "1"
  ipv6prefix = "2003::/64"
}
```


## Argument Reference

* `ipv6prefix` - (Required) Onlink prefixes for RA messages.
* `vlan` - (Required) The VLAN number. Minimum value =  1 Maximum value =  4094


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nd6ravariables_onlinkipv6prefix_binding. It has the same value as the `vlan` and `ipv6prefix` attributes separated by a comma.


## Import

A nd6ravariables_onlinkipv6prefix_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_nd6ravariables_onlinkipv6prefix_binding.tf_nd6ravariables_onlinkipv6prefix_binding 1,2003::/64
```
