---
subcategory: "Network"
---

# Resource: vxlanvlanmap

The vxlanvlanmap resource is used to create vxlanvlanmap.


## Example usage

```hcl
resource "citrixadc_vxlanvlanmap" "tf_vxlanvlanmp" {
  name = "tf_vxlanvlanmp"
}
```


## Argument Reference

* `name` - (Required) Name of the mapping table.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vxlanvlanmap. It has the same value as the `name` attribute.


## Import

A vxlanvlanmap can be imported using its name, e.g.

```shell
terraform import citrixadc_vxlanvlanmap.tf_vxlanvlanmp tf_vxlanvlanmp
```
