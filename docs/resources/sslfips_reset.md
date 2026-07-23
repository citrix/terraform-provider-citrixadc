---
subcategory: "SSL"
---

# Resource: sslfips_reset

This resource is used to reset the FIPS HSM on a Citrix ADC FIPS appliance.

!> **WARNING:** Requires a FIPS appliance. This zeroizes and reinitializes the FIPS card to factory state. Each apply that creates or replaces this resource performs the reset once.


## Example usage

```hcl
resource "citrixadc_sslfips_reset" "tf_sslfips_reset" {
}
```


## Argument Reference

This resource takes no configurable arguments. The reset action operates on the appliance-wide FIPS HSM and carries no input parameters.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The ID of the sslfips_reset resource. It is set to `sslfips_reset`.
