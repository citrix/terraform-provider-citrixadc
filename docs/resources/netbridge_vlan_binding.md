---
subcategory: "Network"
---

# Resource: netbridge_vlan_binding

The netbridge_vlan_binding resource is used to bind vlan to netbridge resource.


## Example usage

```hcl
resource "citrixadc_netbridge" "tf_netbridge" {
  name         = "tf_netbridge"
}
resource "citrixadc_vlan" "tf_vlan" {
  vlanid    = 20
  aliasname = "Management VLAN"
}
resource "citrixadc_netbridge_vlan_binding" "tf_binding" {
  name = citrixadc_netbridge.tf_netbridge.name
  vlan = citrixadc_vlan.tf_vlan.vlanid
}
```


## Argument Reference

* `name` - (Required) The name of the network bridge.
* `vlan` - (Required) The VLAN that is extended by this network bridge. Minimum value =  1 Maximum value =  4094


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the netbridge_vlan_binding. It is the concatenation of `name` and `vlan` attributes separated by comma.


## Import

A netbridge_vlan_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_netbridge_vlan_binding.tf_binding tf_netbridge,20
```
