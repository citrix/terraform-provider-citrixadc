---
subcategory: "Policy"
---

# Resource: policytracing_clear

This resource is used to clear the policy tracing records on the Citrix ADC.


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
