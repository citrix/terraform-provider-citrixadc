---
subcategory: "VideoOptimization"
---

# Data Source: videooptimizationglobaldetection_videooptimizationdetectionpolicy_binding

The videooptimizationglobaldetection_videooptimizationdetectionpolicy_binding data source allows you to retrieve information about a video optimization detection policy bound to the global detection bind point on the Citrix ADC.


## Example usage

```terraform
data "citrixadc_videooptimizationglobaldetection_videooptimizationdetectionpolicy_binding" "tf_binding" {
  policyname = "tf_detectionpolicy"
  priority   = 100
}

output "gotopriorityexpression" {
  value = data.citrixadc_videooptimizationglobaldetection_videooptimizationdetectionpolicy_binding.tf_binding.gotopriorityexpression
}
```


## Argument Reference

* `policyname` - (Required) Name of the globally bound video optimization detection policy to look up.
* `priority` - (Required) Priority of the policy at the global bind point.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the videooptimizationglobaldetection_videooptimizationdetectionpolicy_binding. It is a composite of the unique identifying attributes expressed as comma-separated `key:value` pairs in the form `policyname:<policyname>,priority:<priority>,type:<type>`.
* `type` - The global bind point at which the policy is evaluated.
* `globalbindtype` - The global bind point type assigned by the Citrix ADC (defaults to `SYSTEM_GLOBAL`).
* `gotopriorityexpression` - Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `invoke` - If the current policy evaluates to TRUE, terminate evaluation of policies bound to the current policy label, and then forward the request to the specified virtual server or evaluate the specified policy label.
* `labeltype` - Type of invocation. Possible values: [ vserver, policylabel ]
* `labelname` - Name of the policy label to invoke when `invoke` is set and `labeltype` is `policylabel`.
