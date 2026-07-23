---
subcategory: "VideoOptimization"
---

# Data Source: videooptimizationpacingpolicylabel_videooptimizationpacingpolicy_binding

The videooptimizationpacingpolicylabel_videooptimizationpacingpolicy_binding data source allows you to retrieve information about a video optimization pacing policy bound to a pacing policy label.

~> **Note:** Video optimization pacing is deprecated on the Citrix ADC; retained for backward compatibility.


## Example usage

```terraform
data "citrixadc_videooptimizationpacingpolicylabel_videooptimizationpacingpolicy_binding" "tf_binding" {
  labelname  = "tf_pacingpolicylabel"
  policyname = "tf_pacingpolicy"
  priority   = 100
}

output "gotopriorityexpression" {
  value = data.citrixadc_videooptimizationpacingpolicylabel_videooptimizationpacingpolicy_binding.tf_binding.gotopriorityexpression
}
```


## Argument Reference

* `labelname` - (Required) Name of the video optimization pacing policy label whose binding to look up.
* `policyname` - (Required) Name of the bound video optimization pacing policy.
* `priority` - (Required) Priority of the policy within the policy label.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the videooptimizationpacingpolicylabel_videooptimizationpacingpolicy_binding. It is a composite of the unique identifying attributes expressed as comma-separated `key:value` pairs in the form `labelname:<labelname>,policyname:<policyname>,priority:<priority>`.
* `gotopriorityexpression` - Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `invoke` - If the current policy evaluates to TRUE, terminate evaluation of policies bound to the current policy label and evaluate the specified policy label.
* `labeltype` - Type of policy label to invoke. Possible values: [ vserver, policylabel ]
* `invoke_labelname` - Name of the label to invoke. If `labeltype` is `policylabel`, the name of the policy label to invoke; if `labeltype` is `vserver`, the name of the virtual server.
```
