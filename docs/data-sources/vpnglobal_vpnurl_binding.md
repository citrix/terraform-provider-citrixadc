---
subcategory: "VPN"
---

# Data Source: vpnglobal_vpnurl_binding

The vpnglobal_vpnurl_binding data source allows you to retrieve information about a vpn url binding to the global configuration.


## Example usage

```terraform
data "citrixadc_vpnglobal_vpnurl_binding" "tf_bind" {
  urlname = "Firsturl"
}

output "urlname" {
  value = data.citrixadc_vpnglobal_vpnurl_binding.tf_bind.urlname
}

output "gotopriorityexpression" {
  value = data.citrixadc_vpnglobal_vpnurl_binding.tf_bind.gotopriorityexpression
}
```


## Argument Reference

* `urlname` - (Required) The intranet url.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vpnglobal_vpnurl_binding. It is a system-generated identifier.
* `gotopriorityexpression` - Applicable only to advance vpn session policy. An expression or other value specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
