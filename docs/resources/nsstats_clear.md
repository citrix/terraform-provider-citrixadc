---
subcategory: "NS"
---

# Resource: nsstats_clear

This resource is used to clear the statistics counters on the Citrix ADC.


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
