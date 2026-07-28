---
subcategory: "VideoOptimization"
---

# Resource: videooptimizationpacingpolicylabel

This resource is used to manage video optimization pacing policy labels on the Citrix ADC.


## Example usage

```hcl
resource "citrixadc_videooptimizationpacingpolicylabel" "tf_pacingpolicylabel" {
  labelname       = "tf_pacingpolicylabel"
  policylabeltype = "videoopt_req"
  comment         = "Reusable video optimization request pacing policies"
}
```


## Argument Reference

* `labelname` - (Required) Name for the video optimization pacing policy label. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.), hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the policy label is added. Changing this attribute forces a new resource to be created.
* `policylabeltype` - (Optional) Type of policies that the policy label can contain, which determines the traffic direction the bound pacing policies evaluate. Changing this attribute forces a new resource to be created. Possible values: [ videoopt_req, videoopt_res ]
* `comment` - (Optional) Any comments to preserve information about this video optimization pacing policy label. Changing this attribute forces a new resource to be created.
* `newname` - (Optional) New name for the video optimization pacing policy label. Used only to rename an existing policy label; must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.), hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Changing this attribute forces a new resource to be created.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the videooptimizationpacingpolicylabel. It has the same value as the `labelname` attribute.


## Import

A videooptimizationpacingpolicylabel can be imported using its labelname, e.g.

```shell
terraform import citrixadc_videooptimizationpacingpolicylabel.tf_pacingpolicylabel tf_pacingpolicylabel
```
