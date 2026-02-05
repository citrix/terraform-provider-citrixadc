---
subcategory: "NS"
---

# Data Source `nspbr6`

The nspbr6 data source allows you to retrieve information about IPv6 Policy Based Routes (PBR6).


## Example usage

```terraform
resource "citrixadc_nspbr6" "my_nspbr6" {
  name      = "my_nspbr6"
  action    = "ALLOW"
  srcipv6   = true
  srcipv6val = "2001:db8::/32"
  destipv6  = true
  destipv6val = "2001:db8:1::/48"
  nexthop   = true
  nexthopval = "2001:db8:2::1"
  priority  = 10
}

data "citrixadc_nspbr6" "my_nspbr6_data" {
  name   = citrixadc_nspbr6.my_nspbr6.name
  detail = false
}

output "nspbr6_action" {
  value = data.citrixadc_nspbr6.my_nspbr6_data.action
}

output "nspbr6_priority" {
  value = data.citrixadc_nspbr6.my_nspbr6_data.priority
}
```


## Argument Reference

* `name` - (Required) Name for the PBR6. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.
* `detail` - (Required) To get a detailed view.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `action` - Action to perform on the outgoing IPv6 packets that match the PBR6. Available settings: ALLOW - The Citrix ADC sends the packet to the designated next-hop router. DENY - The Citrix ADC applies the routing table for normal destination-based routing.
* `destipop` - Either the equals (=) or does not equal (!=) logical operator.
* `destipv6` - IP address or range of IP addresses to match against the destination IP address of an outgoing IPv6 packet.
* `destipv6val` - IP address or range of IP addresses to match against the destination IP address of an outgoing IPv6 packet.
* `destport` - Port number or range of port numbers to match against the destination port number of an outgoing IPv6 packet.
* `destportop` - Either the equals (=) or does not equal (!=) logical operator.
* `destportval` - Destination port (range).
* `interface` - ID of an interface. The Citrix ADC compares the PBR6 only to the outgoing packets on the specified interface.
* `iptunnel` - The iptunnel name where packets need to be forwarded upon.
* `monitor` - The name of the monitor (Can be only of type ping or ARP).
* `msr` - Monitor the route specified by the Next Hop parameter.
* `nexthop` - IP address of the next hop router to which to send matching packets if action is set to ALLOW.
* `nexthopval` - The Next Hop IPv6 address.
* `nexthopvlan` - VLAN number to be used for link local nexthop.
* `ownergroup` - The owner node group in a Cluster for this pbr rule.
* `priority` - Priority of the PBR6, which determines the order in which it is evaluated relative to the other PBR6s.
* `protocol` - Protocol, identified by protocol name, to match against the protocol of an outgoing IPv6 packet.
* `protocolnumber` - Protocol, identified by protocol number, to match against the protocol of an outgoing IPv6 packet.
* `srcipop` - Either the equals (=) or does not equal (!=) logical operator.
* `srcipv6` - IP address or range of IP addresses to match against the source IP address of an outgoing IPv6 packet.
* `srcipv6val` - IP address or range of IP addresses to match against the source IP address of an outgoing IPv6 packet.
* `srcmac` - MAC address to match against the source MAC address of an outgoing IPv6 packet.
* `srcmacmask` - Used to define range of Source MAC address.
* `srcport` - Port number or range of port numbers to match against the source port number of an outgoing IPv6 packet.
* `srcportop` - Either the equals (=) or does not equal (!=) logical operator.
* `srcportval` - Source port (range).
* `state` - Enable or disable the PBR6 rule.
* `td` - Integer value that uniquely identifies the traffic domain in which you want to configure the entity.
* `vlan` - VLAN ID of the IPv6 packet that has to match with the PBR6.
* `vxlan` - VXLAN ID of the IPv6 packet that has to match with the PBR6.
* `vxlanvlanmap` - The vlan to vxlan mapping to be applied for incoming packets over this pbr tunnel.
* `id` - The id of the nspbr6. It has the same value as the `name` attribute.
