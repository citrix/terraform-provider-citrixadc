---
subcategory: "VideoOptimization"
---

# Resource: videooptimizationdetectionpolicylabel

The videooptimizationdetectionpolicylabel resource defines a named policy label on the Citrix ADC that groups a set of video optimization detection policies for invocation as a unit. Bind detection policies to the label and invoke the label from a video optimization detection policy so a common set of detection rules can be reused and evaluated together for request or response traffic.


## Example usage

```hcl
resource "citrixadc_videooptimizationdetectionpolicylabel" "tf_detectionpolicylabel" {
  labelname       = "tf_detectionpolicylabel"
  policylabeltype = "videoopt_req"
  comment         = "Reusable video optimization request detection policies"
}
```


## Argument Reference

* `labelname` - (Required) Name for the video optimization detection policy label. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.), hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the policy label is added. Changing this attribute forces a new resource to be created.
* `policylabeltype` - (Optional) Type of policies that the policy label can contain, which determines the traffic direction the bound detection policies evaluate. Defaults to `"videoopt_req"`. Changing this attribute forces a new resource to be created. Possible values: [ videoopt_req, videoopt_res ]
* `comment` - (Optional) Any comments to preserve information about this video optimization detection policy label. Changing this attribute forces a new resource to be created.
* `newname` - (Optional) New name for the video optimization detection policy label. Used only to rename an existing policy label; must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.), hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Changing this attribute forces a new resource to be created.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the videooptimizationdetectionpolicylabel. It has the same value as the `labelname` attribute.


## Import

A videooptimizationdetectionpolicylabel can be imported using its labelname, e.g.

```shell
terraform import citrixadc_videooptimizationdetectionpolicylabel.tf_detectionpolicylabel tf_detectionpolicylabel
```
