---
subcategory: "Rewrite"
---

# Data Source: rewriteglobal_rewritepolicy_binding

The citrixadc_rewriteglobal_rewritepolicy_binding data source allows you to retrieve information about a rewrite policy binding to the global configuration.


## Example usage

```terraform
data "citrixadc_rewriteglobal_rewritepolicy_binding" "tf_rewriteglobal_rewritepolicy_binding" {
  policyname = "tf_rewrite_policy"
  type       = "REQ_DEFAULT"
}

output "policyname" {
  value = data.citrixadc_rewriteglobal_rewritepolicy_binding.tf_rewriteglobal_rewritepolicy_binding.policyname
}

output "priority" {
  value = data.citrixadc_rewriteglobal_rewritepolicy_binding.tf_rewriteglobal_rewritepolicy_binding.priority
}
```


## Argument Reference

* `policyname` - (Required) Name of the rewrite policy.
* `type` - (Required) The bindpoint to which to policy is bound.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `globalbindtype` - Global bind type.
* `priority` - Specifies the priority of the policy.
* `gotopriorityexpression` - Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `id` - The id of the rewriteglobal_rewritepolicy_binding. It is a system-generated identifier.
* `invoke` - Terminate evaluation of policies bound to the current policy label, and then forward the request to the specified virtual server or evaluate the specified policy label.
* `labelname` - * If labelType is policylabel, name of the policy label to invoke. \n* If labelType is reqvserver or resvserver, name of the virtual server to which to forward the request of response.
* `labeltype` - Type of invocation. Available settings function as follows:\n* reqvserver - Forward the request to the specified request virtual server.\n* resvserver - Forward the response to the specified response virtual server.\n* policylabel - Invoke the specified policy label.
