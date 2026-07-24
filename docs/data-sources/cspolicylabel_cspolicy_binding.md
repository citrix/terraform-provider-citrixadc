---
subcategory: "Content Switching"
---

# Data Source: cspolicylabel\_cspolicy\_binding

The cspolicylabel_cspolicy_binding data source allows you to retrieve information about a content switching policy bound to a content switching policy label.


## Example usage

```terraform
data "citrixadc_cspolicylabel_cspolicy_binding" "tf_binding" {
  labelname  = "tf_cspolicylabel"
  policyname = "tf_cspolicy"
}

output "binding_priority" {
  value = data.citrixadc_cspolicylabel_cspolicy_binding.tf_binding.priority
}
```


## Argument Reference

* `labelname` - (Required) Name of the policy label to which the content switching policy is bound.
* `policyname` - (Required) Name of the content switching policy.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The ID of the cspolicylabel\_cspolicy\_binding resource, in the form `labelname:<labelname>,policyname:<policyname>`.
* `priority` - Priority of the policy within the label.
* `targetvserver` - Name of the virtual server to which requests that match the policy are forwarded.
* `gotopriorityexpression` - Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `invoke` - Whether a policy label is invoked if the current policy rule evaluates to TRUE.
* `labeltype` - Type of policy label invocation.
* `invoke_labelname` - Name of the label to invoke if the current policy rule evaluates to TRUE.
