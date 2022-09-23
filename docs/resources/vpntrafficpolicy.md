---
subcategory: "VPN"
---

# Resource: vpntrafficpolicy

The vpntrafficpolicy resource is used to create vpn traffic policy.


## Example usage

```hcl
resource "citrixadc_vpntrafficaction" "foo" {

  fta        = "ON"
  hdx        = "ON"
  name       = "Testingaction"
  qual       = "tcp"
  sso        = "ON"
}
resource "citrixadc_vpntrafficpolicy" "tf_vpntrafficpolicy" {
  name   = "tf_vpntrafficpolicy"
  rule   = "HTTP.REQ.HEADER(\"User-Agent\").CONTAINS(\"CitrixReceiver\").NOT"
  action = citrixadc_vpntrafficaction.foo.name
}
```


## Argument Reference

* `name` - (Required) Name for the traffic policy. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the policy is created.  The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my policy" or 'my policy').
* `action` - (Required) Action to apply to traffic that matches the policy.
* `rule` - (Required) Expression, or name of a named expression, against which traffic is evaluated. The following requirements apply only to the Citrix ADC CLI: * If the expression includes one or more spaces, enclose the entire expression in double quotation marks. * If the expression itself includes double quotation marks, escape the quotations by using the \ character.  * Alternatively, you can use single quotation marks to enclose the rule, in which case you do not have to escape the double quotation marks.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vpntrafficpolicy. It has the same value as the `name` attribute.


## Import

A vpntrafficpolicy can be imported using its name, e.g.

```shell
terraform import citrixadc_vpntrafficpolicy.tf_vpntrafficpolicy tf_vpntrafficpolicy
```
