---
subcategory: "Application Firewall"
---

# Data Source: citrixadc_appfwpolicy

The `citrixadc_appfwpolicy` data source is used to retrieve information about an existing Application Firewall Policy configured on the Citrix ADC.

## Example usage

```hcl

# Use the data source to retrieve the policy
data "citrixadc_appfwpolicy" "example_appfwpolicy" {
  name = citrixadc_appfwpolicy.example_appfwpolicy.name
}

# Use the data source outputs
output "appfwpolicy_profilename" {
  value = data.citrixadc_appfwpolicy.example_appfwpolicy.profilename
}

output "appfwpolicy_rule" {
  value = data.citrixadc_appfwpolicy.example_appfwpolicy.rule
}
```

## Argument Reference

The following arguments are required:

* `name` - (Required) Name of the application firewall policy to retrieve. Must begin with a letter, number, or the underscore character \(_\), and must contain only letters, numbers, and the hyphen \(-\), period \(.\) pound \(\#\), space \( \), at (@), equals \(=\), colon \(:\), and underscore characters.

## Attribute Reference

In addition to the argument above, the following attributes are exported:

* `id` - The id of the application firewall policy. It has the same value as the `name` attribute.
* `profilename` - Name of the application firewall profile used by this policy.
* `rule` - The Citrix ADC named rule or expression that the policy uses to determine whether to filter the connection through the application firewall with the designated profile.
* `comment` - Any comments associated with the policy for reference.
* `logaction` - Where to log information for connections that match this policy.
* `newname` - New name for the policy, if it has been renamed.
