---
subcategory: "AAA"
---

# Data Source `aaapreauthenticationpolicy`

The aaapreauthenticationpolicy data source allows you to retrieve information about AAA preauthentication policies.


## Example usage

```terraform
data "citrixadc_aaapreauthenticationpolicy" "tf_aaapreauthenticationpolicy" {
  name = "my_policy"
}

output "reqaction" {
  value = data.citrixadc_aaapreauthenticationpolicy.tf_aaapreauthenticationpolicy.reqaction
}

output "rule" {
  value = data.citrixadc_aaapreauthenticationpolicy.tf_aaapreauthenticationpolicy.rule
}
```


## Argument Reference

* `name` - (Required) Name for the preauthentication policy.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `reqaction` - Name of the action that the policy is to invoke when a connection matches the policy.
* `rule` - Name of the Citrix ADC named rule, or an expression, defining connections that match the policy.

## Attribute Reference

* `id` - The id of the aaapreauthenticationpolicy. It has the same value as the `name` attribute.


## Import

A aaapreauthenticationpolicy can be imported using its name, e.g.

```shell
terraform import citrixadc_aaapreauthenticationpolicy.tf_aaapreauthenticationpolicy my_policy
```
