---
subcategory: "Authentication"
---

# Resource: authenticationsamlidppolicy

The authenticationsamlidppolicyresource is used to create authentication samlidppolicy resource.


## Example usage

```hcl
resource "citrixadc_sslcertkey" "tf_sslcertkey" {
  certkey = "tf_sslcertkey"
  cert    = "/var/tmp/certificate1.crt"
  key     = "/var/tmp/key1.pem"
}
resource "citrixadc_authenticationsamlidpprofile" "tf_samlidpprofile" {
  name                        = "tf_samlidpprofile"
  samlspcertname              = citrixadc_sslcertkey.tf_sslcertkey.certkey
  assertionconsumerserviceurl = "http://www.example.com"
  sendpassword                = "OFF"
  samlissuername              = "new_user"
  rejectunsignedrequests      = "ON"
  signaturealg                = "RSA-SHA1"
  digestmethod                = "SHA1"
  nameidformat                = "Unspecified"
}
resource "citrixadc_authenticationsamlidppolicy" "tf_samlidppolicy" {
  name    = "tf_samlidppolicy"
  rule    = "false"
  action  = citrixadc_authenticationsamlidpprofile.tf_samlidpprofile.name
  comment = "aSimpleTesting"
}
```


## Argument Reference

* `name` - (Required) Name for the SAML Identity Provider (IdP) authentication policy. This is used for configuring Citrix ADC as SAML Identity Provider. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the policy is created.  The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my policy" or 'my policy').
* `rule` - (Required) Expression which is evaluated to choose a profile for authentication.  The following requirements apply only to the Citrix ADC CLI: * If the expression includes one or more spaces, enclose the entire expression in double quotation marks. * If the expression itself includes double quotation marks, escape the quotations by using the \ character.  * Alternatively, you can use single quotation marks to enclose the rule, in which case you do not have to escape the double quotation marks.
* `action` - (Required) Name of the profile to apply to requests or connections that match this policy.
* `comment` - (Optional) Any comments to preserve information about this policy.
* `logaction` - (Optional) Name of messagelog action to use when a request matches this policy.
* `newname` - (Optional) New name for the SAML IdentityProvider policy.  Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.  The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my samlidppolicy policy" or 'my samlidppolicy policy').
* `undefaction` - (Optional) Action to perform if the result of policy evaluation is undefined (UNDEF). An UNDEF event indicates an internal error condition. Only the above built-in actions can be used.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the authenticationsamlidppolicy. It has the same value as the `name` attribute.


## Import

A authenticationsamlidppolicy can be imported using its name, e.g.

```shell
terraform import citrixadc_authenticationsamlidppolicy.tf_samlidppolicy tf_samlidppolicy
```
