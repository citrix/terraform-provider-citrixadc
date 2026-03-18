---
subcategory: "System"
---

# Data Source: systemglobal_authenticationpolicy_binding

The `citrixadc_systemglobal_authenticationpolicy_binding` data source allows you to retrieve information about a specific binding between the system global configuration and an authentication policy.

## Example usage

```terraform
data "citrixadc_systemglobal_authenticationpolicy_binding" "tf_bind" {
  policyname = "tf_authenticationpolicy"
}

output "priority" {
  value = data.citrixadc_systemglobal_authenticationpolicy_binding.tf_bind.priority
}

output "feature" {
  value = data.citrixadc_systemglobal_authenticationpolicy_binding.tf_bind.feature
}
```

## Argument Reference

* `policyname` - (Required) The name of the command policy.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `feature` - The feature to be checked while applying this config.
* `globalbindtype` - The global bind type for the binding.
* `gotopriorityexpression` - Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE. Applicable only for advanced authentication policies.
* `id` - The id of the systemglobal_authenticationpolicy_binding. It has the same value as the `policyname` attribute.
* `nextfactor` - On success invoke label. Applicable for advanced authentication policy binding.
* `priority` - The priority of the command policy.
