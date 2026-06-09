---
subcategory: "Load Balancing"
---

# Resource: lbpersistentsessions

The lbpersistentsessions resource performs the imperative `clear` action on the Citrix ADC, flushing active load balancing persistence sessions. Use it to forcibly drop client-to-server affinity so that subsequent requests are re-balanced — for example after changing back-end membership or during maintenance. Omit `vserver` to flush persistence sessions for all virtual servers, or set `vserver` (optionally with `persistenceparameter`) to flush a targeted subset.

This is an action-only resource. Applying it invokes `clear` once; there is no managed object to read back. The Read and Delete operations are state-only no-ops (Delete simply removes the entry from Terraform state), and there is no Update operation — every argument forces resource replacement, which re-runs the `clear` action. Because nothing is persisted on the ADC as a queryable object, there is nothing to import.


## Example usage

### Flush persistence sessions for all virtual servers

```hcl
resource "citrixadc_lbpersistentsessions" "tf_lbpersistentsessions" {
}
```

### Flush persistence sessions for a specific virtual server

```hcl
resource "citrixadc_lbpersistentsessions" "tf_lbpersistentsessions" {
  vserver              = "lbvserver1"
  persistenceparameter = "192.0.2.10"
}
```


## Argument Reference

* `vserver` - (Optional) The name of the virtual server whose persistence sessions are to be flushed. If omitted, persistence sessions for all virtual servers are flushed. Changing this value forces the `clear` action to re-run (resource replacement).
* `persistenceparameter` - (Optional) The persistence parameter whose persistence sessions are to be flushed (for example, a source IP or rule value). Changing this value forces the `clear` action to re-run (resource replacement).
* `nodeid` - (Optional) Unique number that identifies the cluster node. Note: `nodeid` is a GET-only cluster filter and is **not** included in the `clear` action payload; it is primarily useful on the corresponding `lbpersistentsessions` data source. Changing this value forces resource replacement.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - A synthetic identifier with the fixed value `lbpersistentsessions`. It is purely a Terraform state handle for this action; it is not a server-assigned key and cannot be used to look the resource up on the ADC.
