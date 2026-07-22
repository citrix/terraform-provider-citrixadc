---
subcategory: "High Availability"
---

# Resource: hasync_force

Forces a configuration synchronization between the nodes of a Citrix ADC high-availability (HA) pair, copying the running configuration from the current node to its peer so both nodes stay consistent and the secondary is ready to take over. Use this when you need to trigger a sync on demand rather than relying on automatic HA propagation.

This is an action resource: applying it performs the `Force` synchronization; it does not manage a persistent object, so re-applying re-runs the action. Changing the `force` or `save` argument forces the resource to be re-created, which re-runs the forced sync.

This resource requires an HA setup and operates on an HA node. Applying it against a standalone node returns an error.


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
