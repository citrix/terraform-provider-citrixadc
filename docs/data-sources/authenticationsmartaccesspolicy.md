---
subcategory: "Authentication"
---

# Data Source: authenticationsmartaccesspolicy

The authenticationsmartaccesspolicy data source allows you to retrieve information about Citrix ADC authentication Smartaccess policies.


## Example usage

```terraform
data "citrixadc_authenticationsmartaccesspolicy" "tf_authenticationsmartaccesspolicy" {
  name = "my_smartaccess_policy"
}

output "rule" {
  value = data.citrixadc_authenticationsmartaccesspolicy.tf_authenticationsmartaccesspolicy.rule
}
```


## Argument Reference

* `name` - (Required) Name for the Smartaccess policy.
Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after Smartaccess policy is created.

The following requirement applies only to the Citrix ADC CLI:
If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my authentication policy" or 'my authentication policy').

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `action` - Name of the Smartaccess profile to use if the policy matches.
* `rule` - Name of the Citrix ADC named rule, or an expression.
* `comment` - Any comments to preserve information about this policy.
* `id` - The id of the authenticationsmartaccesspolicy. It has the same value as the `name` attribute.
