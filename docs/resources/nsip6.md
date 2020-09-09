---
subcategory: "NS"
---

# Resource: nsip6

The nsip6 resource is used to create nsip, snip and vip ipv6 addresses for ADC.


## Example usage

```hcl
resource "citrixadc_nsip6" "tf_nsip6" {
    ipv6address = "2002:db8:100::ff/64"
    type = "VIP"
    icmp = "DISABLED"
}
```


## Argument Reference

* `ipv6address` - (Optional) IPv6 address to create on the Citrix ADC.
* `scope` - (Optional) Scope of the IPv6 address to be created. Cannot be changed after the IP address is created. Possible values: [ global, link-local ]
* `type` - (Optional) Type of IP address to be created on the Citrix ADC. Cannot be changed after the IP address is created. Possible values: [ NSIP, VIP, SNIP, GSLBsiteIP, ADNSsvcIP, RADIUSListenersvcIP, CLIP ]
* `vlan` - (Optional) The VLAN number.
* `nd` - (Optional) Respond to Neighbor Discovery (ND) requests for this IP address. Possible values: [ ENABLED, DISABLED ]
* `icmp` - (Optional) Respond to ICMP requests for this IP address. Possible values: [ ENABLED, DISABLED ]
* `vserver` - (Optional) Enable or disable the state of all the virtual servers associated with this VIP6 address. Possible values: [ ENABLED, DISABLED ]
* `telnet` - (Optional) Allow Telnet access to this IP address. Possible values: [ ENABLED, DISABLED ]
* `ftp` - (Optional) Allow File Transfer Protocol (FTP) access to this IP address. Possible values: [ ENABLED, DISABLED ]
* `gui` - (Optional) Allow graphical user interface (GUI) access to this IP address. Possible values: [ ENABLED, SECUREONLY, DISABLED ]
* `ssh` - (Optional) Allow secure Shell (SSH) access to this IP address. Possible values: [ ENABLED, DISABLED ]
* `snmp` - (Optional) Allow Simple Network Management Protocol (SNMP) access to this IP address. Possible values: [ ENABLED, DISABLED ]
* `mgmtaccess` - (Optional) Allow access to management applications on this IP address. Possible values: [ ENABLED, DISABLED ]
* `restrictaccess` - (Optional) Block access to nonmanagement applications on this IP address. This option is applicable forMIP6s, SNIP6s, and NSIP6s, and is disabled by default. Nonmanagement applications can run on the underlying Citrix ADC Free BSD operating system. Possible values: [ ENABLED, DISABLED ]
* `dynamicrouting` - (Optional) Allow dynamic routing on this IP address. Specific to Subnet IPv6 (SNIP6) address. Possible values: [ ENABLED, DISABLED ]
* `decrementhoplimit` - (Optional) Decrement Hop Limit by 1 when ENABLED.This setting is applicable only for UDP traffic. Possible values: [ ENABLED, DISABLED ]
* `hostroute` - (Optional) Option to push the VIP6 to ZebOS routing table for Kernel route redistribution through dynamic routing protocols. Possible values: [ ENABLED, DISABLED ]
* `advertiseondefaultpartition` - (Optional) Advertise VIPs from Shared VLAN on Default Partition. Possible values: [ ENABLED, DISABLED ]
* `networkroute` - (Optional) Option to push the SNIP6 subnet to ZebOS routing table for Kernel route redistribution through dynamic routing protocol. Possible values: [ ENABLED, DISABLED ]
* `tag` - (Optional) Tag value for the network/host route associated with this IP.
* `ip6hostrtgw` - (Optional) IPv6 address of the gateway for the route. If Gateway is not set, VIP uses :: as the gateway.
* `metric` - (Optional) Integer value to add to or subtract from the cost of the route advertised for the VIP6 address.
* `vserverrhilevel` - (Optional) Advertise or do not advertise the route for the Virtual IP (VIP6) address on the basis of the state of the virtual servers associated with that VIP6. * NONE - Advertise the route for the VIP6 address, irrespective of the state of the virtual servers associated with the address. * ONE VSERVER - Advertise the route for the VIP6 address if at least one of the associated virtual servers is in UP state. * ALL VSERVER - Advertise the route for the VIP6 address if all of the associated virtual servers are in UP state. * VSVR_CNTRLD.   Advertise the route for the VIP address according to the  RHIstate (RHI STATE) parameter setting on all the associated virtual servers of the VIP address along with their states. When Vserver RHI Level (RHI) parameter is set to VSVR_CNTRLD, the following are different RHI behaviors for the VIP address on the basis of RHIstate (RHI STATE) settings on the virtual servers associated with the VIP address: * If you set RHI STATE to PASSIVE on all virtual servers, the Citrix ADC always advertises the route for the VIP address. * If you set RHI STATE to ACTIVE on all virtual servers, the Citrix ADC advertises the route for the VIP address if at least one of the associated virtual servers is in UP state. \*If you set RHI STATE to ACTIVE on some and PASSIVE on others, the Citrix ADC advertises the route for the VIP address if at least one of the associated virtual servers, whose RHI STATE set to ACTIVE, is in UP state. Possible values: [ ONE_VSERVER, ALL_VSERVERS, NONE, VSVR_CNTRLD ]
* `ospf6lsatype` - (Optional) Type of LSAs to be used by the IPv6 OSPF protocol, running on the Citrix ADC, for advertising the route for the VIP6 address. Possible values: [ INTRA_AREA, EXTERNAL ]
* `ospfarea` - (Optional) ID of the area in which the Intra-Area-Prefix LSAs are to be advertised for the VIP6 address by the IPv6 OSPF protocol running on the Citrix ADC. When ospfArea is not set, VIP6 is advertised on all areas.
* `state` - (Optional) Enable or disable the IP address. Possible values: [ DISABLED, ENABLED ]
* `map` - (Optional) Mapped IPV4 address for the IPV6 address.
* `vrid6` - (Optional) A positive integer that uniquely identifies a VMAC address for binding to this VIP address. This binding is used to set up Citrix ADCs in an active-active configuration using VRRP.
* `ownernode` - (Optional) ID of the cluster node for which you are adding the IP address. Must be used if you want the IP address to be active only on the specific node. Can be configured only through the cluster IP address. Cannot be changed after the IP address is created.
* `ownerdownresponse` - (Optional) in cluster system, if the owner node is down, whether should it respond to icmp/arp. Possible values: [ YES, NO ]
* `td` - (Optional) Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nsip6. It has the same value as the `ipv6address` attribute.


## Import

A nsip6 can be imported using its ipv6address, e.g.

```shell
terraform import citrixadc_nsip6.tf_nsip6 2002:db8:100::ff/64
```
