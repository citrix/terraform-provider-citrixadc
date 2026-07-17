---
subcategory: "NS"
---

# Resource: nsstats_clear

Clears (resets to zero) the ever-incrementing statistics counters on the Citrix ADC. This is useful when you want to establish a fresh baseline for monitoring after a configuration change, a maintenance window, or a troubleshooting session, without having to reboot the appliance.

This is an **action-only** resource. Applying it triggers the NITRO `clear` action on the `nsstats` endpoint. There is no corresponding ADC object to read back, so this resource performs a one-shot action on create. Read and update are no-ops, and importing it is not meaningful.


## Example usage

```hcl
resource "citrixadc_nsstats_clear" "clear_global" {
  cleanuplevel = "global"
}
```


## Argument Reference

* `cleanuplevel` - (Required) The level of statistics to be cleared. `global` clears global counters only; `all` clears all device counters in addition to the global counters. In both cases only the ever-incrementing ("total") counters are cleared. Possible values: [ global, all ].


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - A synthetic identifier with the fixed value `nsstats_clear`. Because `nsstats_clear` is an action-only resource with no NITRO GET endpoint, the ID does not correspond to a readable ADC object.


## Note

This resource models a one-shot action rather than a persistent ADC object. Applying it clears the selected statistics counters at the moment of apply; the result cannot be read back from the ADC, so subsequent plans will not detect drift. To clear the counters again, taint the resource or change the `cleanuplevel` argument. Destroying the resource only removes it from Terraform state and has no effect on the appliance. Importing this resource is not meaningful.
