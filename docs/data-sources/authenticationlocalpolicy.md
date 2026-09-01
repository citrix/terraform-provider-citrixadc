---
subcategory: "Authentication"
---

# Data Source: authenticationlocalpolicy

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
* `id` - The id of the authenticationlocalpolicy. It has the same value as the `name` attribute.

### Read-only authenticationlocalpolicy metadata

These attributes are returned by the appliance on a GET (they are not configurable on the `citrixadc_authenticationlocalpolicy` resource). They are GET-only/Computed and are `null` when the appliance omits them.

* `reqaction` - The name of the RADIUS action the policy uses.

## Import

A authenticationlocalpolicy can be imported using its name, e.g.

```shell
terraform import citrixadc_authenticationlocalpolicy.tf_authenticationlocalpolicy my_local_policy
```
