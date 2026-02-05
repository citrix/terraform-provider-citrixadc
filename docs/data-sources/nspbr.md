---
subcategory: "NS"
---

# Data Source `nspbr`

The nspbr data source allows you to retrieve information about NetScaler Policy-Based Routes (PBR).


## Example usage

```terraform
resource "citrixadc_nspbr" "my_pbr" {
  name     = "my_pbr"
  action   = "DENY"
  srcip    = true
  srcipop  = "="
  srcipval = "192.0.2.0-192.0.2.255"
  destip   = true
  destipop = "="
  destipval = "203.0.113.0-203.0.113.255"
  priority = 100
}

data "citrixadc_nspbr" "my_pbr_data" {
  name = citrixadc_nspbr.my_pbr.name
}

output "pbr_action" {
  value = data.citrixadc_nspbr.my_pbr_data.action
}

output "pbr_priority" {
  value = data.citrixadc_nspbr.my_pbr_data.priority
}
```


## Argument Reference

* `name` - (Required) Name for the PBR. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the PBR is created.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `action` - Action to perform on the outgoing IPv4 packets that match the PBR. Available settings function as follows: ALLOW - The Citrix ADC sends the packet to the designated next-hop router. DENY - The Citrix ADC applies the routing table for normal destination-based routing.
* `destip` - IP address or range of IP addresses to match against the destination IP address of an outgoing IPv4 packet.  In the command line interface, separate the range with a hyphen.
* `destipop` - Either the equals (=) or does not equal (!=) logical operator.
* `destipval` - IP address or range of IP addresses to match against the destination IP address of an outgoing IPv4 packet. In the command line interface, separate the range with a hyphen.
* `destport` - Port number or range of port numbers to match against the destination port number of an outgoing IPv4 packet. In the command line interface, separate the range with a hyphen. Note: The destination port can be specified only for TCP and UDP protocols.
* `destportop` - Either the equals (=) or does not equal (!=) logical operator.
* `destportval` - Port number or range of port numbers to match against the destination port number of an outgoing IPv4 packet. In the command line interface, separate the range with a hyphen.
* `detail` - To get a detailed view.
* `interface` - ID of an interface. The Citrix ADC compares the PBR only to the outgoing packets on the specified interface. If you do not specify any value, the appliance compares the PBR to the outgoing packets on all interfaces.
* `iptunnel` - The Tunnel name.
* `iptunnelname` - The iptunnel name where packets need to be forwarded upon.
* `monitor` - The name of the monitor.(Can be only of type ping or ARP ).
* `msr` - Monitor the route specified byte Next Hop parameter. This parameter is not applicable if you specify a link load balancing (LLB) virtual server name with the Next Hop parameter.
* `nexthop` - IP address of the next hop router or the name of the link load balancing virtual server to which to send matching packets if action is set to ALLOW. If you specify a link load balancing (LLB) virtual server, which can provide a backup if a next hop link fails, first make sure that the next hops bound to the LLB virtual server are actually next hops that are directly connected to the Citrix ADC. Otherwise, the Citrix ADC throws an error when you attempt to create the PBR. The next hop can be null to represent null routes.
* `nexthopval` - The Next Hop IP address or gateway name.
* `ownergroup` - The owner node group in a Cluster for this pbr rule. If ownernode is not specified then the pbr rule is treated as Striped pbr rule.
* `priority` - Priority of the PBR, which determines the order in which it is evaluated relative to the other PBRs. If you do not specify priorities while creating PBRs, the PBRs are evaluated in the order in which they are created.
* `protocol` - Protocol, identified by protocol name, to match against the protocol of an outgoing IPv4 packet.
* `protocolnumber` - Protocol, identified by protocol number, to match against the protocol of an outgoing IPv4 packet.
* `srcip` - IP address or range of IP addresses to match against the source IP address of an outgoing IPv4 packet. In the command line interface, separate the range with a hyphen.
* `srcipop` - Either the equals (=) or does not equal (!=) logical operator.
* `srcipval` - IP address or range of IP addresses to match against the source IP address of an outgoing IPv4 packet. In the command line interface, separate the range with a hyphen.
* `srcmac` - MAC address to match against the source MAC address of an outgoing IPv4 packet.
* `srcmacmask` - Used to define range of Source MAC address. It takes string of 0 and 1, 0s are for exact match and 1s for wildcard. For matching first 3 bytes of MAC address, srcMacMask value "000000111111".
* `srcport` - Port number or range of port numbers to match against the source port number of an outgoing IPv4 packet. In the command line interface, separate the range with a hyphen. Note: The destination port can be specified only for TCP and UDP protocols.
* `srcportop` - Either the equals (=) or does not equal (!=) logical operator.
* `srcportval` - Port number or range of port numbers to match against the source port number of an outgoing IPv4 packet. In the command line interface, separate the range with a hyphen.
* `state` - Enable or disable the PBR. After you apply the PBRs, the Citrix ADC compares outgoing packets to the enabled PBRs.
* `targettd` - Integer value that uniquely identifies the traffic domain to which you want to send packet to.
* `td` - Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.
* `vlan` - ID of the VLAN. The Citrix ADC compares the PBR only to the outgoing packets on the specified VLAN. If you do not specify any interface ID, the appliance compares the PBR to the outgoing packets on all VLANs.
* `vxlan` - ID of the VXLAN. The Citrix ADC compares the PBR only to the outgoing packets on the specified VXLAN. If you do not specify any interface ID, the appliance compares the PBR to the outgoing packets on all VXLANs.
* `vxlanvlanmap` - The vlan to vxlan mapping to be applied for incoming packets over this pbr tunnel.
* `id` - The id of the nspbr. It has the same value as the `name` attribute.
