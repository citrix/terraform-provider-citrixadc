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

package network

/**
* Configuration for RNAT configured route resource.
*/
type Rnat struct {
	/**
	* The network address defined for the RNAT entry.
	*/
	Network string `json:"network,omitempty"`
	/**
	* The subnet mask for the network address.
	*/
	Netmask string `json:"netmask,omitempty"`
	/**
	* An extended ACL defined for the RNAT entry.
	*/
	Aclname string `json:"aclname,omitempty"`
	/**
	* Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.
	*/
	Td int `json:"td,omitempty"`
	/**
	* The owner node group in a Cluster for this rnat rule.
	*/
	Ownergroup string `json:"ownergroup,omitempty"`
	/**
	* Name for the RNAT4 rule. Must begin with a letter, number, or the underscore character (_), and can consist of letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at sign (@), equals (=), colon (:), and underscore characters. Cannot be changed after the rule is created. Choose a name that helps identify the RNAT4 rule.
	*/
	Name string `json:"name,omitempty"`
	/**
	* Port number to which the IPv4 packets are redirected. Applicable to TCP and UDP protocols.
	*/
	Redirectport int `json:"redirectport,omitempty"`
	/**
	* Any NetScaler-owned IPv4 address except the NSIP address. The NetScaler appliance replaces the source IP addresses of server-generated packets with the IP address specified. The IP address must be a public NetScaler-owned IP address. If you specify multiple addresses for this field, NATIP selection uses the round robin algorithm for each session. By specifying a range of IP addresses, you can specify all NetScaler-owned IP addresses, except the NSIP, that fall within the specified range.
	*/
	Natip string `json:"natip,omitempty"`
	/**
	* Enables the Citrix ADC to use the same NAT IP address for all RNAT sessions initiated from a particular server.
	*/
	Srcippersistency string `json:"srcippersistency,omitempty"`
	/**
	* Enable source port proxying, which enables the Citrix ADC to use the RNAT ips using proxied source port.
	*/
	Useproxyport string `json:"useproxyport,omitempty"`
	/**
	* Synchronize connection information with the secondary appliance in a high availability (HA) pair. That is, synchronize all connection-related information for the RNAT session. In order for this to work, tcpproxy should be DISABLED. To disable tcpproxy use "set rnatparam tcpproxy DISABLED"
	*/
	Connfailover string `json:"connfailover,omitempty"`
	/**
	* New name for the RNAT4 rule. Must begin with an ASCII alphabetic or underscore (_) character, and must contain       only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.
	*/
	Newname string `json:"newname,omitempty"`

}
