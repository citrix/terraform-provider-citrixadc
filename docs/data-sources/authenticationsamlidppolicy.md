---
subcategory: "Authentication"
---

# Data Source `authenticationsamlidppolicy`

The authenticationsamlidppolicy data source allows you to retrieve information about SAML Identity Provider (IdP) authentication policies.


## Example usage

```terraform
data "citrixadc_authenticationsamlidppolicy" "tf_samlidppolicy" {
  name = "tf_samlidppolicy"
}

output "rule" {
  value = data.citrixadc_authenticationsamlidppolicy.tf_samlidppolicy.rule
}

output "action" {
  value = data.citrixadc_authenticationsamlidppolicy.tf_samlidppolicy.action
}
```


## Argument Reference

* `name` - (Required) Name for the SAML Identity Provider (IdP) authentication policy. This is used for configuring Citrix ADC as SAML Identity Provider. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the policy is created.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `action` - Name of the profile to apply to requests or connections that match this policy.
* `comment` - Any comments to preserve information about this policy.
* `logaction` - Name of messagelog action to use when a request matches this policy.
* `newname` - New name for the SAML IdentityProvider policy. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.
* `rule` - Expression which is evaluated to choose a profile for authentication.
* `undefaction` - Action to perform if the result of policy evaluation is undefined (UNDEF). An UNDEF event indicates an internal error condition. Only the above built-in actions can be used.

## Attribute Reference

* `id` - The id of the authenticationsamlidppolicy. It has the same value as the `name` attribute.


## Import

A authenticationsamlidppolicy can be imported using its name, e.g.

```shell
terraform import citrixadc_authenticationsamlidppolicy.tf_samlidppolicy tf_samlidppolicy
```
