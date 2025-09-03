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
* Configuration for "VXLAN" resource.
*/
type Vxlan struct {
	/**
	* A positive integer, which is also called VXLAN Network Identifier (VNI), that uniquely identifies a VXLAN.
	*/
	Id int `json:"id,omitempty"`
	/**
	* ID of VLANs whose traffic is allowed over this VXLAN. If you do not specify any VLAN IDs, the Citrix ADC allows traffic of all VLANs that are not part of any other VXLANs.
	*/
	Vlan int `json:"vlan,omitempty"`
	/**
	* Specifies UDP destination port for VXLAN packets.
	*/
	Port int `json:"port,omitempty"`
	/**
	* Enable dynamic routing on this VXLAN.
	*/
	Dynamicrouting string `json:"dynamicrouting,omitempty"`
	/**
	* Enable all IPv6 dynamic routing protocols on this VXLAN. Note: For the ENABLED setting to work, you must configure IPv6 dynamic routing protocols from the VTYSH command line.
	*/
	Ipv6dynamicrouting string `json:"ipv6dynamicrouting,omitempty"`
	/**
	* VXLAN encapsulation type. VXLAN, VXLANGPE
	*/
	Type string `json:"type,omitempty"`
	/**
	* VXLAN-GPE next protocol. RESERVED, IPv4, IPv6, ETHERNET, NSH
	*/
	Protocol string `json:"protocol,omitempty"`
	/**
	* Specifies whether Citrix ADC should generate VXLAN packets with inner VLAN tag.
	*/
	Innervlantagging string `json:"innervlantagging,omitempty"`

	//------- Read only Parameter ---------;

	Td string `json:"td,omitempty"`
	Partitionname string `json:"partitionname,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
