---
subcategory: "VideoOptimization"
---

# Data Source: videooptimizationdetectionpolicylabel

The videooptimizationdetectionpolicylabel data source allows you to retrieve information about an existing video optimization detection policy label configured on the Citrix ADC.


## Example usage

```terraform
data "citrixadc_videooptimizationdetectionpolicylabel" "tf_detectionpolicylabel" {
  labelname = "tf_detectionpolicylabel"
}

output "policylabeltype" {
  value = data.citrixadc_videooptimizationdetectionpolicylabel.tf_detectionpolicylabel.policylabeltype
}
```


## Argument Reference

* `labelname` - (Required) Name of the video optimization detection policy label to look up.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the videooptimizationdetectionpolicylabel. It has the same value as the `labelname` attribute.
* `policylabeltype` - Type of policies that the policy label can contain, which determines the traffic direction the bound detection policies evaluate. Possible values: [ videoopt_req, videoopt_res ]
* `comment` - Any comments to preserve information about this video optimization detection policy label.
* `newname` - New name for the video optimization detection policy label (rename-only attribute).
