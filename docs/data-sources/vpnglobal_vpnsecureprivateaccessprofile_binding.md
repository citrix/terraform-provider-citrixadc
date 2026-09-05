---
subcategory: "VPN"
---

# Data Source: vpnglobal_vpnsecureprivateaccessprofile_binding

The vpnglobal_vpnsecureprivateaccessprofile_binding data source allows you to retrieve information about a Secure Private Access Profile bound to the global VPN bind point on the Citrix ADC.


## Example usage

```terraform
data "citrixadc_vpnglobal_vpnsecureprivateaccessprofile_binding" "tf_bind" {
  secureprivateaccessprofile = "tf_spa_profile"
}

output "secureprivateaccessprofile" {
  value = data.citrixadc_vpnglobal_vpnsecureprivateaccessprofile_binding.tf_bind.secureprivateaccessprofile
}

output "gotopriorityexpression" {
  value = data.citrixadc_vpnglobal_vpnsecureprivateaccessprofile_binding.tf_bind.gotopriorityexpression
}
```


## Argument Reference

* `secureprivateaccessprofile` - (Required) The name of the Secure Private Access Profile used to look up the binding.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vpnglobal_vpnsecureprivateaccessprofile_binding. It has the same value as the `secureprivateaccessprofile` attribute.
* `gotopriorityexpression` - Applicable only to advance vpn session policy. An expression or other value specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
