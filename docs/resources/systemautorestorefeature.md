---
subcategory: "System"
---

# Resource: systemautorestorefeature

The systemautorestorefeature resource is an enable/disable toggle for the Citrix ADC autorestore feature (used to create restorepoint resources). It is an action-only, zero-attribute resource: NITRO exposes only the `enable` and `disable` actions and there are no configurable arguments.

Applying this resource **enables** the autorestore feature (the `enable` action). Destroying it **disables** the feature (the `disable` action) — the clean inverse of apply.

~> **NOTE** There is no NITRO GET endpoint for `systemautorestorefeature`, so the resource cannot be read back or verified; `Read`/`Update` are no-ops. Because there is no readable object, this resource has no data source.


## Example usage

```hcl
resource "citrixadc_systemautorestorefeature" "tf_systemautorestorefeature" {
}
```


## Argument Reference

This resource has no configurable arguments.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the systemautorestorefeature resource. It is a synthetic value (`systemautorestorefeature-config`), since the NITRO actions expose no readable object.
