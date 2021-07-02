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

package cluster

/**
* Configuration for cluster node resource.
*/
type Clusternode struct {
	/**
	* Unique number that identifies the cluster node.
	*/
	Nodeid uint32 `json:"nodeid,omitempty"`
	/**
	* Citrix ADC IP (NSIP) address of the appliance to add to the cluster. Must be an IPv4 address.
	*/
	Ipaddress string `json:"ipaddress,omitempty"`
	/**
	* Admin state of the cluster node. The available settings function as follows:
		ACTIVE - The node serves traffic.
		SPARE - The node does not serve traffic unless an ACTIVE node goes down.
		PASSIVE - The node does not serve traffic, unless you change its state. PASSIVE state is useful during temporary maintenance activities in which you want the node to take part in the consensus protocol but not to serve traffic.
	*/
	State string `json:"state,omitempty"`
	/**
	* Interface through which the node communicates with the other nodes in the cluster. Must be specified in the three-tuple form n/c/u, where n represents the node ID and c/u refers to the interface on the appliance.
	*/
	Backplane string `json:"backplane,omitempty"`
	/**
	* Preference for selecting a node as the configuration coordinator. The node with the lowest priority value is selected as the configuration coordinator.
		When the current configuration coordinator goes down, the node with the next lowest priority is made the new configuration coordinator. When the original node comes back up, it will preempt the new configuration coordinator and take over as the configuration coordinator.
		Note: When priority is not configured for any of the nodes or if multiple nodes have the same priority, the cluster elects one of the nodes as the configuration coordinator.
	*/
	Priority uint32 `json:"priority,omitempty"`
	/**
	* The default node group in a Cluster system.
	*/
	Nodegroup string `json:"nodegroup,omitempty"`
	/**
	* Applicable for Passive node and node becomes passive after this timeout (in minutes)
	*/
	Delay uint32 `json:"delay,omitempty"`
	/**
	* To set the tunnel mode
	*/
	Tunnelmode string `json:"tunnelmode,omitempty"`
	/**
	* Option to remove nodegroup config
	*/
	Clearnodegroupconfig string `json:"clearnodegroupconfig,omitempty"`

	//------- Read only Parameter ---------;

	Clusterhealth string `json:"clusterhealth,omitempty"`
	Effectivestate string `json:"effectivestate,omitempty"`
	Operationalsyncstate string `json:"operationalsyncstate,omitempty"`
	Syncfailurereason string `json:"syncfailurereason,omitempty"`
	Masterstate string `json:"masterstate,omitempty"`
	Health string `json:"health,omitempty"`
	Syncstate string `json:"syncstate,omitempty"`
	Isconfigurationcoordinator string `json:"isconfigurationcoordinator,omitempty"`
	Islocalnode string `json:"islocalnode,omitempty"`
	Nodersskeymismatch string `json:"nodersskeymismatch,omitempty"`
	Nodelicensemismatch string `json:"nodelicensemismatch,omitempty"`
	Nodejumbonotsupported string `json:"nodejumbonotsupported,omitempty"`
	Nodelist string `json:"nodelist,omitempty"`
	Ifaceslist string `json:"ifaceslist,omitempty"`
	Enabledifaces string `json:"enabledifaces,omitempty"`
	Disabledifaces string `json:"disabledifaces,omitempty"`
	Partialfailifaces string `json:"partialfailifaces,omitempty"`
	Hamonifaces string `json:"hamonifaces,omitempty"`
	Name string `json:"name,omitempty"`
	Cfgflags string `json:"cfgflags,omitempty"`
	Routemonitor string `json:"routemonitor,omitempty"`
	Netmask string `json:"netmask,omitempty"`

}
