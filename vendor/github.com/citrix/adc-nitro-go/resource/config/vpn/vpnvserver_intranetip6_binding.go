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
* Binding class showing the intranetip6 that can be bound to vpnvserver.
*/
type Vpnvserverintranetip6binding struct {
	/**
	* The network id for the range of intranet IP6 addresses or individual intranet ip to be bound to the vserver.
	*/
	Intranetip6 string `json:"intranetip6,omitempty"`
	/**
	* The number of ipv6 addresses
	*/
	Numaddr *int `json:"numaddr,omitempty"`
	Acttype *int `json:"acttype,omitempty"`
	/**
	* Name of the virtual server.
	*/
	Name string `json:"name,omitempty"`


}