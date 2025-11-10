---
subcategory: "High Availability"
---

# Resource: hanode_routemonitor6_binding

The hanode_routemonitor6_binding resource is used to create hanode_routemonitor6_binding.


## Example usage

```hcl
resource "citrixadc_hanode_routemonitor6_binding" "tf_hanode_routemonitor6_binding" {
  hanode_id    = 0
  routemonitor = "fd7f:6bd8:cea9:f32d::/64"
}
```


## Argument Reference

* `hanode_id` - (Required) Number that uniquely identifies the local node. The ID of the local node is always 0.
* `routemonitor` - (Required) The IP address (IPv4 or IPv6).
* `netmask` - (Optional) The netmask.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the hanode_routemonitor6_binding. It is the concatenation of `hanode_id` and `routemonitor` attributes separated by a comma.


## Import

A hanode_routemonitor6_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_hanode_routemonitor6_binding.tf_hanode_routemonitor6_binding 0,fd7f:6bd8:cea9:f32d::/64
```
