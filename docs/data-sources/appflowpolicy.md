---
subcategory: "AppFlow"
---

# Data Source `appflowpolicy`

The appflowpolicy data source allows you to retrieve information about an existing appflowpolicy.


## Example usage

```terraform
data "citrixadc_appflowpolicy" "tf_appflowpolicy" {
  name = "test_policy"
}

output "name" {
  value = data.citrixadc_appflowpolicy.tf_appflowpolicy.name
}

output "action" {
  value = data.citrixadc_appflowpolicy.tf_appflowpolicy.action
}

output "rule" {
  value = data.citrixadc_appflowpolicy.tf_appflowpolicy.rule
}
```


## Argument Reference

* `name` - (Required) Name for the policy. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.

The following requirement applies only to the Citrix ADC CLI:
If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my appflow policy" or 'my appflow policy').

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the appflowpolicy. It has the same value as the `name` attribute.
* `action` - Name of the action to be associated with this policy.
* `comment` - Any comments about this policy.
* `newname` - New name for the policy. Must begin with an ASCII alphabetic or underscore (_)character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.
* `rule` - Expression or other value against which the traffic is evaluated. Must be a Boolean expression.
* `undefaction` - Name of the appflow action to be associated with this policy when an undef event occurs.
