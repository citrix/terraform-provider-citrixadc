---
subcategory: "Authentication"
---

# Data Source `authenticationsamlpolicy`

The authenticationsamlpolicy data source allows you to retrieve information about an existing SAML authentication policy.


## Example usage

```terraform
data "citrixadc_authenticationsamlpolicy" "tf_samlpolicy" {
  name = "my_samlpolicy"
}

output "rule" {
  value = data.citrixadc_authenticationsamlpolicy.tf_samlpolicy.rule
}

output "reqaction" {
  value = data.citrixadc_authenticationsamlpolicy.tf_samlpolicy.reqaction
}
```


## Argument Reference

* `name` - (Required) Name for the SAML policy. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after SAML policy is created.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the authenticationsamlpolicy. It has the same value as the `name` attribute.
* `reqaction` - Name of the SAML authentication action to be performed if the policy matches.
* `rule` - Name of the Citrix ADC named rule, or an expression, that the policy uses to determine whether to attempt to authenticate the user with the SAML server.


## Import

A authenticationsamlpolicy can be imported using its name, e.g.

```shell
terraform import citrixadc_authenticationsamlpolicy.tf_samlpolicy my_samlpolicy
```
