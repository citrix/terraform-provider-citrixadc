---
subcategory: "Network"
---

# Data Source: vrid_trackinterface_binding

The vrid_trackinterface_binding data source allows you to retrieve information about a track interface bound to a VRID.


## Example usage

```terraform
data "citrixadc_vrid_trackinterface_binding" "tf_vrid_trackinterface_binding" {
  vrid_id    = 60
  trackifnum = "1/3"
}

output "vrid_trackinterface_binding_flags" {
  value = data.citrixadc_vrid_trackinterface_binding.tf_vrid_trackinterface_binding.flags
}
```


## Argument Reference

* `vrid_id` - (Required) Integer that uniquely identifies the VMAC address. The generic VMAC address is in the form of `00:00:5e:00:01:<VRID>`. For example, if you add a VRID with a value of 60 and bind it to an interface, the resulting VMAC address is `00:00:5e:00:01:3c`, where `3c` is the hexadecimal representation of 60.
* `trackifnum` - (Required) Interface which needs to be tracked for this VRID, specified in (slot/port) notation (for example, `1/3`).


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vrid_trackinterface_binding. It is the concatenation of the `vrid_id` and `trackifnum` values in the form `id:<vrid_id>,trackifnum:<trackifnum>`.
* `flags` - (read-only) Flags reported by the appliance for this binding.
