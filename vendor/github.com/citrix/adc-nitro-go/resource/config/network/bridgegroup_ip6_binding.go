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
* Binding class showing the ip6 that can be bound to bridgegroup.
*/
type Bridgegroupip6binding struct {
	/**
	* The IP address assigned to the  bridge group.
	*/
	Ipaddress string `json:"ipaddress,omitempty"`
	/**
	* Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.
	*/
	Td uint32 `json:"td,omitempty"`
	/**
	* Temporary flag used for internal purpose.
	*/
	Rnat bool `json:"rnat,omitempty"`
	/**
	* The owner node group in a Cluster for this vlan.
	*/
	Ownergroup string `json:"ownergroup,omitempty"`
	/**
	* The integer that uniquely identifies the bridge group.
	*/
	Id uint32 `json:"id,omitempty"`
	/**
	* A subnet mask associated with the network address.
	*/
	Netmask string `json:"netmask,omitempty"`


}