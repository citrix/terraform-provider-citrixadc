---
subcategory: "System"
---

# Resource: systemautorestorefeature_disable

The systemautorestorefeature_disable resource turns off the Citrix ADC autorestore feature (the NITRO `disable` action). Disabling autorestore stops the appliance from creating and managing restorepoint resources, which is useful when you want to decommission the feature or prevent further automatic restore activity.

It is an action-only, zero-attribute resource: NITRO exposes only the `disable` action for `systemautorestorefeature` and there are no configurable arguments. Applying this resource invokes the disable action.

~> **NOTE** There is no NITRO GET endpoint for `systemautorestorefeature`, so the resource cannot be read back or verified; `Read`/`Update` are no-ops and `Delete` simply removes the resource from Terraform state (there is no inverse API bound to this resource). Because there is no readable object, this resource has no data source. To turn the feature on, use the companion `citrixadc_systemautorestorefeature_enable` resource.


## Example usage

```hcl
resource "citrixadc_systemautorestorefeature_disable" "tf_systemautorestorefeature_disable" {
}
```


## Argument Reference

This resource has no configurable arguments.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the systemautorestorefeature_disable resource. It is a synthetic value (`systemautorestorefeature_disable`), since the NITRO disable action exposes no readable object.
