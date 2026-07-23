---
subcategory: "VideoOptimization"
---

# Resource: videooptimizationglobaldetection_videooptimizationdetectionpolicy_binding

This resource is used to bind a video optimization detection policy to the global detection bind point.


## Example usage

```hcl
resource "citrixadc_videooptimizationglobaldetection_videooptimizationdetectionpolicy_binding" "tf_binding" {
  policyname             = "tf_detectionpolicy"
  priority               = 100
  gotopriorityexpression = "END"
  type                   = "REQ_DEFAULT"
}
```


## Argument Reference

* `policyname` - (Required) Name of the video optimization detection policy to bind globally. Changing this attribute forces a new resource to be created.
* `priority` - (Required) Specifies the priority of the policy at the global bind point. Lower-numbered policies are evaluated first. Changing this attribute forces a new resource to be created.
* `type` - (Optional) Specifies the global bind point at which the policy is evaluated. Changing this attribute forces a new resource to be created.
* `gotopriorityexpression` - (Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE. Changing this attribute forces a new resource to be created.
* `invoke` - (Optional) If the current policy evaluates to TRUE, terminate evaluation of policies bound to the current policy label, and then forward the request to the specified virtual server or evaluate the specified policy label. Changing this attribute forces a new resource to be created.
* `labeltype` - (Optional) Type of invocation. Changing this attribute forces a new resource to be created. Available settings function as follows:
  * `vserver` - Forward the request to the specified virtual server.
  * `policylabel` - Invoke the specified policy label.
* `labelname` - (Optional) Name of the policy label to invoke if the current policy evaluates to TRUE, the `invoke` parameter is set, and `labeltype` is `policylabel`. Changing this attribute forces a new resource to be created.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the videooptimizationglobaldetection_videooptimizationdetectionpolicy_binding. It is a composite of the unique identifying attributes expressed as comma-separated `key:value` pairs in the form `policyname:<policyname>,priority:<priority>,type:<type>` (for example, `policyname:tf_detectionpolicy,priority:100,type:REQ_DEFAULT`).
* `globalbindtype` - The global bind point type assigned by the Citrix ADC. This value is computed by the server and defaults to `SYSTEM_GLOBAL`.


## Import

A videooptimizationglobaldetection_videooptimizationdetectionpolicy_binding can be imported using its id, which is the comma-separated `key:value` composite of `policyname`, `priority`, and `type`, e.g.

```shell
terraform import citrixadc_videooptimizationglobaldetection_videooptimizationdetectionpolicy_binding.tf_binding "policyname:tf_detectionpolicy,priority:100,type:REQ_DEFAULT"
```
