---
subcategory: "Authentication"
---

# Resource: authenticationsamlpolicy

The authenticationsamlpolicy resource is used to create authentication samlprofile.


## Example usage

```hcl
resource "citrixadc_authenticationsamlaction" "tf_samlaction" {
  name                    = "tf_samlaction"
  metadataurl             = "http://www.example.com"
  samltwofactor           = "OFF"
  requestedauthncontext   = "minimum"
  digestmethod            = "SHA1"
  signaturealg            = "RSA-SHA256"
  metadatarefreshinterval = 1
}
resource "citrixadc_authenticationsamlpolicy" "tf_samlpolicy" {
  name      = "tf_samlpolicy"
  rule      = "NS_TRUE"
  reqaction = citrixadc_authenticationsamlaction.tf_samlaction.name
}
```


## Argument Reference

* `name` - (Required) Name for the SAML policy.  Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after SAML policy is created.  The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my authentication policy" or 'my authentication policy').
* `reqaction` - (Required) Name of the SAML authentication action to be performed if the policy matches.
* `rule` - (Required) Name of the Citrix ADC named rule, or an expression, that the policy uses to determine whether to attempt to authenticate the user with the SAML server.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the authenticationsamlpolicy. It has the same value as the `name` attribute.


## Import

A authenticationsamlpolicy can be imported using its name, e.g.

```shell
terraform import citrixadc_authenticationsamlpolicy.tf_samlpolicy tf_samlpolicy
```
