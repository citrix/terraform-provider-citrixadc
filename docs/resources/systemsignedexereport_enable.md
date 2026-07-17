---
subcategory: "System"
---

# Resource: systemsignedexereport_enable

Turns on the Citrix ADC signed executable report, which validates that system binaries are digitally signed and records the outcome. Apply this resource to invoke the NITRO `enable` action when you want signature verification reporting active on the appliance.

This is an action-only, zero-attribute resource: NITRO exposes only the `enable` and `disable` actions for `systemsignedexereport`, and the enable action accepts an empty payload with no configurable arguments. Each apply performs the enable action.

~> **NOTE** There is no NITRO GET endpoint for `systemsignedexereport`, so the resource cannot be read back or verified; `Read`/`Update` are no-ops and `Delete` simply removes the resource from Terraform state (there is no inverse of the enable action on this resource). Because there is no readable object, this resource has no data source. To disable the report, use the separate `citrixadc_systemsignedexereport_disable` resource.


## Example usage

```hcl
resource "citrixadc_systemsignedexereport_enable" "tf_systemsignedexereport_enable" {
}
```


## Argument Reference

This resource has no configurable arguments.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - A synthetic identifier for this action-only resource. It is a fixed string with the value `systemsignedexereport_enable`. It does not correspond to any readable object on the Citrix ADC, since the NITRO actions expose no GET endpoint.
