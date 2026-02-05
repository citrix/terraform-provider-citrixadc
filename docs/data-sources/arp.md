---
subcategory: "Network"
---

# Data Source `arp`

The arp data source allows you to retrieve information about an existing ARP entry.


## Example usage

```terraform
data "citrixadc_arp" "tf_arp" {
  ipaddress = "10.222.74.175"
  ownernode = 0
  td        = 0
}

output "mac" {
  value = data.citrixadc_arp.tf_arp.mac
}

output "vxlan" {
  value = data.citrixadc_arp.tf_arp.vxlan
}
```


## Argument Reference

* `ipaddress` - (Required) IP address of the network device that you want to add to the ARP table.
* `ownernode` - (Required) The owner node for the ARP entry.
* `td` - (Required) Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the arp. It is a composite ID based on the ipaddress, ownernode, and td.
* `all` - Remove all ARP entries from the ARP table of the Citrix ADC.
* `ifnum` - Interface through which the network device is accessible. Specify the interface in (slot/port) notation. For example, 1/3.
* `mac` - MAC address of the network device.
* `nodeid` - Unique number that identifies the cluster node.
* `vlan` - The VLAN ID through which packets are to be sent after matching the ARP entry. This is a numeric value.
* `vtep` - IP address of the VXLAN tunnel endpoint (VTEP) through which the IP address of this ARP entry is reachable.
* `vxlan` - ID of the VXLAN on which the IP address of this ARP entry is reachable.
