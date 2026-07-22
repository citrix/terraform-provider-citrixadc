---
subcategory: "NS"
---

# Resource: nstestlicense_apply

The nstestlicense_apply resource applies a test/eval license on the Citrix ADC. It is an action-only resource: applying it invokes the NITRO `apply` action, which activates the test/eval license and changes the licensed feature set of the appliance.

~> **WARNING** Applying this resource **applies a test/eval license**, changing the licensed feature set and potentially disrupting the running configuration. Use it only for deliberate, operator-initiated purposes on a disposable appliance.

This is an action resource: applying it performs the license apply; it does not manage a persistent object, so re-applying re-runs the action. Use the `citrixadc_nstestlicense` data source to read the (read-only) license feature flags.


## Example usage

The apply action takes no arguments; applying the resource triggers the license apply.

```hcl
resource "citrixadc_nstestlicense_apply" "tf_nstestlicense_apply" {
}
```


## Argument Reference

This resource has no configurable arguments.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nstestlicense_apply resource. It is set to `nstestlicense_apply`.
