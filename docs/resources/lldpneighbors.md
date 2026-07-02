---
subcategory: "LLDP"
---

# Resource: lldpneighbors

The lldpneighbors resource flushes the learned Link Layer Discovery Protocol (LLDP) neighbor table on the Citrix ADC. Applying this resource triggers a NITRO `?action=clear` operation, which discards the LLDP peer information the ADC has learned on its interfaces. Use it to force the ADC to relearn neighbor data (for example, after correcting a cabling or LLDP configuration issue) instead of waiting for the existing entries to age out.

This is an **action-only** resource:

- `apply` performs the clear action once. The clear action takes no arguments.
- Read, Update, and Delete are no-ops; there is no persistent per-neighbor state to reconcile. Removing the resource from the configuration simply drops it from Terraform state and does not re-run the clear.
- To clear the table again, taint the resource (`terraform taint`) or replace it so a fresh `apply` re-issues the clear action.

LLDP must be enabled on the relevant interfaces (see `citrixadc_lldpparam` and the per-interface LLDP mode) for neighbor information to be learned in the first place.


## Example usage

```hcl
resource "citrixadc_lldpneighbors" "tf_lldpneighbors" {
}
```

The clear action operates on the entire learned neighbor table and takes no arguments. To clear the table again on a subsequent run, taint the resource first:

```shell
terraform taint citrixadc_lldpneighbors.tf_lldpneighbors
terraform apply
```


## Argument Reference

The clear action takes no arguments, so this resource has no required configuration. The following optional attributes are carried in the schema for parity with the datasource filters (they do not affect the clear action, which always operates on all learned neighbors):

* `ifnum` - (Optional) Interface name. Retained for parity with the `citrixadc_lldpneighbors` datasource filter; it does not scope the clear action.
* `nodeid` - (Optional) Unique number that identifies the cluster node. Retained for parity with the datasource filter; it does not scope the clear action.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the lldpneighbors resource. It is a fixed synthetic value, `lldpneighbors`, because the resource represents a transient clear action rather than a persistent object on the Citrix ADC.
