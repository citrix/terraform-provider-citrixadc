---
subcategory: "Traffic Management"
---

# Data Source: tmformssoaction

The tmformssoaction data source allows you to retrieve information about a Traffic Management form-based single sign-on action configured on the Citrix ADC.

## Example usage

```terraform
data "citrixadc_tmformssoaction" "tf_tmformssoaction" {
  name = "my_formsso_action"
}

output "actionurl" {
  value = data.citrixadc_tmformssoaction.tf_tmformssoaction.actionurl
}

output "userfield" {
  value = data.citrixadc_tmformssoaction.tf_tmformssoaction.userfield
}
```

## Argument Reference

* `name` - (Required) Name of the form-based single sign-on profile. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `actionurl` - URL to which the completed form is submitted.
* `id` - The id of the tmformssoaction. It has the same value as the `name` attribute.
* `namevaluepair` - Name-value pair attributes to send to the server in addition to sending the username and password.
* `nvtype` - Type of processing of the name-value pair. Possible values: [ STATIC, DYNAMIC ]
* `passwdfield` - Name of the form field in which the user types in the password.
* `responsesize` - Number of bytes, in the response, to parse for extracting the forms.
* `ssosuccessrule` - Expression, that checks to see if single sign-on is successful.
* `submitmethod` - HTTP method used by the single sign-on form to send the logon credentials to the logon server. Possible values: [ GET, POST ]
* `userfield` - Name of the form field in which the user types in the user ID.

## Import

A tmformssoaction can be imported using its name, e.g.

```shell
terraform import citrixadc_tmformssoaction.tf_tmformssoaction my_formsso_action
```
