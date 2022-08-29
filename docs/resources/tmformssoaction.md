---
subcategory: "Traffic Management"
---

# Resource: tmformssoaciton

The tmformssoaciton resource is used to create tmformssoaciton.


## Example usage

```hcl
resource "citrixadc_tmformssoaction" "tf_tmformssoaction" {
  name           = "my_formsso_action"
  actionurl      = "/logon.php"
  userfield      = "loginID"
  passwdfield    = "passwd"
  ssosuccessrule = "HTTP.RES.HEADER(\"Set-Cookie\").CONTAINS(\"LogonID\")"
}

```


## Argument Reference

* `name` - (Required) Name for the new form-based single sign-on profile. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after an SSO action is created. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my action" or 'my action'). Minimum length =  1
* `actionurl` - (Required) URL to which the completed form is submitted. Minimum length =  1
* `userfield` - (Required) Name of the form field in which the user types in the user ID. Minimum length =  1
* `passwdfield` - (Required) Name of the form field in which the user types in the password. Minimum length =  1
* `ssosuccessrule` - (Required) Expression, that checks to see if single sign-on is successful.
* `namevaluepair` - (Optional) Name-value pair attributes to send to the server in addition to sending the username and password. Value names are separated by an ampersand (&) (for example, name1=value1&name2=value2).
* `responsesize` - (Optional) Number of bytes, in the response, to parse for extracting the forms.
* `nvtype` - (Optional) Type of processing of the name-value pair. If you specify STATIC, the values configured by the administrator are used. For DYNAMIC, the response is parsed, and the form is extracted and then submitted. Possible values: [ STATIC, DYNAMIC ]
* `submitmethod` - (Optional) HTTP method used by the single sign-on form to send the logon credentials to the logon server. Applies only to STATIC name-value type. Possible values: [ GET, POST ]


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the tmformssoaciton. It has the same value as the `name` attribute.


## Import

A tmformssoaciton can be imported using its name, e.g.

```shell
terraform import citrixadc_tmformssoaction.tf_tmformssoaction my_formsso_action
```
