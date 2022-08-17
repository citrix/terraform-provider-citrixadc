---
subcategory: "NS"
---

# Resource: nspbr

The nspbr resource is used to create nspbr.


## Example usage

```hcl
resource "citrixadc_nspbr" "tf_nspbr" {
  name       = "my_nspbr"
  action     = "ALLOW"
  nexthop    = "true"
  nexthopval = "10.222.74.128"
}
```


## Argument Reference

* `name` - (Required) Name for the PBR. Must begin with an ASCII alphabetic or underscore \(_\) character, and must contain only ASCII alphanumeric, underscore, hash \(\#\), period \(.\), space, colon \(:\), at \(@\), equals \(=\), and hyphen \(-\) characters. Cannot be changed after the PBR is created. Minimum length =  1
* `action` - (Required) Action to perform on the outgoing IPv4 packets that match the PBR. Available settings function as follows: * ALLOW - The Citrix ADC sends the packet to the designated next-hop router. * DENY - The Citrix ADC applies the routing table for normal destination-based routing. Possible values: [ ALLOW, DENY ]
* `td` - (Optional) Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0. Minimum value =  0 Maximum value =  4094
* `srcip` - (Optional) IP address or range of IP addresses to match against the source IP address of an outgoing IPv4 packet. In the command line interface, separate the range with a hyphen. For example: 10.102.29.30-10.102.29.189.
* `srcipop` - (Optional) Either the equals (=) or does not equal (!=) logical operator. Possible values: [ =, !=, EQ, NEQ ]
* `srcipval` - (Optional) IP address or range of IP addresses to match against the source IP address of an outgoing IPv4 packet. In the command line interface, separate the range with a hyphen. For example: 10.102.29.30-10.102.29.189.
* `srcport` - (Optional) Port number or range of port numbers to match against the source port number of an outgoing IPv4 packet. In the command line interface, separate the range with a hyphen. For example: 40-90. Note: The destination port can be specified only for TCP and UDP protocols.
* `srcportop` - (Optional) Either the equals (=) or does not equal (!=) logical operator. Possible values: [ =, !=, EQ, NEQ ]
* `srcportval` - (Optional) Port number or range of port numbers to match against the source port number of an outgoing IPv4 packet. In the command line interface, separate the range with a hyphen. For example: 40-90. Note: The destination port can be specified only for TCP and UDP protocols. Maximum length =  65535
* `destip` - (Optional) IP address or range of IP addresses to match against the destination IP address of an outgoing IPv4 packet.  In the command line interface, separate the range with a hyphen. For example: 10.102.29.30-10.102.29.189.
* `destipop` - (Optional) Either the equals (=) or does not equal (!=) logical operator. Possible values: [ =, !=, EQ, NEQ ]
* `destipval` - (Optional) IP address or range of IP addresses to match against the destination IP address of an outgoing IPv4 packet. In the command line interface, separate the range with a hyphen. For example: 10.102.29.30-10.102.29.189.
* `destport` - (Optional) Port number or range of port numbers to match against the destination port number of an outgoing IPv4 packet. In the command line interface, separate the range with a hyphen. For example: 40-90. Note: The destination port can be specified only for TCP and UDP protocols.
* `destportop` - (Optional) Either the equals (=) or does not equal (!=) logical operator. Possible values: [ =, !=, EQ, NEQ ]
* `destportval` - (Optional) Port number or range of port numbers to match against the destination port number of an outgoing IPv4 packet. In the command line interface, separate the range with a hyphen. For example: 40-90. Note: The destination port can be specified only for TCP and UDP protocols. Maximum length =  65535
* `nexthop` - (Optional) IP address of the next hop router or the name of the link load balancing virtual server to which to send matching packets if action is set to ALLOW. If you specify a link load balancing (LLB) virtual server, which can provide a backup if a next hop link fails, first make sure that the next hops bound to the LLB virtual server are actually next hops that are directly connected to the Citrix ADC. Otherwise, the Citrix ADC throws an error when you attempt to create the PBR. The next hop can be null to represent null routes.
* `nexthopval` - (Optional) The Next Hop IP address or gateway name.
* `iptunnel` - (Optional) The Tunnel name.
* `iptunnelname` - (Optional) The iptunnel name where packets need to be forwarded upon.
* `vxlanvlanmap` - (Optional) The vlan to vxlan mapping to be applied for incoming packets over this pbr tunnel.
* `srcmac` - (Optional) MAC address to match against the source MAC address of an outgoing IPv4 packet.
* `srcmacmask` - (Optional) Used to define range of Source MAC address. It takes string of 0 and 1, 0s are for exact match and 1s for wildcard. For matching first 3 bytes of MAC address, srcMacMask value "000000111111". .
* `protocol` - (Optional) Protocol, identified by protocol name, to match against the protocol of an outgoing IPv4 packet. Possible values: [ ICMP, IGMP, TCP, EGP, IGP, ARGUS, UDP, RDP, RSVP, EIGRP, L2TP, ISIS, GGP, IPoverIP, ST, CBT, BBN-RCC-M, NVP-II, PUP, EMCON, XNET, CHAOS, MUX, DCN-MEAS, HMP, PRM, XNS-IDP, TRUNK-1, TRUNK-2, LEAF-1, LEAF-2, IRTP, ISO-TP4, NETBLT, MFE-NSP, MERIT-INP, SEP, 3PC, IDPR, XTP, DDP, IDPR-CMTP, TP++, IL, IPv6, SDRP, IPv6-Route, IPv6-Frag, IDRP, GRE, MHRP, BNA, ESP, AH, I-NLSP, SWIPE, NARP, MOBILE, TLSP, SKIP, ICMPV6, IPv6-NoNx, IPv6-Opts, Any-Host-Internal-Protocol, CFTP, Any-Local-Network, SAT-EXPAK, KRYPTOLAN, RVD, IPPC, Any-Distributed-File-System, TFTP, VISA, IPCV, CPNX, CPHB, WSN, PVP, BR-SAT-MO, SUN-ND, WB-MON, WB-EXPAK, ISO-IP, VMTP, SECURE-VM, VINES, TTP, NSFNET-IG, DGP, TCF, OSPFIGP, Sprite-RP, LARP, MTP, AX.25, IPIP, MICP, SCC-SP, ETHERIP, Any-Private-Encryption-Scheme, GMTP, IFMP, PNNI, PIM, ARIS, SCPS, QNX, A/N, IPComp, SNP, Compaq-Pe, IPX-in-IP, VRRP, PGM, Any-0-Hop-Protocol, ENCAP, DDX, IATP, STP, SRP, UTI, SMP, SM, PTP, FIRE, CRTP, CRUDP, SSCOPMCE, IPLT, SPS, PIPE, SCTP, FC, RSVP-E2E-IGNORE, Mobility-Header, UDPLite ]
* `protocolnumber` - (Optional) Protocol, identified by protocol number, to match against the protocol of an outgoing IPv4 packet. Minimum value =  1 Maximum value =  255
* `vlan` - (Optional) ID of the VLAN. The Citrix ADC compares the PBR only to the outgoing packets on the specified VLAN. If you do not specify any interface ID, the appliance compares the PBR to the outgoing packets on all VLANs. Minimum value =  1 Maximum value =  4094
* `vxlan` - (Optional) ID of the VXLAN. The Citrix ADC compares the PBR only to the outgoing packets on the specified VXLAN. If you do not specify any interface ID, the appliance compares the PBR to the outgoing packets on all VXLANs. Minimum value =  1 Maximum value =  16777215
* `Interface` - (Optional) ID of an interface. The Citrix ADC compares the PBR only to the outgoing packets on the specified interface. If you do not specify any value, the appliance compares the PBR to the outgoing packets on all interfaces.
* `priority` - (Optional) Priority of the PBR, which determines the order in which it is evaluated relative to the other PBRs. If you do not specify priorities while creating PBRs, the PBRs are evaluated in the order in which they are created. Minimum value =  1 Maximum value =  81920
* `msr` - (Optional) Monitor the route specified byte Next Hop parameter. This parameter is not applicable if you specify a link load balancing (LLB) virtual server name with the Next Hop parameter. Possible values: [ ENABLED, DISABLED ]
* `monitor` - (Optional) The name of the monitor.(Can be only of type ping or ARP ). Minimum length =  1
* `state` - (Optional) Enable or disable the PBR. After you apply the PBRs, the Citrix ADC compares outgoing packets to the enabled PBRs. Possible values: [ ENABLED, DISABLED ]
* `ownergroup` - (Optional) The owner node group in a Cluster for this pbr rule. If ownernode is not specified then the pbr rule is treated as Striped pbr rule. Minimum length =  1
* `detail` - (Optional) To get a detailed view.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nspbr. It has the same value as the `name` attribute.


## Import

A nspbr can be imported using its name, e.g.

```shell
terraform import citrixadc_nspbr.tf_nspbr my_nspbr
```
