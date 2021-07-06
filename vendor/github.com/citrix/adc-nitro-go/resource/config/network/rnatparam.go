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
* Configuration for RNAT parameter resource.
*/
type Rnatparam struct {
	/**
	* Enable TCP proxy, which enables the Citrix ADC to optimize the RNAT TCP traffic by using Layer 4 features.
	*/
	Tcpproxy string `json:"tcpproxy,omitempty"`
	/**
	* Enable source ip persistency, which enables the Citrix ADC to use the RNAT ips using source ip.
	*/
	Srcippersistency string `json:"srcippersistency,omitempty"`

}
