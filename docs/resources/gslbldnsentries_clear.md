---
subcategory: "GSLB"
---

# Resource: gslbldnsentries_clear

Flushes the learned GSLB local DNS (LDNS) round-trip time (RTT) entries from the Citrix ADC. The Citrix ADC builds an LDNS table by measuring the RTT to the local DNS servers that resolve client queries; this data feeds the dynamic RTT-based GSLB load balancing method. Apply this resource to clear that accumulated table so the appliance relearns LDNS proximity from scratch (for example, after a topology change or stale-measurement cleanup).

This is an action resource: applying it performs the `clear` action, flushing the learned LDNS RTT entries; it does not manage a persistent object, so re-applying re-runs the action. To flush the table again, change `nodeid` (which forces replacement) or taint the resource. The `citrixadc_gslbldnsentries` data source can be used to read and count the entries.


## Example usage

```hcl
# Flush the learned GSLB LDNS RTT entries on the appliance.
resource "citrixadc_gslbldnsentries_clear" "tf_gslbldnsentries_clear" {
}
```

To target a specific cluster node when clearing:

```hcl
resource "citrixadc_gslbldnsentries_clear" "tf_gslbldnsentries_clear" {
  nodeid = 1
}
```


## Argument Reference

* `nodeid` - (Optional) Unique number that identifies the cluster node. Changing it forces resource replacement, which re-runs the `clear` action.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the gslbldnsentries_clear resource. It is set to `gslbldnsentries_clear`.
