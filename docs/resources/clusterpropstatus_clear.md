---
subcategory: "Cluster"
---

# Resource: clusterpropstatus_clear

Resets the configuration-propagation status counters of a Citrix ADC cluster, clearing the record of how many propagation commands have failed and the associated command strings tracked per cluster node.

This is an action-only resource, not a persistent configuration object. Applying it triggers a one-time `clear` action (the equivalent of the `clear cluster propstatus` CLI command); it does not manage any remote state. Because NITRO exposes no add, update, or delete endpoint for this action, Read, Update, and Delete are no-ops (state-only). There is no corresponding data source.

This resource must be applied against the cluster IP (CLIP), so a cluster must already be configured.


## Example usage

Clear the propagation status counters:

```hcl
resource "citrixadc_clusterpropstatus_clear" "clear" {
}
```


## Argument Reference

This action-only resource takes no configurable arguments. The `clear` verb accepts no parameters (verified against the live CLI: `clear cluster propstatus` takes zero arguments). Each apply re-runs the clear action.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - A synthetic identifier with the constant value `clusterpropstatus_clear`. The clear action has no NITRO identity of its own, so this ID does not reference any server-assigned key.
