---
subcategory: "High Availability"
---

# Resource: hanode_routemonitor_binding

The hanode_routemonitor_binding resource is used to create hanode_routemonitor_binding.


## Example usage

```hcl
resource "citrixadc_hanode_routemonitor_binding" "tf_hanode_routemonitor_binding" {
  hanode_id    = 0
  routemonitor = "10.222.74.128"
  netmask      = "255.255.255.192"
}
```


## Argument Reference

* `id` - (Required) Number that uniquely identifies the local node. The ID of the local node is always 0.
* `netmask` - (Required) The netmask.
* `routemonitor` - (Required) The IP address (IPv4 or IPv6) Notice: this adress should be the network adress.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the hanode_routemonitor_binding. It has the same value as the `name` attribute.


## Import

A hanode_routemonitor_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_hanode_routemonitor_binding.tfhanode_routemonitor_binding 0,10.222.74.128
```
