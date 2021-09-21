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
* Configuration for Next Hop Server resource.
*/
type Vpnnexthopserver struct {
	/**
	* Name for the Citrix Gateway appliance in the first DMZ.
	*/
	Name string `json:"name,omitempty"`
	/**
	* IP address of the Citrix Gateway proxy in the second DMZ.
	*/
	Nexthopip string `json:"nexthopip,omitempty"`
	/**
	* FQDN of the Citrix Gateway proxy in the second DMZ.
	*/
	Nexthopfqdn string `json:"nexthopfqdn,omitempty"`
	/**
	* Address Type (IPV4/IPv6) of DNS name of nextHopServer FQDN.
	*/
	Resaddresstype string `json:"resaddresstype,omitempty"`
	/**
	* Port number of the Citrix Gateway proxy in the second DMZ.
	*/
	Nexthopport int `json:"nexthopport,omitempty"`
	/**
	* Use of a secure port, such as 443, for the double-hop configuration.
	*/
	Secure string `json:"secure,omitempty"`

}
