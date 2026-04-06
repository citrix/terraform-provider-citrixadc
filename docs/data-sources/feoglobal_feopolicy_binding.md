---
subcategory: "Front-end-optimization"
---

# Data Source: feoglobal_feopolicy_binding

The feoglobal_feopolicy_binding data source allows you to retrieve information about a specific policy binding to the global front end optimization.

## Example Usage

```terraform
data "citrixadc_feoglobal_feopolicy_binding" "tf_feoglobal_feopolicy_binding" {
  policyname = "tf_feopolicy"
  type       = "REQ_DEFAULT"
}

output "policyname" {
  value = data.citrixadc_feoglobal_feopolicy_binding.tf_feoglobal_feopolicy_binding.policyname
}

output "priority" {
  value = data.citrixadc_feoglobal_feopolicy_binding.tf_feoglobal_feopolicy_binding.priority
}

output "type" {
  value = data.citrixadc_feoglobal_feopolicy_binding.tf_feoglobal_feopolicy_binding.type
}
```

## Argument Reference

* `policyname` - (Required) The name of the globally bound front end optimization policy.
* `type` - (Required) Bindpoint to which the policy is bound.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `globalbindtype` - Global bind type.
* `priority` - Specifies the priority of the policy.
* `gotopriorityexpression` - Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `id` - The id of the feoglobal_feopolicy_binding. It is a system-generated identifier.
