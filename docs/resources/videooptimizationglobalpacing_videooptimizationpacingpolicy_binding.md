---
subcategory: "VideoOptimization"
---

# Resource: videooptimizationglobalpacing_videooptimizationpacingpolicy_binding

This resource is used to bind a video optimization pacing policy to the global pacing bind point.


## Example usage

```hcl
resource "citrixadc_videooptimizationglobalpacing_videooptimizationpacingpolicy_binding" "tf_binding" {
  policyname = "video_pacing_pol1"
  priority   = 100
  type       = "REQ_DEFAULT"
}
```


## Argument Reference

* `policyname` - (Required) Name of the videooptimization pacing policy.
* `priority` - (Required) Specifies the priority of the policy.
* `type` - (Optional) Specifies the bind point whose policies you want to display. Possible values: [ REQ_OVERRIDE, REQ_DEFAULT, RES_OVERRIDE, RES_DEFAULT ]
* `gotopriorityexpression` - (Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `invoke` - (Optional) If the current policy evaluates to TRUE, terminate evaluation of policies bound to the current policy label, and then forward the request to the specified virtual server or evaluate the specified policy label.
* `labeltype` - (Optional) Type of invocation. Available settings function as follows: `vserver` - Forward the request to the specified virtual server; `policylabel` - Invoke the specified policy label.
* `labelname` - (Optional) Name of the policy label to invoke. Applies when the current policy evaluates to TRUE, the `invoke` parameter is set, and `labeltype` is `policylabel`.

~> **Note** All arguments are immutable. Changing any of them forces the binding to be destroyed and recreated.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The ID of the videooptimizationglobalpacing_videooptimizationpacingpolicy_binding. It is a composite of comma-separated `key:value` pairs built from the unique attributes, in the form `policyname:<policyname>,priority:<priority>,type:<type>` (values are URL-encoded).
* `globalbindtype` - The global bind point type for this binding. This is a read-only value assigned by the appliance (defaults to `SYSTEM_GLOBAL`).


## Import

A videooptimizationglobalpacing_videooptimizationpacingpolicy_binding can be imported using its ID, in the `policyname:<policyname>,priority:<priority>,type:<type>` format, e.g.

```shell
terraform import citrixadc_videooptimizationglobalpacing_videooptimizationpacingpolicy_binding.tf_binding "policyname:video_pacing_pol1,priority:100,type:REQ_DEFAULT"
```
