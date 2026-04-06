---
subcategory: "AAA"
---

# Data Source: aaagroup_authorizationpolicy_binding

The aaagroup_authorizationpolicy_binding data source allows you to retrieve information about an authorization policy binding to an AAA group.

## Example Usage

```terraform
data "citrixadc_aaagroup_authorizationpolicy_binding" "tf_aaagroup_authorizationpolicy_binding" {
  groupname = "my_group"
  policy    = "tp-authorize-1"
}

output "priority" {
  value = data.citrixadc_aaagroup_authorizationpolicy_binding.tf_aaagroup_authorizationpolicy_binding.priority
}

output "gotopriorityexpression" {
  value = data.citrixadc_aaagroup_authorizationpolicy_binding.tf_aaagroup_authorizationpolicy_binding.gotopriorityexpression
}
```

## Argument Reference

* `groupname` - (Required) Name of the group that you are binding.
* `policy` - (Required) The policy name.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `gotopriorityexpression` - Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `priority` - Integer specifying the priority of the policy. A lower number indicates a higher priority. Policies are evaluated in the order of their priority numbers. Maximum value for default syntax policies is 2147483647 and for classic policies is 64000.
* `id` - The id of the aaagroup_authorizationpolicy_binding. It is a system-generated identifier.
* `type` - Bindpoint to which the policy is bound.
