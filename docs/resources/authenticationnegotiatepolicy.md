---
subcategory: "Authentication"
---

# Resource: authenticationnegotiatepolicy

The authenticationnegotiatepolicy resource is used to create authentication negotiatepolicy Resource.


## Example usage

```hcl
resource "citrixadc_authenticationnegotiateaction" "tf_negotiateaction" {
  name                       = "tf_negotiateaction"
  domain                     = "DomainName"
  domainuser                 = "usersame"
  domainuserpasswd           = "password"
  ntlmpath                   = "http://www.example.com/"
  defaultauthenticationgroup = "new_grpname"
}
resource "citrixadc_authenticationnegotiatepolicy" "tf_negotiatepolicy" {
  name      = "tf_negotiatepolicy"
  rule      = "ns_true"
  reqaction = citrixadc_authenticationnegotiateaction.tf_negotiateaction.name
}
```


## Argument Reference

* `name` - (Required ) Name for the negotiate authentication policy.  Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after AD KCD (negotiate) policy is created.  The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my authentication policy" or 'my authentication policy').
* `reqaction` - (Required) Name of the negotiate action to perform if the policy matches.
* `rule` - (Required) Name of the Citrix ADC named rule, or an expression, that the policy uses to determine whether to attempt to authenticate the user with the AD KCD server.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the authenticationnegotiatepolicy. It has the same value as the `name` attribute.


## Import

A authenticationnegotiatepolicy can be imported using its name, e.g.

```shell
terraform import citrixadc_authenticationnegotiatepolicy.tf_negotiatepolicy tf_negotiatepolicy
```
