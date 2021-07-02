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
* Binding class showing the node that can be bound to clusterinstance.
*/
type Clusterinstancenodebinding struct {
	/**
	* The unique number that identiies a cluster.
	*/
	Nodeid uint32 `json:"nodeid,omitempty"`
	/**
	* The IP Address of the node.
	*/
	Ipaddress string `json:"ipaddress,omitempty"`
	/**
	* Node Health state.
	*/
	Health string `json:"health,omitempty"`
	/**
	* Node clusterd state.
	*/
	Clusterhealth string `json:"clusterhealth,omitempty"`
	/**
	* Node effective health state.
	*/
	Effectivestate string `json:"effectivestate,omitempty"`
	/**
	* Master state.
	*/
	Masterstate string `json:"masterstate,omitempty"`
	/**
	* Active, Spare or Passive. An active node serves traffic. A spare node serves as a backup for active nodes. A passive node does not serve traffic. This may be useful during temporary maintenance activity where it is desirable that the node takes part in the consensus protocol, but not serve traffic.
	*/
	State string `json:"state,omitempty"`
	/**
	* This argument is used to determine whether the node is configuration coordinator (CCO).
	*/
	Isconfigurationcoordinator bool `json:"isconfigurationcoordinator,omitempty"`
	/**
	* This argument is used to determine whether it is local node.
	*/
	Islocalnode bool `json:"islocalnode,omitempty"`
	/**
	* This argument is used to determine if there is a RSS key mismatch at cluster node level.
	*/
	Nodersskeymismatch bool `json:"nodersskeymismatch,omitempty"`
	/**
	* This argument is used to determine if there is a License mismatch at cluster node level.
	*/
	Nodelicensemismatch bool `json:"nodelicensemismatch,omitempty"`
	/**
	* This argument is used to determine if Jumbo framework not supported at cluster node level.
	*/
	Nodejumbonotsupported bool `json:"nodejumbonotsupported,omitempty"`
	/**
	* Unique number that identifies the cluster.
	*/
	Clid uint32 `json:"clid,omitempty"`


}