---
subcategory: "Spillover"
---

# Resource: spilloveraction

The spilloveraction resource is used to create spilloveraction.


## Example usage

```hcl
resource "citrixadc_spilloveraction" "tf_spilloveraction" {
  name         = "my_spilloveraction"
  action       = "SPILLOVER"
}
```


## Argument Reference

* `action` - (Required) Spillover action. Currently only type SPILLOVER is supported
* `name` - (Required) Name of the spillover action.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the spilloveraction. It has the same value as the `name` attribute.


## Import

A spilloveraction can be imported using its name, e.g.

```shell
terraform import citrixadc_spilloveraction.tf_spilloveraction my_spilloveraction
```
