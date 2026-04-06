---
subcategory: "VPN"
---

# Data Source: vpnglobal_vpneula_binding

The vpnglobal_vpneula_binding data source allows you to retrieve information about a specific vpneula binding to vpnglobal configuration.


## Example Usage

```terraform
data "citrixadc_vpnglobal_vpneula_binding" "example" {
  eula = "tf_vpneula"
}

output "eula" {
  value = data.citrixadc_vpnglobal_vpneula_binding.example.eula
}

output "gotopriorityexpression" {
  value = data.citrixadc_vpnglobal_vpneula_binding.example.gotopriorityexpression
}
```

## Argument Reference

* `eula` - (Required) Name of the EULA bound to vpnglobal


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `gotopriorityexpression` - Applicable only to advance vpn session policy. An expression or other value specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `id` - The id of the vpnglobal_vpneula_binding. It has the same value as the `eula` attribute.

