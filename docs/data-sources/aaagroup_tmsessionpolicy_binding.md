---
subcategory: "AAA"
---

# Data Source: aaagroup_tmsessionpolicy_binding

The aaagroup_tmsessionpolicy_binding data source allows you to retrieve information about a specific binding between an AAA group and a TM session policy.

## Example Usage

```terraform
data "citrixadc_aaagroup_tmsessionpolicy_binding" "example" {
  groupname = "my_group"
  policy    = "my_tmsession_policy"
}

output "priority" {
  value = data.citrixadc_aaagroup_tmsessionpolicy_binding.example.priority
}

output "gotopriorityexpression" {
  value = data.citrixadc_aaagroup_tmsessionpolicy_binding.example.gotopriorityexpression
}
```

## Argument Reference

* `groupname` - (Required) Name of the group that you are binding.
* `policy` - (Required) The policy name.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `gotopriorityexpression` - Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `id` - The id of the aaagroup_tmsessionpolicy_binding. It is a system-generated identifier.
* `priority` - Integer specifying the priority of the policy. A lower number indicates a higher priority. Policies are evaluated in the order of their priority numbers. Maximum value for default syntax policies is 2147483647 and for classic policies is 64000.
* `type` - Bindpoint to which the policy is bound.
