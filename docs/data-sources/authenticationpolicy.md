---
subcategory: "Authentication"
---

# Data Source: citrixadc_authenticationpolicy

The `citrixadc_authenticationpolicy` data source is used to retrieve information about an authentication policy configured on a Citrix ADC.

## Example Usage

```hcl
data "citrixadc_authenticationpolicy" "example" {
  name = "my_auth_policy"
}
```

## Argument Reference

* `name` - (Required) Name of the authentication policy to retrieve.

## Attribute Reference

In addition to the argument above, the following attributes are exported:

* `id` - The ID of the authentication policy (same as `name`).
* `action` - Name of the authentication action to be performed if the policy matches.
* `comment` - Any comments to preserve information about this policy.
* `logaction` - Name of messagelog action to use when a request matches this policy.
* `newname` - New name for the authentication policy, if it has been renamed.
* `rule` - Name of the Citrix ADC named rule, or an expression, that the policy uses to determine whether to attempt to authenticate the user with the AUTHENTICATION server.
* `undefaction` - Action to perform if the result of policy evaluation is undefined (UNDEF).
