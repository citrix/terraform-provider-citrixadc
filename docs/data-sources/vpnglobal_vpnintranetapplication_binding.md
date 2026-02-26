---
subcategory: "VPN"
---

# Data Source: vpnglobal_vpnintranetapplication_binding

The vpnglobal_vpnintranetapplication_binding data source allows you to retrieve information about a vpnintranetapplication binding to vpnglobal configuration.


## Example Usage

```terraform
data "citrixadc_vpnglobal_vpnintranetapplication_binding" "tf_bind" {
  intranetapplication = "tf_vpnintranetapplication"
}

output "intranetapplication" {
  value = data.citrixadc_vpnglobal_vpnintranetapplication_binding.tf_bind.intranetapplication
}

output "gotopriorityexpression" {
  value = data.citrixadc_vpnglobal_vpnintranetapplication_binding.tf_bind.gotopriorityexpression
}
```


## Argument Reference

* `intranetapplication` - (Required) The intranet vpn application.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `gotopriorityexpression` - Applicable only to advance vpn session policy. An expression or other value specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `id` - The id of the vpnglobal_vpnintranetapplication_binding. It has the same value as the `intranetapplication` attribute.
