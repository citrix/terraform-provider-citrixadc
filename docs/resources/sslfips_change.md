---
subcategory: "SSL"
---

# Resource: sslfips_change

The sslfips_change resource updates the firmware on the FIPS Hardware Security Module (HSM) of a Citrix ADC FIPS appliance by pushing a new FIPS firmware image to the card. Use it when you need to upgrade or reload the FIPS card firmware from a firmware file staged on the appliance.

~> **One-shot action.** This resource performs the `change` action (CLI: `update ssl fips`); it does not create a persistent object on the appliance. Each `terraform apply` that creates or replaces this resource performs the firmware change once, and changing `fipsfw` forces a new firmware change (replacement).


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
