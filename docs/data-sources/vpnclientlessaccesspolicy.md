---
subcategory: "VPN"
---

# Data Source: citrixadc_vpnclientlessaccesspolicy

The vpnclientlessaccesspolicy data source allows you to retrieve information about a VPN clientless access policy configured on the Citrix ADC.

## Example usage

```terraform
data "citrixadc_vpnclientlessaccesspolicy" "tf_vpnclientlessaccesspolicy" {
  name = "my_clientless_policy"
}

output "profilename" {
  value = data.citrixadc_vpnclientlessaccesspolicy.tf_vpnclientlessaccesspolicy.profilename
}

output "rule" {
  value = data.citrixadc_vpnclientlessaccesspolicy.tf_vpnclientlessaccesspolicy.rule
}
```

## Argument Reference

* `name` - (Required) Name of the clientless access policy.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vpnclientlessaccesspolicy. It has the same value as the `name` attribute.
* `profilename` - Name of the profile to invoke for the clientless access.
* `rule` - Expression, or name of a named expression, specifying the traffic that matches the policy. The following requirements apply only to the Citrix ADC CLI: * If the expression includes one or more spaces, enclose the entire expression in double quotation marks. * If the expression itself includes double quotation marks, escape the quotations by using the \ character. * Alternatively, you can use single quotation marks to enclose the rule, in which case you do not have to escape the double quotation marks.
