---
subcategory: "Authentication"
---

# Data Source: citrixadc_authenticationdfapolicy

The `citrixadc_authenticationdfapolicy` data source is used to retrieve information about an existing Authentication DFA (Distributed Fingerprint Authentication) Policy configured on a Citrix ADC appliance.

## Example usage

```hcl
# Retrieve an authentication DFA policy by name
data "citrixadc_authenticationdfapolicy" "example" {
  name = "demo_dfapolicy"
}

# Use the retrieved data in other resources
output "policy_rule" {
  value = data.citrixadc_authenticationdfapolicy.example.rule
}

output "policy_action" {
  value = data.citrixadc_authenticationdfapolicy.example.action
}

```

## Argument Reference

The following arguments are required:

* `name` - (Required) Name of the authentication DFA policy to retrieve. This is the unique identifier for the policy. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters.

## Attribute Reference

In addition to the arguments, the following attributes are exported:

* `id` - The ID of the authentication DFA policy. It has the same value as the `name` attribute.
* `action` - Name of the DFA action to perform if the policy matches.
* `rule` - Name of the Citrix ADC named rule, or an expression, that the policy uses to determine whether to attempt to authenticate the user with the Web server.
