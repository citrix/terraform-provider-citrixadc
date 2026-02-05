---
subcategory: "ICA"
---

# Data Source `icapolicy`

The icapolicy data source allows you to retrieve information about ICA policies.


## Example usage

```terraform
data "citrixadc_icapolicy" "tf_icapolicy" {
  name = "my_ica_policy"
}

output "rule" {
  value = data.citrixadc_icapolicy.tf_icapolicy.rule
}

output "action" {
  value = data.citrixadc_icapolicy.tf_icapolicy.action
}
```


## Argument Reference

* `name` - (Required) Name for the policy. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `action` - Name of the ica action to be associated with this policy.
* `comment` - Any type of information about this ICA policy.
* `logaction` - Name of the messagelog action to use for requests that match this policy.
* `newname` - New name for the policy. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.
* `rule` - Expression or other value against which the traffic is evaluated. Must be a Boolean expression.

## Attribute Reference

* `id` - The id of the icapolicy. It has the same value as the `name` attribute.


## Import

A icapolicy can be imported using its name, e.g.

```shell
terraform import citrixadc_icapolicy.tf_icapolicy my_ica_policy
```
