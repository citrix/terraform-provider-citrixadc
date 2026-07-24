---
subcategory: "NS"
---

# Resource: nssourceroutecachetable_flush

This resource is used to flush the source route cache table on the Citrix ADC.


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
