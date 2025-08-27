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
* Configuration for ACL entry resource.
*/
type Nsacl struct {
	/**
	* Name for the extended ACL rule. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.
	*/
	Aclname string `json:"aclname,omitempty"`
	/**
	* Action to perform on incoming IPv4 packets that match the extended ACL rule.
		Available settings function as follows:
		* ALLOW - The Citrix ADC processes the packet.
		* BRIDGE - The Citrix ADC bridges the packet to the destination without processing it.
		* DENY - The Citrix ADC drops the packet.
	*/
	Aclaction string `json:"aclaction,omitempty"`
	/**
	* Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.
	*/
	Td int `json:"td,omitempty"`
	/**
	* IP address or range of IP addresses to match against the source IP address of an incoming IPv4 packet. In the command line interface, separate the range with a hyphen. For example: 10.102.29.30-10.102.29.189.
	*/
	Srcip bool `json:"srcip,omitempty"`
	/**
	* Either the equals (=) or does not equal (!=) logical operator.
	*/
	Srcipop string `json:"srcipop,omitempty"`
	/**
	* IP address or range of IP addresses to match against the source IP address of an incoming IPv4 packet. In the command line interface, separate the range with a hyphen. For example:10.102.29.30-10.102.29.189.
	*/
	Srcipval string `json:"srcipval,omitempty"`
	/**
	* Policy dataset which can have multiple IP ranges bound to it.
	*/
	Srcipdataset string `json:"srcipdataset,omitempty"`
	/**
	* Port number or range of port numbers to match against the source port number of an incoming IPv4 packet. In the command line interface, separate the range with a hyphen. For example: 40-90.
	*/
	Srcport bool `json:"srcport,omitempty"`
	/**
	* Either the equals (=) or does not equal (!=) logical operator.
	*/
	Srcportop string `json:"srcportop,omitempty"`
	/**
	* Port number or range of port numbers to match against the source port number of an incoming IPv4 packet. In the command line interface, separate the range with a hyphen. For example: 40-90.
	*/
	Srcportval string `json:"srcportval,omitempty"`
	/**
	* Policy dataset which can have multiple port ranges bound to it.
	*/
	Srcportdataset string `json:"srcportdataset,omitempty"`
	/**
	* IP address or range of IP addresses to match against the destination IP address of an incoming IPv4 packet.  In the command line interface, separate the range with a hyphen. For example: 10.102.29.30-10.102.29.189.
	*/
	Destip bool `json:"destip,omitempty"`
	/**
	* Either the equals (=) or does not equal (!=) logical operator.
	*/
	Destipop string `json:"destipop,omitempty"`
	/**
	* IP address or range of IP addresses to match against the destination IP address of an incoming IPv4 packet.  In the command line interface, separate the range with a hyphen. For example: 10.102.29.30-10.102.29.189.
	*/
	Destipval string `json:"destipval,omitempty"`
	/**
	* Policy dataset which can have multiple IP ranges bound to it.
	*/
	Destipdataset string `json:"destipdataset,omitempty"`
	/**
	* Port number or range of port numbers to match against the destination port number of an incoming IPv4 packet. In the command line interface, separate the range with a hyphen. For example: 40-90.
		Note: The destination port can be specified only for TCP and UDP protocols.
	*/
	Destport bool `json:"destport,omitempty"`
	/**
	* Either the equals (=) or does not equal (!=) logical operator.
	*/
	Destportop string `json:"destportop,omitempty"`
	/**
	* Port number or range of port numbers to match against the destination port number of an incoming IPv4 packet. In the command line interface, separate the range with a hyphen. For example: 40-90.
		Note: The destination port can be specified only for TCP and UDP protocols.
	*/
	Destportval string `json:"destportval,omitempty"`
	/**
	* Policy dataset which can have multiple port ranges bound to it.
	*/
	Destportdataset string `json:"destportdataset,omitempty"`
	/**
	* Number of seconds, in multiples of four, after which the extended ACL rule expires. If you do not want the extended ACL rule to expire, do not specify a TTL value.
	*/
	Ttl int `json:"ttl,omitempty"`
	/**
	* MAC address to match against the source MAC address of an incoming IPv4 packet.
	*/
	Srcmac string `json:"srcmac,omitempty"`
	/**
	*  Used to define range of Source MAC address. It takes string of 0 and 1, 0s are for exact match and 1s for wildcard. For matching first 3 bytes of MAC address, srcMacMask value "000000111111". 
	*/
	Srcmacmask string `json:"srcmacmask,omitempty"`
	/**
	* Protocol to match against the protocol of an incoming IPv4 packet.
	*/
	Protocol string `json:"protocol,omitempty"`
	/**
	* Protocol to match against the protocol of an incoming IPv4 packet.
	*/
	Protocolnumber int `json:"protocolnumber,omitempty"`
	/**
	* ID of the VLAN. The Citrix ADC applies the ACL rule only to the incoming packets of the specified VLAN. If you do not specify a VLAN ID, the appliance applies the ACL rule to the incoming packets on all VLANs.
	*/
	Vlan int `json:"vlan,omitempty"`
	/**
	* ID of the VXLAN. The Citrix ADC applies the ACL rule only to the incoming packets of the specified VXLAN. If you do not specify a VXLAN ID, the appliance applies the ACL rule to the incoming packets on all VXLANs.
	*/
	Vxlan int `json:"vxlan,omitempty"`
	/**
	* ID of an interface. The Citrix ADC applies the ACL rule only to the incoming packets from the specified interface. If you do not specify any value, the appliance applies the ACL rule to the incoming packets of all interfaces.
	*/
	Interface string `json:"Interface,omitempty"`
	/**
	* Allow only incoming TCP packets that have the ACK or RST bit set, if the action set for the ACL rule is ALLOW and these packets match the other conditions in the ACL rule.
	*/
	Established bool `json:"established,omitempty"`
	/**
	* ICMP Message type to match against the message type of an incoming ICMP packet. For example, to block DESTINATION UNREACHABLE messages, you must specify 3 as the ICMP type.
		Note: This parameter can be specified only for the ICMP protocol.
	*/
	Icmptype int `json:"icmptype,omitempty"`
	/**
	* Code of a particular ICMP message type to match against the ICMP code of an incoming ICMP packet.  For example, to block DESTINATION HOST UNREACHABLE messages, specify 3 as the ICMP type and 1 as the ICMP code.
		If you set this parameter, you must set the ICMP Type parameter.
	*/
	Icmpcode int `json:"icmpcode,omitempty"`
	/**
	* Priority for the extended ACL rule that determines the order in which it is evaluated relative to the other extended ACL rules. If you do not specify priorities while creating extended ACL rules, the ACL rules are evaluated in the order in which they are created.
	*/
	Priority int `json:"priority,omitempty"`
	/**
	* Enable or disable the extended ACL rule. After you apply the extended ACL rules, the Citrix ADC compares incoming packets against the enabled extended ACL rules.
	*/
	State string `json:"state,omitempty"`
	/**
	* Enable or disable logging of events related to the extended ACL rule. The log messages are stored in the configured syslog or auditlog server.
	*/
	Logstate string `json:"logstate,omitempty"`
	/**
	* Maximum number of log messages to be generated per second. If you set this parameter, you must enable the Log State parameter.
	*/
	Ratelimit int `json:"ratelimit,omitempty"`
	/**
	*  Type of the acl ,default will be CLASSIC.
		Available options as follows:
		* CLASSIC - specifies the regular extended acls.
		* DFD - cluster specific acls,specifies hashmethod for steering of the packet in cluster .
	*/
	Type string `json:"type,omitempty"`
	/**
	* Specifies the type hashmethod to be applied, to steer the packet to the FP of the packet.
	*/
	Dfdhash string `json:"dfdhash,omitempty"`
	/**
	* Specifies the NodeId to steer the packet to the provided FP.
	*/
	Nodeid int `json:"nodeid,omitempty"`
	/**
	* If stateful option is enabled, transparent sessions are created for the traffic hitting this ACL and not hitting any other features like LB, INAT etc. 
	*/
	Stateful string `json:"stateful,omitempty"`
	/**
	* New name for the extended ACL rule. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.
	*/
	Newname string `json:"newname,omitempty"`

	//------- Read only Parameter ---------;

	Hits string `json:"hits,omitempty"`
	Kernelstate string `json:"kernelstate,omitempty"`
	Aclassociate string `json:"aclassociate,omitempty"`
	Aclchildcount string `json:"aclchildcount,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
