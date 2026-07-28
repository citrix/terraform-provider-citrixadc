---
subcategory: "Cluster"
---

# Resource: clustersync_force

This resource is used to force a synchronization of the cluster configuration across nodes on the Citrix ADC.

!> **WARNING:** Forcing a cluster sync overwrites the running configuration on cluster nodes with the configuration coordinator's config. For deliberate operator use only.


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
