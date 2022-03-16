---
subcategory: "Authentication"
---

# Resource: authenticationldappolicy

The authenticationldappolicy resource is used to create authentication ldap policy.


## Example usage

```hcl
resource "citrixadc_authenticationldapaction" "tf_authenticationldapaction" {
  name          = "ldapaction"
  serverip      = "1.2.3.4"
  serverport    = 8080
  authtimeout   = 1
  ldaploginname = "username"
}
resource "citrixadc_authenticationldappolicy" "tf_authenticationldappolicy" {
  name      = "tf_authenticationldappolicy"
  rule      = "NS_TRUE"
  reqaction = citrixadc_authenticationldapaction.tf_authenticationldapaction.name
}
```


## Argument Reference

* `name` - (Required) Name for the LDAP policy.  Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after LDAP policy is created.  The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my authentication policy" or 'my authentication policy').
* `rule` - (Required) Name of the Citrix ADC named rule, or an expression, that the policy uses to determine whether to attempt to authenticate the user with the LDAP server.
* `reqaction` - (Optional) Name of the LDAP action to perform if the policy matches.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the authenticationldappolicy. It has the same value as the `name` attribute.


## Import

A authenticationldappolicy can be imported using its name, e.g.

```shell
terraform import citrixadc_authenticationldappolicy.tf_authenticationldappolicy tf_authenticationldappolicy
```
