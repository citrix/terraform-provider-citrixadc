---
subcategory: "Authentication"
---

# Data Source `authenticationoauthidppolicy`

The authenticationoauthidppolicy data source allows you to retrieve information about authentication OAuth Identity Provider (IdP) policies.


## Example usage

```terraform
data "citrixadc_authenticationoauthidppolicy" "tf_authenticationoauthidppolicy" {
  name = "my_oauth_idp_policy"
}

output "rule" {
  value = data.citrixadc_authenticationoauthidppolicy.tf_authenticationoauthidppolicy.rule
}

output "action" {
  value = data.citrixadc_authenticationoauthidppolicy.tf_authenticationoauthidppolicy.action
}
```


## Argument Reference

* `name` - (Required) Name for the OAuth Identity Provider (IdP) authentication policy. This is used for configuring Citrix ADC as OAuth Identity Provider. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.

The following requirement applies only to the Citrix ADC CLI:
If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my policy" or 'my policy').

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `action` - Name of the profile to apply to requests or connections that match this policy.
* `comment` - Any comments to preserve information about this policy.
* `logaction` - Name of messagelog action to use when a request matches this policy.
* `newname` - New name for the OAuth IdentityProvider policy.
Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.
The following requirement applies only to the Citrix ADC CLI:
If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my oauthidppolicy policy" or 'my oauthidppolicy policy').
* `rule` - Expression that the policy uses to determine whether to respond to the specified request.
* `undefaction` - Action to perform if the result of policy evaluation is undefined (UNDEF). An UNDEF event indicates an internal error condition. Only DROP/RESET actions can be used.

## Attribute Reference

* `id` - The id of the authenticationoauthidppolicy. It has the same value as the `name` attribute.


## Import

A authenticationoauthidppolicy can be imported using its name, e.g.

```shell
terraform import citrixadc_authenticationoauthidppolicy.tf_authenticationoauthidppolicy my_oauth_idp_policy
```
