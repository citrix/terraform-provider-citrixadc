---
subcategory: "Network"
---

# Resource: vrid_channel_binding

Binds a link-aggregation channel (LA channel) to a Virtual Router ID (VRID) so that the channel participates in VRRP for that virtual router. Binding a channel associates the VRID's virtual MAC address (VMAC, of the form `00:00:5e:00:01:<VRID>`) with the channel, allowing the channel to send and receive traffic for the virtual router during a VRRP failover.

Because a binding cannot be modified in place, every configurable attribute forces replacement: changing any attribute destroys the existing binding and creates a new one.


## Example usage

```hcl
resource "citrixadc_vrid" "tf_vrid" {
  id = 60
}

resource "citrixadc_vrid_channel_binding" "tf_vrid_channel_binding" {
  vrid_id = citrixadc_vrid.tf_vrid.id
  ifnum   = "LA/1"
}
```


## Argument Reference

* `vrid_id` - (Required) Integer that uniquely identifies the VMAC address. The generic VMAC address is in the form of `00:00:5e:00:01:<VRID>`. For example, if you add a VRID with a value of 60 and bind it to an interface, the resulting VMAC address is `00:00:5e:00:01:3c`, where `3c` is the hexadecimal representation of 60. This is the identifier of the parent `citrixadc_vrid` resource (an integer in the range 1-255). Changing this attribute forces a new resource to be created.
* `ifnum` - (Required) Channel to bind to the VMAC, specified in `LA/x` notation (for example, `LA/1`). Use spaces to separate multiple entries. Changing this attribute forces a new resource to be created.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vrid_channel_binding. It is the concatenation of the `vrid_id` and `ifnum` values in the form `id:<vrid_id>,ifnum:<ifnum>` (the `ifnum` value is URL-encoded because channel names may contain a `/`).


## Import

A vrid_channel_binding can be imported using the composite id in the form `id:<vrid_id>,ifnum:<ifnum>`, e.g.

```shell
terraform import citrixadc_vrid_channel_binding.tf_vrid_channel_binding id:60,ifnum:LA%2F1
```
