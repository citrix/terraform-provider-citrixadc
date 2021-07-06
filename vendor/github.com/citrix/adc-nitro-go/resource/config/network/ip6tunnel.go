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
* Configuration for ip6 Tunnel resource.
*/
type Ip6tunnel struct {
	/**
	* Name for the IPv6 Tunnel. Cannot be changed after the service group is created. Must begin with a number or letter, and can consist of letters, numbers, and the @ _ - . (period) : (colon) # and space ( ) characters.
	*/
	Name string `json:"name,omitempty"`
	/**
	* An IPv6 address of the remote Citrix ADC used to set up the tunnel.
	*/
	Remote string `json:"remote,omitempty"`
	/**
	* An IPv6 address of the local Citrix ADC used to set up the tunnel.
	*/
	Local string `json:"local,omitempty"`
	/**
	* The owner node group in a Cluster for the tunnel.
	*/
	Ownergroup string `json:"ownergroup,omitempty"`

	//------- Read only Parameter ---------;

	Remoteip string `json:"remoteip,omitempty"`
	Type string `json:"type,omitempty"`
	Encapip string `json:"encapip,omitempty"`

}
