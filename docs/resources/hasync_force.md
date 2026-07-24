---
subcategory: "High Availability"
---

# Resource: hasync_force

This resource is used to force a configuration synchronization across a high-availability (HA) pair.


## Example usage

```hcl
resource "citrixadc_hasync_force" "force_sync" {
  force = true
  save  = "YES"
}
```


## Argument Reference

* `force` - (Optional) Force synchronization regardless of the state of HA propagation and HA synchronization on either node. Changing this attribute forces a new resource to be created (the forced sync is re-run).
* `save` - (Optional) After synchronization, automatically save the configuration in the secondary node configuration file (`ns.conf`) without prompting for confirmation. Possible values: [ YES, NO ]. Changing this attribute forces a new resource to be created (the forced sync is re-run).


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the hasync_force resource. It is set to `hasync_force`.
