---
subcategory: "Authentication"
---

# Data Source `authenticationnegotiatepolicy`

The authenticationnegotiatepolicy data source allows you to retrieve information about authentication negotiate policies.


## Example usage

```terraform
data "citrixadc_authenticationnegotiatepolicy" "tf_negotiatepolicy" {
  name = "my_negotiatepolicy"
}

output "rule" {
  value = data.citrixadc_authenticationnegotiatepolicy.tf_negotiatepolicy.rule
}

output "reqaction" {
  value = data.citrixadc_authenticationnegotiatepolicy.tf_negotiatepolicy.reqaction
}
```


## Argument Reference

* `name` - (Required) Name for the negotiate authentication policy. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after AD KCD (negotiate) policy is created.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `reqaction` - Name of the negotiate action to perform if the policy matches.
* `rule` - Name of the Citrix ADC named rule, or an expression, that the policy uses to determine whether to attempt to authenticate the user with the AD KCD server.

## Attribute Reference

* `id` - The id of the authenticationnegotiatepolicy. It has the same value as the `name` attribute.


## Import

A authenticationnegotiatepolicy can be imported using its name, e.g.

```shell
terraform import citrixadc_authenticationnegotiatepolicy.tf_negotiatepolicy my_negotiatepolicy
```
