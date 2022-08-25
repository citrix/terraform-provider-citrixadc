---
subcategory: "Ica"
---

# Resource: icapolicy

The icapolicy resource is used to create icapolicy.


## Example usage

```hcl
resource "citrixadc_icapolicy" "tf_icapolicy" {
  name   = "my_ica_policy"
  rule   = true
  action = "my_ica_action"
}
```


## Argument Reference

* `name` - (Required) Name for the policy. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my ica policy" or 'my ica policy').
* `rule` - (Required) Expression or other value against which the traffic is evaluated. Must be a Boolean expression. The following requirements apply only to the Citrix ADC CLI: * If the expression includes one or more spaces, enclose the entire expression in double quotation marks. * If the expression itself includes double quotation marks, escape the quotations by using the \ character. * Alternatively, you can use single quotation marks to enclose the rule, in which case you do not have to escape the double quotation marks.
* `action` - (Required) Name of the ica action to be associated with this policy.
* `comment` - (Optional) Any type of information about this ICA policy.
* `logaction` - (Optional) Name of the messagelog action to use for requests that match this policy.
* `newname` - (Optional) New name for the policy. Must begin with an ASCII alphabetic or underscore (_)character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), s pace, colon (:), at (@), equals (=), and hyphen (-) characters. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my ica policy" or 'my ica policy'). Minimum length =  1


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the icapolicy. It has the same value as the `name` attribute.


## Import

A icapolicy can be imported using its name, e.g.

```shell
terraform import citrixadc_icapolicy.tf_icapolicy my_ica_policy
```
