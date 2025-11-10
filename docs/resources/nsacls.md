---
subcategory: "NS"
---

# Resource: nsacls

The nsacls resource is used to create a set of ADC acls.


## Example usage

```hcl
resource "citrixadc_nsacls" "foo" {
 acl  {
    aclname = "restricttcp2"
    protocol = "TCP"
    aclaction = "DENY"
    destipval = "192.168.199.52"
    srcportval = "149-1524"
    priority = "25"
  }

  acl  {
    aclname = "restrictvlan"
    aclaction = "DENY"
    vlan = "2000"
    priority = "250"
  }
}

resource "citrixadc_nsacls" "foo_1" {
  acls_apply_trigger = "No"
  acl  {
    aclname = "allowudp"
    protocol = "UDP"
    aclaction = "ALLOW"
    destipval = "192.168.45.55"
    srcportval = "490-1024"
    priority = "100"
  }
}

```


## Argument Reference


* `aclsname` - (Optional) Name for the acls resource.
* `acls_apply_trigger` - (Optional) Trigger to force `apply` of ACLs. Set to "Yes" to apply all configured ACLs to the NetScaler. This acts as a toggle mechanism. Valid values are "Yes" and "No". Default: "No". (The value automatically resets to "No" after each read operation, allowing subsequent plans to detect changes when set to "Yes" again.)
* `type` - (Optional) Type of the acl ,default will be CLASSIC. Available options as follows: * CLASSIC - specifies the regular extended acls. * DFD - cluster specific acls,specifies hashmethod for steering of the packet in cluster.
* `acl` - (Optional) A set of block defining acls.
The acl block supports the following:

