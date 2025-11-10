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
* Binding class showing the channel that can be bound to vlan.
*/
type Vlanchannelbinding struct {
	/**
	* The interface to be bound to the VLAN, specified in slot/port notation (for example, 1/3).
	*/
	Ifnum string `json:"ifnum,omitempty"`
	/**
	* Make the interface an 802.1q tagged interface. Packets sent on this interface on this VLAN have an additional 4-byte 802.1q tag, which identifies the VLAN. To use 802.1q tagging, you must also configure the switch connected to the appliance's interfaces.
	*/
	Tagged bool `json:"tagged,omitempty"`
	/**
	* Specifies the virtual LAN ID.
	*/
	Id *int `json:"id,omitempty"`
	/**
	* The owner node group in a Cluster for this vlan.
	*/
	Ownergroup string `json:"ownergroup,omitempty"`


}