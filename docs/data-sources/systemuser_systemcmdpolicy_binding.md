---
subcategory: "System"
---

# Data Source: systemuser_systemcmdpolicy_binding

The systemuser_systemcmdpolicy_binding data source allows you to retrieve information about a specific binding between a system user and a system command policy.

## Example usage

```terraform
data "citrixadc_systemuser_systemcmdpolicy_binding" "tf_bind" {
  username   = "tf_user"
  policyname = "tf_policy"
}

output "priority" {
  value = data.citrixadc_systemuser_systemcmdpolicy_binding.tf_bind.priority
}

output "username" {
  value = data.citrixadc_systemuser_systemcmdpolicy_binding.tf_bind.username
}

output "policyname" {
  value = data.citrixadc_systemuser_systemcmdpolicy_binding.tf_bind.policyname
}
```

## Argument Reference

* `username` - (Required) Name of the system-user entry to which to bind the command policy.
* `policyname` - (Required) The name of command policy.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the systemuser_systemcmdpolicy_binding. It is a system-generated identifier.
* `priority` - The priority of the policy.
