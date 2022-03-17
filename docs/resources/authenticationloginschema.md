---
subcategory: "Authentication"
---

# Resource: authenticationloginschema

The authenticationloginschema resource is used to create authentication loginschema Resource.


## Example usage

```hcl
resource "citrixadc_authenticationloginschema" "tf_loginschema" {
  name                    = "tf_loginschema"
  authenticationschema    = "LoginSchema/SingleAuth.xml"
  ssocredentials          = "YES"
  authenticationstrength  = "30"
  passwordcredentialindex = "10"
}
```


## Argument Reference

* `name` - (Required) Name for the new login schema. Login schema defines the way login form is rendered. It provides a way to customize the fields that are shown to the user. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after an action is created.  The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my action" or 'my action').
* `authenticationschema` - (Required) Name of the file for reading authentication schema to be sent for Login Page UI. This file should contain xml definition of elements as per Citrix Forms Authentication Protocol to be able to render login form. If administrator does not want to prompt users for additional credentials but continue with previously obtained credentials, then "noschema" can be given as argument. Please note that this applies only to loginSchemas that are used with user-defined factors, and not the vserver factor.
* `authenticationstrength` - (Optional) Weight of the current authentication
* `passwdexpression` - (Optional) Expression for password extraction during login. This can be any relevant advanced policy expression.
* `passwordcredentialindex` - (Optional) The index at which user entered password should be stored in session.
* `ssocredentials` - (Optional) This option indicates whether current factor credentials are the default SSO (SingleSignOn) credentials.
* `usercredentialindex` - (Optional) The index at which user entered username should be stored in session.
* `userexpression` - (Optional) Expression for username extraction during login. This can be any relevant advanced policy expression.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the authenticationloginschema. It has the same value as the `name` attribute.


## Import

A authenticationloginschema can be imported using its name, e.g.

```shell
terraform import citrixadc_authenticationloginschema.tf_loginschema tf_loginschema
```
