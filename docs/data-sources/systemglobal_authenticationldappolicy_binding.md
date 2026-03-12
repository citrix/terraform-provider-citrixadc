---
subcategory: "System"
---

# Data Source: systemglobal_authenticationldappolicy_binding

The `citrixadc_systemglobal_authenticationldappolicy_binding` data source allows you to retrieve information about a specific binding between the system global configuration and an authentication LDAP policy.

## Example usage

```terraform
data "citrixadc_systemglobal_authenticationldappolicy_binding" "tf_bind" {
  policyname = "tf_authenticationldappolicy"
}

output "priority" {
  value = data.citrixadc_systemglobal_authenticationldappolicy_binding.tf_bind.priority
}

output "feature" {
  value = data.citrixadc_systemglobal_authenticationldappolicy_binding.tf_bind.feature
}
```

## Argument Reference

* `policyname` - (Required) The name of the command policy.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `feature` - The feature to be checked while applying this config.
* `globalbindtype` - The global bind type for the binding.
* `gotopriorityexpression` - Applicable only to advance authentication policy. Expression or other value specifying the next policy to be evaluated if the current policy evaluates to TRUE. Specify one of the following values: NEXT - Evaluate the policy with the next higher priority number. END - End policy evaluation.
* `id` - The id of the systemglobal_authenticationldappolicy_binding. It has the same value as the `policyname` attribute.
* `nextfactor` - On success invoke label. Applicable for advanced authentication policy binding.
* `priority` - The priority of the command policy.
