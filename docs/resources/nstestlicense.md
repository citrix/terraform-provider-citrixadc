---
subcategory: "NS"
---

# Resource: nstestlicense

The nstestlicense resource performs the NITRO `apply` action, which applies a test/eval license on the Citrix ADC. It is an action-only, zero-attribute resource: applying it triggers the license apply.

~> **WARNING** Applying this resource **applies a test/eval license**, changing the licensed feature set and potentially disrupting the running configuration. Use it only for deliberate, operator-initiated purposes on a disposable appliance. There is no add/set/delete endpoint; `Read`/`Update`/`Delete` are no-ops. Use the `citrixadc_nstestlicense` data source to read the (read-only) license feature flags via `get (all)`.


## Example usage

```hcl
resource "citrixadc_nstestlicense" "tf_nstestlicense" {
}
```


## Argument Reference

This resource has no configurable arguments.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nstestlicense resource. It is a synthetic value (`nstestlicense-config`).
