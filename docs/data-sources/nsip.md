---
subcategory: "NS"
---

# Data Source `nsip`

The nsip data source allows you to retrieve information about NetScaler owned IP addresses (NSIP, SNIP, VIP, GSLBsiteIP, and CLIP).


## Example usage

```terraform
data "citrixadc_nsip" "my_nsip" {
  ipaddress = "10.222.74.149"
  td        = 0
}

output "netmask" {
  value = data.citrixadc_nsip.my_nsip.netmask
}

output "type" {
  value = data.citrixadc_nsip.my_nsip.type
}
```


## Argument Reference

* `ipaddress` - (Required) IPv4 address to retrieve from the Citrix ADC.
* `td` - (Optional) Integer value that uniquely identifies the traffic domain in which you want to configure the entity. Default is 0.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `advertiseondefaultpartition` - Advertise VIPs from Shared VLAN on Default Partition.
* `arp` - Respond to ARP requests for this IP address.
* `arpowner` - The arp owner in a Cluster for this IP address. It can vary from 0 to 31.
* `arpresponse` - Respond to ARP requests for a Virtual IP (VIP) address on the basis of the states of the virtual servers associated with that VIP.
* `bgp` - Use this option to enable or disable BGP on this IP address for the entity.
* `decrementttl` - Decrement TTL by 1 when ENABLED. This setting is applicable only for UDP traffic.
* `dynamicrouting` - Allow dynamic routing on this IP address. Specific to Subnet IP (SNIP) address.
* `ftp` - Allow File Transfer Protocol (FTP) access to this IP address.
* `gui` - Allow graphical user interface (GUI) access to this IP address.
* `hostroute` - Option to push the VIP to ZebOS routing table for Kernel route redistribution through dynamic routing protocols.
* `hostrtgw` - IP address of the gateway of the route for this VIP address.
* `icmp` - Respond to ICMP requests for this IP address.
* `icmpresponse` - Respond to ICMP requests for a Virtual IP (VIP) address on the basis of the states of the virtual servers associated with that VIP.
* `metric` - Integer value to add to or subtract from the cost of the route advertised for the VIP address.
* `mgmtaccess` - Allow access to management applications on this IP address.
* `mptcpadvertise` - If enabled, this IP will be advertised by Citrix ADC to MPTCP enabled clients as part of ADD_ADDR option.
* `netmask` - Subnet mask associated with the IP address.
* `networkroute` - Option to push the SNIP subnet to ZebOS routing table for Kernel route redistribution through dynamic routing protocol.
* `ospf` - Use this option to enable or disable OSPF on this IP address for the entity.
* `ospfarea` - ID of the area in which the type 1 link-state advertisements (LSAs) are to be advertised for this virtual IP (VIP) address by the OSPF protocol running on the Citrix ADC.
* `ospflsatype` - Type of LSAs to be used by the OSPF protocol, running on the Citrix ADC, for advertising the route for this VIP address.
* `ownerdownresponse` - In a cluster setup, if the owner node is down, whether to allow a different node to respond to the ping or ARP request.
* `ownernode` - The owner node in a Cluster for this IP address. Owner node can vary from 0 to 31.
* `restrictaccess` - Block access to nonmanagement applications on this IP. This option is applicable for MIPs, SNIPs, and NSIP, and is disabled by default.
* `rip` - Use this option to enable or disable RIP on this IP address for the entity.
* `snmp` - Allow Simple Network Management Protocol (SNMP) access to this IP address.
* `ssh` - Allow secure shell (SSH) access to this IP address.
* `state` - Enable or disable the IP address.
* `tag` - Tag value for the network/host route associated with this IP.
* `telnet` - Allow Telnet access to this IP address.
* `type` - Type of the IP address to create on the Citrix ADC. Possible values: SNIP, VIP, NSIP, GSLBsiteIP, CLIP.
* `vrid` - A positive integer that uniquely identifies a VMAC address for binding to this VIP address. This binding is used to set up Citrix ADCs in an active-active configuration using VRRP.
* `vserver` - Use this option to set (enable or disable) the virtual server attribute for this IP address.
* `vserverrhilevel` - Advertise the route for the Virtual IP (VIP) address on the basis of the state of the virtual servers associated with that VIP.
* `id` - The id of the nsip. It has the same value as the `ipaddress` attribute.
