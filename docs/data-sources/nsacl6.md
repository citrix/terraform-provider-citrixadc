---
subcategory: "NS"
---

# Data Source `nsacl6`

The nsacl6 data source allows you to retrieve information about IPv6 ACL (Access Control List) rules configured on Citrix ADC.


## Example usage

```terraform
data "citrixadc_nsacl6" "tf_nsacl6" {
  acl6name = "my_acl6"
  type     = "CLASSIC"
}

output "acl6_action" {
  value = data.citrixadc_nsacl6.tf_nsacl6.acl6action
}

output "acl6_state" {
  value = data.citrixadc_nsacl6.tf_nsacl6.state
}
```


## Argument Reference

* `acl6name` - (Required) Name of the ACL6 rule.
* `type` - (Required) Type of the ACL6. Possible values: `CLASSIC` (regular extended ACLs) or `DFD` (cluster specific ACLs with hashmethod for steering packets in cluster).

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nsacl6. It is a combination of `acl6name` and `type`.
* `interface` - ID of an interface. The Citrix ADC applies the ACL6 rule only to the incoming packets from the specified interface.
* `acl6action` - Action to perform on the incoming IPv6 packets that match the ACL6 rule. Possible values: `ALLOW`, `BRIDGE`, `DENY`.
* `aclaction` - Action associated with the ACL6.
* `destipop` - Either the equals (=) or does not equal (!=) logical operator.
* `destipv6` - IP address or range of IP addresses to match against the destination IP address of an incoming IPv6 packet.
* `destipv6val` - Destination IPv6 address (range).
* `destport` - Port number or range of port numbers to match against the destination port number of an incoming IPv6 packet.
* `destportop` - Either the equals (=) or does not equal (!=) logical operator.
* `destportval` - Destination port (range).
* `dfdhash` - Type of hashmethod to be applied to steer the packet to the FP of the packet.
* `dfdprefix` - Hashprefix to be applied to SIP/DIP to generate rsshash FP.
* `established` - Allow only incoming TCP packets that have the ACK or RST bit set.
* `icmpcode` - Code of a particular ICMP message type to match against the ICMP code of an incoming IPv6 ICMP packet.
* `icmptype` - ICMP Message type to match against the message type of an incoming IPv6 ICMP packet.
* `logstate` - Enable or disable logging of events related to the ACL6 rule.
* `newname` - New name for the ACL6 rule.
* `nodeid` - NodeId to steer the packet to the provided FP.
* `priority` - Priority for the ACL6 rule, which determines the order in which it is evaluated relative to other ACL6 rules.
* `protocol` - Protocol, identified by protocol name, to match against the protocol of an incoming IPv6 packet.
* `protocolnumber` - Protocol, identified by protocol number, to match against the protocol of an incoming IPv6 packet.
* `ratelimit` - Maximum number of log messages to be generated per second.
* `srcipop` - Either the equals (=) or does not equal (!=) logical operator.
* `srcipv6` - IP address or range of IP addresses to match against the source IP address of an incoming IPv6 packet.
* `srcipv6val` - Source IPv6 address (range).
* `srcmac` - MAC address to match against the source MAC address of an incoming IPv6 packet.
* `srcmacmask` - Used to define range of Source MAC address.
* `srcport` - Port number or range of port numbers to match against the source port number of an incoming IPv6 packet.
* `srcportop` - Either the equals (=) or does not equal (!=) logical operator.
* `srcportval` - Source port (range).
* `state` - State of the ACL6. Possible values: `ENABLED`, `DISABLED`.
* `stateful` - If stateful option is enabled, transparent sessions are created for the traffic hitting this ACL6.
* `td` - Integer value that uniquely identifies the traffic domain in which you want to configure the entity.
* `ttl` - Time to expire this ACL6 (in seconds).
* `vlan` - ID of the VLAN. The Citrix ADC applies the ACL6 rule only to the incoming packets on the specified VLAN.
* `vxlan` - ID of the VXLAN. The Citrix ADC applies the ACL6 rule only to the incoming packets on the specified VXLAN.
