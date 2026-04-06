---
subcategory: "VPN"
---

# Data Source: vpnglobal_sharefileserver_binding

The vpnglobal_sharefileserver_binding data source allows you to retrieve information about a vpnglobal_sharefileserver_binding.

## Example Usage

```terraform
data "citrixadc_vpnglobal_sharefileserver_binding" "tf_vpnglobal_sharefileserver_binding" {
  sharefile = "3.4.5.2:8080"
}

output "sharefile" {
  value = data.citrixadc_vpnglobal_sharefileserver_binding.tf_vpnglobal_sharefileserver_binding.sharefile
}

output "gotopriorityexpression" {
  value = data.citrixadc_vpnglobal_sharefileserver_binding.tf_vpnglobal_sharefileserver_binding.gotopriorityexpression
}
```

## Argument Reference

* `sharefile` - (Required) Configured Sharefile server, in the format IP:PORT / FQDN:PORT

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `gotopriorityexpression` - Applicable only to advance vpn session policy. An expression or other value specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `id` - The id of the vpnglobal_sharefileserver_binding. It is a system-generated identifier.
