---
subcategory: "Authentication"
---

# Resource: authenticationlocalpolicy

The authenticationlocalpolicy resource is used to create authentication localpolicy resource.


## Example usage

```hcl
resource "citrixadc_authenticationlocalpolicy" "tf_authenticationlocalpolicy" {
  name   = "tf_authenticationlocalpolicy"
  rule   = "ns_true"
}
```


## Argument Reference

* `name` - (Required) Name for the local authentication policy.  Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after local policy is created.  The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my authentication policy" or 'my authentication policy').
* `rule` - (Required) Name of the Citrix ADC named rule, or an expression, that the policy uses to perform the authentication.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the authenticationlocalpolicy. It has the same value as the `name` attribute.


## Import

A authenticationlocalpolicy can be imported using its name, e.g.

```shell
terraform import citrixadc_authenticationlocalpolicy.tf_authenticationlocalpolicy tf_authenticationlocalpolicy
```
