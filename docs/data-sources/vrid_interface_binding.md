---
subcategory: "Network"
---

# Data Source: vrid_interface_binding

The vrid_interface_binding data source allows you to retrieve information about a network interface that is bound to a Virtual Router ID (VRID) for VRRP, including the read-only VLAN and flags reported by the appliance for that binding.


## Example usage

```terraform
data "citrixadc_vrid_interface_binding" "tf_vrid_interface_binding" {
  vrid_id = 60
  ifnum   = "1/2"
}

output "vrid_interface_binding_vlan" {
  value = data.citrixadc_vrid_interface_binding.tf_vrid_interface_binding.vlan
}
```


## Argument Reference

* `vrid_id` - (Required) Integer that uniquely identifies the VMAC address. The generic VMAC address is in the form of `00:00:5e:00:01:<VRID>`. For example, if you add a VRID with a value of 60 and bind it to an interface, the resulting VMAC address is `00:00:5e:00:01:3c`, where `3c` is the hexadecimal representation of 60.
* `ifnum` - (Required) Interface to bind to the VMAC, specified in (slot/port) notation (for example, `1/2`). Use spaces to separate multiple entries.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vrid_interface_binding. It is the concatenation of the `vrid_id` and `ifnum` values in the form `id:<vrid_id>,ifnum:<ifnum>`.
* `flags` - (read-only) Flags reported by the appliance for this binding.
* `vlan` - (read-only) The VLAN in which this VRID resides.
