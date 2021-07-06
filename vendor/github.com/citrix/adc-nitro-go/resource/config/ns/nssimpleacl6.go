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
* Configuration for simple ACL6 resource.
*/
type Nssimpleacl6 struct {
	/**
	* Name for the simple ACL6 rule. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the simple ACL6 rule is created.
	*/
	Aclname string `json:"aclname,omitempty"`
	/**
	* Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.
	*/
	Td int `json:"td,omitempty"`
	/**
	* Drop incoming IPv6 packets that match the simple ACL6 rule.
	*/
	Aclaction string `json:"aclaction,omitempty"`
	/**
	* IP address to match against the source IP address of an incoming IPv6 packet.
	*/
	Srcipv6 string `json:"srcipv6,omitempty"`
	/**
	* Port number to match against the destination port number of an incoming IPv6 packet.
		DestPort is mandatory while setting Protocol. Omitting the port number and protocol creates an all-ports  and all protocol simple ACL6 rule, which matches any port and any protocol. In that case, you cannot create another simple ACL6 rule specifying a specific port and the same source IPv6 address.
	*/
	Destport int `json:"destport,omitempty"`
	/**
	* Protocol to match against the protocol of an incoming IPv6 packet. You must set this parameter if you set the Destination Port parameter.
	*/
	Protocol string `json:"protocol,omitempty"`
	/**
	* Number of seconds, in multiples of four, after which the simple ACL6 rule expires. If you do not want the simple ACL6 rule to expire, do not specify a TTL value.
	*/
	Ttl int `json:"ttl,omitempty"`
	Estsessions bool `json:"estsessions,omitempty"`

	//------- Read only Parameter ---------;

	Hits string `json:"hits,omitempty"`

}
