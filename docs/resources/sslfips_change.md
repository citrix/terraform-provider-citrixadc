---
subcategory: "SSL"
---

# Resource: sslfips_change

This resource is used to update the firmware on the FIPS HSM of a Citrix ADC FIPS appliance.

~> **One-shot action:** Requires a FIPS appliance. Each apply that creates or replaces this resource performs the firmware change once; changing `fipsfw` forces a new change.


## Example usage

```hcl
resource "citrixadc_sslfips_change" "tf_sslfips_change" {
  fipsfw = "FIPS-140-2-level-3"
}
```


## Argument Reference

* `fipsfw` - (Required) Path to the FIPS firmware file to apply to the HSM. Maximum length: 63. Changing this value forces the resource to be recreated (re-running the firmware change action with the new firmware file).


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The ID of the sslfips_change resource. It has the format `sslfips_change-<fipsfw>` (for example, `sslfips_change-FIPS-140-2-level-3`).
