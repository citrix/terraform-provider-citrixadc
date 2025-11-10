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

package vpn

/**
* Binding class showing the intranetip that can be bound to vpnvserver.
*/
type Vpnvserverintranetipbinding struct {
	/**
	* The network ID for the range of intranet IP addresses or individual intranet IP addresses to be bound to the virtual server.
	*/
	Intranetip string `json:"intranetip,omitempty"`
	/**
	* The netmask of the intranet IP address or range.
	*/
	Netmask string `json:"netmask,omitempty"`
	/**
	* Whether or not mapped IP addresses are ON or OFF. Mapped IP addresses are source IP addresses
		for the virtual servers running on the Citrix ADC. Mapped IP addresses are used by the system to connect to the backend servers.
	*/
	Map string `json:"map,omitempty"`
	Acttype *int `json:"acttype,omitempty"`
	/**
	* Name of the virtual server.
	*/
	Name string `json:"name,omitempty"`


}