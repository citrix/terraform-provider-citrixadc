---
subcategory: "Vpn"
---

# Resource: vpnsessionpolicy

The vpnsessionpolicy resource is used to create vpn session policy.


## Example usage

```hcl
resource "citrixadc_vpnsessionaction" "tf_vpnsessionaction" {
  name                       = "newsession"
  sesstimeout                = "10"
  defaultauthorizationaction = "ALLOW"
}
resource "citrixadc_vpnsessionpolicy" "tf_vpnsessionpolicy" {
  name   = "tf_vpnsessionpolicy"
  rule   = "HTTP.REQ.HEADER(\"User-Agent\").CONTAINS(\"CitrixReceiver\").NOT"
  action = citrixadc_vpnsessionaction.tf_vpnsessionaction.name
}
```


## Argument Reference

* `name` - (Required) Name for the new session policy that is applied after the user logs on to Citrix Gateway.
* `action` - (Required) Action to be applied by the new session policy if the rule criteria are met.
* `rule` - (Required) Expression, or name of a named expression, specifying the traffic that matches the policy.  The following requirements apply only to the Citrix ADC CLI: * If the expression includes one or more spaces, enclose the entire expression in double quotation marks. * If the expression itself includes double quotation marks, escape the quotations by using the \ character.  * Alternatively, you can use single quotation marks to enclose the rule, in which case you do not have to escape the double quotation marks.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vpnsessionpolicy. It has the same value as the `name` attribute.


## Import

A vpnsessionpolicy can be imported using its name, e.g.

```shell
terraform import citrixadc_vpnsessionpolicy.tf_vpnsessionpolicy tf_vpnsessionpolicy
```
