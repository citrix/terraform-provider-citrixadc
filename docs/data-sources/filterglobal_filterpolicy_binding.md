---
subcategory: "Filter"
---

# Data Source: filterglobal_filterpolicy_binding

The citrixadc_filterglobal_filterpolicy_binding data source allows you to retrieve information about a filter policy binding to the global configuration.


## Example usage

```terraform
data "citrixadc_filterglobal_filterpolicy_binding" "tf_filterglobal" {
  policyname = "tf_filterpolicy"
}

output "policyname" {
  value = data.citrixadc_filterglobal_filterpolicy_binding.tf_filterglobal.policyname
}

output "priority" {
  value = data.citrixadc_filterglobal_filterpolicy_binding.tf_filterglobal.priority
}
```


## Argument Reference

* `policyname` - (Required) Name of the filter policy.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the filterglobal_filterpolicy_binding. It has the format `policyname:<policyname>`.
* `priority` - Specifies the priority of the policy.
* `state` - State of the binding. Possible values: [ ENABLED, DISABLED ]
