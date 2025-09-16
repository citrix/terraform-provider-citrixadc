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
* Configuration for Policy Based Routing(PBR) entry resource.
*/
type Nspbr struct {
	/**
	* Name for the PBR. Must begin with an ASCII alphabetic or underscore \(_\) character, and must contain only ASCII alphanumeric, underscore, hash \(\#\), period \(.\), space, colon \(:\), at \(@\), equals \(=\), and hyphen \(-\) characters. Cannot be changed after the PBR is created.
	*/
	Name string `json:"name,omitempty"`
	/**
	* Action to perform on the outgoing IPv4 packets that match the PBR.
		Available settings function as follows:
		* ALLOW - The Citrix ADC sends the packet to the designated next-hop router.
		* DENY - The Citrix ADC applies the routing table for normal destination-based routing.
	*/
	Action string `json:"action,omitempty"`
	/**
	* Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.
	*/
	Td int `json:"td,omitempty"`
	/**
	* IP address or range of IP addresses to match against the source IP address of an outgoing IPv4 packet. In the command line interface, separate the range with a hyphen. For example: 10.102.29.30-10.102.29.189.
	*/
	Srcip bool `json:"srcip,omitempty"`
	/**
	* Either the equals (=) or does not equal (!=) logical operator.
	*/
	Srcipop string `json:"srcipop,omitempty"`
	/**
	* IP address or range of IP addresses to match against the source IP address of an outgoing IPv4 packet. In the command line interface, separate the range with a hyphen. For example: 10.102.29.30-10.102.29.189.
	*/
	Srcipval string `json:"srcipval,omitempty"`
	/**
	* Port number or range of port numbers to match against the source port number of an outgoing IPv4 packet. In the command line interface, separate the range with a hyphen. For example: 40-90.
		Note: The destination port can be specified only for TCP and UDP protocols.
	*/
	Srcport bool `json:"srcport,omitempty"`
	/**
	* Either the equals (=) or does not equal (!=) logical operator.
	*/
	Srcportop string `json:"srcportop,omitempty"`
	/**
	* Port number or range of port numbers to match against the source port number of an outgoing IPv4 packet. In the command line interface, separate the range with a hyphen. For example: 40-90.
		Note: The destination port can be specified only for TCP and UDP protocols.
	*/
	Srcportval string `json:"srcportval,omitempty"`
	/**
	* IP address or range of IP addresses to match against the destination IP address of an outgoing IPv4 packet.  In the command line interface, separate the range with a hyphen. For example: 10.102.29.30-10.102.29.189.
	*/
	Destip bool `json:"destip,omitempty"`
	/**
	* Either the equals (=) or does not equal (!=) logical operator.
	*/
	Destipop string `json:"destipop,omitempty"`
	/**
	* IP address or range of IP addresses to match against the destination IP address of an outgoing IPv4 packet. In the command line interface, separate the range with a hyphen. For example: 10.102.29.30-10.102.29.189.
	*/
	Destipval string `json:"destipval,omitempty"`
	/**
	* Port number or range of port numbers to match against the destination port number of an outgoing IPv4 packet. In the command line interface, separate the range with a hyphen. For example: 40-90.
		Note: The destination port can be specified only for TCP and UDP protocols.
	*/
	Destport bool `json:"destport,omitempty"`
	/**
	* Either the equals (=) or does not equal (!=) logical operator.
	*/
	Destportop string `json:"destportop,omitempty"`
	/**
	* Port number or range of port numbers to match against the destination port number of an outgoing IPv4 packet. In the command line interface, separate the range with a hyphen. For example: 40-90.
		Note: The destination port can be specified only for TCP and UDP protocols.
	*/
	Destportval string `json:"destportval,omitempty"`
	/**
	* IP address of the next hop router or the name of the link load balancing virtual server to which to send matching packets if action is set to ALLOW.
		If you specify a link load balancing (LLB) virtual server, which can provide a backup if a next hop link fails, first make sure that the next hops bound to the LLB virtual server are actually next hops that are directly connected to the Citrix ADC. Otherwise, the Citrix ADC throws an error when you attempt to create the PBR. The next hop can be null to represent null routes
	*/
	Nexthop bool `json:"nexthop,omitempty"`
	/**
	* The Next Hop IP address or gateway name.
	*/
	Nexthopval string `json:"nexthopval,omitempty"`
	/**
	* The Tunnel name.
	*/
	Iptunnel bool `json:"iptunnel,omitempty"`
	/**
	* The iptunnel name where packets need to be forwarded upon.
	*/
	Iptunnelname string `json:"iptunnelname,omitempty"`
	/**
	* The vlan to vxlan mapping to be applied for incoming packets over this pbr tunnel
	*/
	Vxlanvlanmap string `json:"vxlanvlanmap,omitempty"`
	/**
	* Integer value that uniquely identifies the traffic domain to which you want to send packet to.
	*/
	Targettd int `json:"targettd,omitempty"`
	/**
	* MAC address to match against the source MAC address of an outgoing IPv4 packet.
	*/
	Srcmac string `json:"srcmac,omitempty"`
	/**
	*  Used to define range of Source MAC address. It takes string of 0 and 1, 0s are for exact match and 1s for wildcard. For matching first 3 bytes of MAC address, srcMacMask value "000000111111". 
	*/
	Srcmacmask string `json:"srcmacmask,omitempty"`
	/**
	* Protocol, identified by protocol name, to match against the protocol of an outgoing IPv4 packet.
	*/
	Protocol string `json:"protocol,omitempty"`
	/**
	* Protocol, identified by protocol number, to match against the protocol of an outgoing IPv4 packet.
	*/
	Protocolnumber int `json:"protocolnumber,omitempty"`
	/**
	* ID of the VLAN. The Citrix ADC compares the PBR only to the outgoing packets on the specified VLAN. If you do not specify any interface ID, the appliance compares the PBR to the outgoing packets on all VLANs.
	*/
	Vlan int `json:"vlan,omitempty"`
	/**
	* ID of the VXLAN. The Citrix ADC compares the PBR only to the outgoing packets on the specified VXLAN. If you do not specify any interface ID, the appliance compares the PBR to the outgoing packets on all VXLANs.
	*/
	Vxlan int `json:"vxlan,omitempty"`
	/**
	* ID of an interface. The Citrix ADC compares the PBR only to the outgoing packets on the specified interface. If you do not specify any value, the appliance compares the PBR to the outgoing packets on all interfaces.
	*/
	Interface string `json:"Interface,omitempty"`
	/**
	* Priority of the PBR, which determines the order in which it is evaluated relative to the other PBRs. If you do not specify priorities while creating PBRs, the PBRs are evaluated in the order in which they are created.
	*/
	Priority int `json:"priority,omitempty"`
	/**
	* Monitor the route specified byte Next Hop parameter. This parameter is not applicable if you specify a link load balancing (LLB) virtual server name with the Next Hop parameter.
	*/
	Msr string `json:"msr,omitempty"`
	/**
	* The name of the monitor.(Can be only of type ping or ARP )
	*/
	Monitor string `json:"monitor,omitempty"`
	/**
	* Enable or disable the PBR. After you apply the PBRs, the Citrix ADC compares outgoing packets to the enabled PBRs.
	*/
	State string `json:"state,omitempty"`
	/**
	* The owner node group in a Cluster for this pbr rule. If ownernode is not specified then the pbr rule is treated as Striped pbr rule.
	*/
	Ownergroup string `json:"ownergroup,omitempty"`
	/**
	* To get a detailed view.
	*/
	Detail bool `json:"detail,omitempty"`

	//------- Read only Parameter ---------;

	Hits string `json:"hits,omitempty"`
	Kernelstate string `json:"kernelstate,omitempty"`
	Curstate string `json:"curstate,omitempty"`
	Totalprobes string `json:"totalprobes,omitempty"`
	Totalfailedprobes string `json:"totalfailedprobes,omitempty"`
	Failedprobes string `json:"failedprobes,omitempty"`
	Monstatcode string `json:"monstatcode,omitempty"`
	Monstatparam1 string `json:"monstatparam1,omitempty"`
	Monstatparam2 string `json:"monstatparam2,omitempty"`
	Monstatparam3 string `json:"monstatparam3,omitempty"`
	Data string `json:"data,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
