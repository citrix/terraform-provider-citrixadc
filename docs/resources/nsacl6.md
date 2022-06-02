---
subcategory: "NS"
---

# Resource: nsacl6

The nsacl6 resource is used to create ACL6 entry resource.


## Example usage

```hcl
resource "citrixadc_nstrafficdomain" "tf_trafficdomain" {
  td        = 2
  aliasname = "tf_trafficdomain"
  vmac      = "ENABLED"
}
resource "citrixadc_nsacl6" "tf_nsacl6" {
  acl6name   = "tf_nsacl6"
  acl6action = "ALLOW"
  td         = citrixadc_nstrafficdomain.tf_trafficdomain.td
  logstate   = "ENABLED"
  stateful   = "NO"
  ratelimit  = 120
  state      = "ENABLED"
  priority   = 20
  protocol   = "TCP"
}
```


## Argument Reference

* `acl6name` - (Required) Name for the ACL6 rule. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Minimum length =  1
* `acl6action` - (Required) Action to perform on the incoming IPv6 packets that match the ACL6 rule. Available settings function as follows: * ALLOW - The Citrix ADC processes the packet. * BRIDGE - The Citrix ADC bridges the packet to the destination without processing it. * DENY - The Citrix ADC drops the packet. Possible values: [ BRIDGE, DENY, ALLOW ]
* `td` - (Optional) Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0. Minimum value =  0 Maximum value =  4094
* `srcipv6` - (Optional) IP address or range of IP addresses to match against the source IP address of an incoming IPv6 packet. In the command line interface, separate the range with a hyphen.
* `srcipop` - (Optional) Either the equals (=) or does not equal (!=) logical operator. Possible values: [ =, !=, EQ, NEQ ]
* `srcipv6val` - (Optional) Source IPv6 address (range).
* `srcport` - (Optional) Port number or range of port numbers to match against the source port number of an incoming IPv6 packet. In the command line interface, separate the range with a hyphen. For example: 40-90. Note: The destination port can be specified only for TCP and UDP protocols.
* `srcportop` - (Optional) Either the equals (=) or does not equal (!=) logical operator. Possible values: [ =, !=, EQ, NEQ ]
* `srcportval` - (Optional) Source port (range). Maximum length =  65535
* `destipv6` - (Optional) IP address or range of IP addresses to match against the destination IP address of an incoming IPv6 packet.  In the command line interface, separate the range with a hyphen.
* `destipop` - (Optional) Either the equals (=) or does not equal (!=) logical operator. Possible values: [ =, !=, EQ, NEQ ]
* `destipv6val` - (Optional) Destination IPv6 address (range).
* `destport` - (Optional) Port number or range of port numbers to match against the destination port number of an incoming IPv6 packet. In the command line interface, separate the range with a hyphen. For example: 40-90. Note: The destination port can be specified only for TCP and UDP protocols.
* `destportop` - (Optional) Either the equals (=) or does not equal (!=) logical operator. Possible values: [ =, !=, EQ, NEQ ]
* `destportval` - (Optional) Destination port (range). Maximum length =  65535
* `ttl` - (Optional) Time to expire this ACL6 (in seconds). Minimum value =  1 Maximum value =  2147483647
* `srcmac` - (Optional) MAC address to match against the source MAC address of an incoming IPv6 packet.
* `srcmacmask` - (Optional) Used to define range of Source MAC address. It takes string of 0 and 1, 0s are for exact match and 1s for wildcard. For matching first 3 bytes of MAC address, srcMacMask value "000000111111". .
* `protocol` - (Optional) Protocol, identified by protocol name, to match against the protocol of an incoming IPv6 packet. Possible values: [ ICMPV6, TCP, UDP, ICMP, IGMP, EGP, IGP, ARGUS, RDP, RSVP, EIGRP, L2TP, ISIS, GGP, IPoverIP, ST, CBT, BBN-RCC-M, NVP-II, PUP, EMCON, XNET, CHAOS, MUX, DCN-MEAS, HMP, PRM, XNS-IDP, TRUNK-1, TRUNK-2, LEAF-1, LEAF-2, IRTP, ISO-TP4, NETBLT, MFE-NSP, MERIT-INP, SEP, 3PC, IDPR, XTP, DDP, IDPR-CMTP, TP++, IL, IPv6, SDRP, IPv6-Route, IPv6-Frag, IDRP, GRE, MHRP, BNA, ESP, AH, I-NLSP, SWIPE, NARP, MOBILE, TLSP, SKIP, IPv6-NoNx, IPv6-Opts, Any-Host-Internal-Protocol, CFTP, Any-Local-Network, SAT-EXPAK, KRYPTOLAN, RVD, IPPC, Any-Distributed-File-System, TFTP, VISA, IPCV, CPNX, CPHB, WSN, PVP, BR-SAT-MO, SUN-ND, WB-MON, WB-EXPAK, ISO-IP, VMTP, SECURE-VM, VINES, TTP, NSFNET-IG, DGP, TCF, OSPFIGP, Sprite-RP, LARP, MTP, AX.25, IPIP, MICP, SCC-SP, ETHERIP, Any-Private-Encryption-Scheme, GMTP, IFMP, PNNI, PIM, ARIS, SCPS, QNX, A/N, IPComp, SNP, Compaq-Pe, IPX-in-IP, VRRP, PGM, Any-0-Hop-Protocol, ENCAP, DDX, IATP, STP, SRP, UTI, SMP, SM, PTP, FIRE, CRTP, CRUDP, SSCOPMCE, IPLT, SPS, PIPE, SCTP, FC, RSVP-E2E-IGNORE, Mobility-Header, UDPLite ]
* `protocolnumber` - (Optional) Protocol, identified by protocol number, to match against the protocol of an incoming IPv6 packet. Minimum value =  1 Maximum value =  255
* `vlan` - (Optional) ID of the VLAN. The Citrix ADC applies the ACL6 rule only to the incoming packets on the specified VLAN. If you do not specify a VLAN ID, the appliance applies the ACL6 rule to the incoming packets on all VLANs. Minimum value =  1 Maximum value =  4094
* `vxlan` - (Optional) ID of the VXLAN. The Citrix ADC applies the ACL6 rule only to the incoming packets on the specified VXLAN. If you do not specify a VXLAN ID, the appliance applies the ACL6 rule to the incoming packets on all VXLANs. Minimum value =  1 Maximum value =  16777215
* `Interface` - (Optional) ID of an interface. The Citrix ADC applies the ACL6 rule only to the incoming packets from the specified interface. If you do not specify any value, the appliance applies the ACL6 rule to the incoming packets from all interfaces.
* `established` - (Optional) Allow only incoming TCP packets that have the ACK or RST bit set if the action set for the ACL6 rule is ALLOW and these packets match the other conditions in the ACL6 rule.
* `icmptype` - (Optional) ICMP Message type to match against the message type of an incoming IPv6 ICMP packet. For example, to block DESTINATION UNREACHABLE messages, you must specify 3 as the ICMP type. Note: This parameter can be specified only for the ICMP protocol. Minimum value =  0 Maximum value =  65536
* `icmpcode` - (Optional) Code of a particular ICMP message type to match against the ICMP code of an incoming IPv6 ICMP packet.  For example, to block DESTINATION HOST UNREACHABLE messages, specify 3 as the ICMP type and 1 as the ICMP code. If you set this parameter, you must set the ICMP Type parameter. Minimum value =  0 Maximum value =  65536
* `priority` - (Optional) Priority for the ACL6 rule, which determines the order in which it is evaluated relative to the other ACL6 rules. If you do not specify priorities while creating ACL6 rules, the ACL6 rules are evaluated in the order in which they are created. Minimum value =  1 Maximum value =  81920
* `state` - (Optional) State of the ACL6. Possible values: [ ENABLED, DISABLED ]
* `type` - (Optional) Type of the acl6 ,default will be CLASSIC. Available options as follows: * CLASSIC - specifies the regular extended acls. * DFD - cluster specific acls,specifies hashmethod for steering of the packet in cluster . Possible values: [ CLASSIC, DFD ]
* `dfdhash` - (Optional) Specifies the type of hashmethod to be applied, to steer the packet to the FP of the packet. Possible values: [ SIP-SPORT-DIP-DPORT, SIP, DIP, SIP-DIP, SIP-SPORT, DIP-DPORT ]
* `dfdprefix` - (Optional) hashprefix to be applied to SIP/DIP to generate rsshash FP.eg 128 => hash calculated on the complete IP. Minimum value =  1 Maximum value =  128
* `stateful` - (Optional) If stateful option is enabled, transparent sessions are created for the traffic hitting this ACL6 and not hitting any other features like LB, INAT etc. . Possible values: [ YES, NO ]
* `logstate` - (Optional) Enable or disable logging of events related to the ACL6 rule. The log messages are stored in the configured syslog or auditlog server. Possible values: [ ENABLED, DISABLED ]
* `ratelimit` - (Optional) Maximum number of log messages to be generated per second. If you set this parameter, you must enable the Log State parameter. Minimum value =  1 Maximum value =  10000
* `aclaction` - (Optional) Action associated with the ACL6. Possible values: [ BRIDGE, DENY, ALLOW ]
* `newname` - (Optional) New name for the ACL6 rule. Must begin with an ASCII alphabetic or underscore \(_\) character, and must contain only ASCII alphanumeric, underscore, hash \(\#\), period \(.\), space, colon \(:\), at \(@\), equals \(=\), and hyphen \(-\) characters. Minimum length =  1


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nsacl6. It has the same value as the `acl6name` attribute.


## Import

A nsacl6 can be imported using its acl6name, e.g.

```shell
terraform import citrixadc_nsacl6.tf_nsacl6 tf_nsacl6
```
