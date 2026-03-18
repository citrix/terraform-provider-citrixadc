---
subcategory: "Rewrite"
---

# Data Source: rewritepolicylabel_rewritepolicy_binding

The rewritepolicylabel_rewritepolicy_binding data source allows you to retrieve information about a specific binding between a rewrite policy and a rewrite policy label.


## Example Usage

```terraform
data "citrixadc_rewritepolicylabel_rewritepolicy_binding" "tf_rewritepolicylabel_rewritepolicy_binding" {
  labelname  = "tf_rewritepolicylabel"
  policyname = "tf_rewrite_policy"
}

output "gotopriorityexpression" {
  value = data.citrixadc_rewritepolicylabel_rewritepolicy_binding.tf_rewritepolicylabel_rewritepolicy_binding.gotopriorityexpression
}

output "labeltype" {
  value = data.citrixadc_rewritepolicylabel_rewritepolicy_binding.tf_rewritepolicylabel_rewritepolicy_binding.labeltype
}
```


## Argument Reference

* `labelname` - (Required) Name of the rewrite policy label to which the policy is bound.
* `policyname` - (Required) Name of the rewrite policy that is bound to the policy label.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `gotopriorityexpression` - Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `priority` - Priority of the policy binding.
* `id` - The id of the rewritepolicylabel_rewritepolicy_binding. It is a system-generated identifier.
* `invoke` - Suspend evaluation of policies bound to the current policy label, and then forward the request to the specified virtual server or evaluate the specified policy label.
* `invoke_labelname` - If labelType is policylabel, name of the policy label to invoke. If labelType is reqvserver or resvserver, name of the virtual server to which to forward the request or response.
* `labeltype` - Type of invocation. Available settings function as follows: reqvserver - Forward the request to the specified request virtual server. resvserver - Forward the response to the specified response virtual server. policylabel - Invoke the specified policy label. Possible values: [ reqvserver, resvserver, policylabel ]
