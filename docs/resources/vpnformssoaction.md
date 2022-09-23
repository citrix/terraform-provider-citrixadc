---
subcategory: "VPN"
---

# Resource: vpnformssoaction

The vpnformssoaction resource is used to create a form-based single sign-on profile.


## Example usage

```hcl
resource "citrixadc_vpnformssoaction" "tf_vpnformssoaction" {
	name = "tf_vpnformssoaction"
	actionurl = "/home"
	userfield = "username"
	passwdfield = "password"
	ssosuccessrule = "true"
	namevaluepair = "name1=value1&name2=value2"
	nvtype = "STATIC"
	responsesize = "150"
	submitmethod = "POST"
}
```


## Argument Reference

* `name` - (Required) Name for the form based single sign-on profile.
* `actionurl` - (Required) Root-relative URL to which the completed form is submitted.
* `userfield` - (Required) Name of the form field in which the user types in the user ID.
* `passwdfield` - (Required) Name of the form field in which the user types in the password.
* `ssosuccessrule` - (Required) Expression that defines the criteria for SSO success. Expression such as checking for cookie in the response is a common example.
* `namevaluepair` - (Optional) Other name-value pair attributes to send to the server, in addition to sending the user name and password. Value names are separated by an ampersand (&), such as in name1=value1&name2=value2.
* `responsesize` - (Optional) 
* `nvtype` - (Optional) How to process the name-value pair. Available settings function as follows: * STATIC - The administrator-configured values are used. * DYNAMIC - The response is parsed, the form is extracted, and then submitted. Possible values: [ STATIC, DYNAMIC ]
* `submitmethod` - (Optional) HTTP method (GET or POST) used by the single sign-on form to send the logon credentials to the logon server. Possible values: [ GET, POST ]


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vpnformssoaction. It has the same value as the `name` attribute.


## Import

A vpnformssoaction can be imported using its name, e.g.

```shell
terraform import citrixadc_vpnformssoaction.tf_vpnformssoaction tf_vpnformssoaction
```
