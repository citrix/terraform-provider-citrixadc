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
* Configuration for bridge table entry resource.
*/
type Bridgetable struct {
	/**
	* The MAC address of the target.
	*/
	Mac string `json:"mac,omitempty"`
	/**
	* The VXLAN to which this address is associated.
	*/
	Vxlan uint32 `json:"vxlan,omitempty"`
	/**
	* The IP address of the destination VXLAN tunnel endpoint where the Ethernet MAC ADDRESS resides.
	*/
	Vtep string `json:"vtep,omitempty"`
	/**
	* The VXLAN VNI Network Identifier (or VXLAN Segment ID) to use to connect to the remote VXLAN tunnel endpoint.  If omitted the value specified as vxlan will be used.
	*/
	Vni uint32 `json:"vni,omitempty"`
	/**
	* The vlan on which to send multicast packets when the VXLAN tunnel endpoint is a muticast group address.
	*/
	Devicevlan uint32 `json:"devicevlan,omitempty"`
	/**
	* Time-out value for the bridge table entries, in seconds. The new value applies only to the entries that are dynamically learned after the new value is set. Previously existing bridge table entries expire after the previously configured time-out value.
	*/
	Bridgeage uint32 `json:"bridgeage,omitempty"`
	/**
	* Unique number that identifies the cluster node.
	*/
	Nodeid uint32 `json:"nodeid,omitempty"`
	/**
	* VLAN  whose entries are to be removed.
	*/
	Vlan uint32 `json:"vlan,omitempty"`
	/**
	* INTERFACE  whose entries are to be removed.
	*/
	Ifnum string `json:"ifnum,omitempty"`

	//------- Read only Parameter ---------;

	Flags string `json:"flags,omitempty"`
	Type string `json:"type,omitempty"`
	Channel string `json:"channel,omitempty"`
	Controlplane string `json:"controlplane,omitempty"`

}
