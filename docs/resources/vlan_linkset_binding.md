---
subcategory: "Network"
---

# Resource: vlan\_linkset\_binding

Binds a network interface to a VLAN on the Citrix ADC, optionally as an 802.1q
tagged member. Despite the resource name, this binding does **not** reference a
linkset object: the underlying NITRO endpoint binds a single interface (`ifnum`)
to a VLAN identified by `vlanid`. Use it to control which physical or logical
interfaces carry traffic for a given VLAN.


## Example usage

```hcl
resource "citrixadc_vlan" "tf_vlan" {
  vlanid    = 40
  aliasname = "Management VLAN"
}

resource "citrixadc_vlan_linkset_binding" "tf_bind" {
  vlanid = citrixadc_vlan.tf_vlan.vlanid
  ifnum  = "1/3"
  tagged = true
}
```


## Argument Reference

* `vlanid` - (Required) Specifies the virtual LAN ID. Changing this value forces a new resource to be created.
* `ifnum` - (Optional) The interface to be bound to the VLAN, specified in slot/port notation (for example, `1/3`). Changing this value forces a new resource to be created.
* `tagged` - (Optional) Make the interface an 802.1q tagged interface. Packets sent on this interface on this VLAN have an additional 4-byte 802.1q tag, which identifies the VLAN. To use 802.1q tagging, you must also configure the switch connected to the appliance's interfaces. Changing this value forces a new resource to be created.
* `ownergroup` - (Optional) The owner node group in a Cluster for this VLAN. Defaults to `"DEFAULT_NG"`. Changing this value forces a new resource to be created.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vlan\_linkset\_binding. It is a composite key of the form `vlanid:<vlanid>,ifnum:<ifnum>`, where the `ifnum` value is URL-encoded.


## Import

A vlan\_linkset\_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_vlan_linkset_binding.tf_bind vlanid:40,ifnum:1%2F3
```
