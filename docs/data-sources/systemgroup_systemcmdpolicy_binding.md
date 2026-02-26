---
subcategory: "System"
---

# Data Source: systemgroup_systemcmdpolicy_binding

This data source is used to retrieve information about a specific `systemgroup_systemcmdpolicy_binding` configuration.

## Example usage

```hcl
data "citrixadc_systemgroup_systemcmdpolicy_binding" "example" {
  groupname  = "tf_systemgroup"
  policyname = "tf_policy"
}

output "binding_id" {
  value = data.citrixadc_systemgroup_systemcmdpolicy_binding.example.id
}

output "binding_priority" {
  value = data.citrixadc_systemgroup_systemcmdpolicy_binding.example.priority
}
```

## Argument Reference

* `groupname` - (Required) Name of the system group.
* `policyname` - (Required) The name of command policy.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The ID of the binding.
* `priority` - The priority of the command policy.
