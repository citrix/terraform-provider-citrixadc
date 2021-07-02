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
* Configuration for inbound nat resource.
*/
type Inat struct {
	/**
	* Name for the Inbound NAT (INAT) entry. Leading character must be a number or letter. Other characters allowed, after the first character, are @ _ - . (period) : (colon) # and space ( ).
	*/
	Name string `json:"name,omitempty"`
	/**
	* Public IP address of packets received on the Citrix ADC. Can be aNetScaler-owned VIP or VIP6 address.
	*/
	Publicip string `json:"publicip,omitempty"`
	/**
	* IP address of the server to which the packet is sent by the Citrix ADC. Can be an IPv4 or IPv6 address.
	*/
	Privateip string `json:"privateip,omitempty"`
	/**
	* Stateless translation.
	*/
	Mode string `json:"mode,omitempty"`
	/**
	* Enable TCP proxy, which enables the Citrix ADC to optimize the RNAT TCP traffic by using Layer 4 features.
	*/
	Tcpproxy string `json:"tcpproxy,omitempty"`
	/**
	* Enable the FTP protocol on the server for transferring files between the client and the server.
	*/
	Ftp string `json:"ftp,omitempty"`
	/**
	* To enable/disable TFTP (Default DISABLED).
	*/
	Tftp string `json:"tftp,omitempty"`
	/**
	* Enable the Citrix ADC to retain the source IP address of packets before sending the packets to the server.
	*/
	Usip string `json:"usip,omitempty"`
	/**
	* Enable the Citrix ADC to use a SNIP address as the source IP address of packets before sending the packets to the server.
	*/
	Usnip string `json:"usnip,omitempty"`
	/**
	* Unique IP address used as the source IP address in packets sent to the server. Must be a MIP or SNIP address.
	*/
	Proxyip string `json:"proxyip,omitempty"`
	/**
	* Enable the Citrix ADC to proxy the source port of packets before sending the packets to the server.
	*/
	Useproxyport string `json:"useproxyport,omitempty"`
	/**
	* Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.
	*/
	Td uint32 `json:"td,omitempty"`
	/**
	* Synchronize connection information with the secondary appliance in a high availability (HA) pair. That is, synchronize all connection-related information for the INAT session
	*/
	Connfailover string `json:"connfailover,omitempty"`

	//------- Read only Parameter ---------;

	Flags string `json:"flags,omitempty"`

}
