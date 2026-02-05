---
subcategory: "Authentication"
---

# Data Source `authenticationtacacspolicy`

The authenticationtacacspolicy data source allows you to retrieve information about authentication TACACS+ policies.


## Example usage

```terraform
data "citrixadc_authenticationtacacspolicy" "tf_authenticationtacacspolicy" {
  name = "my_tacacs_policy"
}

output "rule" {
  value = data.citrixadc_authenticationtacacspolicy.tf_authenticationtacacspolicy.rule
}

output "reqaction" {
  value = data.citrixadc_authenticationtacacspolicy.tf_authenticationtacacspolicy.reqaction
}
```


## Argument Reference

* `name` - (Required) Name for the TACACS+ policy.
Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after TACACS+ policy is created.

The following requirement applies only to the Citrix ADC CLI:
If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my authentication policy" or 'my authentication policy').

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `reqaction` - Name of the TACACS+ action to perform if the policy matches.
* `rule` - Name of the Citrix ADC named rule, or an expression, that the policy uses to determine whether to attempt to authenticate the user with the TACACS+ server.

## Attribute Reference

* `id` - The id of the authenticationtacacspolicy. It has the same value as the `name` attribute.


## Import

A authenticationtacacspolicy can be imported using its name, e.g.

```shell
terraform import citrixadc_authenticationtacacspolicy.tf_authenticationtacacspolicy my_tacacs_policy
```
