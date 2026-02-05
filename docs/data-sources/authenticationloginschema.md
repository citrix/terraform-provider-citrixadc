---
subcategory: "Authentication"
---

# Data Source: citrixadc_authenticationloginschema

The `citrixadc_authenticationloginschema` data source is used to retrieve information about an existing Authentication Login Schema configured on a Citrix ADC appliance.

## Example usage

```hcl
# Retrieve an authentication login schema by name
data "citrixadc_authenticationloginschema" "example" {
  name = "demo_loginschema"
}

# Use the retrieved data in other resources
output "schema_file" {
  value = data.citrixadc_authenticationloginschema.example.authenticationschema
}

output "sso_credentials" {
  value = data.citrixadc_authenticationloginschema.example.ssocredentials
}

output "auth_strength" {
  value = data.citrixadc_authenticationloginschema.example.authenticationstrength
}

```

## Argument Reference

The following arguments are required:

* `name` - (Required) Name of the login schema to retrieve. Login schema defines the way login form is rendered. It provides a way to customize the fields that are shown to the user. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.

## Attribute Reference

In addition to the arguments, the following attributes are exported:

* `id` - The ID of the authentication login schema. It has the same value as the `name` attribute.
* `authenticationschema` - Name of the file for reading authentication schema to be sent for Login Page UI. This file should contain xml definition of elements as per Citrix Forms Authentication Protocol to be able to render login form. If administrator does not want to prompt users for additional credentials but continue with previously obtained credentials, then "noschema" can be given as argument. Please note that this applies only to loginSchemas that are used with user-defined factors, and not the vserver factor.
* `authenticationstrength` - Weight of the current authentication.
* `passwdexpression` - Expression for password extraction during login. This can be any relevant advanced policy expression.
* `passwordcredentialindex` - The index at which user entered password should be stored in session.
* `ssocredentials` - This option indicates whether current factor credentials are the default SSO (SingleSignOn) credentials.
* `usercredentialindex` - The index at which user entered username should be stored in session.
* `userexpression` - Expression for username extraction during login. This can be any relevant advanced policy expression.
