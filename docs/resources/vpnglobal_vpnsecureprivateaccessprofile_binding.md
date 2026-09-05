---
subcategory: "VPN"
---

# Resource: vpnglobal_vpnsecureprivateaccessprofile_binding

This resource is used to bind a Secure Private Access Profile to the global VPN bind point.


## Example usage

```hcl
resource "citrixadc_vpnglobal_vpnsecureprivateaccessprofile_binding" "tf_bind" {
  secureprivateaccessprofile = "tf_spa_profile"
  gotopriorityexpression     = "END"
}
```

~> **NOTE:** The Secure Private Access Profile referenced by `secureprivateaccessprofile` is a
feature-gated object created by the Secure Private Access (SPA) subsystem and must already exist on
the ADC before this binding can be created.


## Argument Reference

* `secureprivateaccessprofile` - (Required) The name of the Secure Private Access Profile bound to vpn global. Changing this value forces a new binding to be created.
* `gotopriorityexpression` - (Optional) Applicable only to advance vpn session policy. An expression or other value specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE. Changing this value forces a new binding to be created.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vpnglobal_vpnsecureprivateaccessprofile_binding. It has the same value as the `secureprivateaccessprofile` attribute.


## Import

A vpnglobal_vpnsecureprivateaccessprofile_binding can be imported using its secureprivateaccessprofile value (which is also the resource id), e.g.

```shell
terraform import citrixadc_vpnglobal_vpnsecureprivateaccessprofile_binding.tf_bind "tf_spa_profile"
```
