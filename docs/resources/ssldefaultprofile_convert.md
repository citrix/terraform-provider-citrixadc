---
subcategory: "SSL"
---

# Resource: ssldefaultprofile_convert

This resource is used to convert the Citrix ADC to SSL default profile mode.

!> **WARNING:** Converting to SSL default profile mode changes how SSL settings are applied globally on the appliance. For deliberate operator use only.


## Example usage

```hcl
resource "citrixadc_ssldefaultprofile_convert" "tf_ssldefaultprofile_convert" {
}
```


## Argument Reference

This resource has no configurable arguments.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the ssldefaultprofile_convert resource. It is set to `ssldefaultprofile_convert`.
