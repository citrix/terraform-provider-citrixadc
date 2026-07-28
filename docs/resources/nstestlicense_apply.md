---
subcategory: "NS"
---

# Resource: nstestlicense_apply

This resource is used to apply a test/eval license on the Citrix ADC.

!> **WARNING:** Applying this resource changes the licensed feature set and may disrupt the running configuration. Use only on a disposable appliance.


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
