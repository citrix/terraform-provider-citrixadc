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
* Configuration for ACL6 entry resource.
*/
type Nsacl6 struct {
	/**
	* Name for the ACL6 rule. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.
	*/
	Acl6name string `json:"acl6name,omitempty"`
	/**
	* Action to perform on the incoming IPv6 packets that match the ACL6 rule.
		Available settings function as follows:
		* ALLOW - The Citrix ADC processes the packet.
		* BRIDGE - The Citrix ADC bridges the packet to the destination without processing it.
		* DENY - The Citrix ADC drops the packet.
	*/
	Acl6action string `json:"acl6action,omitempty"`
	/**
	* Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.
	*/
	Td int `json:"td,omitempty"`
	/**
	* IP address or range of IP addresses to match against the source IP address of an incoming IPv6 packet. In the command line interface, separate the range with a hyphen.
	*/
	Srcipv6 bool `json:"srcipv6,omitempty"`
	/**
	* Either the equals (=) or does not equal (!=) logical operator.
	*/
	Srcipop string `json:"srcipop,omitempty"`
	/**
	* Source IPv6 address (range).
	*/
	Srcipv6val string `json:"srcipv6val,omitempty"`
	/**
	* Port number or range of port numbers to match against the source port number of an incoming IPv6 packet. In the command line interface, separate the range with a hyphen. For example: 40-90.
		Note: The destination port can be specified only for TCP and UDP protocols.
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
	* IP address or range of IP addresses to match against the destination IP address of an incoming IPv6 packet.  In the command line interface, separate the range with a hyphen.
	*/
	Destipv6 bool `json:"destipv6,omitempty"`
	/**
	* Either the equals (=) or does not equal (!=) logical operator.
	*/
	Destipop string `json:"destipop,omitempty"`
	/**
	* Destination IPv6 address (range).
	*/
	Destipv6val string `json:"destipv6val,omitempty"`
	/**
	* Port number or range of port numbers to match against the destination port number of an incoming IPv6 packet. In the command line interface, separate the range with a hyphen. For example: 40-90.
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
	* Time to expire this ACL6 (in seconds).
	*/
	Ttl int `json:"ttl,omitempty"`
	/**
	* MAC address to match against the source MAC address of an incoming IPv6 packet.
	*/
	Srcmac string `json:"srcmac,omitempty"`
	/**
	*  Used to define range of Source MAC address. It takes string of 0 and 1, 0s are for exact match and 1s for wildcard. For matching first 3 bytes of MAC address, srcMacMask value "000000111111". 
	*/
	Srcmacmask string `json:"srcmacmask,omitempty"`
	/**
	* Protocol, identified by protocol name, to match against the protocol of an incoming IPv6 packet.
	*/
	Protocol string `json:"protocol,omitempty"`
	/**
	* Protocol, identified by protocol number, to match against the protocol of an incoming IPv6 packet.
	*/
	Protocolnumber int `json:"protocolnumber,omitempty"`
	/**
	* ID of the VLAN. The Citrix ADC applies the ACL6 rule only to the incoming packets on the specified VLAN. If you do not specify a VLAN ID, the appliance applies the ACL6 rule to the incoming packets on all VLANs.
	*/
	Vlan int `json:"vlan,omitempty"`
	/**
	* ID of the VXLAN. The Citrix ADC applies the ACL6 rule only to the incoming packets on the specified VXLAN. If you do not specify a VXLAN ID, the appliance applies the ACL6 rule to the incoming packets on all VXLANs.
	*/
	Vxlan int `json:"vxlan,omitempty"`
	/**
	* ID of an interface. The Citrix ADC applies the ACL6 rule only to the incoming packets from the specified interface. If you do not specify any value, the appliance applies the ACL6 rule to the incoming packets from all interfaces.
	*/
	Interface string `json:"Interface,omitempty"`
	/**
	* Allow only incoming TCP packets that have the ACK or RST bit set if the action set for the ACL6 rule is ALLOW and these packets match the other conditions in the ACL6 rule.
	*/
	Established bool `json:"established,omitempty"`
	/**
	* ICMP Message type to match against the message type of an incoming IPv6 ICMP packet. For example, to block DESTINATION UNREACHABLE messages, you must specify 3 as the ICMP type.
		Note: This parameter can be specified only for the ICMP protocol.
	*/
	Icmptype int `json:"icmptype,omitempty"`
	/**
	* Code of a particular ICMP message type to match against the ICMP code of an incoming IPv6 ICMP packet.  For example, to block DESTINATION HOST UNREACHABLE messages, specify 3 as the ICMP type and 1 as the ICMP code.
		If you set this parameter, you must set the ICMP Type parameter.
	*/
	Icmpcode int `json:"icmpcode,omitempty"`
	/**
	* Priority for the ACL6 rule, which determines the order in which it is evaluated relative to the other ACL6 rules. If you do not specify priorities while creating ACL6 rules, the ACL6 rules are evaluated in the order in which they are created.
	*/
	Priority int `json:"priority,omitempty"`
	/**
	* State of the ACL6.
	*/
	State string `json:"state,omitempty"`
	/**
	*  Type of the acl6 ,default will be CLASSIC.
		Available options as follows:
		* CLASSIC - specifies the regular extended acls.
		* DFD - cluster specific acls,specifies hashmethod for steering of the packet in cluster .
	*/
	Type string `json:"type,omitempty"`
	/**
	* Specifies the type of hashmethod to be applied, to steer the packet to the FP of the packet.
	*/
	Dfdhash string `json:"dfdhash,omitempty"`
	/**
	* hashprefix to be applied to SIP/DIP to generate rsshash FP.eg 128 => hash calculated on the complete IP
	*/
	Dfdprefix int `json:"dfdprefix,omitempty"`
	/**
	* If stateful option is enabled, transparent sessions are created for the traffic hitting this ACL6 and not hitting any other features like LB, INAT etc. 
	*/
	Stateful string `json:"stateful,omitempty"`
	/**
	* Enable or disable logging of events related to the ACL6 rule. The log messages are stored in the configured syslog or auditlog server.
	*/
	Logstate string `json:"logstate,omitempty"`
	/**
	* Maximum number of log messages to be generated per second. If you set this parameter, you must enable the Log State parameter.
	*/
	Ratelimit int `json:"ratelimit,omitempty"`
	/**
	* Action associated with the ACL6.
	*/
	Aclaction string `json:"aclaction,omitempty"`
	/**
	* New name for the ACL6 rule. Must begin with an ASCII alphabetic or underscore \(_\) character, and must contain only ASCII alphanumeric, underscore, hash \(\#\), period \(.\), space, colon \(:\), at \(@\), equals \(=\), and hyphen \(-\) characters.
	*/
	Newname string `json:"newname,omitempty"`

	//------- Read only Parameter ---------;

	Kernelstate string `json:"kernelstate,omitempty"`
	Hits string `json:"hits,omitempty"`
	Aclassociate string `json:"aclassociate,omitempty"`

}
