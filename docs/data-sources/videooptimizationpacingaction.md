---
subcategory: "VideoOptimization"
---

# Data Source: videooptimizationpacingaction

The videooptimizationpacingaction data source allows you to retrieve information about video optimization pacing actions.

## Example usage

```terraform
data "citrixadc_videooptimizationpacingaction" "tf_pacingaction" {
  name = "tf_pacingaction"
}

output "rate" {
  value = data.citrixadc_videooptimizationpacingaction.tf_pacingaction.rate
}

output "comment" {
  value = data.citrixadc_videooptimizationpacingaction.tf_pacingaction.comment
}
```

## Argument Reference

* `name` - (Required) Name for the video optimization pacing action. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `comment` - Comment. Any type of information about this video optimization detection action.
* `id` - The id of the videooptimizationpacingaction. It has the same value as the `name` attribute.
* `newname` - New name for the videooptimization pacing action. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.
* `rate` - ABR Video Optimization Pacing Rate (in Kbps)

## Import

A videooptimizationpacingaction can be imported using its name, e.g.

```shell
terraform import citrixadc_videooptimizationpacingaction.tf_pacingaction tf_pacingaction
```
