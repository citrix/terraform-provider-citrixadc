---
subcategory: "SSL"
---

# Resource: sslfips_reset

The sslfips_reset resource resets the FIPS Hardware Security Module (HSM) on a Citrix ADC FIPS appliance. Use it when you need to zeroize and reinitialize the FIPS card back to a known factory state (for example, before re-provisioning the HSM, recovering from an inconsistent FIPS configuration, or as part of a documented FIPS key-management procedure).

~> **One-shot action.** This resource performs the `reset` action (CLI: `reset ssl fips`); it does not create a persistent object on the appliance. Each `terraform apply` that creates or replaces this resource performs the reset once.


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
