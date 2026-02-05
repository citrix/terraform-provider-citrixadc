---
subcategory: "Authentication"
---

# Data Source `authenticationradiuspolicy`

The authenticationradiuspolicy data source allows you to retrieve information about an existing RADIUS authentication policy.


## Example usage

```terraform
data "citrixadc_authenticationradiuspolicy" "tf_radiuspolicy" {
  name = "my_radiuspolicy"
}

output "rule" {
  value = data.citrixadc_authenticationradiuspolicy.tf_radiuspolicy.rule
}

output "reqaction" {
  value = data.citrixadc_authenticationradiuspolicy.tf_radiuspolicy.reqaction
}
```


## Argument Reference

* `name` - (Required) Name for the RADIUS authentication policy. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after RADIUS policy is created.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the authenticationradiuspolicy. It has the same value as the `name` attribute.
* `reqaction` - Name of the RADIUS action to perform if the policy matches.
* `rule` - Name of the Citrix ADC named rule, or an expression, that the policy uses to determine whether to attempt to authenticate the user with the RADIUS server.


## Import

A authenticationradiuspolicy can be imported using its name, e.g.

```shell
terraform import citrixadc_authenticationradiuspolicy.tf_radiuspolicy my_radiuspolicy
```
