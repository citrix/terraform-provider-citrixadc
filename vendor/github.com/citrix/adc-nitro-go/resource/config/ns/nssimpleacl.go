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
* Configuration for simple ACL resource.
*/
type Nssimpleacl struct {
	/**
	* Name for the simple ACL rule. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the simple ACL rule is created.
	*/
	Aclname string `json:"aclname,omitempty"`
	/**
	* Drop incoming IPv4 packets that match the simple ACL rule.
	*/
	Aclaction string `json:"aclaction,omitempty"`
	/**
	* Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.
	*/
	Td *int `json:"td,omitempty"`
	/**
	* IP address to match against the source IP address of an incoming IPv4 packet.
	*/
	Srcip string `json:"srcip,omitempty"`
	/**
	* Port number to match against the destination port number of an incoming IPv4 packet.
		DestPort is mandatory while setting Protocol. Omitting the port number and protocol creates an all-ports  and all protocols simple ACL rule, which matches any port and any protocol. In that case, you cannot create another simple ACL rule specifying a specific port and the same source IPv4 address.
	*/
	Destport *int `json:"destport,omitempty"`
	/**
	* Protocol to match against the protocol of an incoming IPv4 packet. You must set this parameter if you have set the Destination Port parameter.
	*/
	Protocol string `json:"protocol,omitempty"`
	/**
	* Number of seconds, in multiples of four, after which the simple ACL rule expires. If you do not want the simple ACL rule to expire, do not specify a TTL value.
	*/
	Ttl *int `json:"ttl,omitempty"`
	Estsessions bool `json:"estsessions,omitempty"`

	//------- Read only Parameter ---------;

	Hits string `json:"hits,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
