---
subcategory: "VideoOptimization"
---

# Data Source: videooptimizationdetectionpolicylabel_videooptimizationdetectionpolicy_binding

The videooptimizationdetectionpolicylabel_videooptimizationdetectionpolicy_binding data source allows you to retrieve information about a video optimization detection policy bound to a video optimization detection policy label on the Citrix ADC.


## Example usage

```terraform
data "citrixadc_videooptimizationdetectionpolicylabel_videooptimizationdetectionpolicy_binding" "tf_binding" {
  labelname  = "tf_detectionpolicylabel"
  policyname = "tf_detectionpolicy"
  priority   = 100
}

output "gotopriorityexpression" {
  value = data.citrixadc_videooptimizationdetectionpolicylabel_videooptimizationdetectionpolicy_binding.tf_binding.gotopriorityexpression
}
```


## Argument Reference

* `labelname` - (Required) Name of the video optimization detection policy label whose binding to look up.
* `policyname` - (Required) Name of the bound video optimization detection policy.
* `priority` - (Required) Priority of the policy within the policy label.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the videooptimizationdetectionpolicylabel_videooptimizationdetectionpolicy_binding. It is a composite of the unique identifying attributes expressed as comma-separated `key:value` pairs in the form `labelname:<labelname>,policyname:<policyname>,priority:<priority>`.
* `gotopriorityexpression` - Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `invoke` - If the current policy evaluates to TRUE, terminate evaluation of policies bound to the current policy label and evaluate the specified policy label.
* `labeltype` - Type of policy label to invoke. Possible values: [ vserver, policylabel ]
* `invoke_labelname` - Name of the label to invoke. If `labeltype` is `policylabel`, the name of the policy label to invoke; if `labeltype` is `vserver`, the name of the virtual server.
