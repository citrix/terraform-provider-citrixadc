---
subcategory: "Network"
---

# Data Source: vrid6_trackinterface_binding

The vrid6_trackinterface_binding data source allows you to retrieve information about a tracked interface bound to an IPv6 Virtual Router ID (VRID6), including the read-only flags reported by the Citrix ADC.


## Example usage

```terraform
data "citrixadc_vrid6_trackinterface_binding" "tf_vrid6_trackinterface_binding" {
  vrid_id    = 100
  trackifnum = "1/3"
}

output "vrid6_trackinterface_binding_flags" {
  value = data.citrixadc_vrid6_trackinterface_binding.tf_vrid6_trackinterface_binding.flags
}
```


## Argument Reference

* `vrid_id` - (Required) Integer value that uniquely identifies a VMAC6 address. This is the ID of the parent VRID6.
* `trackifnum` - (Required) Interface which needs to be tracked for this vrID, specified in (slot/port) notation (for example, `1/3`).


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vrid6_trackinterface_binding. It is the concatenation of the `vrid_id` and `trackifnum` attributes in the form `id:<vrid_id>,trackifnum:<trackifnum>` (each value URL-encoded).
* `flags` - Flags reported for this binding.
