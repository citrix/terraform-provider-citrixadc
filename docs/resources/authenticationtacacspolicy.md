---
subcategory: "Authentication"
---

# Resource: authenticationtacacspolicy

The authenticationtacacspolicy resource is used to create authentication tacacspolicy resource.


## Example usage

```hcl
resource "citrixadc_authenticationtacacsaction" "tf_tacacsaction" {
  name            = "tf_tacacsaction"
  serverip        = "1.2.3.4"
  serverport      = 8080
  authtimeout     = 5
  authorization   = "ON"
  accounting      = "ON"
  auditfailedcmds = "ON"
  groupattrname   = "group"
}
resource "citrixadc_authenticationtacacspolicy" "tf_tacacspolicy" {
  name= "tf_tacacspolicy"
  rule= "NS_FALSE"
  reqaction= citrixadc_authenticationtacacsaction.tf_tacacsaction.name
  
}
```


## Argument Reference

* `name` - (Required) Name for the TACACS+ policy.  Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after TACACS+ policy is created.  The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my authentication policy" or 'my authentication policy').
* `rule` - (Required) Name of the Citrix ADC named rule, or an expression, that the policy uses to determine whether to attempt to authenticate the user with the TACACS+ server.
* `reqaction` - (Optional) Name of the TACACS+ action to perform if the policy matches.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the authenticationtacacspolicy. It has the same value as the `name` attribute.


## Import

A authenticationtacacspolicy can be imported using its name, e.g.

```shell
terraform import citrixadc_authenticationtacacspolicy.tf_tacacspolicy tf_tacacspolicy
```
