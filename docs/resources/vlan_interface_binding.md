---
subcategory: "Network"
---

# Resource: vlan\_interface\_binding

The vlan\_interface\_binding resource is used to bind a vlan to an interface.


## Example usage

```hcl
resource "citrixadc_vlan" "tf_vlan" {
    vlanid = 40
    aliasname = "Management VLAN"
}

resource "citrixadc_vlan_interface_binding" "tf_bind" {
    vlanid = citrixadc_vlan.tf_vlan.vlanid
    ifnum = "1/1"
}
```


## Argument Reference

* `ifnum` - (Required) The interface to be bound to the VLAN, specified in slot/port notation (for example, 1/3).
* `tagged` - (Optional) Make the interface an 802.1q tagged interface. Packets sent on this interface on this VLAN have an additional 4-byte 802.1q tag, which identifies the VLAN. To use 802.1q tagging, you must also configure the switch connected to the appliance's interfaces.
* `vlanid` - (Required) Specifies the virtual LAN ID.
* `ownergroup` - (Optional) The owner node group in a Cluster for this vlan.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vlan\_interface\_binding. It is the concatenation of the vlanid and ifnum attributes separated by a comma.


## Import

A vlan\_interface\_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_vlan_interface_binding.tf_bind 40,1/1
```
