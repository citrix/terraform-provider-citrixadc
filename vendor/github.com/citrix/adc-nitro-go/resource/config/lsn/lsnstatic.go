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

package lsn

/**
* Configuration for static mapping resource.
*/
type Lsnstatic struct {
	/**
	* Name for the LSN static mapping entry. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the LSN group is created. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "lsn static1" or 'lsn static1').
	*/
	Name string `json:"name,omitempty"`
	/**
	* Protocol for the LSN mapping entry.
	*/
	Transportprotocol string `json:"transportprotocol,omitempty"`
	/**
	* IPv4(NAT44 & DS-Lite)/IPv6(NAT64) address of an LSN subscriber for the LSN static mapping entry.
	*/
	Subscrip string `json:"subscrip,omitempty"`
	/**
	* Port of the LSN subscriber for the LSN mapping entry. * represents all ports being used. Used in case of static wildcard
	*/
	Subscrport int `json:"subscrport,omitempty"`
	/**
	* B4 address in DS-Lite setup
	*/
	Network6 string `json:"network6,omitempty"`
	/**
	* ID of the traffic domain to which the subscriber belongs. 
		If you do not specify an ID, the subscriber is assumed to be a part of the default traffic domain.
	*/
	Td int `json:"td"`
	/**
	* IPv4 address, already existing on the Citrix ADC as type LSN, to be used as NAT IP address for this mapping entry.
	*/
	Natip string `json:"natip,omitempty"`
	/**
	* NAT port for this LSN mapping entry. * represents all ports being used. Used in case of static wildcard
	*/
	Natport int `json:"natport,omitempty"`
	/**
	* Destination IP address for the LSN mapping entry.
	*/
	Destip string `json:"destip,omitempty"`
	/**
	* ID of the traffic domain through which the destination IP address for this LSN mapping entry is reachable from the Citrix ADC.
		If you do not specify an ID, the destination IP address is assumed to be reachable through the default traffic domain, which has an ID of 0.
	*/
	Dsttd int `json:"dsttd,omitempty"`
	/**
	* Type of sessions to be displayed.
	*/
	Nattype string `json:"nattype,omitempty"`

	//------- Read only Parameter ---------;

	Status string `json:"status,omitempty"`

}
