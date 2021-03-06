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
* Binding class showing the vlan that can be bound to bridgegroup.
*/
type Bridgegroupvlanbinding struct {
	/**
	* Names of all member VLANs.
	*/
	Vlan int `json:"vlan,omitempty"`
	/**
	* Temporary flag used for internal purpose.
	*/
	Rnat bool `json:"rnat,omitempty"`
	/**
	* The integer that uniquely identifies the bridge group.
	*/
	Id int `json:"id,omitempty"`


}