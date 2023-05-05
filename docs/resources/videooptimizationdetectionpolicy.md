---
subcategory: "VideoOptimization"
---

# Resource: videooptimizationdetectionpolicy

The videooptimizationdetectionpolicy resource is used to create videooptimization detectionpolicy resource.


## Example usage

```hcl
resource "citrixadc_videooptimizationdetectionaction" "tf_detectionaction" {
  name = "tf_videooptimizationdetectionaction"
  type = "clear_text_abr"
}

resource "citrixadc_videooptimizationdetectionpolicy" "tf_detectionpolicy" {
  name   = "tf_videooptimizationdetectionpolicy"
  rule   = "true"
  action = citrixadc_videooptimizationdetectionaction.tf_detectionaction.name
}

```


## Argument Reference

* `name` - (Required) Name for the videooptimization detection policy. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters.Can be modified, removed or renamed.
* `rule` - (Required) Expression that determines which request or response match the video optimization detection policy. The following requirements apply only to the Citrix ADC CLI: * If the expression includes one or more spaces, enclose the entire expression in double quotation marks. * If the expression itself includes double quotation marks, escape the quotations by using the \ character. * Alternatively, you can use single quotation marks to enclose the rule, in which case you do not have to escape the double quotation marks.
* `action` - (Required) Name of the videooptimization detection action to perform if the request matches this videooptimization detection policy. Built-in actions should be used. These are: * DETECT_CLEARTEXT_PD - Cleartext PD is detected and increment related counters. * DETECT_CLEARTEXT_ABR - Cleartext ABR is detected and increment related counters. * DETECT_ENCRYPTED_ABR - Encrypted ABR is detected and increment related counters. * TRIGGER_ENC_ABR_DETECTION - This is potentially encrypted ABR. Internal traffic heuristics algorithms will further process traffic to confirm detection. * TRIGGER_CT_ABR_BODY_DETECTION -  This is potentially cleartext ABR. Internal traffic heuristics algorithms will further process traffic to confirm detection. * RESET - Reset the client connection by closing it. * DROP - Drop the connection without sending a response.
* `undefaction` - (Optional) Action to perform if the result of policy evaluation is undefined (UNDEF). An UNDEF event indicates an internal error condition. Only the above built-in actions can be used.
* `comment` - (Optional) Any type of information about this videooptimization detection policy.
* `logaction` - (Optional) Name of the messagelog action to use for requests that match this policy.
* `newname` - (Optional) New name for the videooptimization detection policy. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Minimum length =  1


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the videooptimizationdetectionpolicy. It has the same value as the `name` attribute.


## Import

A videooptimizationdetectionpolicy can be imported using its name, e.g.

```shell
terraform import citrixadc_videooptimizationdetectionpolicy.tf_detectionpolicy tf_videooptimizationdetectionpolicy
```
