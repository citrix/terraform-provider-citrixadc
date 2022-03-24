---
subcategory: "Authentication"
---

# Resource: authenticationnegotiateaction

The authenticationnegotiateaction resource is used to create authentication negotiate action resource.


## Example usage

```hcl
resource "citrixadc_authenticationnegotiateaction" "tf_negotiateaction" {
  name                       = "tf_negotiateaction"
  domain                     = "DomainName"
  domainuser                 = "username"
  domainuserpasswd           = "password"
  ntlmpath                   = "http://www.example.com/"
  defaultauthenticationgroup = "new_grpname"
}
```


## Argument Reference

* `name` - (Required) Name for the AD KDC server profile (negotiate action).  Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after AD KDC server profile is created.  The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my authentication action" or 'my authentication action').
* `defaultauthenticationgroup` - (Optional) This is the default group that is chosen when the authentication succeeds in addition to extracted groups.
* `domain` - (Optional) Domain name of the service principal that represnts Citrix ADC.
* `domainuser` - (Optional) User name of the account that is mapped with Citrix ADC principal. This can be given along with domain and password when keytab file is not available. If username is given along with keytab file, then that keytab file will be searched for this user's credentials.
* `domainuserpasswd` - (Optional) Password of the account that is mapped to the Citrix ADC principal.
* `keytab` - (Optional) The path to the keytab file that is used to decrypt kerberos tickets presented to Citrix ADC. If keytab is not available, domain/username/password can be specified in the negotiate action configuration
* `ntlmpath` - (Optional) The path to the site that is enabled for NTLM authentication, including FQDN of the server. This is used when clients fallback to NTLM.
* `ou` - (Optional) Active Directory organizational units (OU) attribute.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the authenticationnegotiateaction. It has the same value as the `name` attribute.


## Import

A authenticationnegotiateaction can be imported using its name, e.g.

```shell
terraform import citrixadc_authenticationnegotiateaction.tf_negotiateaction tf_negotiateaction
```
