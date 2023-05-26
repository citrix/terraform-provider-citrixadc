---
subcategory: "VideoOptimization"
---

# Resource: videooptimizationpacingpolicy

The videooptimizationpacingpolicy resource is used to Configure videooptimization pacingpolicy resource.


## Example usage

```hcl
resource "citrixadc_videooptimizationpacingaction" "tf_pacingaction" {
  name = "tf_pacingaction"
  rate = 10
}

resource "citrixadc_videooptimizationpacingpolicy" "tf_pacingpolicy" {
  name   = "tf_pacingpolicy"
  rule   = "true"
  action = citrixadc_videooptimizationpacingaction.tf_pacingaction.name
}
```


## Argument Reference

* `name` - (Required) Name for the videooptimization pacing policy. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters.Can be modified, removed or renamed.
* `rule` - (Required) Expression that determines which request or response match the video optimization pacing policy. The following requirements apply only to the Citrix ADC CLI: * If the expression includes one or more spaces, enclose the entire expression in double quotation marks. * If the expression itself includes double quotation marks, escape the quotations by using the \ character. * Alternatively, you can use single quotation marks to enclose the rule, in which case you do not have to escape the double quotation marks.
* `action` - (Required) Name of the videooptimization pacing action to perform if the request matches this videooptimization pacing policy.
* `undefaction` - (Optional) Action to perform if the result of policy evaluation is undefined (UNDEF). An UNDEF event indicates an internal error condition. Only the above built-in actions can be used.
* `comment` - (Optional) Any type of information about this videooptimization pacing policy.
* `logaction` - (Optional) Name of the messagelog action to use for requests that match this policy.
* `newname` - (Optional) New name for the videooptimization pacing policy. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Minimum length =  1


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the videooptimizationpacingpolicy. It has the same value as the `name` attribute.


## Import

A videooptimizationpacingpolicy can be imported using its name, e.g.

```shell
terraform import citrixadc_videooptimizationpacingpolicy.tf_pacingpolicy tf_pacingpolicy
```
