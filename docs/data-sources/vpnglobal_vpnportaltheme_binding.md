---
subcategory: "VPN"
---

# Data Source: vpnglobal_vpnportaltheme_binding

The vpnglobal_vpnportaltheme_binding data source allows you to retrieve information about a VPN portal theme binding to the global VPN configuration.

## Example Usage

```terraform
data "citrixadc_vpnglobal_vpnportaltheme_binding" "tf_bind" {
  portaltheme = "tf_vpnportaltheme"
}

output "portaltheme" {
  value = data.citrixadc_vpnglobal_vpnportaltheme_binding.tf_bind.portaltheme
}

output "gotopriorityexpression" {
  value = data.citrixadc_vpnglobal_vpnportaltheme_binding.tf_bind.gotopriorityexpression
}
```

## Argument Reference

* `portaltheme` - (Required) Name of the portal theme bound to vpnglobal.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `gotopriorityexpression` - Applicable only to advance vpn session policy. An expression or other value specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `id` - The id of the vpnglobal_vpnportaltheme_binding. It is a system-generated identifier.
