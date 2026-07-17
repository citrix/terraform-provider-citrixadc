---
subcategory: "NS"
---

# Resource: nslimitsessions_clear

The nslimitsessions_clear resource flushes (clears) the active rate-limit sessions tracked on the Citrix ADC for a given rate-limit identifier. Use it when you want to reset the accumulated hit/drop counters and per-selectlet session state for a limit identifier so that throttling decisions start from a clean slate.

This resource maps to the NITRO `clear` action (`POST ?action=clear`); it does not create, read, or manage a persistent object on the appliance. There is no NITRO GET endpoint for cleared sessions, so there is no corresponding data source. Each apply performs the clear once; changing `limitidentifier` forces a new clear action to be performed (replacement).


## Example usage

```hcl
resource "citrixadc_nslimitsessions_clear" "tf_nslimitsessions_clear" {
  limitidentifier = "myratelimit"
}
```


## Argument Reference

* `limitidentifier` - (Required) Name of the rate limit identifier whose sessions you want to clear. Changing this value forces the resource to be replaced (re-running the clear action against the new identifier).


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - A synthetic identifier for this action-only resource. It is a fixed string with the value `nslimitsessions_clear`. It does not correspond to any object on the Citrix ADC.
