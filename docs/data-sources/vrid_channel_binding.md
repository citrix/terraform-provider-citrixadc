---
subcategory: "Network"
---

# Data Source: vrid_channel_binding

The vrid_channel_binding data source allows you to retrieve information about a channel bound to a VRID.


## Example usage

```terraform
data "citrixadc_vrid_channel_binding" "tf_vrid_channel_binding" {
  vrid_id = 60
  ifnum   = "LA/1"
}

output "vrid_channel_binding_vlan" {
  value = data.citrixadc_vrid_channel_binding.tf_vrid_channel_binding.vlan
}
```


## Argument Reference

* `vrid_id` - (Required) Integer that uniquely identifies the VMAC address. The generic VMAC address is in the form of `00:00:5e:00:01:<VRID>`. For example, if you add a VRID with a value of 60 and bind it to an interface, the resulting VMAC address is `00:00:5e:00:01:3c`, where `3c` is the hexadecimal representation of 60.
* `ifnum` - (Required) Channel to bind to the VMAC, specified in `LA/x` notation (for example, `LA/1`). Use spaces to separate multiple entries.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vrid_channel_binding. It is the concatenation of the `vrid_id` and `ifnum` values in the form `id:<vrid_id>,ifnum:<ifnum>`.
* `flags` - (read-only) Flags reported by the appliance for this binding.
* `vlan` - (read-only) The VLAN in which this VRID resides.
