---
subcategory: "Ns"
---

# Resource: nstimer

The nstimer resource is used to create Timer resource.


## Example usage

```hcl
resource "citrixadc_nstimer" "tf_nstimer" {
  name     = "tf_nstimer"
  interval = 10
  unit     = "SEC"
  comment  = "Testing"
}
```


## Argument Reference

* `name` - (Required) Timer name. Minimum length =  1
* `interval` - (Optional) The frequency at which the policies bound to this timer are invoked. The minimum value is 20 msec. The maximum value is 20940 in seconds and 349 in minutes. Minimum value =  1 Maximum value =  20940000
* `unit` - (Optional) Timer interval unit. Possible values: [ SEC, MIN ]
* `comment` - (Optional) Comments associated with this timer.
* `newname` - (Optional) The new name of the timer. Minimum length =  1


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nstimer. It has the same value as the `name` attribute.


## Import

A nstimer can be imported using its name, e.g.

```shell
terraform import citrixadc_nstimer.tf_nstimer tf_nstimer
```
