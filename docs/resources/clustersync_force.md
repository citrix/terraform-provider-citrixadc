---
subcategory: "Cluster"
---

# Resource: clustersync_force

The clustersync_force resource performs the NITRO `clustersync` `Force` action, which forces a synchronization of the cluster configuration across nodes. It is an action-only, zero-attribute resource: applying it triggers the forced sync, and there are no configurable arguments.

~> **WARNING** Forcing a cluster sync overwrites the running configuration on cluster nodes with the configuration coordinator's config. It is intended for deliberate, operator-initiated use only.


## Example usage

```hcl
resource "citrixadc_clustersync_force" "tf_clustersync_force" {
}
```


## Argument Reference

This resource has no configurable arguments.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the clustersync_force resource. It is set to `clustersync_force`.
