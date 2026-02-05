---
subcategory: "Authentication"
---

# Data Source `authenticationlocalpolicy`

The authenticationlocalpolicy data source allows you to retrieve information about authentication local policies.


## Example usage

```terraform
data "citrixadc_authenticationlocalpolicy" "tf_authenticationlocalpolicy" {
  name = "my_local_policy"
}

output "rule" {
  value = data.citrixadc_authenticationlocalpolicy.tf_authenticationlocalpolicy.rule
}
```


## Argument Reference

* `name` - (Required) Name for the local authentication policy.
Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after local policy is created.

The following requirement applies only to the Citrix ADC CLI:
If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my authentication policy" or 'my authentication policy').

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `rule` - Name of the Citrix ADC named rule, or an expression, that the policy uses to perform the authentication.

## Attribute Reference

* `id` - The id of the authenticationlocalpolicy. It has the same value as the `name` attribute.


## Import

A authenticationlocalpolicy can be imported using its name, e.g.

```shell
terraform import citrixadc_authenticationlocalpolicy.tf_authenticationlocalpolicy my_local_policy
```
