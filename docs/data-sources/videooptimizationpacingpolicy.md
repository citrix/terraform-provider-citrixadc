---
subcategory: "VideoOptimization"
---

# Data Source: videooptimizationpacingpolicy

The videooptimizationpacingpolicy data source allows you to retrieve information about video optimization pacing policies.

## Example usage

```terraform
data "citrixadc_videooptimizationpacingpolicy" "tf_policy" {
  name = "tf_policy"
}

output "action" {
  value = data.citrixadc_videooptimizationpacingpolicy.tf_policy.action
}

output "rule" {
  value = data.citrixadc_videooptimizationpacingpolicy.tf_policy.rule
}
```

## Argument Reference

* `name` - (Required) Name for the videooptimization pacing policy. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters.Can be modified, removed or renamed.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `action` - Name of the videooptimization pacing action to perform if the request matches this videooptimization pacing policy.
* `comment` - Any type of information about this videooptimization pacing policy.
* `id` - The id of the videooptimizationpacingpolicy. It has the same value as the `name` attribute.
* `logaction` - Name of the messagelog action to use for requests that match this policy.
* `newname` - New name for the videooptimization pacing policy. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.
* `rule` - Expression that determines which request or response match the video optimization pacing policy. The following requirements apply only to the Citrix ADC CLI: If the expression includes one or more spaces, enclose the entire expression in double quotation marks. If the expression itself includes double quotation marks, escape the quotations by using the \ character. Alternatively, you can use single quotation marks to enclose the rule, in which case you do not have to escape the double quotation marks.
* `undefaction` - Action to perform if the result of policy evaluation is undefined (UNDEF). An UNDEF event indicates an internal error condition. Only the above built-in actions can be used.

## Import

A videooptimizationpacingpolicy can be imported using its name, e.g.

```shell
terraform import citrixadc_videooptimizationpacingpolicy.tf_policy tf_policy
```
