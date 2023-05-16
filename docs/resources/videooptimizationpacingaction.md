---
subcategory: "VideoOptimization"
---

# Resource: videooptimizationpacingaction

The videooptimizationpacingaction resource is used to Configure videooptimization pacingaction resource.


## Example usage

```hcl
resource "citrixadc_videooptimizationpacingaction" "tf_pacingaction" {
  name = "tf_pacingaction"
  rate = 10
}
```


## Argument Reference

* `name` - (Required) Name for the video optimization pacing action. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.
* `rate` - (Required) ABR Video Optimization Pacing Rate (in Kbps). Minimum value =  1 Maximum value =  2147483647
* `comment` - (Optional) Comment. Any type of information about this video optimization detection action.
* `newname` - (Optional) New name for the videooptimization pacing action. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Minimum length =  1


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the videooptimizationpacingaction. It has the same value as the `name` attribute.


## Import

A videooptimizationpacingaction can be imported using its name, e.g.

```shell
terraform import citrixadc_videooptimizationpacingaction.tf_pacingaction tf_pacingaction
```
