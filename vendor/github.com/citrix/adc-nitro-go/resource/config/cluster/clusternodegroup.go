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
* Configuration for Node group object type resource.
*/
type Clusternodegroup struct {
	/**
	* Name of the nodegroup. The name uniquely identifies the nodegroup on the cluster.
	*/
	Name string `json:"name,omitempty"`
	/**
	* Specifies whether cluster nodes, that are not part of the nodegroup, will be used as backup for the nodegroup.
		* Enabled - When one of the nodes goes down, no other cluster node is picked up to replace it. When the node comes up, it will continue being part of the nodegroup.
		* Disabled - When one of the nodes goes down, a non-nodegroup cluster node is picked up and acts as part of the nodegroup. When the original node of the nodegroup comes up, the backup node will be replaced.
	*/
	Strict string `json:"strict,omitempty"`
	/**
	* Only one node can be bound to nodegroup with this option enabled. It specifies whether to prempt the traffic for the entities bound to nodegroup when owner node goes down and rejoins the cluster.
		* Enabled - When owner node goes down, backup node will become the owner node and takes the traffic for the entities bound to the nodegroup. When bound node rejoins the cluster, traffic for the entities bound to nodegroup will not be steered back to this bound node. Current owner will have the ownership till it goes down.
		* Disabled - When one of the nodes goes down, a non-nodegroup cluster node is picked up and acts as part of the nodegroup. When the original node of the nodegroup comes up, the backup node will be replaced.
	*/
	Sticky string `json:"sticky,omitempty"`
	/**
	* State of the nodegroup. All the nodes binding to this nodegroup must have the same state. ACTIVE/SPARE/PASSIVE
	*/
	State string `json:"state,omitempty"`
	/**
	* Priority of Nodegroup. This priority is used for all the nodes bound to the nodegroup for Nodegroup selection
	*/
	Priority int `json:"priority,omitempty"`

	//------- Read only Parameter ---------;

	Currentnodemask string `json:"currentnodemask,omitempty"`
	Backupnodemask string `json:"backupnodemask,omitempty"`
	Boundedentitiescntfrompe string `json:"boundedentitiescntfrompe,omitempty"`
	Activelist string `json:"activelist,omitempty"`
	Backuplist string `json:"backuplist,omitempty"`

}
