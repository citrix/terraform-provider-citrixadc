---
subcategory: "Vpn"
---

# Resource: vpnurlpolicy

The vpnurlpolicy resource is used to create vpn url policy.


## Example usage

```hcl
resource "citrixadc_vpnurlaction" "tf_vpnurlaction" {
  name             = "tf_vpnurlaction"
  linkname         = "new_link"
  actualurl        = "www.citrix.com"
  applicationtype  = "CVPN"
  clientlessaccess = "OFF"
  comment          = "Testing"
  ssotype          = "unifiedgateway"
  vservername      = "vserver1"
}
resource "citrixadc_vpnurlpolicy" "tf_vpnurlpolicy" {
  name = "new_policy"
  rule = "true"
  action = citrixadc_vpnurlaction.tf_vpnurlaction.name
}
```


## Argument Reference

* `name` - (Required) Name for the new urlPolicy.
* `action` - (Required) Action to be applied by the new urlPolicy if the rule criteria are met.
* `rule` - (Required) Expression, or name of a named expression, specifying the traffic that matches the policy.  The following requirements apply only to the NetScaler CLI: * If the expression includes one or more spaces, enclose the entire expression in double quotation marks. * If the expression itself includes double quotation marks, escape the quotations by using the \ character. * Alternatively, you can use single quotation marks to enclose the rule, in which case you do not have to escape the double quotation marks.
* `comment` - (Optional) Any comments to preserve information about this policy.
* `logaction` - (Optional) Name of messagelog action to use when a request matches this policy.
* `newname` - (Optional) New name for the vpn urlPolicy. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.  The following requirement applies only to the NetScaler CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my vpnurl policy" or 'my vpnurl policy').


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vpnurlpolicy. It has the same value as the `name` attribute.


## Import

A vpnurlpolicy can be imported using its name, e.g.

```shell
terraform import citrixadc_vpnurlpolicy.tf_vpnurlpolicy new_policy
```
