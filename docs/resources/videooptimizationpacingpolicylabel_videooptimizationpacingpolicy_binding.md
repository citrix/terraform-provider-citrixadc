---
subcategory: "VideoOptimization"
---

# Resource: videooptimizationpacingpolicylabel_videooptimizationpacingpolicy_binding

Binds a video optimization pacing policy to a video optimization pacing policy label on the Citrix ADC. Use this resource to add a pacing policy to a reusable policy label at a specific evaluation priority, and to control the flow of evaluation (next-priority expression and label invocation) when that policy's rule matches request or response traffic.

~> **Note** The Citrix ADC CLI marks the video optimization pacing feature as deprecated. This binding resource is retained for backward compatibility.


## Example usage

```hcl
resource "citrixadc_videooptimizationpacingpolicylabel_videooptimizationpacingpolicy_binding" "tf_binding" {
  labelname              = "tf_pacingpolicylabel"
  policyname             = "tf_pacingpolicy"
  priority               = 100
  gotopriorityexpression = "NEXT"
}
```


## Argument Reference

* `labelname` - (Required) Name of the video optimization pacing policy label to which to bind the policy. Changing this attribute forces a new resource to be created.
* `policyname` - (Required) Name of the video optimization pacing policy to bind to the policy label. Changing this attribute forces a new resource to be created.
* `priority` - (Required) Specifies the priority of the policy within the policy label. Lower-numbered policies are evaluated first. Changing this attribute forces a new resource to be created.
* `gotopriorityexpression` - (Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE. Changing this attribute forces a new resource to be created.
* `invoke` - (Optional) If the current policy evaluates to TRUE, terminate evaluation of policies bound to the current policy label and evaluate the specified policy label. Changing this attribute forces a new resource to be created.
* `labeltype` - (Optional) Type of policy label to invoke. Changing this attribute forces a new resource to be created. Possible values: [ vserver, policylabel ]
  * `vserver` - Invoke an unnamed policy label associated with a virtual server.
  * `policylabel` - Invoke a user-defined policy label.
* `invoke_labelname` - (Optional) Name of the label to invoke if the current policy evaluates to TRUE. If `labeltype` is `policylabel`, this is the name of the policy label to invoke; if `labeltype` is `vserver`, this is the name of the virtual server. Changing this attribute forces a new resource to be created.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the videooptimizationpacingpolicylabel_videooptimizationpacingpolicy_binding. It is a composite of the unique identifying attributes expressed as comma-separated `key:value` pairs in the form `labelname:<labelname>,policyname:<policyname>,priority:<priority>` (for example, `labelname:tf_pacingpolicylabel,policyname:tf_pacingpolicy,priority:100`).


## Import

A videooptimizationpacingpolicylabel_videooptimizationpacingpolicy_binding can be imported using its id, which is the comma-separated `key:value` composite of `labelname`, `policyname`, and `priority`, e.g.

```shell
terraform import citrixadc_videooptimizationpacingpolicylabel_videooptimizationpacingpolicy_binding.tf_binding "labelname:tf_pacingpolicylabel,policyname:tf_pacingpolicy,priority:100"
```
