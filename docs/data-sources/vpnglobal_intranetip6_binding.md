---
subcategory: "VPN"
---

# Data Source: vpnglobal_intranetip6_binding

The vpnglobal_intranetip6_binding data source allows you to retrieve information about a vpnglobal_intranetip6_binding.

## Example Usage

```terraform
data "citrixadc_vpnglobal_intranetip6_binding" "tf_bind" {
  intranetip6 = "2.3.4.5"
}

output "intranetip6" {
  value = data.citrixadc_vpnglobal_intranetip6_binding.tf_bind.intranetip6
}

output "numaddr" {
  value = data.citrixadc_vpnglobal_intranetip6_binding.tf_bind.numaddr
}

output "gotopriorityexpression" {
  value = data.citrixadc_vpnglobal_intranetip6_binding.tf_bind.gotopriorityexpression
}
```

## Argument Reference

* `intranetip6` - (Required) The intranet ip address or range.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `gotopriorityexpression` - Applicable only to advance vpn session policy. An expression or other value specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `numaddr` - The intranet ip address or range's netmask.
* `id` - The id of the vpnglobal_intranetip6_binding. It is a system-generated identifier.
