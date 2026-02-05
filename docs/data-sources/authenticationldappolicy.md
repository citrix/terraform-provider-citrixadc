---
subcategory: "Authentication"
---

# Data Source: citrixadc_authenticationldappolicy

The `citrixadc_authenticationldappolicy` data source is used to retrieve information about an existing Authentication LDAP (Lightweight Directory Access Protocol) Policy configured on a Citrix ADC appliance.

## Example usage

```hcl
# Retrieve an authentication LDAP policy by name
data "citrixadc_authenticationldappolicy" "example" {
  name = "demo_ldappolicy"
}

# Use the retrieved data in other resources
output "policy_rule" {
  value = data.citrixadc_authenticationldappolicy.example.rule
}

output "policy_reqaction" {
  value = data.citrixadc_authenticationldappolicy.example.reqaction
}

```

## Argument Reference

The following arguments are required:

* `name` - (Required) Name of the authentication LDAP policy to retrieve. This is the unique identifier for the policy. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters.

## Attribute Reference

In addition to the arguments, the following attributes are exported:

* `id` - The ID of the authentication LDAP policy. It has the same value as the `name` attribute.
* `reqaction` - Name of the LDAP action to perform if the policy matches.
* `rule` - Name of the Citrix ADC named rule, or an expression, that the policy uses to determine whether to attempt to authenticate the user with the LDAP server.
