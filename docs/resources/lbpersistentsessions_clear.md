---
subcategory: "Load Balancing"
---

# Resource: lbpersistentsessions_clear

This resource is used to clear load balancing persistence sessions on the Citrix ADC.


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
