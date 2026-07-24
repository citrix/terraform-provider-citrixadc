---
subcategory: "Network"
---

# Data Source: vrid6_channel_binding

The vrid6_channel_binding data source allows you to retrieve information about a channel interface bound to an IPv6 Virtual Router ID (VRID6).


## Example usage

```terraform
data "citrixadc_vrid6_channel_binding" "tf_vrid6_channel_binding" {
  vrid_id = 100
  ifnum   = "LA/1"
}

output "vrid6_channel_binding_vlan" {
  value = data.citrixadc_vrid6_channel_binding.tf_vrid6_channel_binding.vlan
}
```


## Argument Reference

* `vrid_id` - (Required) Integer value that uniquely identifies a VMAC6 address. This is the ID of the parent VRID6.
* `ifnum` - (Required) Channel interface bound to the VMAC6, specified in (slot/port) notation (for example, `LA/1`).


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vrid6_channel_binding. It is the concatenation of the `vrid_id` and `ifnum` attributes in the form `id:<vrid_id>,ifnum:<ifnum>` (each value URL-encoded).
* `flags` - Flags reported for this binding.
* `vlan` - The VLAN in which this VRID resides.
