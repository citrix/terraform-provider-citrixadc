---
subcategory: "NS"
---

# Resource: nsmemrecovery_start

This resource is used to recover memory from the freepools on the Citrix ADC appliance.

This is an action-only resource: applying it triggers a one-time memory recovery from the freepools. There is no corresponding data source, and removing the resource from your configuration does not undo the memory recovery.


## Example usage

```hcl
resource "citrixadc_nsmemrecovery_start" "tf_nsmemrecovery_start" {
  percentage = 10
}
```


## Argument Reference

The following arguments are supported:

* `percentage` - (Optional) Percentage of memory to be recovered from freepools. Default value: `10`. Minimum value = `5`. Maximum value = `90`.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nsmemrecovery_start resource. It is set to `nsmemrecovery-config`.
