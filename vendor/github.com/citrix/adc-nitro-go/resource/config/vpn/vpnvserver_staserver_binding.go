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
* Binding class showing the staserver that can be bound to vpnvserver.
*/
type Vpnvserverstaserverbinding struct {
	/**
	* Configured Secure Ticketing Authority (STA) server.
	*/
	Staserver string `json:"staserver,omitempty"`
	/**
	* Authority ID of the STA Server. Authority ID is used to match incoming STA tickets in the SOCKS/CGP protocol with the right STA server.
	*/
	Staauthid string `json:"staauthid,omitempty"`
	/**
	* State of the STA Server. If Authority ID is set then STA Server is UP else DOWN.
	*/
	Stastate string `json:"stastate,omitempty"`
	Acttype *int `json:"acttype,omitempty"`
	/**
	* Type of the STA server address(ipv4/v6).
	*/
	Staaddresstype string `json:"staaddresstype,omitempty"`
	/**
	* Name of the virtual server.
	*/
	Name string `json:"name,omitempty"`


}