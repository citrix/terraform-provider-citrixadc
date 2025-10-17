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
* Binding class showing the nsip6 that can be bound to vxlan.
*/
type Vxlannsip6binding struct {
	/**
	* The IP address assigned to the VXLAN.
	*/
	Ipaddress string `json:"ipaddress,omitempty"`
	/**
	* A positive integer, which is also called VXLAN Network Identifier (VNI), that uniquely identifies a VXLAN.
	*/
	Id *int `json:"id,omitempty"`
	/**
	* Subnet mask for the network address defined for this VXLAN.
	*/
	Netmask string `json:"netmask,omitempty"`


}