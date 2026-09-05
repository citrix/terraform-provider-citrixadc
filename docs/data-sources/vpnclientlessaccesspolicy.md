---
subcategory: "VPN"
---

# Data Source: vpnclientlessaccesspolicy

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

### Read-only vpnclientlessaccesspolicy metadata

These attributes are returned by the appliance on a GET (they are not configurable on the `citrixadc_vpnclientlessaccesspolicy` resource). They are GET-only/Computed, and any attribute the appliance does not return is `null`.

* `undefaction` - The UNDEF action.
* `hits` - The number of times the policy evaluated to true.
* `undefhits` - The number of times the policy evaluation resulted in undefined processing.
* `description` - Description of the clientless access policy.
* `isdefault` - A value of true is returned if it is a default vpnclientlessrwpolicy.
* `builtin` - Flag to determine if the vpn clientless rewrite policy is built-in or not (for example `MODIFIABLE`, `DELETABLE`, `IMMUTABLE`, `PARTITION_ALL`). A list of strings.
* `feature` - The feature to be checked while applying this config.
