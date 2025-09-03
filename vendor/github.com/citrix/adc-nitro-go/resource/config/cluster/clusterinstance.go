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
* Configuration for cluster instance resource.
*/
type Clusterinstance struct {
	/**
	* Unique number that identifies the cluster.
	*/
	Clid int `json:"clid,omitempty"`
	/**
	* Amount of time, in seconds, after which nodes that do not respond to the heartbeats are assumed to be down.If the value is less than 3 sec, set the helloInterval parameter to 200 msec
	*/
	Deadinterval int `json:"deadinterval,omitempty"`
	/**
	* Interval, in milliseconds, at which heartbeats are sent to each cluster node to check the health status.Set the value to 200 msec, if the deadInterval parameter is less than 3 sec
	*/
	Hellointerval int `json:"hellointerval,omitempty"`
	/**
	* Preempt a cluster node that is configured as a SPARE if an ACTIVE node becomes available.
	*/
	Preemption string `json:"preemption,omitempty"`
	/**
	* Quorum Configuration Choices  - "Majority" (recommended) requires majority of nodes to be online for the cluster to be UP. "None" relaxes this requirement.
	*/
	Quorumtype string `json:"quorumtype,omitempty"`
	/**
	* This option is required if the cluster nodes reside on different networks.
	*/
	Inc string `json:"inc,omitempty"`
	/**
	* By turning on this option packets destined to a service in a cluster will not under go any steering.
	*/
	Processlocal string `json:"processlocal,omitempty"`
	/**
	* This option enables you to retain existing connections on a node joining a Cluster system or when a node is being configured for passive timeout. By default, this option is disabled.
	*/
	Retainconnectionsoncluster string `json:"retainconnectionsoncluster,omitempty"`
	/**
	* View based on heartbeat only on bkplane interface
	*/
	Backplanebasedview string `json:"backplanebasedview,omitempty"`
	/**
	* strict mode for sync status of cluster. Depending on the the mode if there are any errors while applying config, sync status is displayed accordingly. By default the flag is disabled.
	*/
	Syncstatusstrictmode string `json:"syncstatusstrictmode,omitempty"`
	/**
	* flag to add ext l2 header during steering. By default the flag is disabled.
	*/
	Dfdretainl2params string `json:"dfdretainl2params,omitempty"`
	/**
	* This field controls the proxy arp feature in cluster. By default the flag is enabled.
	*/
	Clusterproxyarp string `json:"clusterproxyarp,omitempty"`
	/**
	* By turning on this option cluster heartbeats will have security enabled.
	*/
	Secureheartbeats string `json:"secureheartbeats,omitempty"`
	/**
	* The node group in a Cluster system used for transition from L2 to L3.
	*/
	Nodegroup string `json:"nodegroup,omitempty"`

	//------- Read only Parameter ---------;

	Adminstate string `json:"adminstate,omitempty"`
	Propstate string `json:"propstate,omitempty"`
	Validmtu string `json:"validmtu,omitempty"`
	Heterogeneousflag string `json:"heterogeneousflag,omitempty"`
	Operationalstate string `json:"operationalstate,omitempty"`
	Status string `json:"status,omitempty"`
	Rsskeymismatch string `json:"rsskeymismatch,omitempty"`
	Penummismatch string `json:"penummismatch,omitempty"`
	Nodegroupstatewarning string `json:"nodegroupstatewarning,omitempty"`
	Licensemismatch string `json:"licensemismatch,omitempty"`
	Jumbonotsupported string `json:"jumbonotsupported,omitempty"`
	Clustertunnelmodemismatch string `json:"clustertunnelmodemismatch,omitempty"`
	Clusternoheartbeatonnode string `json:"clusternoheartbeatonnode,omitempty"`
	Clusternolinksetmbf string `json:"clusternolinksetmbf,omitempty"`
	Clusternospottedip string `json:"clusternospottedip,omitempty"`
	Clusterclipfailure string `json:"clusterclipfailure,omitempty"`
	Clusterhbhmacerrordetected string `json:"clusterhbhmacerrordetected,omitempty"`
	Nodepenummismatch string `json:"nodepenummismatch,omitempty"`
	Operationalpropstate string `json:"operationalpropstate,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
