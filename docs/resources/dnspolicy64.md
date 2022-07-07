---
subcategory: "DNS"
---

# Resource: dnspolicy64

The dnspolicy64 resource is used to create DNS policy64.


## Example usage

```hcl
resource "citrixadc_dnspolicy64" "dnspolicy64" {
  name  = "policy_1"
  rule = "dns.req.question.type.ne(aaaa)"
  action = "default_DNS64_action"
}
```


## Argument Reference

* `name` - (Required) Name for the DNS64 policy.
* `rule` - (Required) Expression against which DNS traffic is evaluated. Note: * On the command line interface, if the expression includes blank spaces, the entire expression must be enclosed in double quotation marks. * If the expression itself includes double quotation marks, you must escape the quotations by using the  character.  * Alternatively, you can use single quotation marks to enclose the rule, in which case you do not have to escape the double quotation marks.  Example: CLIENT.IP.SRC.IN_SUBENT(23.34.0.0/16)
* `action` - (Optional) Name of the DNS64 action to perform when the rule evaluates to TRUE. The built in actions function as follows: * A default dns64 action with prefix and mapped and exclude are any  You can create custom actions by using the add dns action command in the CLI or the DNS64 > Actions > Create DNS64 Action dialog box in the Citrix ADC configuration utility.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the dnspolicy64. It has the same value as the `name` attribute.


## Import

A <resource> can be imported using its name, e.g.

```shell
terraform import citrixadc_dnspolicy64.dnspolicy64 policy_1
```
