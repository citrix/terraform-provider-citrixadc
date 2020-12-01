---
subcategory: "NS"
---

# Resource: nsip

The nsip resource is used to create nsip, snip and vip, ipv4 addresses for the ADC.


## Example usage

```hcl
resource "citrixadc_nsip" "tf_nsip" {
    ipaddress = "192.168.2.55"
    type = "VIP"
    netmask = "255.255.255.0"
    icmp = "ENABLED"
}
```


## Argument Reference

* `ipaddress` - (Optional) IPv4 address to create on the Citrix ADC. Cannot be changed after the IP address is created.
* `netmask` - (Optional) Subnet mask associated with the IP address.
* `type` - (Optional) Type of the IP address to create on the Citrix ADC. Cannot be changed after the IP address is created. The following are the different types of Citrix ADC owned IP addresses: * A Subnet IP (SNIP) address is used by the Citrix ADC to communicate with the servers. The Citrix ADC also uses the subnet IP address when generating its own packets, such as packets related to dynamic routing protocols, or to send monitor probes to check the health of the servers. * A Virtual IP (VIP) address is the IP address associated with a virtual server. It is the IP address to which clients connect. An appliance managing a wide range of traffic may have many VIPs configured. Some of the attributes of the VIP address are customized to meet the requirements of the virtual server. * A GSLB site IP (GSLBIP) address is associated with a GSLB site. It is not mandatory to specify a GSLBIP address when you initially configure the Citrix ADC. A GSLBIP address is used only when you create a GSLB site. * A Cluster IP (CLIP) address is the management address of the cluster. All cluster configurations must be performed by accessing the cluster through this IP address. Possible values: [ SNIP, VIP, NSIP, GSLBsiteIP, CLIP ]
* `arp` - (Optional) Respond to ARP requests for this IP address. Possible values: [ ENABLED, DISABLED ]
* `icmp` - (Optional) Respond to ICMP requests for this IP address. Possible values: [ ENABLED, DISABLED ]
* `vserver` - (Optional) Use this option to set (enable or disable) the virtual server attribute for this IP address. Possible values: [ ENABLED, DISABLED ]
* `telnet` - (Optional) Allow Telnet access to this IP address. Possible values: [ ENABLED, DISABLED ]
* `ftp` - (Optional) Allow File Transfer Protocol (FTP) access to this IP address. Possible values: [ ENABLED, DISABLED ]
* `gui` - (Optional) Allow graphical user interface (GUI) access to this IP address. Possible values: [ ENABLED, SECUREONLY, DISABLED ]
* `ssh` - (Optional) Allow secure shell (SSH) access to this IP address. Possible values: [ ENABLED, DISABLED ]
* `snmp` - (Optional) Allow Simple Network Management Protocol (SNMP) access to this IP address. Possible values: [ ENABLED, DISABLED ]
* `mgmtaccess` - (Optional) Allow access to management applications on this IP address. Possible values: [ ENABLED, DISABLED ]
* `restrictaccess` - (Optional) Block access to nonmanagement applications on this IP. This option is applicable for MIPs, SNIPs, and NSIP, and is disabled by default. Nonmanagement applications can run on the underlying Citrix ADC Free BSD operating system. Possible values: [ ENABLED, DISABLED ]
* `dynamicrouting` - (Optional) Allow dynamic routing on this IP address. Specific to Subnet IP (SNIP) address. Possible values: [ ENABLED, DISABLED ]
* `decrementttl` - (Optional) Decrement TTL by 1 when ENABLED.This setting is applicable only for UDP traffic. Possible values: [ ENABLED, DISABLED ]
* `ospf` - (Optional) Use this option to enable or disable OSPF on this IP address for the entity. Possible values: [ ENABLED, DISABLED ]
* `bgp` - (Optional) Use this option to enable or disable BGP on this IP address for the entity. Possible values: [ ENABLED, DISABLED ]
* `rip` - (Optional) Use this option to enable or disable RIP on this IP address for the entity. Possible values: [ ENABLED, DISABLED ]
* `hostroute` - (Optional) Option to push the VIP to ZebOS routing table for Kernel route redistribution through dynamic routing protocols. Possible values: [ ENABLED, DISABLED ]
* `advertiseondefaultpartition` - (Optional) Advertise VIPs from Shared VLAN on Default Partition. Possible values: [ ENABLED, DISABLED ]
* `networkroute` - (Optional) Option to push the SNIP subnet to ZebOS routing table for Kernel route redistribution through dynamic routing protocol. Possible values: [ ENABLED, DISABLED ]
* `tag` - (Optional) Tag value for the network/host route associated with this IP.
* `hostrtgw` - (Optional) IP address of the gateway of the route for this VIP address.
* `metric` - (Optional) Integer value to add to or subtract from the cost of the route advertised for the VIP address.
* `vserverrhilevel` - (Optional) Advertise the route for the Virtual IP (VIP) address on the basis of the state of the virtual servers associated with that VIP. * NONE - Advertise the route for the VIP address, regardless of the state of the virtual servers associated with the address. * ONE VSERVER - Advertise the route for the VIP address if at least one of the associated virtual servers is in UP state. * ALL VSERVER - Advertise the route for the VIP address if all of the associated virtual servers are in UP state. * VSVR_CNTRLD - Advertise the route for the VIP address according to the  RHIstate (RHI STATE) parameter setting on all the associated virtual servers of the VIP address along with their states. When Vserver RHI Level (RHI) parameter is set to VSVR_CNTRLD, the following are different RHI behaviors for the VIP address on the basis of RHIstate (RHI STATE) settings on the virtual servers associated with the VIP address: * If you set RHI STATE to PASSIVE on all virtual servers, the Citrix ADC always advertises the route for the VIP address. * If you set RHI STATE to ACTIVE on all virtual servers, the Citrix ADC advertises the route for the VIP address if at least one of the associated virtual servers is in UP state. \*If you set RHI STATE to ACTIVE on some and PASSIVE on others, the Citrix ADC advertises the route for the VIP address if at least one of the associated virtual servers, whose RHI STATE set to ACTIVE, is in UP state. Possible values: [ ONE_VSERVER, ALL_VSERVERS, NONE, VSVR_CNTRLD ]
* `ospflsatype` - (Optional) Type of LSAs to be used by the OSPF protocol, running on the Citrix ADC, for advertising the route for this VIP address. Possible values: [ TYPE1, TYPE5 ]
* `ospfarea` - (Optional) ID of the area in which the type1 link-state advertisements (LSAs) are to be advertised for this virtual IP (VIP)  address by the OSPF protocol running on the Citrix ADC.  When this parameter is not set, the VIP is advertised on all areas.
* `state` - (Optional) Enable or disable the IP address. Possible values: [ ENABLED, DISABLED ]
* `vrid` - (Optional) A positive integer that uniquely identifies a VMAC address for binding to this VIP address. This binding is used to set up Citrix ADCs in an active-active configuration using VRRP.
* `icmpresponse` - (Optional) Respond to ICMP requests for a Virtual IP (VIP) address on the basis of the states of the virtual servers associated with that VIP. Available settings function as follows: * NONE - The Citrix ADC responds to any ICMP request for the VIP address, irrespective of the states of the virtual servers associated with the address. * ONE VSERVER - The Citrix ADC responds to any ICMP request for the VIP address if at least one of the associated virtual servers is in UP state. * ALL VSERVER - The Citrix ADC responds to any ICMP request for the VIP address if all of the associated virtual servers are in UP state. * VSVR_CNTRLD - The behavior depends on the ICMP VSERVER RESPONSE setting on all the associated virtual servers. The following settings can be made for the ICMP VSERVER RESPONSE parameter on a virtual server: * If you set ICMP VSERVER RESPONSE to PASSIVE on all virtual servers, Citrix ADC always responds. * If you set ICMP VSERVER RESPONSE to ACTIVE on all virtual servers, Citrix ADC responds if even one virtual server is UP. * When you set ICMP VSERVER RESPONSE to ACTIVE on some and PASSIVE on others, Citrix ADC responds if even one virtual server set to ACTIVE is UP. Possible values: [ NONE, ONE_VSERVER, ALL_VSERVERS, VSVR_CNTRLD ]
* `ownernode` - (Optional) The owner node in a Cluster for this IP address. Owner node can vary from 0 to 31. If ownernode is not specified then the IP is treated as Striped IP.
* `arpresponse` - (Optional) Respond to ARP requests for a Virtual IP (VIP) address on the basis of the states of the virtual servers associated with that VIP. Available settings function as follows: * NONE - The Citrix ADC responds to any ARP request for the VIP address, irrespective of the states of the virtual servers associated with the address. * ONE VSERVER - The Citrix ADC responds to any ARP request for the VIP address if at least one of the associated virtual servers is in UP state. * ALL VSERVER - The Citrix ADC responds to any ARP request for the VIP address if all of the associated virtual servers are in UP state. Possible values: [ NONE, ONE_VSERVER, ALL_VSERVERS ]
* `ownerdownresponse` - (Optional) in cluster system, if the owner node is down, whether should it respond to icmp/arp. Possible values: [ YES, NO ]
* `td` - (Optional) Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0. TD id 4095 is used reserved for  LSN use  .


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nsip. It has the same value as the `ipaddress` attribute.


## Import

A nsip can be imported using its name, e.g.

```shell
terraform import citrixadc_nsip.tf_nsip 192.168.2.55
```
