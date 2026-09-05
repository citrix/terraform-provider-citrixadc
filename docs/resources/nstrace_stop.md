---
subcategory: "Basic"
---

# Resource: nstrace_stop

This resource stops a running nstrace packet capture on the Citrix ADC.

Starting and stopping a trace are separate action resources: `citrixadc_nstrace_start` and `citrixadc_nstrace_stop`. Stopping is idempotent — applying it when no trace is running succeeds. Use the `citrixadc_nstrace` data source to read the live trace state.


## Example usage

```hcl
resource "citrixadc_nstrace_stop" "stop" {
}
```


## Argument Reference

This resource has no configurable arguments.


## Attribute Reference

* `id` - The id of the nstrace_stop resource. It is set to `nstrace_stop`.
