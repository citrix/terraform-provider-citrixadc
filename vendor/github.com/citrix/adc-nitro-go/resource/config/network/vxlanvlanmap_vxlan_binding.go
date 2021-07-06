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
* Binding class showing the vxlan that can be bound to vxlanvlanmap.
*/
type Vxlanvlanmapvxlanbinding struct {
	/**
	* The VXLAN assigned to the vlan inside the cloud.
	*/
	Vxlan int `json:"vxlan,omitempty"`
	/**
	* The vlan id or the range of vlan ids in the on-premise network.
	*/
	Vlan []string `json:"vlan,omitempty"`
	/**
	* Name of the mapping table.
	*/
	Name string `json:"name,omitempty"`


}