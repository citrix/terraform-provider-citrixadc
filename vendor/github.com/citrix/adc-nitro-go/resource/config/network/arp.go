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
* Configuration for arp resource.
*/
type Arp struct {
	/**
	* IP address of the network device that you want to add to the ARP table.
	*/
	Ipaddress string `json:"ipaddress,omitempty"`
	/**
	* Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.
	*/
	Td uint32 `json:"td,omitempty"`
	/**
	* MAC address of the network device.
	*/
	Mac string `json:"mac,omitempty"`
	/**
	* Interface through which the network device is accessible. Specify the interface in (slot/port) notation. For example, 1/3.
	*/
	Ifnum string `json:"ifnum,omitempty"`
	/**
	* ID of the VXLAN on which the IP address of this ARP entry is reachable.
	*/
	Vxlan uint32 `json:"vxlan,omitempty"`
	/**
	* IP address of the VXLAN tunnel endpoint (VTEP) through which the IP address of this ARP entry is reachable.
	*/
	Vtep string `json:"vtep,omitempty"`
	/**
	* The VLAN ID through which packets are to be sent after matching the ARP entry. This is a numeric value.
	*/
	Vlan uint32 `json:"vlan,omitempty"`
	/**
	* The owner node for the Arp entry.
	*/
	Ownernode uint32 `json:"ownernode,omitempty"`
	/**
	* Remove all ARP entries from the ARP table of the Citrix ADC.
	*/
	All bool `json:"all,omitempty"`
	/**
	* Unique number that identifies the cluster node.
	*/
	Nodeid uint32 `json:"nodeid,omitempty"`

	//------- Read only Parameter ---------;

	Timeout string `json:"timeout,omitempty"`
	State string `json:"state,omitempty"`
	Flags string `json:"flags,omitempty"`
	Type string `json:"type,omitempty"`
	Channel string `json:"channel,omitempty"`
	Controlplane string `json:"controlplane,omitempty"`

}
