---
subcategory: "System"
---

# Resource: systemautorestorefeature_disable

The systemautorestorefeature_disable resource turns off the Citrix ADC autorestore feature (the NITRO `disable` action). Disabling autorestore stops the appliance from creating and managing restorepoint resources, which is useful when you want to decommission the feature or prevent further automatic restore activity.

It is an action-only, zero-attribute resource with no configurable arguments. Applying this resource invokes the disable action.

~> **NOTE** To turn the feature on, use the companion `citrixadc_systemautorestorefeature_enable` resource.


## Example usage

```hcl
resource "citrixadc_systemautorestorefeature_disable" "tf_systemautorestorefeature_disable" {
}
```


## Argument Reference

This resource has no configurable arguments.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the systemautorestorefeature_disable resource. It is set to `systemautorestorefeature_disable`.
