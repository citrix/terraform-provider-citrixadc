---
subcategory: "NS"
---

# Resource: nssourceroutecachetable_flush

The nssourceroutecachetable_flush resource flushes the source-route (Source IP MAC) cache table on the Citrix ADC. It is an action-only resource: applying it invokes the NITRO `flush` action, which clears all cached source-route entries so that subsequent traffic re-populates the table from the current network state. This is useful for clearing stale cache entries after topology or routing changes, or for troubleshooting.

This is an action resource: applying it performs the flush; it does not manage a persistent object, so re-applying re-runs the action. To read the cache-table entries, use the `citrixadc_nssourceroutecachetable` data source.


## Example usage

The flush action takes no arguments; applying the resource triggers the flush.

```hcl
resource "citrixadc_nssourceroutecachetable_flush" "tf_nssourceroutecachetable_flush" {
}
```


## Argument Reference

This resource has no configurable arguments.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nssourceroutecachetable_flush resource. It is set to `nssourceroutecachetable_flush`.
