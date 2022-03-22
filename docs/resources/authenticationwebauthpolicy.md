---
subcategory: "Authentication"
---

# Resource: authenticationwebauthpolicy

The authenticationwebauthpolicy resource is used to create authentication webauthpolicy resource.


## Example usage

```hcl
resource "citrixadc_authenticationwebauthaction" "tf_webauthaction" {
  name                       = "tf_webauthaction"
  serverip                   = "1.2.3.4"
  serverport                 = 8080
  fullreqexpr                = "TRUE"
  scheme                     = "http"
  successrule                = "http.RES.STATUS.EQ(200)"
  defaultauthenticationgroup = "new_group"
}
resource "citrixadc_authenticationwebauthpolicy" "tf_webauthpolicy" {
  name   = "tf_webauthpolicy"
  rule   = "NS_TRUE"
  action = citrixadc_authenticationwebauthaction.tf_webauthaction.name
}
```


## Argument Reference

* `name` - (Required) Name for the WebAuth policy.  Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after LDAP policy is created.  The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my authentication policy" or 'my authentication policy').
* `action` - (Required) Name of the WebAuth action to perform if the policy matches.
* `rule` - (Required) Name of the Citrix ADC named rule, or an expression, that the policy uses to determine whether to attempt to authenticate the user with the Web server.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the authenticationwebauthpolicy. It has the same value as the `name` attribute.


## Import

A authenticationwebauthpolicy can be imported using its name, e.g.

```shell
terraform import citrixadc_authenticationwebauthpolicy.tf_webauthpolicy tf_webauthpolicy
```
