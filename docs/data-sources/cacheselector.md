---
subcategory: "Integrated Caching"
---

# Data Source `cacheselector`

The cacheselector data source allows you to retrieve information about an existing cacheselector.


## Example usage

```terraform
data "citrixadc_cacheselector" "tf_cacheselector" {
  selectorname = "my_cacheselector"
}

output "selectorname" {
  value = data.citrixadc_cacheselector.tf_cacheselector.selectorname
}

output "rule" {
  value = data.citrixadc_cacheselector.tf_cacheselector.rule
}
```


## Argument Reference

* `selectorname` - (Required) Name for the selector.  Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the cacheselector. It has the same value as the `selectorname` attribute.
* `rule` - One or multiple PIXL expressions for evaluating an HTTP request or response.
