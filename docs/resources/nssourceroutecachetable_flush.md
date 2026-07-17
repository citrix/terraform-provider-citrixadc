---
subcategory: "NS"
---

# Resource: nssourceroutecachetable_flush

The nssourceroutecachetable_flush resource flushes the source-route (Source IP MAC) cache table on the Citrix ADC. It is an action-only resource: applying it invokes the NITRO `flush` action, which clears all cached source-route entries so that subsequent traffic re-populates the table from the current network state. This is useful for clearing stale cache entries after topology or routing changes, or for troubleshooting.

This resource does not create, read, or manage a persistent object on the appliance. There is no add/set/delete endpoint for the flush action, so `Create` performs the flush while `Read` and `Update` are no-ops and `Delete` only removes the resource from Terraform state. To read the cache-table entries, use the `citrixadc_nssourceroutecachetable` data source (`get (all)`).


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

* `id` - A synthetic identifier for this action-only resource. It is a fixed string with the value `nssourceroutecachetable_flush`. It does not correspond to any object on the Citrix ADC.
