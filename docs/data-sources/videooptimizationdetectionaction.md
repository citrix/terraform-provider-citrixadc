---
subcategory: "VideoOptimization"
---

# Data Source: videooptimizationdetectionaction

The videooptimizationdetectionaction data source allows you to retrieve information about a video optimization detection action.

## Example usage

```terraform
data "citrixadc_videooptimizationdetectionaction" "tf_detectionaction" {
  name = "tf_videooptimizationdetectionaction"
}

output "type" {
  value = data.citrixadc_videooptimizationdetectionaction.tf_detectionaction.type
}

output "comment" {
  value = data.citrixadc_videooptimizationdetectionaction.tf_detectionaction.comment
}
```

## Argument Reference

* `name` - (Required) Name for the video optimization detection action. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the videooptimizationdetectionaction. It has the same value as the `name` attribute.
* `comment` - Comment. Any type of information about this video optimization detection action.
* `newname` - New name for the videooptimization detection action. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.
* `type` - Type of video optimization action. Available settings function as follows:
  * clear_text_pd - Cleartext PD type is detected.
  * clear_text_abr - Cleartext ABR is detected.
  * encrypted_abr - Encrypted ABR is detected.
  * trigger_enc_abr - Possible encrypted ABR is detected.
  * trigger_body_detection - Possible cleartext ABR is detected. Triggers body content detection.

### Read-only videooptimizationdetectionaction metadata

These attributes are returned by the appliance on a GET (they are not configurable on the `citrixadc_videooptimizationdetectionaction` resource). They are Computed/GET-only, and any attribute the appliance does not return is `null`.

* `hits` - The number of times the action has been taken.
* `referencecount` - The number of references to the action.
* `undefhits` - The number of times the action resulted in UNDEF.
* `builtin` - Flag to determine whether the video optimization detection action is built-in or not. A list of strings.
* `feature` - The feature to be checked while applying this config.
