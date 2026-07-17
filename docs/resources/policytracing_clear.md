---
subcategory: "Policy"
---

# Resource: policytracing_clear

The policytracing_clear resource clears the policy-tracing records collected on the Citrix ADC. Policy tracing records the policy-evaluation history for transactions; over time these records accumulate in memory, and you may want to discard them before starting a fresh capture. Applying this resource invokes the NITRO `clear` action (`POST ?action=clear`), removing the collected policy-tracing data from memory.

This is an action-only resource. Applying it performs a one-shot side effect and does not create, read, or manage a persistent object on the appliance. There is no NITRO GET endpoint for the clear state, so Read and Update are no-ops and there is no corresponding data source for this action. There is no inverse NITRO endpoint either, so destroying the resource only removes it from Terraform state — it does not undo or re-apply the clear. Re-running `terraform apply` after a destroy (or after tainting the resource) clears the records again.


## Example usage

Applying an empty `citrixadc_policytracing_clear` block clears the policy-tracing records on apply:

```hcl
resource "citrixadc_policytracing_clear" "tf_policytracing_clear" {
}
```


## Argument Reference

This resource has no configurable arguments. The `clear` action takes no parameters; the request body is always empty.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - A synthetic identifier for this action-only resource. It is a fixed string with the value `policytracing_clear`. The ID is purely a Terraform state handle, not a NITRO lookup key.
