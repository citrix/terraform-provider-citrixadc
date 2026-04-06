---
subcategory: "AAA"
---

# Data Source: aaauser_intranetip6_binding

The aaauser_intranetip6_binding data source allows you to retrieve information about a specific binding between an AAA user and an intranet IPv6 address.

## Example Usage

```terraform
data "citrixadc_aaauser_intranetip6_binding" "example" {
  username    = "user1"
  intranetip6 = "2003:db8:100::fb/128"
}

output "username" {
  value = data.citrixadc_aaauser_intranetip6_binding.example.username
}

output "intranetip6" {
  value = data.citrixadc_aaauser_intranetip6_binding.example.intranetip6
}

output "gotopriorityexpression" {
  value = data.citrixadc_aaauser_intranetip6_binding.example.gotopriorityexpression
}
```

## Argument Reference

* `username` - (Required) User account to which to bind the policy.
* `intranetip6` - (Required) The Intranet IP6 bound to the user.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `gotopriorityexpression` - Expression or other value specifying the next policy to evaluate if the current policy evaluates to TRUE. Specify one of the following values: NEXT - Evaluate the policy with the next higher priority number. END - End policy evaluation. USE_INVOCATION_RESULT - Applicable if this policy invokes another policy label. If the final goto in the invoked policy label has a value of END, the evaluation stops. If the final goto is anything other than END, the current policy label performs a NEXT. An expression that evaluates to a number.
* `id` - The id of the aaauser_intranetip6_binding. It is a system-generated identifier.
* `numaddr` - Numbers of ipv6 address bound starting with intranetip6.
