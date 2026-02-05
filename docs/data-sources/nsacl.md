---
subcategory: "NS"
---

# Data Source `nsacl`

The nsacl data source allows you to retrieve information about Citrix ADC extended ACL (Access Control List) rules.


## Example usage

```terraform
data "citrixadc_nsacl" "tf_nsacl" {
  aclname = "test_acl"
  type    = "CLASSIC"
}

output "aclaction" {
  value = data.citrixadc_nsacl.tf_nsacl.aclaction
}

output "destipval" {
  value = data.citrixadc_nsacl.tf_nsacl.destipval
}
```


## Argument Reference

* `aclname` - (Required) Name for the extended ACL rule. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.
* `type` - (Required) Type of the acl, default will be CLASSIC. Available options as follows: CLASSIC - specifies the regular extended acls. DFD - cluster specific acls, specifies hashmethod for steering of the packet in cluster.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nsacl.
* `interface` - ID of an interface. The Citrix ADC applies the ACL rule only to the incoming packets from the specified interface. If you do not specify any value, the appliance applies the ACL rule to the incoming packets of all interfaces.
* `aclaction` - Action to perform on incoming IPv4 packets that match the extended ACL rule. Available settings function as follows: ALLOW - The Citrix ADC processes the packet. BRIDGE - The Citrix ADC bridges the packet to the destination without processing it. DENY - The Citrix ADC drops the packet.
* `destip` - IP address or range of IP addresses to match against the destination IP address of an incoming IPv4 packet. In the command line interface, separate the range with a hyphen.
* `destipdataset` - Policy dataset which can have multiple IP ranges bound to it.
* `destipop` - Either the equals (=) or does not equal (!=) logical operator.
* `destipval` - IP address or range of IP addresses to match against the destination IP address of an incoming IPv4 packet. In the command line interface, separate the range with a hyphen. For example: 10.102.29.30-10.102.29.189.
* `destport` - Port number or range of port numbers to match against the destination port number of an incoming IPv4 packet. In the command line interface, separate the range with a hyphen. For example: 40-90. Note: The destination port can be specified only for TCP and UDP protocols.
* `destportdataset` - Policy dataset which can have multiple port ranges bound to it.
* `destportop` - Either the equals (=) or does not equal (!=) logical operator.
* `destportval` - Port number or range of port numbers to match against the destination port number of an incoming IPv4 packet. In the command line interface, separate the range with a hyphen. For example: 40-90. Note: The destination port can be specified only for TCP and UDP protocols.
* `dfdhash` - Specifies the type hashmethod to be applied, to steer the packet to the FP of the packet.
* `established` - Allow only incoming TCP packets that have the ACK or RST bit set, if the action set for the ACL rule is ALLOW and these packets match the other conditions in the ACL rule.
* `icmpcode` - Code of a particular ICMP message type to match against the ICMP code of an incoming ICMP packet. For example, to block DESTINATION HOST UNREACHABLE messages, specify 3 as the ICMP type and 1 as the ICMP code. If you set this parameter, you must set the ICMP Type parameter.
* `icmptype` - ICMP Message type to match against the message type of an incoming ICMP packet. For example, to block DESTINATION UNREACHABLE messages, you must specify 3 as the ICMP type. Note: This parameter can be specified only for the ICMP protocol.
* `logstate` - Enable or disable logging of events related to the extended ACL rule. The log messages are stored in the configured syslog or auditlog server.
* `newname` - New name for the extended ACL rule. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.
* `nodeid` - Specifies the NodeId to steer the packet to the provided FP.
* `priority` - Priority for the extended ACL rule that determines the order in which it is evaluated relative to the other extended ACL rules. If you do not specify priorities while creating extended ACL rules, the ACL rules are evaluated in the order in which they are created.
* `protocol` - Protocol to match against the protocol of an incoming IPv4 packet.
* `protocolnumber` - Protocol to match against the protocol of an incoming IPv4 packet.
* `ratelimit` - Maximum number of log messages to be generated per second. If you set this parameter, you must enable the Log State parameter.
* `srcip` - IP address or range of IP addresses to match against the source IP address of an incoming IPv4 packet. In the command line interface, separate the range with a hyphen. For example: 10.102.29.30-10.102.29.189.
* `srcipdataset` - Policy dataset which can have multiple IP ranges bound to it.
* `srcipop` - Either the equals (=) or does not equal (!=) logical operator.
* `srcipval` - IP address or range of IP addresses to match against the source IP address of an incoming IPv4 packet. In the command line interface, separate the range with a hyphen. For example:10.102.29.30-10.102.29.189.
* `srcmac` - MAC address to match against the source MAC address of an incoming IPv4 packet.
* `srcmacmask` - Used to define range of Source MAC address. It takes string of 0 and 1, 0s are for exact match and 1s for wildcard. For matching first 3 bytes of MAC address, srcMacMask value "000000111111".
* `srcport` - Port number or range of port numbers to match against the source port number of an incoming IPv4 packet. In the command line interface, separate the range with a hyphen. For example: 40-90.
* `srcportdataset` - Policy dataset which can have multiple port ranges bound to it.
* `srcportop` - Either the equals (=) or does not equal (!=) logical operator.
* `srcportval` - Port number or range of port numbers to match against the source port number of an incoming IPv4 packet. In the command line interface, separate the range with a hyphen. For example: 40-90.
* `state` - Enable or disable the extended ACL rule. After you apply the extended ACL rules, the Citrix ADC compares incoming packets against the enabled extended ACL rules.
* `stateful` - If stateful option is enabled, transparent sessions are created for the traffic hitting this ACL and not hitting any other features like LB, INAT etc.
* `td` - Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.
* `ttl` - Number of seconds, in multiples of four, after which the extended ACL rule expires. If you do not want the extended ACL rule to expire, do not specify a TTL value.
* `vlan` - ID of the VLAN. The Citrix ADC applies the ACL rule only to the incoming packets of the specified VLAN. If you do not specify a VLAN ID, the appliance applies the ACL rule to the incoming packets on all VLANs.
