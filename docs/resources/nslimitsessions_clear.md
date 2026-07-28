---
subcategory: "NS"
---

# Resource: nslimitsessions_clear

This resource is used to clear the active rate-limit sessions for a rate-limit identifier on the Citrix ADC.


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
