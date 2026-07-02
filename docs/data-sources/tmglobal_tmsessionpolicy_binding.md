---
subcategory: "Traffic Management"
---

# Data Source: tmglobal_tmsessionpolicy_binding

The tmglobal_tmsessionpolicy_binding data source allows you to retrieve information about a TM session policy bound to the global TM bind point, identified by its policy name.


## Example Usage

```terraform
data "citrixadc_tmglobal_tmsessionpolicy_binding" "tf_tmglobal_tmsessionpolicy_binding" {
  policyname = "tf_tmsessionpolicy"
}

output "priority" {
  value = data.citrixadc_tmglobal_tmsessionpolicy_binding.tf_tmglobal_tmsessionpolicy_binding.priority
}

output "feature" {
  value = data.citrixadc_tmglobal_tmsessionpolicy_binding.tf_tmglobal_tmsessionpolicy_binding.feature
}
```


## Argument Reference

* `policyname` - (Required) The name of the TM session policy bound to TM global.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the tmglobal_tmsessionpolicy_binding. It has the same value as the `policyname` attribute.
* `priority` - The priority of the policy.
* `gotopriorityexpression` - Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `feature` - The feature to be checked while applying this config.
