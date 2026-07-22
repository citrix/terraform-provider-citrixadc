---
subcategory: "Cluster"
---

# Resource: clusterpropstatus_clear

Resets the configuration-propagation status counters of a Citrix ADC cluster, clearing the record of how many propagation commands have failed and the associated command strings tracked per cluster node.

This is an action resource: applying it performs the `clear` action (the equivalent of the `clear cluster propstatus` CLI command); it does not manage a persistent object, so re-applying re-runs the action.

This resource must be applied against the cluster IP (CLIP), so a cluster must already be configured.


## Example usage

Clear the propagation status counters:

```hcl
resource "citrixadc_clusterpropstatus_clear" "clear" {
}
```


## Argument Reference

This action-only resource takes no configurable arguments. The `clear` verb accepts no parameters. Each apply re-runs the clear action.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the clusterpropstatus_clear resource. It is set to `clusterpropstatus_clear`.
