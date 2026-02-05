---
subcategory: "NS"
---

# Data Source `nsip6`

The nsip6 data source allows you to retrieve information about NetScaler owned IPv6 addresses (NSIP6, SNIP6, VIP6, GSLBsiteIP6, and CLIP6).


## Example usage

```terraform
data "citrixadc_nsip6" "my_nsip6" {
  ipv6address = "2002:db8:100::aa/64"
  td          = 0
}

output "type" {
  value = data.citrixadc_nsip6.my_nsip6.type
}

output "icmp" {
  value = data.citrixadc_nsip6.my_nsip6.icmp
}

output "state" {
  value = data.citrixadc_nsip6.my_nsip6.state
}
```


## Argument Reference

* `ipv6address` - (Required) IPv6 address to retrieve from the Citrix ADC.
* `td` - (Required) Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `advertiseondefaultpartition` - Advertise VIPs from Shared VLAN on Default Partition.
* `decrementhoplimit` - Decrement Hop Limit by 1 when ENABLED. This setting is applicable only for UDP traffic.
* `dynamicrouting` - Allow dynamic routing on this IP address. Specific to Subnet IPv6 (SNIP6) address.
* `ftp` - Allow File Transfer Protocol (FTP) access to this IP address.
* `gui` - Allow graphical user interface (GUI) access to this IP address.
* `hostroute` - Option to push the VIP6 to ZebOS routing table for Kernel route redistribution through dynamic routing protocols.
* `icmp` - Respond to ICMP requests for this IP address.
* `icmpresponse` - Respond to ICMPv6 requests for a Virtual IP (VIP) address on the basis of the states of the virtual servers associated with that VIP.
* `ip6hostrtgw` - IPv6 address of the gateway for the route. If Gateway is not set, VIP uses :: as the gateway.
* `map` - Mapped IPV4 address for the IPV6 address.
* `metric` - Integer value to add to or subtract from the cost of the route advertised for the VIP6 address.
* `mgmtaccess` - Allow access to management applications on this IP address.
* `mptcpadvertise` - If enabled, this IP will be advertised by Citrix ADC to MPTCP enabled clients as part of ADD_ADDR option.
* `nd` - Respond to Neighbor Discovery (ND) requests for this IP address.
* `ndowner` - NdOwner in Cluster for VIPS and Striped SNIPS.
* `networkroute` - Option to push the SNIP6 subnet to ZebOS routing table for Kernel route redistribution through dynamic routing protocol.
* `ospf6lsatype` - Type of LSAs to be used by the IPv6 OSPF protocol, running on the Citrix ADC, for advertising the route for the VIP6 address.
* `ospfarea` - ID of the area in which the Intra-Area-Prefix LSAs are to be advertised for the VIP6 address by the IPv6 OSPF protocol running on the Citrix ADC. When ospfArea is not set, VIP6 is advertised on all areas.
* `ownerdownresponse` - In cluster system, if the owner node is down, whether should it respond to icmp/arp.
* `ownernode` - ID of the cluster node for which you are adding the IP address. Must be used if you want the IP address to be active only on the specific node. Can be configured only through the cluster IP address. Cannot be changed after the IP address is created.
* `restrictaccess` - Block access to nonmanagement applications on this IP address. This option is applicable for MIP6s, SNIP6s, and NSIP6s, and is disabled by default. Nonmanagement applications can run on the underlying Citrix ADC Free BSD operating system.
* `scope` - Scope of the IPv6 address to be created. Cannot be changed after the IP address is created.
* `snmp` - Allow Simple Network Management Protocol (SNMP) access to this IP address.
* `ssh` - Allow secure Shell (SSH) access to this IP address.
* `state` - Enable or disable the IP address.
* `tag` - Tag value for the network/host route associated with this IP.
* `telnet` - Allow Telnet access to this IP address.
* `type` - Type of IP address to be created on the Citrix ADC. Cannot be changed after the IP address is created. Possible values: SNIP6, VIP6, NSIP6, GSLBsiteIP6, CLIP6.
* `vlan` - The VLAN number.
* `vrid6` - A positive integer that uniquely identifies a VMAC address for binding to this VIP address. This binding is used to set up Citrix ADCs in an active-active configuration using VRRP.
* `vserver` - Enable or disable the state of all the virtual servers associated with this VIP6 address.
* `vserverrhilevel` - Advertise or do not advertise the route for the Virtual IP (VIP6) address on the basis of the state of the virtual servers associated with that VIP6.
* `id` - The id of the nsip6. It is a comma-separated value of `ipv6address` and `td`.

