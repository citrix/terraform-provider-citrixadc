---
subcategory: "network"
---

# Resource: arp

The arp resource is used to create arp.


## Example usage

```hcl
resource "citrixadc_arp" "tf_arp" {
  ipaddress = "10.222.74.175"
  mac       = "3B:FD:37:27:A1:F8"
  vxlan     =  4
}
```


## Argument Reference

* `ipaddress` - (Required) IP address of the network device that you want to add to the ARP table. Minimum length =  1
* `mac` - (Required) MAC address of the network device.
* `td` - (Optional) Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0. Minimum value =  0 Maximum value =  4094
* `ifnum` - (Optional) Interface through which the network device is accessible. Specify the interface in (slot/port) notation. For example, 1/3.
* `vxlan` - (Optional) ID of the VXLAN on which the IP address of this ARP entry is reachable. Minimum value =  1 Maximum value =  16777215
* `vtep` - (Optional) IP address of the VXLAN tunnel endpoint (VTEP) through which the IP address of this ARP entry is reachable. Minimum length =  1
* `vlan` - (Optional) The VLAN ID through which packets are to be sent after matching the ARP entry. This is a numeric value.
* `ownernode` - (Optional) The owner node for the Arp entry. Minimum value =  0 Maximum value =  31
* `all` - (Optional) Remove all ARP entries from the ARP table of the Citrix ADC.
* `nodeid` - (Optional) Unique number that identifies the cluster node. Minimum value =  0 Maximum value =  31


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the arp. It has the same value as the `ipaddress` attribute.