* `aclname` - (Required) Name for the extended ACL rule. 
* `aclaction` - (Optional) Action to perform on incoming IPv4 packets that match the extended ACL rule. Available settings function as follows: * ALLOW - The Citrix ADC processes the packet. * BRIDGE - The Citrix ADC bridges the packet to the destination without processing it. * DENY - The Citrix ADC drops the packet. Possible values: [ BRIDGE, DENY, ALLOW ]
* `td` - (Optional) Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.
* `srcipop` - (Optional) Either the equals (=) or does not equal (!=) logical operator. Possible values: [ =, !=, EQ, NEQ ]
* `srcipval` - (Optional) IP address or range of IP addresses to match against the source IP address of an incoming IPv4 packet. In the command line interface, separate the range with a hyphen. For example:10.102.29.30-10.102.29.189.
* `srcportop` - (Optional) Either the equals (=) or does not equal (!=) logical operator. Possible values: [ =, !=, EQ, NEQ ]
* `srcportval` - (Optional) Port number or range of port numbers to match against the source port number of an incoming IPv4 packet. In the command line interface, separate the range with a hyphen. For example: 40-90.
* `destipop` - (Optional) Either the equals (=) or does not equal (!=) logical operator. Possible values: [ =, !=, EQ, NEQ ]
* `destipval` - (Optional) IP address or range of IP addresses to match against the destination IP address of an incoming IPv4 packet.  In the command line interface, separate the range with a hyphen. For example: 10.102.29.30-10.102.29.189.
* `destportop` - (Optional) Either the equals (=) or does not equal (!=) logical operator. Possible values: [ =, !=, EQ, NEQ ]
* `destportval` - (Optional) Port number or range of port numbers to match against the destination port number of an incoming IPv4 packet. In the command line interface, separate the range with a hyphen. For example: 40-90. Note: The destination port can be specified only for TCP and UDP protocols.
* `ttl` - (Optional) Number of seconds, in multiples of four, after which the extended ACL rule expires. If you do not want the extended ACL rule to expire, do not specify a TTL value.
* `srcmac` - (Optional) MAC address to match against the source MAC address of an incoming IPv4 packet.
* `protocol` - (Optional) Protocol to match against the protocol of an incoming IPv4 packet. Possible values: [ ICMP, IGMP, TCP, EGP, IGP, ARGUS, UDP, RDP, RSVP, EIGRP, L2TP, ISIS, GGP, IPoverIP, ST, CBT, BBN-RCC-M, NVP-II, PUP, EMCON, XNET, CHAOS, MUX, DCN-MEAS, HMP, PRM, XNS-IDP, TRUNK-1, TRUNK-2, LEAF-1, LEAF-2, IRTP, ISO-TP4, NETBLT, MFE-NSP, MERIT-INP, SEP, 3PC, IDPR, XTP, DDP, IDPR-CMTP, TP++, IL, IPv6, SDRP, IPv6-Route, IPv6-Frag, IDRP, GRE, MHRP, BNA, ESP, AH, I-NLSP, SWIPE, NARP, MOBILE, TLSP, SKIP, ICMPV6, IPv6-NoNx, IPv6-Opts, Any-Host-Internal-Protocol, CFTP, Any-Local-Network, SAT-EXPAK, KRYPTOLAN, RVD, IPPC, Any-Distributed-File-System, TFTP, VISA, IPCV, CPNX, CPHB, WSN, PVP, BR-SAT-MO, SUN-ND, WB-MON, WB-EXPAK, ISO-IP, VMTP, SECURE-VM, VINES, TTP, NSFNET-IG, DGP, TCF, OSPFIGP, Sprite-RP, LARP, MTP, AX.25, IPIP, MICP, SCC-SP, ETHERIP, Any-Private-Encryption-Scheme, GMTP, IFMP, PNNI, PIM, ARIS, SCPS, QNX, A/N, IPComp, SNP, Compaq-Pe, IPX-in-IP, VRRP, PGM, Any-0-Hop-Protocol, ENCAP, DDX, IATP, STP, SRP, UTI, SMP, SM, PTP, FIRE, CRTP, CRUDP, SSCOPMCE, IPLT, SPS, PIPE, SCTP, FC, RSVP-E2E-IGNORE, Mobility-Header, UDPLite ]
* `protocolnumber` - (Optional) Protocol to match against the protocol of an incoming IPv4 packet.
* `vlan` - (Optional) ID of the VLAN. The Citrix ADC applies the ACL rule only to the incoming packets of the specified VLAN. If you do not specify a VLAN ID, the appliance applies the ACL rule to the incoming packets on all VLANs.
* `interface` - (Optional) ID of an interface. The Citrix ADC applies the ACL rule only to the incoming packets from the specified interface. If you do not specify any value, the appliance applies the ACL rule to the incoming packets of all interfaces.
* `established` - (Optional) Allow only incoming TCP packets that have the ACK or RST bit set, if the action set for the ACL rule is ALLOW and these packets match the other conditions in the ACL rule.
* `icmptype` - (Optional) ICMP Message type to match against the message type of an incoming ICMP packet. For example, to block DESTINATION UNREACHABLE messages, you must specify 3 as the ICMP type. Note: This parameter can be specified only for the ICMP protocol.
* `icmpcode` - (Optional) Code of a particular ICMP message type to match against the ICMP code of an incoming ICMP packet.  For example, to block DESTINATION HOST UNREACHABLE messages, specify 3 as the ICMP type and 1 as the ICMP code. If you set this parameter, you must set the ICMP Type parameter.
* `priority` - (Optional) Priority for the extended ACL rule that determines the order in which it is evaluated relative to the other extended ACL rules. If you do not specify priorities while creating extended ACL rules, the ACL rules are evaluated in the order in which they are created.
* `state` - (Optional) Enable or disable the extended ACL rule. After you apply the extended ACL rules, the Citrix ADC compares incoming packets against the enabled extended ACL rules. Possible values: [ ENABLED, DISABLED ]
* `logstate` - (Optional) Enable or disable logging of events related to the extended ACL rule. The log messages are stored in the configured syslog or auditlog server. Possible values: [ ENABLED, DISABLED ]
* `ratelimit` - (Optional) 


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nsacls. It has the same value as the `aclsname` attribute.
