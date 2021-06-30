---
subcategory: "Network"
---

# Resource: vlan

The vlan resource is used to create vlans.


## Example usage

```hcl
resource "citrixadc_vlan" "tf_vlan" {
    vlanid = 40
    aliasname = "Management VLAN"
}
```


## Argument Reference

* `vlanid` - (Requiredc) A positive integer that uniquely identifies a VLAN.
* `aliasname` - (Optional) A name for the VLAN. Must begin with a letter, a number, or the underscore symbol, and can consist of from 1 to 31 letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at sign (@), equals (=), colon (:), and underscore (\_) characters. You should choose a name that helps identify the VLAN. However, you cannot perform any VLAN operation by specifying this name instead of the VLAN ID.
* `dynamicrouting` - (Optional) Enable dynamic routing on this VLAN. Possible values: [ ENABLED, DISABLED ]
* `ipv6dynamicrouting` - (Optional) Enable all IPv6 dynamic routing protocols on this VLAN. Note: For the ENABLED setting to work, you must configure IPv6 dynamic routing protocols from the VTYSH command line. Possible values: [ ENABLED, DISABLED ]
* `mtu` - (Optional) Specifies the maximum transmission unit (MTU), in bytes. The MTU is the largest packet size, excluding 14 bytes of ethernet header and 4 bytes of crc, that can be transmitted and received over this VLAN.
* `sharing` - (Optional) If sharing is enabled, then this vlan can be shared across multiple partitions by binding it to all those partitions. If sharing is disabled, then this vlan can be bound to only one of the partitions. Possible values: [ ENABLED, DISABLED ]


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vlan. It has the same value as the `vlanid` attribute.


## Import

A  can be imported using its `vlanid`, e.g.

```shell
terraform import citrixadc_vlan.tf_vlan 40
```
