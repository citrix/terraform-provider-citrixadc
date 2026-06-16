---
subcategory: "Compression"
---

# Data Source: cmpglobal_cmppolicy_binding

The cmpglobal_cmppolicy_binding data source allows you to retrieve information about a specific cmppolicy binding to cmpglobal configuration.

## Example Usage

```terraform
data "citrixadc_cmpglobal_cmppolicy_binding" "tf_cmpglobal_cmppolicy_binding" {
  policyname = "tf_cmppolicy_ds"
  type       = "RES_OVERRIDE"
}

output "policyname" {
  value = data.citrixadc_cmpglobal_cmppolicy_binding.tf_cmpglobal_cmppolicy_binding.policyname
}

output "priority" {
  value = data.citrixadc_cmpglobal_cmppolicy_binding.tf_cmpglobal_cmppolicy_binding.priority
}

output "type" {
  value = data.citrixadc_cmpglobal_cmppolicy_binding.tf_cmpglobal_cmppolicy_binding.type
}
```

## Argument Reference

* `policyname` - (Required) The name of the globally bound HTTP compression policy.
* `type` - (Required) Bind point to which the policy is bound.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `globalbindtype` - The global bind type.
* `priority` - Positive integer specifying the priority of the policy. The lower the number, the higher the priority. By default, polices within a label are evaluated in the order of their priority numbers.
* `gotopriorityexpression` - Expression or other value specifying the priority of the next policy, within the policy label, to evaluate if the current policy evaluates to TRUE.
* `id` - The id of the cmpglobal_cmppolicy_binding. It is the concatenation of the `policyname` and `type` attributes separated by a comma.
* `invoke` - Invoke policies bound to a virtual server or a policy label. After the invoked policies are evaluated, the flow returns to the policy with the next priority.
* `labelname` - Name of the label to invoke if the current policy rule evaluates to TRUE.
* `labeltype` - Type of policy label invocation.
