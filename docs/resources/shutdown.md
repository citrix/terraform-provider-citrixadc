---
subcategory: "NS"
---

# Resource: shutdown

The shutdown resource performs the NITRO `shutdown` action, which shuts down the Citrix ADC appliance. It is an action-only, zero-attribute resource: applying it triggers the shutdown, and there are no configurable arguments.

~> **WARNING** Applying this resource **shuts down the appliance**, making it unreachable. It is intended for deliberate, operator-initiated use only.


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
