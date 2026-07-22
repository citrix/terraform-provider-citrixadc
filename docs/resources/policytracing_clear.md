---
subcategory: "Policy"
---

# Resource: policytracing_clear

The policytracing_clear resource clears the policy-tracing records collected on the Citrix ADC. Policy tracing records the policy-evaluation history for transactions; over time these records accumulate in memory, and you may want to discard them before starting a fresh capture. Applying this resource invokes the NITRO `clear` action (`POST ?action=clear`), removing the collected policy-tracing data from memory.

This is an action resource: applying it performs a one-shot clear and does not manage a persistent object. Re-running `terraform apply` after a destroy (or after tainting the resource) clears the records again.


## Example usage

Applying an empty `citrixadc_policytracing_clear` block clears the policy-tracing records on apply:

```hcl
resource "citrixadc_policytracing_clear" "tf_policytracing_clear" {
}
```


## Argument Reference

This resource has no configurable arguments.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the policytracing_clear resource. It is set to `policytracing_clear`.
