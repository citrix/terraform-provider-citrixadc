---
subcategory: "VPN"
---

# Resource: vpnclientlessaccesspolicy

The vpnclientlessaccesspolicy resource, when configured for a resource on the Citrix ADC appliance, allows end-users to access the resource without using the Citrix Gateway client software.


## Example usage

```hcl
resource "citrixadc_vpnclientlessaccesspolicy" "tf_vpnclientlessaccesspolicy" {
	name = "tf_vpnclientlessaccesspolicy"
	profilename = "ns_cvpn_default_profile"
	rule = "true"
}
```


## Argument Reference

* `name` - (Required) Name of the new clientless access policy.
* `rule` - (Optional) Expression, or name of a named expression, specifying the traffic that matches the policy. The following requirements apply only to the Citrix ADC CLI: * If the expression includes one or more spaces, enclose the entire expression in double quotation marks. * If the expression itself includes double quotation marks, escape the quotations by using the \ character. * Alternatively, you can use single quotation marks to enclose the rule, in which case you do not have to escape the double quotation marks.
* `profilename` - (Optional) Name of the profile to invoke for the clientless access.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vpnclientlessaccesspolicy. It has the same value as the `name` attribute.


## Import

A vpnclientlessaccesspolicy can be imported using its name, e.g.

```shell
terraform import citrixadc_vpnclientlessaccesspolicy.tf_vpnclientlessaccesspolicy tf_vpnclientlessaccesspolicy
```
