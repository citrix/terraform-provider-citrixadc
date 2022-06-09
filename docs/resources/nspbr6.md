---
subcategory: "NS"
---

# Resource: nspbr6

The nspbr6 resource is used to create PBR6 entry resource.


## Example usage

```hcl
resource "citrixadc_iptunnel" "tf_iptunnel" {
  name             = "tf_iptunnel"
  remote           = "66.0.0.11"
  remotesubnetmask = "255.255.255.255"
  local            = "*"
}
resource "citrixadc_nspbr6" "tf_nspbr6" {
  name     = "tf_nspbr6"
  action   = "ALLOW"
  protocol = "ICMPV6"
  priority = 20
  state    = "ENABLED"
  iptunnel = citrixadc_iptunnel.tf_iptunnel.name
}
```


## Argument Reference

* `name` - (Required) Name for the PBR6. Must begin with an ASCII alphabetic or underscore \(_\) character, and must contain only ASCII alphanumeric, underscore, hash \(\#\), period \(.\), space, colon \(:\), at \(@\), equals \(=\), and hyphen \(-\) characters. Cannot be changed after the PBR6 is created. Minimum length =  1
* `action` - (Required) Action to perform on the outgoing IPv6 packets that match the PBR6. Available settings function as follows: * ALLOW - The Citrix ADC sends the packet to the designated next-hop router. * DENY - The Citrix ADC applies the routing table for normal destination-based routing. Possible values: [ ALLOW, DENY ]
* `td` - (Optional) Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0. Minimum value =  0 Maximum value =  4094
* `srcipv6` - (Optional) IP address or range of IP addresses to match against the source IP address of an outgoing IPv6 packet. In the command line interface, separate the range with a hyphen.
* `srcipop` - (Optional) Either the equals (=) or does not equal (!=) logical operator. Possible values: [ =, !=, EQ, NEQ ]
* `srcipv6val` - (Optional) IP address or range of IP addresses to match against the source IP address of an outgoing IPv6 packet. In the command line interface, separate the range with a hyphen.
* `srcport` - (Optional) Port number or range of port numbers to match against the source port number of an outgoing IPv6 packet. In the command line interface, separate the range with a hyphen. For example: 40-90.
* `srcportop` - (Optional) Either the equals (=) or does not equal (!=) logical operator. Possible values: [ =, !=, EQ, NEQ ]
* `srcportval` - (Optional) Source port (range). Maximum length =  65535
* `destipv6` - (Optional) IP address or range of IP addresses to match against the destination IP address of an outgoing IPv6 packet.  In the command line interface, separate the range with a hyphen.
* `destipop` - (Optional) Either the equals (=) or does not equal (!=) logical operator. Possible values: [ =, !=, EQ, NEQ ]
* `destipv6val` - (Optional) IP address or range of IP addresses to match against the destination IP address of an outgoing IPv6 packet.  In the command line interface, separate the range with a hyphen.
* `destport` - (Optional) Port number or range of port numbers to match against the destination port number of an outgoing IPv6 packet. In the command line interface, separate the range with a hyphen. For example: 40-90. Note: The destination port can be specified only for TCP and UDP protocols.
* `destportop` - (Optional) Either the equals (=) or does not equal (!=) logical operator. Possible values: [ =, !=, EQ, NEQ ]
* `destportval` - (Optional) Destination port (range). Maximum length =  65535
* `srcmac` - (Optional) MAC address to match against the source MAC address of an outgoing IPv6 packet.
* `srcmacmask` - (Optional) Used to define range of Source MAC address. It takes string of 0 and 1, 0s are for exact match and 1s for wildcard. For matching first 3 bytes of MAC address, srcMacMask value "000000111111". .
* `protocol` - (Optional) Protocol, identified by protocol name, to match against the protocol of an outgoing IPv6 packet. Possible values: [ ICMPV6, TCP, UDP, ICMP, IGMP, EGP, IGP, ARGUS, RDP, RSVP, EIGRP, L2TP, ISIS, GGP, IPoverIP, ST, CBT, BBN-RCC-M, NVP-II, PUP, EMCON, XNET, CHAOS, MUX, DCN-MEAS, HMP, PRM, XNS-IDP, TRUNK-1, TRUNK-2, LEAF-1, LEAF-2, IRTP, ISO-TP4, NETBLT, MFE-NSP, MERIT-INP, SEP, 3PC, IDPR, XTP, DDP, IDPR-CMTP, TP++, IL, IPv6, SDRP, IPv6-Route, IPv6-Frag, IDRP, GRE, MHRP, BNA, ESP, AH, I-NLSP, SWIPE, NARP, MOBILE, TLSP, SKIP, IPv6-NoNx, IPv6-Opts, Any-Host-Internal-Protocol, CFTP, Any-Local-Network, SAT-EXPAK, KRYPTOLAN, RVD, IPPC, Any-Distributed-File-System, TFTP, VISA, IPCV, CPNX, CPHB, WSN, PVP, BR-SAT-MO, SUN-ND, WB-MON, WB-EXPAK, ISO-IP, VMTP, SECURE-VM, VINES, TTP, NSFNET-IG, DGP, TCF, OSPFIGP, Sprite-RP, LARP, MTP, AX.25, IPIP, MICP, SCC-SP, ETHERIP, Any-Private-Encryption-Scheme, GMTP, IFMP, PNNI, PIM, ARIS, SCPS, QNX, A/N, IPComp, SNP, Compaq-Pe, IPX-in-IP, VRRP, PGM, Any-0-Hop-Protocol, ENCAP, DDX, IATP, STP, SRP, UTI, SMP, SM, PTP, FIRE, CRTP, CRUDP, SSCOPMCE, IPLT, SPS, PIPE, SCTP, FC, RSVP-E2E-IGNORE, Mobility-Header, UDPLite ]
* `protocolnumber` - (Optional) Protocol, identified by protocol number, to match against the protocol of an outgoing IPv6 packet. Minimum value =  1 Maximum value =  255
* `vlan` - (Optional) ID of the VLAN. The Citrix ADC compares the PBR6 only to the outgoing packets on the specified VLAN. If you do not specify an interface ID, the appliance compares the PBR6 to the outgoing packets on all VLANs. Minimum value =  1 Maximum value =  4094
* `vxlan` - (Optional) ID of the VXLAN. The Citrix ADC compares the PBR6 only to the outgoing packets on the specified VXLAN. If you do not specify an interface ID, the appliance compares the PBR6 to the outgoing packets on all VXLANs. Minimum value =  1 Maximum value =  16777215
* `Interface` - (Optional) ID of an interface. The Citrix ADC compares the PBR6 only to the outgoing packets on the specified interface. If you do not specify a value, the appliance compares the PBR6 to the outgoing packets on all interfaces.
* `priority` - (Optional) Priority of the PBR6, which determines the order in which it is evaluated relative to the other PBR6s. If you do not specify priorities while creating PBR6s, the PBR6s are evaluated in the order in which they are created. Minimum value =  1 Maximum value =  81920
* `state` - (Optional) Enable or disable the PBR6. After you apply the PBR6s, the Citrix ADC compares outgoing packets to the enabled PBR6s. Possible values: [ ENABLED, DISABLED ]
* `msr` - (Optional) Monitor the route specified by the Next Hop parameter. Possible values: [ ENABLED, DISABLED ]
* `monitor` - (Optional) The name of the monitor.(Can be only of type ping or ARP ). Minimum length =  1
* `nexthop` - (Optional) IP address of the next hop router to which to send matching packets if action is set to ALLOW. This next hop should be directly reachable from the appliance.
* `nexthopval` - (Optional) The Next Hop IPv6 address.
* `iptunnel` - (Optional) The iptunnel name where packets need to be forwarded upon.
* `vxlanvlanmap` - (Optional) The vlan to vxlan mapping to be applied for incoming packets over this pbr tunnel.
* `nexthopvlan` - (Optional) VLAN number to be used for link local nexthop . Minimum value =  1 Maximum value =  4094
* `ownergroup` - (Optional) The owner node group in a Cluster for this pbr rule. If owner node group is not specified then the pbr rule is treated as Striped pbr rule. Minimum length =  1
* `detail` - (Optional) To get a detailed view.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nspbr6. It has the same value as the `name` attribute.


## Import

A nspbr6 can be imported using its name, e.g.

```shell
terraform import citrixadc_nspbr6.tf_nspbr6 tf_nspbr6
```
