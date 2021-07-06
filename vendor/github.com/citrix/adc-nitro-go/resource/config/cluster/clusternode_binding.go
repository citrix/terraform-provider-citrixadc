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
* Binding object which returns the resources bound to clusternode_binding. 
*/
type Clusternodebinding struct {
	/**
	* ID of the cluster node for which to display information. If an ID is not provided, information about all nodes is shown.<br/>Default value: 255<br/>Minimum value =  0<br/>Maximum value =  31
	*/
	Nodeid int `json:"nodeid,omitempty"`


}