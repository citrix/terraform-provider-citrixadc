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
* Configuration for IPv6 RNAT configured route resource.
*/
type Rnat6 struct {
	/**
	* Name for the RNAT6 rule. Must begin with a letter, number, or the underscore character (_), and can consist of letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at sign (@), equals (=), colon (:), and underscore characters. Cannot be changed after the rule is created. Choose a name that helps identify the RNAT6 rule.
	*/
	Name string `json:"name,omitempty"`
	/**
	* IPv6 address of the network on whose traffic you want the Citrix ADC to do RNAT processing.
	*/
	Network string `json:"network,omitempty"`
	/**
	* Name of any configured ACL6 whose action is ALLOW. The rule of the ACL6 is used as an RNAT6 rule.
	*/
	Acl6name string `json:"acl6name,omitempty"`
	/**
	* Port number to which the IPv6 packets are redirected. Applicable to TCP and UDP protocols.
	*/
	Redirectport *int `json:"redirectport,omitempty"`
	/**
	* Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.
	*/
	Td *int `json:"td,omitempty"`
	/**
	* Enable source ip persistency, which enables the Citrix ADC to use the RNAT ips using source ip.
	*/
	Srcippersistency string `json:"srcippersistency,omitempty"`
	/**
	* The owner node group in a Cluster for this rnat rule.
	*/
	Ownergroup string `json:"ownergroup,omitempty"`

	//------- Read only Parameter ---------;

	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
