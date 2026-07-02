---
subcategory: "SSL"
---

# Data Source: sslfips

The sslfips data source allows you to retrieve the current FIPS Hardware Security Module (HSM) configuration of a Citrix ADC FIPS appliance, such as the HSM label and initialization level.

~> **WARNING: FIPS / HSM hardware required.**
> This data source queries the on-board Hardware Security Module of a dedicated FIPS appliance. It is **not supported on non-FIPS appliances** (including VPX/CPX/standard MPX models) and the read will fail there. Note that the secret password attributes are never returned by the NITRO API.

## Example usage

sslfips is a singleton; no lookup attribute is required.

```terraform
data "citrixadc_sslfips" "example" {
}

output "sslfips_hsmlabel" {
  value = data.citrixadc_sslfips.example.hsmlabel
}
```

## Argument Reference

This data source takes no arguments. The FIPS configuration is a singleton on the appliance.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslfips data source.
* `inithsm` - FIPS initialization level. The appliance currently supports Level-2 (FIPS 140-2).
* `hsmlabel` - Label to identify the Hardware Security Module (HSM).
* `fipsfw` - Path to the FIPS firmware file.

Note: the security-officer and user password attributes (`sopassword`, `oldsopassword`, `userpassword`, and their write-only variants) are secret and are never returned by the NITRO API.
