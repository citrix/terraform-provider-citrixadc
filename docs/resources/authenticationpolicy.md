---
subcategory: "Authentication"
---

# Resource: authenticationpolicy

The authenticationpolicy resource is used to create Authentication Policy.


## Example usage

```hcl
resource "citrixadc_authenticationldapaction" "tf_authenticationldapaction" {
  name          = "ldapaction"
  serverip      = "1.2.3.4"
  serverport    = 8080
  authtimeout   = 1
  ldaploginname = "username"
}
resource "citrixadc_authenticationpolicy" "tf_authenticationpolicy" {
  name   = "tf_authenticationpolicy"
  rule   = "true"
  action = citrixadc_authenticationldapaction.tf_authenticationldapaction.name
}
```


## Argument Reference

* `name` - (Required) Name for the advance AUTHENTICATION policy.  Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after AUTHENTICATION policy is created.  The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my authentication policy" or 'my authentication policy').
* `action` - (Required) Name of the authentication action to be performed if the policy matches.
* `rule` - (Required) Name of the Citrix ADC named rule, or an expression, that the policy uses to determine whether to attempt to authenticate the user with the AUTHENTICATION server.
* `comment` - (Optional) Any comments to preserve information about this policy.
* `logaction` - (Optional) Name of messagelog action to use when a request matches this policy.
* `newname` - (Optional) New name for the authentication policy. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.   The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my authentication policy" or 'my authentication policy').
* `undefaction` - (Optional) Action to perform if the result of policy evaluation is undefined (UNDEF). An UNDEF event indicates an internal error condition. Only the above built-in actions can be used.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the authenticationpolicy. It has the same value as the `name` attribute.


## Import

A authenticationpolicy can be imported using its name, e.g.

```shell
terraform import citrixadc_authenticationpolicy.tf_authenticationpolicy tf_authenticationpolicy
```
