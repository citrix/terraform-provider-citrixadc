---
subcategory: "NS"
---

# Resource: nstrafficdomain_vlan_binding

The nstrafficdomain_vlan_binding resource is used to bind vlan to the nstrafficdomain.


## Example usage

```hcl
resource "citrixadc_nstrafficdomain" "tf_trafficdomain" {
  td        = 2
  aliasname = "tf_trafficdomain"
  vmac      = "DISABLED"
}
resource "citrixadc_vlan" "tf_vlan" {
  vlanid    = 20
  aliasname = "Management VLAN"
}
resource "citrixadc_nstrafficdomain_vlan_binding" "tf_binding" {
  td   = citrixadc_nstrafficdomain.tf_trafficdomain.td
  vlan = citrixadc_vlan.tf_vlan.vlanid
}
```


## Argument Reference

* `td` - (Required) Integer value that uniquely identifies a traffic domain. Minimum value =  1 Maximum value =  4094
* `vlan` - (Required) ID of the VLAN to bind to this traffic domain. More than one VLAN can be bound to a traffic domain, but the same VLAN cannot be a part of multiple traffic domains. Minimum value =  1 Maximum value =  4094


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nstrafficdomain_vlan_binding. It is the concatenation of `td` and `vlan` separated by comma ",".


## Import

A nstrafficdomain_vlan_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_nstrafficdomain_vlan_binding.tf_binding 2,20
```
