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

package snmp

/**
* Configuration for manager resource.
*/
type Snmpmanager struct {
	/**
	* IP address of the SNMP manager. Can be an IPv4 or IPv6 address. You can instead specify an IPv4 network address or IPv6 network prefix if you want the Citrix ADC to respond to SNMP queries from any device on the specified network. Alternatively, instead of an IPv4 address, you can specify a host name that has been assigned to an SNMP manager. If you do so, you must add a DNS name server that resolves the host name of the SNMP manager to its IP address. 
		Note: The Citrix ADC does not support host names for SNMP managers that have IPv6 addresses.
	*/
	Ipaddress string `json:"ipaddress,omitempty"`
	/**
	* Subnet mask associated with an IPv4 network address. If the IP address specifies the address or host name of a specific host, accept the default value of 255.255.255.255.
	*/
	Netmask string `json:"netmask,omitempty"`
	/**
	* Amount of time, in seconds, for which the Citrix ADC waits before sending another DNS query to resolve the host name of the SNMP manager if the last query failed. This parameter is valid for host-name based SNMP managers only. After a query succeeds, the TTL determines the wait time. The minimum and default value is 5.
	*/
	Domainresolveretry int `json:"domainresolveretry,omitempty"`

	//------- Read only Parameter ---------;

	Ip string `json:"ip,omitempty"`
	Domain string `json:"domain,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
