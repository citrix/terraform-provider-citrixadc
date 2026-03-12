---
subcategory: "VPN"
---

# Data Source: vpnglobal_vpnnexthopserver_binding

The vpnglobal_vpnnexthopserver_binding data source allows you to retrieve information about a vpnnexthopserver binding to vpnglobal configuration.


## Example Usage

```terraform
data "citrixadc_vpnglobal_vpnnexthopserver_binding" "tf_bind" {
  nexthopserver = "tf_vpnnexthopserver"
}

output "nexthopserver" {
  value = data.citrixadc_vpnglobal_vpnnexthopserver_binding.tf_bind.nexthopserver
}

output "gotopriorityexpression" {
  value = data.citrixadc_vpnglobal_vpnnexthopserver_binding.tf_bind.gotopriorityexpression
}
```


## Argument Reference

* `nexthopserver` - (Required) The name of the next hop server bound to vpn global.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `gotopriorityexpression` - Applicable only to advance vpn session policy. An expression or other value specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `id` - The id of the vpnglobal_vpnnexthopserver_binding. It has the same value as the `nexthopserver` attribute.
