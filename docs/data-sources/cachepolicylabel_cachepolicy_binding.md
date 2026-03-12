---
subcategory: "Integrated Caching"
---

# Data Source: cachepolicylabel_cachepolicy_binding

The cachepolicylabel_cachepolicy_binding data source allows you to retrieve information about a specific binding between a cache policy label and a cache policy.

## Example Usage

```terraform
data "citrixadc_cachepolicylabel_cachepolicy_binding" "example" {
  labelname  = "my_cachepolicylabel"
  policyname = "my_cachepolicy"
}

output "gotopriorityexpression" {
  value = data.citrixadc_cachepolicylabel_cachepolicy_binding.example.gotopriorityexpression
}

output "invoke" {
  value = data.citrixadc_cachepolicylabel_cachepolicy_binding.example.invoke
}
```

## Argument Reference

* `labelname` - (Required) Name of the cache policy label to which to bind the policy.
* `policyname` - (Required) Name of the cache policy to bind to the policy label.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `gotopriorityexpression` - Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `id` - The id of the cachepolicylabel_cachepolicy_binding. It is a system-generated identifier.
* `priority` - Specifies the priority of the policy.
* `invoke` - Invoke policies bound to a virtual server or a user-defined policy label. After the invoked policies are evaluated, the flow returns to the policy with the next-lower priority.
* `invoke_labelname` - Name of the policy label to invoke if the current policy rule evaluates to TRUE.
* `labeltype` - Type of policy label to invoke: an unnamed label associated with a virtual server, or user-defined policy label.
