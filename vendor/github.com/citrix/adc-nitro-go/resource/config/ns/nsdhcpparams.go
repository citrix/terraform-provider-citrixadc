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

package ns

/**
* Configuration for DHCP parameters resource.
*/
type Nsdhcpparams struct {
	/**
	* Enables DHCP client to acquire IP address from the DHCP server in the next boot. When set to OFF, disables the DHCP client in the next boot.
	*/
	Dhcpclient string `json:"dhcpclient,omitempty"`
	/**
	* DHCP acquired routes are saved on the Citrix ADC.
	*/
	Saveroute string `json:"saveroute,omitempty"`

	//------- Read only Parameter ---------;

	Ipaddress string `json:"ipaddress,omitempty"`
	Netmask string `json:"netmask,omitempty"`
	Hostrtgw string `json:"hostrtgw,omitempty"`
	Running string `json:"running,omitempty"`

}
