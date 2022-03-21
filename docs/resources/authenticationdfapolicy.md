---
subcategory: "Authentication"
---

# Resource: authenticationdfapolicy

The authenticationdfapolicy resource is used to create authentication dfapolicy resource.


## Example usage

```hcl
resource "citrixadc_authenticationdfaaction" "tf_dfaaction" {
  name       = "tf_dfaaction"
  serverurl  = "https://example.com/"
  clientid   = "cliId"
  passphrase = "secret"
}
resource "citrixadc_authenticationdfapolicy" "td_dfapolicy" {
  name   = "td_dfapolicy"
  rule   = "NS_TRUE"
  action = citrixadc_authenticationdfaaction.tf_dfaaction.name
}
```


## Argument Reference

* `name` - (Required) Name for the DFA policy.  Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after DFA policy is created.  The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my authentication policy" or 'my authentication policy').
* `action` - (Required) Name of the DFA action to perform if the policy matches.
* `rule` - (Required) Name of the Citrix ADC named rule, or an expression, that the policy uses to determine whether to attempt to authenticate the user with the Web server.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the authenticationdfapolicy. It has the same value as the `name` attribute.


## Import

A authenticationdfapolicy can be imported using its name, e.g.

```shell
terraform import citrixadc_authenticationdfapolicy.td_dfapolicy td_dfapolicy
```
