---
subcategory: "VideoOptimization"
---

# Data Source: videooptimizationglobalpacing_videooptimizationpacingpolicy_binding

The videooptimizationglobalpacing_videooptimizationpacingpolicy_binding data source allows you to retrieve information about a video optimization pacing policy bound to the global pacing bind point.

~> **Note:** The video pacing feature is deprecated in current NetScaler releases; retained for backward compatibility.


## Example usage

```terraform
data "citrixadc_videooptimizationglobalpacing_videooptimizationpacingpolicy_binding" "example" {
  policyname = "video_pacing_pol1"
  priority   = 100
}

output "binding_type" {
  value = data.citrixadc_videooptimizationglobalpacing_videooptimizationpacingpolicy_binding.example.type
}
```


## Argument Reference

* `policyname` - (Required) Name of the videooptimization pacing policy.
* `priority` - (Required) Specifies the priority of the policy.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The composite ID of the binding, in the `policyname:<policyname>,priority:<priority>,type:<type>` format.
* `type` - Specifies the bind point whose policies you want to display.
* `globalbindtype` - The global bind point type for this binding (defaults to `SYSTEM_GLOBAL`).
* `gotopriorityexpression` - Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `invoke` - Whether, if the current policy evaluates to TRUE, evaluation of policies bound to the current policy label is terminated and the request is forwarded to the specified virtual server or policy label.
* `labeltype` - Type of invocation. Available settings function as follows: `vserver` - Forward the request to the specified virtual server; `policylabel` - Invoke the specified policy label.
* `labelname` - Name of the policy label to invoke when `invoke` is set and `labeltype` is `policylabel`.
