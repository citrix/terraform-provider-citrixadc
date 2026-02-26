---
subcategory: "Ica"
---

# Data Source: icaglobal_icapolicy_binding

The icaglobal_icapolicy_binding data source allows you to retrieve information about ICA global policy bindings.


## Example Usage

```terraform
data "citrixadc_icaglobal_icapolicy_binding" "tf_icaglobal_icapolicy_binding" {
  policyname = "tf_icapolicy"
  priority   = 100
  type       = "ICA_REQ_DEFAULT"
}

output "globalbindtype" {
  value = data.citrixadc_icaglobal_icapolicy_binding.tf_icaglobal_icapolicy_binding.globalbindtype
}

output "gotopriorityexpression" {
  value = data.citrixadc_icaglobal_icapolicy_binding.tf_icaglobal_icapolicy_binding.gotopriorityexpression
}
```


## Argument Reference

* `policyname` - (Required) Name of the ICA policy.
* `type` - (Required) Global bind point for which to show detailed information about the policies bound to the bind point.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the icaglobal_icapolicy_binding. It is a system-generated identifier.
* `globalbindtype` - Global bind type.
* `priority` - Specifies the priority of the policy.
* `gotopriorityexpression` - Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
