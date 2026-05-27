---
subcategory: "Authentication"
---

# Resource: authenticationsmartaccesspolicy

The authenticationsmartaccesspolicy resource is used to create and manage Citrix ADC authentication Smartaccess policies.


## Example usage

```hcl
resource "citrixadc_authenticationsmartaccesspolicy" "tf_authenticationsmartaccesspolicy" {
  name   = "tf_authenticationsmartaccesspolicy"
  action = "my_smartaccess_profile"
  rule   = "ns_true"
}
```


## Argument Reference

* `name` - (Required) Name for the Smartaccess policy. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after Smartaccess policy is created. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my authentication policy" or 'my authentication policy').
* `action` - (Required) Name of the Smartaccess profile to use if the policy matches.
* `rule` - (Required) Name of the Citrix ADC named rule, or an expression.
* `comment` - (Optional) Any comments to preserve information about this policy.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the authenticationsmartaccesspolicy. It has the same value as the `name` attribute.


## Import

An authenticationsmartaccesspolicy can be imported using its name, e.g.

```shell
terraform import citrixadc_authenticationsmartaccesspolicy.tf_authenticationsmartaccesspolicy tf_authenticationsmartaccesspolicy
```
