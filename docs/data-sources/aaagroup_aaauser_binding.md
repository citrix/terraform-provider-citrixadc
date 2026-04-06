---
subcategory: "AAA"
---

# Data Source: aaagroup_aaauser_binding

The `aaagroup_aaauser_binding` data source allows you to retrieve information about a specific binding between an AAA group and a user. This binding represents the association of a user to a group for authentication, authorization, and auditing purposes.

## Example usage

```terraform
data "citrixadc_aaagroup_aaauser_binding" "tf_aaagroup_aaauser_binding" {
  groupname = "my_group"
  username  = "user1"
}

output "groupname" {
  value = data.citrixadc_aaagroup_aaauser_binding.tf_aaagroup_aaauser_binding.groupname
}

output "username" {
  value = data.citrixadc_aaagroup_aaauser_binding.tf_aaagroup_aaauser_binding.username
}
```

## Argument Reference

* `groupname` - (Required) Name of the group that you are binding.
* `username` - (Required) The user name.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `gotopriorityexpression` - Expression or other value specifying the next policy to evaluate if the current policy evaluates to TRUE. Specify one of the following values:
  * `NEXT` - Evaluate the policy with the next higher priority number.
  * `END` - End policy evaluation.
  * `USE_INVOCATION_RESULT` - Applicable if this policy invokes another policy label. If the final goto in the invoked policy label has a value of END, the evaluation stops. If the final goto is anything other than END, the current policy label performs a NEXT.
  * An expression that evaluates to a number - If you specify an expression, the number to which it evaluates determines the next policy to evaluate.
* `id` - The id of the aaagroup_aaauser_binding. It is a system-generated identifier.
