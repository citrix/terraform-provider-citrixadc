---
subcategory: "SSL"
---

# Resource: ssldefaultprofile_convert

The ssldefaultprofile_convert resource performs the NITRO `ssldefaultprofile` `convert` action, which converts the appliance to the SSL default profile mode. It is an action-only, zero-attribute resource: applying it triggers the conversion, and there are no configurable arguments.

~> **WARNING** Converting to the SSL default profile mode changes how SSL settings are applied globally on the appliance. It is intended for deliberate, operator-initiated use only.


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
