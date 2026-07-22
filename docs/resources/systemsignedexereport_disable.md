---
subcategory: "System"
---

# Resource: systemsignedexereport_disable

Turns off the Citrix ADC signed executable report, stopping the appliance from validating and reporting on the digital signatures of system binaries. Apply this resource to invoke the NITRO `disable` action when you want to switch off signature verification reporting.

This is an action-only, zero-attribute resource. Each apply performs the disable action.

~> **NOTE** To enable the report, use the separate `citrixadc_systemsignedexereport_enable` resource.


## Example usage

```hcl
resource "citrixadc_systemsignedexereport_disable" "tf_systemsignedexereport_disable" {
}
```


## Argument Reference

This resource has no configurable arguments.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the systemsignedexereport_disable resource. It is set to `systemsignedexereport_disable`.
