---
subcategory: "Integrated Caching"
---

# Data Source: cacheglobal_cachepolicy_binding

The cacheglobal_cachepolicy_binding data source allows you to retrieve information about a specific cachepolicy binding to cacheglobal configuration.

## Example Usage

```terraform
data "citrixadc_cacheglobal_cachepolicy_binding" "tf_cacheglobal_cachepolicy_binding" {
  policy   = "my_cachepolicy"
  type     = "REQ_DEFAULT"
}

output "policy" {
  value = data.citrixadc_cacheglobal_cachepolicy_binding.tf_cacheglobal_cachepolicy_binding.policy
}

output "priority" {
  value = data.citrixadc_cacheglobal_cachepolicy_binding.tf_cacheglobal_cachepolicy_binding.priority
}

output "type" {
  value = data.citrixadc_cacheglobal_cachepolicy_binding.tf_cacheglobal_cachepolicy_binding.type
}
```

## Argument Reference

* `policy` - (Required) Name of the cache policy.
* `type` - (Required) The bind point to which policy is bound.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `globalbindtype` - Global bind type.
* `gotopriorityexpression` - Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `id` - The id of the cacheglobal_cachepolicy_binding. It is a system-generated identifier.
* `priority` - Specifies the priority of the policy.
* `invoke` - Invoke policies bound to a virtual server or a user-defined policy label. After the invoked policies are evaluated, the flow returns to the policy with the next priority. Applicable only to default-syntax policies.
* `labelname` - Name of the label to invoke if the current policy rule evaluates to TRUE. (To invoke a label associated with a virtual server, specify the name of the virtual server.)
* `labeltype` - Type of policy label to invoke.
* `precededefrules` - Specify whether this policy should be evaluated.
