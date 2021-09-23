---
subcategory: "Vpn"
---

# Resource: vpneula

The vpneula resource is used to create a new eula entity.


## Example usage

```hcl
resource "citrixadc_vpneula" "tf_vpneula" {
	name = "tf_vpneula"	
}
```


## Argument Reference

* `name` - (Required) Name for the eula.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vpneula. It has the same value as the `name` attribute.


## Import

A vpneula can be imported using its name, e.g.

```shell
terraform import citrixadc_vpneula.tf_vpneula tf_vpneula
```
