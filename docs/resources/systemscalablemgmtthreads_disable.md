---
subcategory: "System"
---

# Resource: systemscalablemgmtthreads_disable

This resource **disables** the Scalable Management Threads feature on the Citrix ADC.

Enabling and disabling the feature are separate action resources: `citrixadc_systemscalablemgmtthreads_enable` and `citrixadc_systemscalablemgmtthreads_disable`. This resource takes no configurable arguments. Removing it from your configuration does **not** re-enable the feature — use `citrixadc_systemscalablemgmtthreads_enable` for that. Use the `citrixadc_systemscalablemgmtthreads` data source to read the live feature state.

~> **NOTE:** This feature is not supported on all platforms; disabling it fails on appliances that do not support it.


## Example usage

```hcl
resource "citrixadc_systemscalablemgmtthreads_disable" "tf_disable" {
}
```


## Argument Reference

This resource has no configurable arguments.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the systemscalablemgmtthreads_disable resource. It is set to `systemscalablemgmtthreads_disable`.
