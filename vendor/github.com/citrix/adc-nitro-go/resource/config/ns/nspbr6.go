/*
* Copyright (c) 2021 Citrix Systems, Inc.
*
*   Licensed under the Apache License, Version 2.0 (the "License");
*   you may not use this file except in compliance with the License.
*   You may obtain a copy of the License at
*
*       http://www.apache.org/licenses/LICENSE-2.0
*
*  Unless required by applicable law or agreed to in writing, software
*   distributed under the License is distributed on an "AS IS" BASIS,
*   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
*   See the License for the specific language governing permissions and
*   limitations under the License.
*/

package ns

/**
* Configuration for PBR6 entry resource.
*/
type Nspbr6 struct {
	/**
	* Name for the PBR6. Must begin with an ASCII alphabetic or underscore \(_\) character, and must contain only ASCII alphanumeric, underscore, hash \(\#\), period \(.\), space, colon \(:\), at \(@\), equals \(=\), and hyphen \(-\) characters. Cannot be changed after the PBR6 is created.
	*/
	Name string `json:"name,omitempty"`
	/**
	* Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.
	*/
	Td uint32 `json:"td,omitempty"`
	/**
	* Action to perform on the outgoing IPv6 packets that match the PBR6.
		Available settings function as follows:
		* ALLOW - The Citrix ADC sends the packet to the designated next-hop router.
		* DENY - The Citrix ADC applies the routing table for normal destination-based routing.
	*/
	Action string `json:"action,omitempty"`
	/**
	* IP address or range of IP addresses to match against the source IP address of an outgoing IPv6 packet. In the command line interface, separate the range with a hyphen.
	*/
	Srcipv6 bool `json:"srcipv6,omitempty"`
	/**
	* Either the equals (=) or does not equal (!=) logical operator.
	*/
	Srcipop string `json:"srcipop,omitempty"`
	/**
	* IP address or range of IP addresses to match against the source IP address of an outgoing IPv6 packet. In the command line interface, separate the range with a hyphen.
	*/
	Srcipv6val string `json:"srcipv6val,omitempty"`
	/**
	* Port number or range of port numbers to match against the source port number of an outgoing IPv6 packet. In the command line interface, separate the range with a hyphen. For example: 40-90.
	*/
	Srcport bool `json:"srcport,omitempty"`
	/**
	* Either the equals (=) or does not equal (!=) logical operator.
	*/
	Srcportop string `json:"srcportop,omitempty"`
	/**
	* Source port (range).
	*/
	Srcportval string `json:"srcportval,omitempty"`
	/**
	* IP address or range of IP addresses to match against the destination IP address of an outgoing IPv6 packet.  In the command line interface, separate the range with a hyphen.
	*/
	Destipv6 bool `json:"destipv6,omitempty"`
	/**
	* Either the equals (=) or does not equal (!=) logical operator.
	*/
	Destipop string `json:"destipop,omitempty"`
	/**
	* IP address or range of IP addresses to match against the destination IP address of an outgoing IPv6 packet.  In the command line interface, separate the range with a hyphen.
	*/
	Destipv6val string `json:"destipv6val,omitempty"`
	/**
	* Port number or range of port numbers to match against the destination port number of an outgoing IPv6 packet. In the command line interface, separate the range with a hyphen. For example: 40-90.
		Note: The destination port can be specified only for TCP and UDP protocols.
	*/
	Destport bool `json:"destport,omitempty"`
	/**
	* Either the equals (=) or does not equal (!=) logical operator.
	*/
	Destportop string `json:"destportop,omitempty"`
	/**
	* Destination port (range).
	*/
	Destportval string `json:"destportval,omitempty"`
	/**
	* MAC address to match against the source MAC address of an outgoing IPv6 packet.
	*/
	Srcmac string `json:"srcmac,omitempty"`
	/**
	*  Used to define range of Source MAC address. It takes string of 0 and 1, 0s are for exact match and 1s for wildcard. For matching first 3 bytes of MAC address, srcMacMask value "000000111111". 
	*/
	Srcmacmask string `json:"srcmacmask,omitempty"`
	/**
	* Protocol, identified by protocol name, to match against the protocol of an outgoing IPv6 packet.
	*/
	Protocol string `json:"protocol,omitempty"`
	/**
	* Protocol, identified by protocol number, to match against the protocol of an outgoing IPv6 packet.
	*/
	Protocolnumber uint32 `json:"protocolnumber,omitempty"`
	/**
	* ID of the VLAN. The Citrix ADC compares the PBR6 only to the outgoing packets on the specified VLAN. If you do not specify an interface ID, the appliance compares the PBR6 to the outgoing packets on all VLANs.
	*/
	Vlan uint32 `json:"vlan,omitempty"`
	/**
	* ID of the VXLAN. The Citrix ADC compares the PBR6 only to the outgoing packets on the specified VXLAN. If you do not specify an interface ID, the appliance compares the PBR6 to the outgoing packets on all VXLANs.
	*/
	Vxlan uint32 `json:"vxlan,omitempty"`
	/**
	* ID of an interface. The Citrix ADC compares the PBR6 only to the outgoing packets on the specified interface. If you do not specify a value, the appliance compares the PBR6 to the outgoing packets on all interfaces.
	*/
	Interface string `json:"Interface,omitempty"`
	/**
	* Priority of the PBR6, which determines the order in which it is evaluated relative to the other PBR6s. If you do not specify priorities while creating PBR6s, the PBR6s are evaluated in the order in which they are created.
	*/
	Priority uint32 `json:"priority,omitempty"`
	/**
	* Enable or disable the PBR6. After you apply the PBR6s, the Citrix ADC compares outgoing packets to the enabled PBR6s.
	*/
	State string `json:"state,omitempty"`
	/**
	* Monitor the route specified by the Next Hop parameter.
	*/
	Msr string `json:"msr,omitempty"`
	/**
	* The name of the monitor.(Can be only of type ping or ARP )
	*/
	Monitor string `json:"monitor,omitempty"`
	/**
	* IP address of the next hop router to which to send matching packets if action is set to ALLOW. This next hop should be directly reachable from the appliance.
	*/
	Nexthop bool `json:"nexthop,omitempty"`
	/**
	* The Next Hop IPv6 address.
	*/
	Nexthopval string `json:"nexthopval,omitempty"`
	/**
	* The iptunnel name where packets need to be forwarded upon.
	*/
	Iptunnel string `json:"iptunnel,omitempty"`
	/**
	* The vlan to vxlan mapping to be applied for incoming packets over this pbr tunnel.
	*/
	Vxlanvlanmap string `json:"vxlanvlanmap,omitempty"`
	/**
	* VLAN number to be used for link local nexthop .
	*/
	Nexthopvlan uint32 `json:"nexthopvlan,omitempty"`
	/**
	* The owner node group in a Cluster for this pbr rule. If owner node group is not specified then the pbr rule is treated as Striped pbr rule.
	*/
	Ownergroup string `json:"ownergroup,omitempty"`
	/**
	* To get a detailed view.
	*/
	Detail bool `json:"detail,omitempty"`

	//------- Read only Parameter ---------;

	Kernelstate string `json:"kernelstate,omitempty"`
	Hits string `json:"hits,omitempty"`
	Curstate string `json:"curstate,omitempty"`
	Totalprobes string `json:"totalprobes,omitempty"`
	Totalfailedprobes string `json:"totalfailedprobes,omitempty"`
	Failedprobes string `json:"failedprobes,omitempty"`
	Monstatcode string `json:"monstatcode,omitempty"`
	Monstatparam1 string `json:"monstatparam1,omitempty"`
	Monstatparam2 string `json:"monstatparam2,omitempty"`
	Monstatparam3 string `json:"monstatparam3,omitempty"`
	Data string `json:"data,omitempty"`

}
