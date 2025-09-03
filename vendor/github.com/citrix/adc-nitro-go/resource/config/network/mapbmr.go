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
* Configuration for MAP-T Basic Mapping rule resource.
*/
type Mapbmr struct {
	/**
	* Name for the Basic Mapping Rule. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the  MAP Basic Mapping Rule is created. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "add network MapBmr bmr1 -natprefix 2005::/64 -EAbitLength 16 -psidoffset 6 -portsharingratio 8" ).
		The Basic Mapping Rule information allows a MAP BR to determine source IPv4 address from the IPv6 packet sent from MAP CE device.
		Also it allows to determine destination IPv6 address of MAP CE before sending packets to MAP CE
	*/
	Name string `json:"name,omitempty"`
	/**
	* IPv6 prefix of Customer Edge(CE) device.MAP-T CE will send ipv6 packets with this ipv6 prefix as source ipv6 address prefix
	*/
	Ruleipv6prefix string `json:"ruleipv6prefix,omitempty"`
	/**
	* Start bit position  of Port Set Identifier(PSID) value in Embedded Address (EA) bits.
	*/
	Psidoffset int `json:"psidoffset,omitempty"`
	/**
	* The Embedded Address (EA) bit field encodes the CE-specific IPv4 address and port information.  The EA bit field, which is unique for a
		given Rule IPv6 prefix.
	*/
	Eabitlength int `json:"eabitlength,omitempty"`
	/**
	* Length of Port Set IdentifierPort Set Identifier(PSID) in Embedded Address (EA) bits
	*/
	Psidlength int `json:"psidlength,omitempty"`

	//------- Read only Parameter ---------;

	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
