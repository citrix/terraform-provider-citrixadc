---
subcategory: "NS"
---

# Resource: shutdown

This resource is used to shut down the Citrix ADC appliance.

!> **WARNING:** Applying this resource shuts down the appliance, making it unreachable. For deliberate operator use only.


## Example usage

```hcl
resource "citrixadc_shutdown" "tf_shutdown" {
}
```


## Argument Reference

This resource has no configurable arguments.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the shutdown resource. It is set to `shutdown-config`.
