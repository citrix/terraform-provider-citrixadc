---
subcategory: "Cluster"
---

# Resource: clusterpropstatus

Resets the configuration-propagation status counters of a Citrix ADC cluster, clearing the record of how many propagation commands have failed and the associated command strings tracked per cluster node.

This is an imperative action resource, not a persistent configuration object. Applying it triggers a one-time `clear` action (the equivalent of the `clear cluster propstatus` CLI command); it does not manage any remote state. Because NITRO exposes no add, update, or delete endpoint for this action, Read and Delete are no-ops (state-only) and changing the `nodeid` argument forces the resource to be re-created, which re-runs the clear.

This resource must be applied against the cluster IP (CLIP), so a cluster must already be configured.


## Example usage

Clear the propagation status counters on all cluster nodes:

```hcl
resource "citrixadc_clusterpropstatus" "clear_all" {
}
```

Clear the propagation status counters on a single cluster node:

```hcl
resource "citrixadc_clusterpropstatus" "clear_node1" {
  nodeid = 1
}
```


## Argument Reference

* `nodeid` - (Optional) Unique number that identifies the cluster node whose propagation status counters are cleared. Omit this argument to clear the counters on all nodes. Minimum value = 0. Maximum value = 31. Changing this value forces the resource to be re-created (re-running the clear action).


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - A synthetic identifier with the constant value `clusterpropstatus`. The clear action has no NITRO identity of its own, so this ID does not reference any server-assigned key.
