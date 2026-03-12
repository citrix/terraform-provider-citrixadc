---
subcategory: "Network"
---

# Data Source: vlan_channel_binding

The vlan_channel_binding data source allows you to retrieve information about a specific VLAN channel binding.

## Example Usage

```terraform
data "citrixadc_vlan_channel_binding" "tf_vlan_channel_binding" {
  vlanid = 2
  ifnum  = "LA/3"
}

output "id" {
  value = data.citrixadc_vlan_channel_binding.tf_vlan_channel_binding.id
}

output "tagged" {
  value = data.citrixadc_vlan_channel_binding.tf_vlan_channel_binding.tagged
}

output "ownergroup" {
  value = data.citrixadc_vlan_channel_binding.tf_vlan_channel_binding.ownergroup
}
```

## Argument Reference

* `vlanid` - (Required) Specifies the virtual LAN ID.
* `ifnum` - (Required) The interface to be bound to the VLAN, specified in slot/port notation (for example, 1/3).

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vlan_channel_binding. It has the same value as the `vlanid` and `ifnum` attributes separated by a comma.
* `ownergroup` - The owner node group in a Cluster for this vlan.
* `tagged` - Indicates whether the interface is an 802.1q tagged interface. Packets sent on this interface on this VLAN have an additional 4-byte 802.1q tag, which identifies the VLAN.
* `ownergroup` - The owner node group in a Cluster for this vlan. Use this to filter the binding by ownergroup.
* `tagged` - Make the interface an 802.1q tagged interface. Use this to filter the binding by tagged status.
