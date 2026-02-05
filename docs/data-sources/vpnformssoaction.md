---
subcategory: "VPN"
---

# Data Source `vpnformssoaction`

The vpnformssoaction data source allows you to retrieve information about a form-based single sign-on profile.


## Example usage

```terraform
data "citrixadc_vpnformssoaction" "tf_vpnformssoaction" {
  name = "tf_vpnformssoaction"
}

output "actionurl" {
  value = data.citrixadc_vpnformssoaction.tf_vpnformssoaction.actionurl
}

output "userfield" {
  value = data.citrixadc_vpnformssoaction.tf_vpnformssoaction.userfield
}
```


## Argument Reference

* `name` - (Required) Name for the form based single sign-on profile.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `actionurl` - Root-relative URL to which the completed form is submitted.
* `namevaluepair` - Other name-value pair attributes to send to the server, in addition to sending the user name and password. Value names are separated by an ampersand (&), such as in name1=value1&name2=value2.
* `nvtype` - How to process the name-value pair. Available settings function as follows: * STATIC - The administrator-configured values are used. * DYNAMIC - The response is parsed, the form is extracted, and then submitted.
* `passwdfield` - Name of the form field in which the user types in the password.
* `responsesize` - Maximum number of bytes to allow in the response size. Specifies the number of bytes in the response to be parsed for extracting the forms.
* `ssosuccessrule` - Expression that defines the criteria for SSO success. Expression such as checking for cookie in the response is a common example.
* `submitmethod` - HTTP method (GET or POST) used by the single sign-on form to send the logon credentials to the logon server.
* `userfield` - Name of the form field in which the user types in the user ID.
* `id` - The id of the vpnformssoaction. It has the same value as the `name` attribute.


## Import

A vpnformssoaction can be imported using its name, e.g.

```shell
terraform import citrixadc_vpnformssoaction.tf_vpnformssoaction tf_vpnformssoaction
```
