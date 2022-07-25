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
* Configuration for nd6 resource.
*/
type Nd6 struct {
	/**
	* Link-local IPv6 address of the adjacent network device to add to the ND6 table.
	*/
	Neighbor string `json:"neighbor,omitempty"`
	/**
	* MAC address of the adjacent network device.
	*/
	Mac string `json:"mac,omitempty"`
	/**
	* Interface through which the adjacent network device is available, specified in slot/port notation (for example, 1/3). Use spaces to separate multiple entries.
	*/
	Ifnum string `json:"ifnum,omitempty"`
	/**
	* Integer value that uniquely identifies the VLAN on which the adjacent network device exists.
	*/
	Vlan int `json:"vlan,omitempty"`
	/**
	* ID of the VXLAN on which the IPv6 address of this ND6 entry is reachable.
	*/
	Vxlan int `json:"vxlan,omitempty"`
	/**
	* IP address of the VXLAN tunnel endpoint (VTEP) through which the IPv6 address of this ND6 entry is reachable.
	*/
	Vtep string `json:"vtep,omitempty"`
	/**
	* Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.
	*/
	Td int `json:"td"` // Zero is a valid value
	/**
	* Unique number that identifies the cluster node.
	*/
	Nodeid int `json:"nodeid"` // Zero is a valid value

	//------- Read only Parameter ---------;

	State string `json:"state,omitempty"`
	Timeout string `json:"timeout,omitempty"`
	Flags string `json:"flags,omitempty"`
	Controlplane string `json:"controlplane,omitempty"`
	Channel string `json:"channel,omitempty"`

}
