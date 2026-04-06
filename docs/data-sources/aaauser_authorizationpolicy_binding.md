---
subcategory: "AAA"
---

# Data Source: aaauser_authorizationpolicy_binding

The aaauser_authorizationpolicy_binding data source allows you to retrieve information about a specific binding between an AAA user and an authorization policy.

## Example Usage

```terraform
data "citrixadc_aaauser_authorizationpolicy_binding" "example" {
  username = "user1"
  policy   = "tp-authorize-1"
}

output "priority" {
  value = data.citrixadc_aaauser_authorizationpolicy_binding.example.priority
}

output "gotopriorityexpression" {
  value = data.citrixadc_aaauser_authorizationpolicy_binding.example.gotopriorityexpression
}
```

## Argument Reference

* `username` - (Required) User account to which to bind the policy.
* `policy` - (Required) The policy Name.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `gotopriorityexpression` - Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `id` - The id of the aaauser_authorizationpolicy_binding. It is a system-generated identifier.
* `type` - Bindpoint to which the policy is bound. Possible values: [ REQUEST, UDP_REQUEST, DNS_REQUEST, ICMP_REQUEST ]
* `priority` - Integer specifying the priority of the policy. A lower number indicates a higher priority. Policies are evaluated in the order of their priority numbers. Maximum value for default syntax policies is 2147483647 and for classic policies max priority is 64000.
