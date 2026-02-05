---
subcategory: "Authentication"
---

# Data Source `authenticationwebauthpolicy`

The authenticationwebauthpolicy data source allows you to retrieve information about an existing WebAuth authentication policy.


## Example usage

```terraform
data "citrixadc_authenticationwebauthpolicy" "tf_webauthpolicy" {
  name = "my_webauthpolicy"
}

output "rule" {
  value = data.citrixadc_authenticationwebauthpolicy.tf_webauthpolicy.rule
}

output "action" {
  value = data.citrixadc_authenticationwebauthpolicy.tf_webauthpolicy.action
}
```


## Argument Reference

* `name` - (Required) Name for the WebAuth policy. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after LDAP policy is created.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the authenticationwebauthpolicy. It has the same value as the `name` attribute.
* `action` - Name of the WebAuth action to perform if the policy matches.
* `rule` - Name of the Citrix ADC named rule, or an expression, that the policy uses to determine whether to attempt to authenticate the user with the Web server.


## Import

A authenticationwebauthpolicy can be imported using its name, e.g.

```shell
terraform import citrixadc_authenticationwebauthpolicy.tf_webauthpolicy my_webauthpolicy
```
