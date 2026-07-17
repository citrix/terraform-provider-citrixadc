---
subcategory: "NS"
---

# Resource: nstestlicense_apply

The nstestlicense_apply resource applies a test/eval license on the Citrix ADC. It is an action-only resource: applying it invokes the NITRO `apply` action (`POST /config/nstestlicense?action=apply`) with an empty payload, which activates the test/eval license and changes the licensed feature set of the appliance.

~> **WARNING** Applying this resource **applies a test/eval license**, changing the licensed feature set and potentially disrupting the running configuration. Use it only for deliberate, operator-initiated purposes on a disposable appliance.

This resource does not create, read, or manage a persistent object on the appliance. There is no add/set/delete endpoint for the apply action, so `Read`, `Update`, and `Delete` are no-ops (`Delete` simply removes the resource from Terraform state). Use the `citrixadc_nstestlicense` data source to read the (read-only) license feature flags via `get (all)`.


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

* `id` - A synthetic identifier for this action-only resource. It is a fixed string with the value `nstestlicense_apply`. It does not correspond to any object on the Citrix ADC.
