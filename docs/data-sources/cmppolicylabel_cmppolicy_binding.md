---
subcategory: "Compression"
---

# Data Source: cmppolicylabel_cmppolicy_binding

The cmppolicylabel_cmppolicy_binding data source allows you to retrieve information about a specific cmppolicy binding to cmppolicylabel configuration.

## Example Usage

```terraform
data "citrixadc_cmppolicylabel_cmppolicy_binding" "tf_cmppolicylabel_cmppolicy_binding" {
  labelname  = "tf_cmppolicylabel_ds"
  policyname = "tf_cmppolicy_ds"
}

output "labelname" {
  value = data.citrixadc_cmppolicylabel_cmppolicy_binding.tf_cmppolicylabel_cmppolicy_binding.labelname
}

output "policyname" {
  value = data.citrixadc_cmppolicylabel_cmppolicy_binding.tf_cmppolicylabel_cmppolicy_binding.policyname
}

output "priority" {
  value = data.citrixadc_cmppolicylabel_cmppolicy_binding.tf_cmppolicylabel_cmppolicy_binding.priority
}
```

## Argument Reference

* `labelname` - (Required) Name of the HTTP compression policy label to which to bind the policy.
* `policyname` - (Required) The compression policy name.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `gotopriorityexpression` - Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `priority` - Specifies the priority of the policy.
* `id` - The id of the cmppolicylabel_cmppolicy_binding. It is a system-generated identifier.
* `invoke` - Invoke policies bound to a virtual server or a user-defined policy label. After the invoked policies are evaluated, the flow returns to the policy with the next higher priority number in the original label.
* `invoke_labelname` - Name of the label to invoke if the current policy evaluates to TRUE.
* `labeltype` - Type of policy label invocation.
