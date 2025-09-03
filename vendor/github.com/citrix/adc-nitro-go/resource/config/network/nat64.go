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
* Configuration for nat64 config resource.
*/
type Nat64 struct {
	/**
	* Name for the NAT64 rule. Must begin with a letter, number, or the underscore character (_), and can consist of letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at sign (@), equals (=), colon (:), and underscore characters. Cannot be changed after the rule is created. Choose a name that helps identify the NAT64 rule.
	*/
	Name string `json:"name,omitempty"`
	/**
	* Name of any configured ACL6 whose action is ALLOW.  IPv6 Packets matching the condition of this ACL6 rule and destination IP address of these packets matching the NAT64 IPv6 prefix are considered for NAT64 translation.
	*/
	Acl6name string `json:"acl6name,omitempty"`
	/**
	* Name of the configured netprofile. The Citrix ADC selects one of the IP address in the netprofile as the source IP address of the translated IPv4 packet to be sent to the IPv4 server.
	*/
	Netprofile string `json:"netprofile,omitempty"`

	//------- Read only Parameter ---------;

	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
