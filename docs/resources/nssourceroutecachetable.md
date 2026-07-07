---
subcategory: "NS"
---

# Resource: nssourceroutecachetable

The nssourceroutecachetable resource performs the NITRO `flush` action, which flushes the Source IP MAC cache table on the Citrix ADC. It is an action-only, zero-attribute resource: applying it triggers the flush.

~> **NOTE** There is no add/set/delete endpoint for `nssourceroutecachetable`. `Create` performs the flush, and `Read`/`Update`/`Delete` are no-ops (`Delete` only removes the resource from Terraform state). Use the `citrixadc_nssourceroutecachetable` data source to read the cache-table entries via `get (all)`.


## Example usage

```hcl
resource "citrixadc_nssourceroutecachetable" "tf_nssourceroutecachetable" {
}
```


## Argument Reference

This resource has no configurable arguments.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nssourceroutecachetable resource. It is a synthetic value (`nssourceroutecachetable-config`).
