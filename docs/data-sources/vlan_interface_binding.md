---
subcategory: "Network"
---

# Data Source: vlan_interface_binding

The vlan_interface_binding data source allows you to retrieve information about a VLAN to interface binding.

## Example Usage

```terraform
data "citrixadc_vlan_interface_binding" "tf_bind" {
  vlanid = 50
  ifnum  = "1/1"
}

output "vlanid" {
  value = data.citrixadc_vlan_interface_binding.tf_bind.vlanid
}

output "ifnum" {
  value = data.citrixadc_vlan_interface_binding.tf_bind.ifnum
}

output "tagged" {
  value = data.citrixadc_vlan_interface_binding.tf_bind.tagged
}
```

## Argument Reference

* `vlanid` - (Required) Specifies the virtual LAN ID.
* `ifnum` - (Required) The interface to be bound to the VLAN, specified in slot/port notation (for example, 1/3).

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vlan_interface_binding. It is the concatenation of the vlanid and ifnum attributes separated by a comma.
* `ownergroup` - The owner node group in a Cluster for this vlan.
* `tagged` - Whether the interface is an 802.1q tagged interface.
