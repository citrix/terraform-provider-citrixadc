---
subcategory: "VideoOptimization"
---

# Data Source: videooptimizationpacingpolicylabel

The videooptimizationpacingpolicylabel data source allows you to retrieve information about an existing video optimization pacing policy label configured on the Citrix ADC.

~> **Note** Video optimization pacing functionality is deprecated on the Citrix ADC (NITRO/CLI). This data source is retained for compatibility with existing configurations.


## Example usage

```terraform
data "citrixadc_videooptimizationpacingpolicylabel" "tf_pacingpolicylabel" {
  labelname = "tf_pacingpolicylabel"
}

output "policylabeltype" {
  value = data.citrixadc_videooptimizationpacingpolicylabel.tf_pacingpolicylabel.policylabeltype
}
```


## Argument Reference

* `labelname` - (Required) Name of the video optimization pacing policy label to look up.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the videooptimizationpacingpolicylabel. It has the same value as the `labelname` attribute.
* `policylabeltype` - Type of policies that the policy label can contain, which determines the traffic direction the bound pacing policies evaluate. Possible values: [ videoopt_req, videoopt_res ]
* `comment` - Any comments to preserve information about this video optimization pacing policy label.
* `newname` - New name for the video optimization pacing policy label (rename-only attribute).
