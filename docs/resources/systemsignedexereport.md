---
subcategory: "System"
---

# Resource: systemsignedexereport

The systemsignedexereport resource is an enable/disable toggle for the Citrix ADC system signed executable report. It is an action-only, zero-attribute resource: NITRO exposes only the `enable` and `disable` actions and there are no configurable arguments.

Applying this resource **enables** the signed executable report (the `enable` action). Destroying it **disables** it (the `disable` action) — the clean inverse of apply.

~> **NOTE** There is no NITRO GET endpoint for `systemsignedexereport`, so the resource cannot be read back or verified; `Read`/`Update` are no-ops. Because there is no readable object, this resource has no data source.


## Example usage

```hcl
resource "citrixadc_systemsignedexereport" "tf_systemsignedexereport" {
}
```


## Argument Reference

This resource has no configurable arguments.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the systemsignedexereport resource. It is a synthetic value (`systemsignedexereport-config`), since the NITRO actions expose no readable object.
