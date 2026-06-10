---
subcategory: "GSLB"
---

# Resource: gslbldnsentries

Flushes the learned GSLB local DNS (LDNS) round-trip time (RTT) entries from the Citrix ADC. The Citrix ADC builds an LDNS table by measuring the RTT to the local DNS servers that resolve client queries; this data feeds the dynamic RTT-based GSLB load balancing method. Apply this resource to clear that accumulated table so the appliance relearns LDNS proximity from scratch (for example, after a topology change or stale-measurement cleanup).

This is an **action-only** resource that performs the imperative NITRO `clear` action:

* **Apply (create)** invokes `clear`, flushing the learned LDNS RTT entries. The `clear` action takes no arguments.
* **Read** is a no-op. The cleared entries are not a persistent managed object, so the provider preserves state rather than re-fetching (which would otherwise cause perpetual drift).
* **Update** is a no-op. There is no NITRO update endpoint, and the only argument (`nodeid`) forces replacement.
* **Delete** is a state-only removal. `clear` has no inverse NITRO endpoint, so destroying the resource simply removes it from Terraform state without contacting the appliance.

Because the action has no return value, re-running `clear` is the way to flush the table again; change `nodeid` (which forces replacement) or taint the resource to re-trigger the action.


## Example usage

```hcl
# Flush the learned GSLB LDNS RTT entries on the appliance.
resource "citrixadc_gslbldnsentries" "tf_gslbldnsentries" {
}
```

To target a specific cluster node when clearing:

```hcl
resource "citrixadc_gslbldnsentries" "tf_gslbldnsentries" {
  nodeid = 1
}
```


## Argument Reference

* `nodeid` - (Optional) Unique number that identifies the cluster node. This is a GET-only filter on the appliance and is not sent in the `clear` action payload; the provider tracks it so that changing it forces the `clear` action to re-run (it forces resource replacement).


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - A synthetic identifier with the constant value `gslbldnsentries`. The cleared LDNS entries are not a queryable managed object, so this ID is purely a Terraform state handle and does not map to any server-assigned key.
