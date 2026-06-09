---
subcategory: "High Availability"
---

# Resource: hasync

Forces a configuration synchronization between the nodes of a Citrix ADC high-availability (HA) pair, copying the running configuration from the current node to its peer so both nodes stay consistent and the secondary is ready to take over. Use this when you need to trigger a sync on demand rather than relying on automatic HA propagation.

This is an **action-only** resource: applying it triggers a one-time `Force` synchronization action on the ADC. It does **not** manage persistent server-side state. Because NITRO exposes no GET endpoint for this action, the resource performs no read-back (drift cannot be detected), there is no update, and Read/Delete are state-only no-ops. Changing the `force` or `save` argument forces the resource to be re-created, which re-runs the forced sync.

This resource requires an HA setup and operates on an HA node. Applying it against a standalone node returns an error.


## Example usage

```hcl
resource "citrixadc_hasync" "force_sync" {
  force = true
  save  = "YES"
}
```


## Argument Reference

* `force` - (Optional) Force synchronization regardless of the state of HA propagation and HA synchronization on either node. Changing this attribute forces a new resource to be created (the forced sync is re-run).
* `save` - (Optional) After synchronization, automatically save the configuration in the secondary node configuration file (`ns.conf`) without prompting for confirmation. Possible values: [ YES, NO ]. Changing this attribute forces a new resource to be created (the forced sync is re-run).


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - A synthetic identifier with the constant value `"hasync"`. This action-only resource has no server-assigned identity.
