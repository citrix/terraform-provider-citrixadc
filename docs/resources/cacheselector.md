---
subcategory: "Integrated Caching"
---

# Resource: cacheselector

The cacheselector resource is used to create a cacheselector.


## Example usage

```hcl
resource "citrixadc_cacheselector" "selector1" {
    selectorname = "selector1"
    rule = ["true"]
}
```


## Argument Reference

* `rule` - (Optional) One or multiple PIXL expressions for evaluating an HTTP request or response.
* `selectorname` - (Optional) Name for the selector.  Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the cacheselector. It has the same value as the `name` attribute.


## Import

A cacheselector can be imported using its name, e.g.

```shell
terraform import citrixadc_csaction.tf_csaction tf_csaction
```
