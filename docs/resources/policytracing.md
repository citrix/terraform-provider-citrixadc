---
subcategory: "Policy"
---

# Resource: policytracing

The policytracing resource clears the policy-tracing records captured on the Citrix ADC. Policy tracing records the policy-evaluation history for transactions; over time these records accumulate and you may want to discard them before starting a fresh capture. Applying this resource invokes the NITRO `clear` action (`POST ?action=clear`) and resets those captured records.

~>
* This is an **action-only** resource. Applying it performs a one-shot side effect (clearing the captured policy-tracing records) and has no readable managed object behind it.
* The `clear` action takes no arguments, so the resource has no configurable attributes.
* Read and Update are no-ops; the captured records cannot be read back through this resource. To inspect captured policy-tracing records, use the `citrixadc_policytracing` **data source** instead.
* There is no inverse NITRO endpoint for `clear`, so destroying the resource only removes it from Terraform state — it does not undo or re-apply the clear. Re-running `terraform apply` after a destroy (or after tainting the resource) clears the records again.
* Import is not meaningful for this resource because there is no underlying queryable object.


## Example usage

Applying an empty `citrixadc_policytracing` block clears the policy-tracing records on apply:

```hcl
resource "citrixadc_policytracing" "tf_policytracing" {
}
```


## Argument Reference

This resource has no configurable arguments. The `clear` action takes no parameters.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the policytracing resource. It is a synthetic constant string `"policytracing"`. The ID is purely a Terraform state handle, not a NITRO lookup key.
