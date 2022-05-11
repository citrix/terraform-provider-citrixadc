---
subcategory: "Network"
---

# Resource: netbridge

The netbridge resource is used to create network bridge resource.


## Example usage

```hcl
resource "citrixadc_vxlanvlanmap" "tf_vxlanvlanmp" {
  name = "tf_vxlanvlanmp"
}
resource "citrixadc_netbridge" "tf_netbridge" {
  name         = "tf_netbridge"
  vxlanvlanmap = citrixadc_vxlanvlanmap.tf_vxlanvlanmp.name
}
```


## Argument Reference

* `name` - (Required) The name of the network bridge.
* `vxlanvlanmap` - (Optional) The vlan to vxlan mapping to be applied to this netbridge.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the netbridge. It has the same value as the `name` attribute.


## Import

A netbridge can be imported using its name, e.g.

```shell
terraform import citrixadc_netbridge.tf_netbridge tf_netbridge
```
