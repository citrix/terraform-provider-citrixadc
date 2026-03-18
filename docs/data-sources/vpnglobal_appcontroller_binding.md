---
subcategory: "VPN"
---

# Data Source: vpnglobal_appcontroller_binding

The vpnglobal_appcontroller_binding data source allows you to retrieve information about a vpnglobal_appcontroller_binding.

## Example Usage

```terraform
data "citrixadc_vpnglobal_appcontroller_binding" "tf_vpnglobal_appcontroller_binding" {
  appcontroller = "http://www.citrix.com"
}

output "appcontroller" {
  value = data.citrixadc_vpnglobal_appcontroller_binding.tf_vpnglobal_appcontroller_binding.appcontroller
}

output "gotopriorityexpression" {
  value = data.citrixadc_vpnglobal_appcontroller_binding.tf_vpnglobal_appcontroller_binding.gotopriorityexpression
}
```

## Argument Reference

* `appcontroller` - (Required) Configured App Controller server.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `gotopriorityexpression` - Applicable only to advance vpn session policy. An expression or other value specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `id` - The id of the vpnglobal_appcontroller_binding. It is a system-generated identifier.
