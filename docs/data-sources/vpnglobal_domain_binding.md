---
subcategory: "VPN"
---

# Data Source: vpnglobal_domain_binding

The vpnglobal_domain_binding data source allows you to retrieve information about a vpnglobal_domain_binding.

## Example Usage

```terraform
data "citrixadc_vpnglobal_domain_binding" "tf_bind" {
  intranetdomain = "http://www.example.com/"
}

output "intranetdomain" {
  value = data.citrixadc_vpnglobal_domain_binding.tf_bind.intranetdomain
}

output "gotopriorityexpression" {
  value = data.citrixadc_vpnglobal_domain_binding.tf_bind.gotopriorityexpression
}
```

## Argument Reference

* `intranetdomain` - (Required) The conflicting intranet domain name.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `gotopriorityexpression` - Applicable only to advance vpn session policy. An expression or other value specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `id` - The id of the vpnglobal_domain_binding. It is a system-generated identifier.
