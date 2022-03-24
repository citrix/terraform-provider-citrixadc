---
subcategory: "Authentication"
---

# Resource:authenticationoauthidpprofile

The authenticationoauthidpprofile resource is used to create authenticationOauthIdp Policy resource.


## Example usage

```hcl
resource "citrixadc_authenticationoauthidpprofile" "tf_idpprofile" {
  name         = "tf_idpprofile"
  clientid     = "cliId"
  clientsecret = "secret"
  redirecturl  = "http://www.example.com/1/"
}
resource "citrixadc_authenticationoauthidppolicy" "tf_idppolicy" {
  name    = "tf_idppolicy"
  rule    = "true"
  action  = citrixadc_authenticationoauthidpprofile.tf_idpprofile.name
  comment = "aboutpolicy"
}
```


## Argument Reference

* `name` - (Required) Name for the OAuth Identity Provider (IdP) authentication policy. This is used for configuring Citrix ADC as OAuth Identity Provider. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.  The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my policy" or 'my policy').
* `rule` - (Required) Expression that the policy uses to determine whether to respond to the specified request.
* `action` - (Required) Name of the profile to apply to requests or connections that match this policy.
* `comment` - (Optional) Any comments to preserve information about this policy.
* `logaction` - (Optional) Name of messagelog action to use when a request matches this policy.
* `newname` - (Optional) New name for the OAuth IdentityProvider policy. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my oauthidppolicy policy" or 'my oauthidppolicy policy').
* `undefaction` - (Optional) Action to perform if the result of policy evaluation is undefined (UNDEF). An UNDEF event indicates an internal error condition. Only DROP/RESET actions can be used.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the authenticationoauthidpprofile. It has the same value as the `name` attribute.


## Import

A authenticationoauthidpprofile can be imported using its name, e.g.

```shell
terraform import citrixadc_authenticationoauthidppolicy.tf_idppolicy tf_idppolicy
```
