---
subcategory: "NS"
---

# Resource: nsstats_clear

Clears (resets to zero) the ever-incrementing statistics counters on the Citrix ADC. This is useful when you want to establish a fresh baseline for monitoring after a configuration change, a maintenance window, or a troubleshooting session, without having to reboot the appliance.

This is an action resource: applying it performs the clear; it does not manage a persistent object, so re-applying re-runs the action.


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

* `id` - The id of the nsstats_clear resource. It is set to `nsstats_clear`.


## Note

This resource models a one-shot action rather than a persistent ADC object. Applying it clears the selected statistics counters at the moment of apply. To clear the counters again, taint the resource or change the `cleanuplevel` argument. Importing this resource is not meaningful.
