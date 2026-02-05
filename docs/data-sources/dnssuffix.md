---
subcategory: "DNS"
---

# Data Source `dnssuffix`

The dnssuffix data source allows you to retrieve information about DNS suffixes configured on the Citrix ADC.


## Example usage

```terraform
data "citrixadc_dnssuffix" "tf_dnssuffix" {
  dnssuffix = "example.com"
}

output "dnssuffix" {
  value = data.citrixadc_dnssuffix.tf_dnssuffix.dnssuffix
}
```


## Argument Reference

* `dnssuffix` - (Optional) Suffix to be appended when resolving domain names that are not fully qualified.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the dnssuffix. It has the same value as the `dnssuffix` attribute.
* `dnssuffix` - Suffix to be appended when resolving domain names that are not fully qualified.


## Import

A dnssuffix can be imported using its name, e.g.

```shell
terraform import citrixadc_dnssuffix.tf_dnssuffix example.com
```
