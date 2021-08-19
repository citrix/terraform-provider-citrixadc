---
subcategory: "Load Balancing"
---

# Resource: lbmetrictable

The lbmetrictable resource is used to create metric table.


## Example usage

```hcl
resource "citrixadc_lbmetrictable" "Table-Custom-1" {
  metrictable = "Table-Custom-1"
}

```


## Argument Reference

* `metrictable` - (Optional) Name for the metric table. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. CLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my metrictable" or 'my metrictable').

## Import

An lbmetrictable can be imported using its id, e.g.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the lbmetrictable. It has the same value as the `metrictable` attribute.


```shell
terraform import citrixadc_metrictable.tf_metrictable tf_metrictable
```
