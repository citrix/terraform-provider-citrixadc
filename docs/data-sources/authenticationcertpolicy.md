---
subcategory: "Authentication"
---

# Data Source `authenticationcertpolicy`

The authenticationcertpolicy data source allows you to retrieve information about authentication certificate policies.


## Example usage

```terraform
data "citrixadc_authenticationcertpolicy" "tf_certpolicy" {
  name = "example_certpolicy"
}

output "rule" {
  value = data.citrixadc_authenticationcertpolicy.tf_certpolicy.rule
}

output "reqaction" {
  value = data.citrixadc_authenticationcertpolicy.tf_certpolicy.reqaction
}
```


## Argument Reference

* `name` - (Required) Name for the client certificate authentication policy. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after cert authentication policy is created.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `reqaction` - Name of the client cert authentication action to be performed if the policy matches.
* `rule` - Name of the Citrix ADC named rule, or an expression, that the policy uses to determine whether to attempt to authenticate the user with the authentication server.
* `id` - The id of the authenticationcertpolicy. It has the same value as the `name` attribute.


## Import

A authenticationcertpolicy can be imported using its name, e.g.

```shell
terraform import citrixadc_authenticationcertpolicy.tf_certpolicy example_certpolicy
```
