---
subcategory: "System"
---

# Resource: systemautorestorefeature_enable

The systemautorestorefeature_enable resource turns on the Citrix ADC autorestore feature (the NITRO `enable` action). The autorestore feature is what allows the appliance to create and manage restorepoint resources, so this action must be applied before autorestore-dependent configuration can be used.

It is an action-only, zero-attribute resource with no configurable arguments. Applying this resource invokes the enable action.

~> **NOTE** To turn the feature off, use the companion `citrixadc_systemautorestorefeature_disable` resource.


## Example usage

```hcl
resource "citrixadc_systemautorestorefeature_enable" "tf_systemautorestorefeature_enable" {
}
```


## Argument Reference

This resource has no configurable arguments.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the systemautorestorefeature_enable resource. It is set to `systemautorestorefeature_enable`.
