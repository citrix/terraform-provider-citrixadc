---
subcategory: "Load Balancing"
---

# Resource: lbpersistentsessions_clear

The lbpersistentsessions_clear resource performs the imperative `clear` action on the Citrix ADC, flushing active load balancing persistence sessions. Use it to forcibly drop client-to-server affinity so that subsequent requests are re-balanced — for example after changing back-end membership or during maintenance. Omit `vserver` to flush persistence sessions for all virtual servers, or set `vserver` (optionally with `persistenceparameter`) to flush a targeted subset.

This is an action resource: applying it performs the `clear`; it does not manage a persistent object, so re-applying re-runs the action. Every argument forces resource replacement, which re-runs the `clear` action.


## Example usage

### Flush persistence sessions for all virtual servers

With no arguments, persistence sessions for all virtual servers are flushed.

```hcl
resource "citrixadc_lbpersistentsessions_clear" "tf_lbpersistentsessions_clear" {
}
```

### Flush persistence sessions for a specific virtual server

```hcl
resource "citrixadc_lbpersistentsessions_clear" "tf_lbpersistentsessions_clear" {
  vserver              = "lbvserver1"
  persistenceparameter = "192.0.2.10"
}
```


## Argument Reference

* `vserver` - (Optional) The name of the virtual server whose persistence sessions are to be flushed. If omitted, persistence sessions for all virtual servers are flushed. Changing this value forces the `clear` action to re-run (resource replacement).
* `persistenceparameter` - (Optional) The persistence parameter whose persistence sessions are to be flushed (for example, a source IP or rule value). Changing this value forces the `clear` action to re-run (resource replacement).
* `nodeid` - (Optional) Unique number that identifies the cluster node. Changing this value forces resource replacement.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the lbpersistentsessions_clear resource. It is set to `lbpersistentsessions_clear`.
