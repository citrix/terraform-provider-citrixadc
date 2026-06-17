---
subcategory: "NS"
---

# Data Source: nslimitselector

The nslimitselector data source allows you to retrieve information about a configured rate limit selector on the Citrix ADC.


## Example usage

```terraform
data "citrixadc_nslimitselector" "tf_nslimitselector" {
  selectorname = "tf_limitselector"
}

output "rule" {
  value = data.citrixadc_nslimitselector.tf_nslimitselector.rule
}
```


## Argument Reference

* `selectorname` - (Required) Name of the rate limit selector to look up.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `rule` - List of default-syntax expressions that identify the request fields tracked by the selector.
* `id` - The id of the nslimitselector. It has the same value as the `selectorname` attribute.
