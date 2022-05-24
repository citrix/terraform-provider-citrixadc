---
subcategory: "NS"
---

# Resource: nstrafficdomain_vxlan_binding

The nstrafficdomain_vxlan_binding resource is used to bind vxlan to nstrafficdomain resource.


## Example usage

```hcl
resource "citrixadc_nstrafficdomain" "tf_trafficdomain" {
  td        = 2
  aliasname = "tf_trafficdomain"
}
resource "citrixadc_vxlan" "tf_vxlan" {
  vxlanid            = 123
  port               = 33
  dynamicrouting     = "DISABLED"
  ipv6dynamicrouting = "DISABLED"
  innervlantagging   = "ENABLED"
}
resource "citrixadc_nstrafficdomain_vxlan_binding" "tf_binding" {
  td    = citrixadc_nstrafficdomain.tf_trafficdomain.td
  vxlan = citrixadc_vxlan.tf_vxlan.vxlanid
}
```


## Argument Reference

* `td` - (Required) Integer value that uniquely identifies a traffic domain. Minimum value =  1 Maximum value =  4094
* `vxlan` - (Required) ID of the VXLAN to bind to this traffic domain. More than one VXLAN can be bound to a traffic domain, but the same VXLAN cannot be a part of multiple traffic domains. Minimum value =  1 Maximum value =  16777215


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nstrafficdomain_vxlan_binding. It is the concatenation of `td` and `vxlan` attributes separated by comma.


## Import

A nstrafficdomain_vxlan_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_nstrafficdomain_vxlan_binding.tf_binding 2,123
```
