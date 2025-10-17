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
* Configuration for SSLVPN intranet application resource.
*/
type Vpnintranetapplication struct {
	/**
	* Name of the intranet application.
	*/
	Intranetapplication string `json:"intranetapplication,omitempty"`
	/**
	* Protocol used by the intranet application. If protocol is set to BOTH, TCP and UDP traffic is allowed.
	*/
	Protocol string `json:"protocol,omitempty"`
	/**
	* Destination IP address, IP range, or host name of the intranet application. This address is the server IP address.
	*/
	Destip string `json:"destip,omitempty"`
	/**
	* Destination subnet mask for the intranet application.
	*/
	Netmask string `json:"netmask,omitempty"`
	/**
	* If you have multiple servers in your network, such as web, email, and file shares, configure an intranet application that includes the IP range for all the network applications. This allows users to access all the intranet applications contained in the IP address range.
	*/
	Iprange string `json:"iprange,omitempty"`
	/**
	* Name of the host for which to configure interception. The names are resolved during interception when users log on with the Citrix Gateway Plug-in.
	*/
	Hostname string `json:"hostname,omitempty"`
	/**
	* Names of the client applications, such as PuTTY and Xshell.
	*/
	Clientapplication []string `json:"clientapplication,omitempty"`
	/**
	* IP address that the intranet application will use to route the connection through the virtual adapter.
	*/
	Spoofiip string `json:"spoofiip,omitempty"`
	/**
	* Destination TCP or UDP port number for the intranet application. Use a hyphen to specify a range of port numbers, for example 90-95.
	*/
	Destport string `json:"destport,omitempty"`
	/**
	* Interception mode for the intranet application or resource. Correct value depends on the type of client software used to make connections. If the interception mode is set to TRANSPARENT, users connect with the Citrix Gateway Plug-in for Windows. With the PROXY setting, users connect with the Citrix Gateway Plug-in for Java.
	*/
	Interception string `json:"interception,omitempty"`
	/**
	* Source IP address. Required if interception mode is set to PROXY. Default is the loopback address, 127.0.0.1.
	*/
	Srcip string `json:"srcip,omitempty"`
	/**
	* Source port for the application for which the Citrix Gateway virtual server proxies the traffic. If users are connecting from a device that uses the Citrix Gateway Plug-in for Java, applications must be configured manually by using the source IP address and TCP port values specified in the intranet application profile. If a port value is not set, the destination port value is used.
	*/
	Srcport *int `json:"srcport,omitempty"`

	//------- Read only Parameter ---------;

	Ipaddress string `json:"ipaddress,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
