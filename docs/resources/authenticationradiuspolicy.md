---
subcategory: "Authentication"
---

# Resource: authenticationradiuspolicy

The authenticationradiuspolicy resource is used to create authentication radiuspolicy resource.


## Example usage

```hcl
resource "citrixadc_authenticationradiusaction" "tf_radiusaction" {
  name         = "tf_radiusaction"
  radkey       = "secret"
  serverip     = "1.2.3.4"
  serverport   = 8080
  authtimeout  = 2
  radnasip     = "DISABLED"
  passencoding = "chap"
}
resource "citrixadc_authenticationradiuspolicy" "tf_radiuspolicy" {
  name      = "tf_radiuspolicy"
  rule      = "NS_TRUE"
  reqaction = citrixadc_authenticationradiusaction.tf_radiusaction.name
}
```


## Argument Reference

* `name` - (Required) Name for the RADIUS authentication policy.  Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after RADIUS policy is created.  The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my authentication policy" or 'my authentication policy').
* `rule` - (Required) Name of the Citrix ADC named rule, or an expression, that the policy uses to determine whether to attempt to authenticate the user with the RADIUS server.
* `reqaction` - (Optional) Name of the RADIUS action to perform if the policy matches.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the authenticationradiuspolicy. It has the same value as the `name` attribute.


## Import

A <authenticationradiuspolicy can be imported using its name, e.g.

```shell
terraform import citrixadc_authenticationradiuspolicy.tf_radiuspolicy tf_radiuspolicy
```
