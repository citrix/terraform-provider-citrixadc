---
subcategory: "Traffic Management"
---

# Data Source: tmglobal_tmtrafficpolicy_binding

The `citrixadc_tmglobal_tmtrafficpolicy_binding` data source allows you to retrieve information about a specific binding between the global traffic management configuration and a traffic policy. This binding determines which traffic policies are applied globally and their priority order.

## Example usage

```terraform
data "citrixadc_tmglobal_tmtrafficpolicy_binding" "tf_tmglobal_tmtrafficpolicy_binding" {
  policyname = "my_tmtrafficpolicy"
}

output "policy_name" {
  value = data.citrixadc_tmglobal_tmtrafficpolicy_binding.tf_tmglobal_tmtrafficpolicy_binding.policyname
}

output "policy_priority" {
  value = data.citrixadc_tmglobal_tmtrafficpolicy_binding.tf_tmglobal_tmtrafficpolicy_binding.priority
}
```

## Argument Reference

* `policyname` - (Required) The name of the policy.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the tmglobal_tmtrafficpolicy_binding. It has the same value as the `policyname` attribute.
* `globalbindtype` - The global bind point to which the policy is bound.
* `gotopriorityexpression` - Applicable only to advance tmsession policy. Expression or other value specifying the next policy to be evaluated if the current policy evaluates to TRUE. Specify one of the following values: NEXT - Evaluate the policy with the next higher priority number. END - End policy evaluation. An expression that evaluates to a number.
* `priority` - The priority of the policy.
* `type` - Bind point to which the policy is bound.
