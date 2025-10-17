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
* Configuration for session forward resource.
*/
type Forwardingsession struct {
	/**
	* Name for the forwarding session rule. Can begin with a letter, number, or the underscore character (_), and can consist of letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the rule is created.
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my rule" or 'my rule').
	*/
	Name string `json:"name,omitempty"`
	/**
	* An IPv4 network address or IPv6 prefix of a network from which the forwarded traffic originates or to which it is destined.
	*/
	Network string `json:"network,omitempty"`
	/**
	* Subnet mask associated with the network.
	*/
	Netmask string `json:"netmask,omitempty"`
	/**
	* Name of any configured ACL6 whose action is ALLOW. The rule of the ACL6 is used as a forwarding session rule.
	*/
	Acl6name string `json:"acl6name,omitempty"`
	/**
	* Name of any configured ACL whose action is ALLOW. The rule of the ACL is used as a forwarding session rule.
	*/
	Aclname string `json:"aclname,omitempty"`
	/**
	* Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.
	*/
	Td *int `json:"td,omitempty"`
	/**
	* Synchronize connection information with the secondary appliance in a high availability (HA) pair. That is, synchronize all connection-related information for the forwarding session.
	*/
	Connfailover string `json:"connfailover,omitempty"`
	/**
	* Cache the source ip address and mac address of the DA servers.
	*/
	Sourceroutecache string `json:"sourceroutecache,omitempty"`
	/**
	* Enabling this option on forwarding session will not steer the packet to flow processor. Instead, packet will be routed.
	*/
	Processlocal string `json:"processlocal,omitempty"`

	//------- Read only Parameter ---------;

	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
