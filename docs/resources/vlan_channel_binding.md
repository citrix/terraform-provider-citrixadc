---
subcategory: "Network"
---

# Resource: vlan_channel_binding

The vlan_channel_binding resource is used to bind a channel (link aggregation) interface to a VLAN.


## Example usage

```hcl
resource "citrixadc_vlan" "tf_vlan" {
  vlanid    = 2
  aliasname = "Management VLAN"
}

resource "citrixadc_vlan_channel_binding" "tf_vlan_channel_binding" {
  vlanid = citrixadc_vlan.tf_vlan.vlanid
  ifnum  = "LA/2"
  tagged = false
}
```


## Argument Reference

* `vlanid` - (Required) Specifies the virtual LAN ID.
* `ifnum` - (Required) The interface to be bound to the VLAN, specified in slot/port notation (for example, 1/3).
* `tagged` - (Optional) Make the interface an 802.1q tagged interface. Packets sent on this interface on this VLAN have an additional 4-byte 802.1q tag, which identifies the VLAN. To use 802.1q tagging, you must also configure the switch connected to the appliance's interfaces.
* `ownergroup` - (Optional) The owner node group in a Cluster for this vlan. Defaults to `"DEFAULT_NG"`.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vlan_channel_binding. It is the concatenation of the `vlanid` and `ifnum` attributes separated by a comma.


## Import

A vlan_channel_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_vlan_channel_binding.tf_vlan_channel_binding 2,LA/2
```
