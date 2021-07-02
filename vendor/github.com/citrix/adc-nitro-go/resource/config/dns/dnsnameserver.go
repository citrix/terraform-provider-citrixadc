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

package dns

/**
* Configuration for name server resource.
*/
type Dnsnameserver struct {
	/**
	* IP address of an external name server or, if the Local parameter is set, IP address of a local DNS server (LDNS).
	*/
	Ip string `json:"ip,omitempty"`
	/**
	* Name of a DNS virtual server. Overrides any IP address-based name servers configured on the Citrix ADC.
	*/
	Dnsvservername string `json:"dnsvservername,omitempty"`
	/**
	* Mark the IP address as one that belongs to a local recursive DNS server on the Citrix ADC. The appliance recursively resolves queries received on an IP address that is marked as being local. For recursive resolution to work, the global DNS parameter, Recursion, must also be set. 
		If no name server is marked as being local, the appliance functions as a stub resolver and load balances the name servers.
	*/
	Local bool `json:"local,omitempty"`
	/**
	* Administrative state of the name server.
	*/
	State string `json:"state,omitempty"`
	/**
	* Protocol used by the name server. UDP_TCP is not valid if the name server is a DNS virtual server configured on the appliance.
	*/
	Type string `json:"type,omitempty"`
	/**
	* Name of the DNS profile to be associated with the name server
	*/
	Dnsprofilename string `json:"dnsprofilename,omitempty"`

	//------- Read only Parameter ---------;

	Servicename string `json:"servicename,omitempty"`
	Port string `json:"port,omitempty"`
	Nameserverstate string `json:"nameserverstate,omitempty"`
	Clmonowner string `json:"clmonowner,omitempty"`
	Clmonview string `json:"clmonview,omitempty"`

}
