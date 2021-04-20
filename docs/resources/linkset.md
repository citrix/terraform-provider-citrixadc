---
subcategory: "Network"
---

# Resource: linkset

The linkset resource is used to create Link sets.


## Example usage

```hcl
resource "citrixadc_linkset" "tf_linkset" {
    linkset_id = "LS/1"

    interfacebinding = [
        "1/1/1",
        "2/1/1",
    ]
}
```


## Argument Reference

* `linkset_id` - (Required) Unique identifier for the linkset. Must be of the form LS/x, where x can be an integer from 1 to 32.
* `interfacebinding` - (Optional) A set of interfaces that are bound to the linkset.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the linkset. It has the same value as the `linkset_id` attribute.


## Import

A linkset can be imported using its `linskset_id` attribute, e.g.

```shell
terraform import citrixadc_linkset.tf_linkset LS/1
```
