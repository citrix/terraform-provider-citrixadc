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
* Configuration for MAP-T Default Mapping rule resource.
*/
type Mapdmr struct {
	/**
	* Name for the Default Mapping Rule. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the MAP Default Mapping Rule is created. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "add network MapDmr map1 -BRIpv6Prefix 2003::/96").
		Default Mapping Rule (DMR) is defined in terms of the IPv6 prefix advertised by one or more BRs, which provide external connectivity.  A typical MAP-T CE will install an IPv4 default route using this rule.  A BR will use this rule when translating all outside IPv4 source addresses to the IPv6 MAP domain.
	*/
	Name string `json:"name,omitempty"`
	/**
	* IPv6 prefix of Border Relay (Citrix ADC) device.MAP-T CE will send ipv6 packets to this ipv6 prefix.The DMR IPv6 prefix length SHOULD be 64 bits long by default and in any case MUST NOT exceed 96 bits
	*/
	Bripv6prefix string `json:"bripv6prefix,omitempty"`

	//------- Read only Parameter ---------;

	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
