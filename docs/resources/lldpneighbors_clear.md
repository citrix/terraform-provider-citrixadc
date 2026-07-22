---
subcategory: "LLDP"
---

# Resource: lldpneighbors_clear

The lldpneighbors_clear resource flushes the learned Link Layer Discovery Protocol (LLDP) neighbor table on the Citrix ADC. Applying it invokes the NITRO `?action=clear` operation, which discards the LLDP peer information the ADC has learned on its interfaces. Use it to force the ADC to relearn neighbor data (for example, after correcting a cabling or LLDP configuration issue) instead of waiting for the existing entries to age out.

This is an action resource: applying it performs the clear; it does not manage a persistent object, so re-applying re-runs the action. The clear action takes no arguments. To clear the table again, taint the resource (`terraform taint`) or change one of its arguments so a fresh `apply` re-issues the clear action.

LLDP must be enabled on the relevant interfaces (see `citrixadc_lldpparam` and the per-interface LLDP mode) for neighbor information to be learned in the first place. Read-only LLDP neighbor telemetry is available through the `citrixadc_lldpneighbors` data source.


## Example usage

### Clear the learned neighbor table

The clear action operates on the entire learned neighbor table and takes no arguments.

```hcl
resource "citrixadc_lldpneighbors_clear" "tf_lldpneighbors_clear" {
}
```

To clear the table again on a subsequent run, taint the resource first:

```shell
terraform taint citrixadc_lldpneighbors_clear.tf_lldpneighbors_clear
terraform apply
```


## Argument Reference

The clear action takes no arguments, so this resource has no required configuration. The following optional attributes are carried in the schema for parity with the `citrixadc_lldpneighbors` data source filters. They do **not** scope the clear action (which always operates on all learned neighbors). Changing either attribute forces a new clear action to be performed.

* `ifnum` - (Optional) Interface name. Retained for parity with the `citrixadc_lldpneighbors` data source filter; it does not scope the clear action. Changing this attribute re-triggers the clear.
* `nodeid` - (Optional) Unique number that identifies the cluster node. Retained for parity with the data source filter; it does not scope the clear action. Changing this attribute re-triggers the clear.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the lldpneighbors_clear resource. It is set to `lldpneighbors_clear`.
