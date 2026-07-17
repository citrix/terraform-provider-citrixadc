---
subcategory: "System"
---

# Resource: systemsignedexereport_disable

Turns off the Citrix ADC signed executable report, stopping the appliance from validating and reporting on the digital signatures of system binaries. Apply this resource to invoke the NITRO `disable` action when you want to switch off signature verification reporting.

This is an action-only, zero-attribute resource: NITRO exposes only the `enable` and `disable` actions for `systemsignedexereport`, and the disable action accepts an empty payload with no configurable arguments. Each apply performs the disable action.

~> **NOTE** There is no NITRO GET endpoint for `systemsignedexereport`, so the resource cannot be read back or verified; `Read`/`Update` are no-ops and `Delete` simply removes the resource from Terraform state (there is no inverse of the disable action on this resource). Because there is no readable object, this resource has no data source. To enable the report, use the separate `citrixadc_systemsignedexereport_enable` resource.


## Example usage

```hcl
resource "citrixadc_systemsignedexereport_disable" "tf_systemsignedexereport_disable" {
}
```


## Argument Reference

This resource has no configurable arguments.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - A synthetic identifier for this action-only resource. It is a fixed string with the value `systemsignedexereport_disable`. It does not correspond to any readable object on the Citrix ADC, since the NITRO actions expose no GET endpoint.
