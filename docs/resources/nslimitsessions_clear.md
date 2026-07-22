---
subcategory: "NS"
---

# Resource: nslimitsessions_clear

The nslimitsessions_clear resource flushes (clears) the active rate-limit sessions tracked on the Citrix ADC for a given rate-limit identifier. Use it when you want to reset the accumulated hit/drop counters and per-selectlet session state for a limit identifier so that throttling decisions start from a clean slate.

This is an action resource: applying it flushes the active rate-limit sessions; it does not manage a persistent object, so re-applying re-runs the clear. Changing `limitidentifier` forces a new clear action to be performed (replacement).


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

* `id` - The id of the nslimitsessions_clear resource. It is set to `nslimitsessions_clear`.
