---
subcategory: "System"
---

# Resource: systemsignedexereport_enable

Turns on the Citrix ADC signed executable report, which validates that system binaries are digitally signed and records the outcome. Apply this resource to invoke the NITRO `enable` action when you want signature verification reporting active on the appliance.

This is an action-only, zero-attribute resource. Each apply performs the enable action.

~> **NOTE** To disable the report, use the separate `citrixadc_systemsignedexereport_disable` resource.


## Example usage

```hcl
resource "citrixadc_systemsignedexereport_enable" "tf_systemsignedexereport_enable" {
}
```


## Argument Reference

This resource has no configurable arguments.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the systemsignedexereport_enable resource. It is set to `systemsignedexereport_enable`.
