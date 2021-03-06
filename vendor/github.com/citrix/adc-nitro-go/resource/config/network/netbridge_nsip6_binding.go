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
* Binding class showing the nsip6 that can be bound to netbridge.
*/
type Netbridgensip6binding struct {
	/**
	* The subnet that is extended by this network bridge.
	*/
	Ipaddress string `json:"ipaddress,omitempty"`
	/**
	* The network mask for the subnet.
	*/
	Netmask string `json:"netmask,omitempty"`
	/**
	* The name of the network bridge.
	*/
	Name string `json:"name,omitempty"`


}