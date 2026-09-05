---
subcategory: "SSL"
---

# Resource: ssldynamicclientcertcache_flush

This resource is used to flush the dynamic client certificate cache on the Citrix ADC.

This is an action-only resource that takes no attributes: applying it triggers a one-time flush of the dynamic client certificate cache. There is no corresponding data source, and removing the resource from your configuration does not undo the flush.


## Example usage

The flush action takes no arguments; applying the resource triggers the flush.

```hcl
resource "citrixadc_ssldynamicclientcertcache_flush" "tf_ssldynamicclientcertcache_flush" {
}
```


## Argument Reference

This resource has no configurable arguments.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the ssldynamicclientcertcache_flush resource. It is set to `ssldynamicclientcertcache`.
