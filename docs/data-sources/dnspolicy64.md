---
subcategory: "DNS"
---

# Data Source `dnspolicy64`

The dnspolicy64 data source allows you to retrieve information about DNS64 policies.


## Example usage

```terraform
data "citrixadc_dnspolicy64" "tf_dnspolicy64" {
  name = "my_dnspolicy64"
}

output "action" {
  value = data.citrixadc_dnspolicy64.tf_dnspolicy64.action
}

output "rule" {
  value = data.citrixadc_dnspolicy64.tf_dnspolicy64.rule
}
```


## Argument Reference

* `name` - (Required) Name for the DNS64 policy.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `action` - Name of the DNS64 action to perform when the rule evaluates to TRUE. The built in actions function as follows: A default dns64 action with prefix <default prefix> and mapped and exclude are any. You can create custom actions by using the add dns action command in the CLI or the DNS64 > Actions > Create DNS64 Action dialog box in the Citrix ADC configuration utility.
* `rule` - Expression against which DNS traffic is evaluated. Note: On the command line interface, if the expression includes blank spaces, the entire expression must be enclosed in double quotation marks. If the expression itself includes double quotation marks, you must escape the quotations by using the character. Alternatively, you can use single quotation marks to enclose the rule, in which case you do not have to escape the double quotation marks. Example: CLIENT.IP.SRC.IN_SUBENT(23.34.0.0/16)
* `id` - The id of the dnspolicy64. It has the same value as the `name` attribute.
