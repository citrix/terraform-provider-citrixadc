---
subcategory: "Vpn"
---

# Resource: vpnglobal_vpnportaltheme_binding

The vpnglobal_vpnportaltheme_binding resource is used to bind the vpnportaltheme to the vpnglobal configuration.


## Example usage

```hcl
resource "citrixadc_vpnportaltheme" "tf_vpnportaltheme" {
  name      = "tf_vpnportaltheme"
  basetheme = "X1"
}
resource "citrixadc_vpnglobal_vpnportaltheme_binding" "tf_bind" {
  portaltheme = citrixadc_vpnportaltheme.tf_vpnportaltheme.name
}
```


## Argument Reference

* `portaltheme` - (Required) Name of the portal theme bound to vpnglobal
* `gotopriorityexpression` - (Optional) Applicable only to advance vpn session policy. An expression or other value specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vpnglobal_vpnportaltheme_binding. It has the same value as the `portaltheme` attribute.


## Import

A vpnglobal_vpnportaltheme_binding can be imported using its portaltheme, e.g.

```shell
terraform import citrixadc_vpnglobal_vpnportaltheme_binding.tf_bind tf_vpnportaltheme
```
