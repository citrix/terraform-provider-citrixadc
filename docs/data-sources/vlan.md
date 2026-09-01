---
subcategory: "Network"
---

# Data Source: vlan

The vlan data source allows you to retrieve information about VLAN configurations.

## Example usage

```terraform
data "citrixadc_vlan" "tf_vlan" {
  vlanid = 40
}

output "aliasname" {
  value = data.citrixadc_vlan.tf_vlan.aliasname
}

output "mtu" {
  value = data.citrixadc_vlan.tf_vlan.mtu
}
```

## Argument Reference

* `vlanid` - (Required) A positive integer that uniquely identifies a VLAN.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `aliasname` - A name for the VLAN. Must begin with a letter, a number, or the underscore symbol, and can consist of from 1 to 31 letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at sign (@), equals (=), colon (:), and underscore (_) characters. You should choose a name that helps identify the VLAN. However, you cannot perform any VLAN operation by specifying this name instead of the VLAN ID.
* `dynamicrouting` - Enable dynamic routing on this VLAN.
* `id` - The id of the vlan. It has the same value as the `vlanid` attribute.
* `ipv6dynamicrouting` - Enable all IPv6 dynamic routing protocols on this VLAN. Note: For the ENABLED setting to work, you must configure IPv6 dynamic routing protocols from the VTYSH command line.
* `mtu` - Specifies the maximum transmission unit (MTU), in bytes. The MTU is the largest packet size, excluding 14 bytes of ethernet header and 4 bytes of crc, that can be transmitted and received over this VLAN.
* `sharing` - If sharing is enabled, then this vlan can be shared across multiple partitions by binding it to all those partitions. If sharing is disabled, then this vlan can be bound to only one of the partitions.

### Read-only vlan metadata

These attributes are returned by the appliance on a GET (they are not configurable on the `citrixadc_vlan` resource). They are GET-only/Computed, and any attribute the appliance does not return is `null`.

* `linklocalipv6addr` - The link-local IP address assigned to the VLAN.
* `rnat` - Temporary flag used for internal purpose.
* `portbitmap` - Member interfaces of this vlan.
* `lsbitmap` - Member linksets of this vlan.
* `tagbitmap` - Tagged members of this vlan.
* `lstagbitmap` - Tagged linksets of this vlan.
* `ifaces` - Names of all member interfaces of this vlan.
* `tagifaces` - Names of all tagged member interfaces of this vlan.
* `ifnum` - The interface bound to the VLAN, specified in slot/port notation (for example, 1/3).
* `tagged` - Whether the interface is an 802.1q tagged interface.
* `vlantd` - Traffic domain associated with vlan.
* `sdxvlan` - SDX vlan (for example `YES`, `NO`).
* `partitionname` - Name of the Partition to which this vlan is bound.
* `vxlan` - The VXLAN that extends this vlan.

## Import

A vlan can be imported using its vlanid, e.g.

```shell
terraform import citrixadc_vlan.tf_vlan 40
```
