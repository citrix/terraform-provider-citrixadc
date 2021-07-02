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
* Configuration for Citrix ADC parameters resource.
*/
type Nsparam struct {
	/**
	* HTTP ports on the web server. This allows the system to perform connection off-load for any client request that has a destination port matching one of these configured ports.
	*/
	Httpport []int32 `json:"httpport,omitempty"`
	/**
	* Maximum number of connections that will be made from the appliance to the web server(s) attached to it. The value entered here is applied globally to all attached servers.
	*/
	Maxconn uint32 `json:"maxconn,omitempty"`
	/**
	* Maximum number of requests that the system can pass on a particular connection between the appliance and a server attached to it. Setting this value to 0 allows an unlimited number of requests to be passed. This value is overridden by the maximum number of requests configured on the individual service.
	*/
	Maxreq uint32 `json:"maxreq,omitempty"`
	/**
	* Enable or disable the insertion of the actual client IP address into the HTTP header request passed from the client to one, some, or all servers attached to the system. The passed address can then be accessed through a minor modification to the server.
		* If the CIP header is specified, it will be used as the client IP header.
		* If the CIP header is not specified, the value that has been set will be used as the client IP header.
	*/
	Cip string `json:"cip,omitempty"`
	/**
	* Text that will be used as the client IP address header.
	*/
	Cipheader string `json:"cipheader,omitempty"`
	/**
	* Version of the cookie inserted by the system.
	*/
	Cookieversion string `json:"cookieversion,omitempty"`
	/**
	* Enable or disable secure flag for persistence cookie.
	*/
	Securecookie string `json:"securecookie,omitempty"`
	/**
	* Minimum path MTU value that Citrix ADC will process in the ICMP fragmentation needed message. If the ICMP message contains a value less than this value, then this value is used instead.
	*/
	Pmtumin uint32 `json:"pmtumin,omitempty"`
	/**
	* Interval, in minutes, for flushing the PMTU entries.
	*/
	Pmtutimeout uint32 `json:"pmtutimeout,omitempty"`
	/**
	* Minimum and maximum port (port range) that FTP services are allowed to use.
	*/
	Ftpportrange string `json:"ftpportrange,omitempty"`
	/**
	* Port range for cache redirection services.
	*/
	Crportrange string `json:"crportrange,omitempty"`
	/**
	* Time zone for the Citrix ADC. Name of the time zone should be specified as argument.
	*/
	Timezone string `json:"timezone,omitempty"`
	/**
	* Percentage of shared quota to be granted at a time for maxClient.
	*/
	Grantquotamaxclient uint32 `json:"grantquotamaxclient,omitempty"`
	/**
	* Percentage of maxClient to be given to PEs.
	*/
	Exclusivequotamaxclient uint32 `json:"exclusivequotamaxclient,omitempty"`
	/**
	* Percentage of shared quota to be granted at a time for spillover.
	*/
	Grantquotaspillover uint32 `json:"grantquotaspillover,omitempty"`
	/**
	* Percentage of maximum limit to be given to PEs.
	*/
	Exclusivequotaspillover uint32 `json:"exclusivequotaspillover,omitempty"`
	/**
	* Enable/Disable use_proxy_port setting
	*/
	Useproxyport string `json:"useproxyport,omitempty"`
	/**
	* Enables/disables the internal user from logging in to the appliance. Before disabling internal user login, you must have key-based authentication set up on the appliance. The file name for the key pair must be "ns_comm_key".
	*/
	Internaluserlogin string `json:"internaluserlogin,omitempty"`
	/**
	* Allow the FTP server to come from a random source port for active FTP data connections
	*/
	Aftpallowrandomsourceport string `json:"aftpallowrandomsourceport,omitempty"`
	/**
	* The ICA ports on the Web server. This allows the system to perform connection off-load for any
		client request that has a destination port matching one of these configured ports.
	*/
	Icaports []int32 `json:"icaports,omitempty"`
	/**
	* Enable or disable the insertion of the client TCP/IP header in TCP payload passed from the client to one, some, or all servers attached to the system. The passed address can then be accessed through a minor modification to the server.
	*/
	Tcpcip string `json:"tcpcip,omitempty"`
	/**
	* VLAN on which the subscriber traffic arrives on the appliance.
	*/
	Servicepathingressvlan uint32 `json:"servicepathingressvlan,omitempty"`
	/**
	* The Secure ICA ports on the Web server. This allows the system to perform connection off-load for any
		client request that has a destination port matching one of these configured ports.
	*/
	Secureicaports []int32 `json:"secureicaports,omitempty"`
	/**
	* This allow the configuration of management HTTP port.
	*/
	Mgmthttpport int32 `json:"mgmthttpport,omitempty"`
	/**
	* This allows the configuration of management HTTPS port.
	*/
	Mgmthttpsport int32 `json:"mgmthttpsport,omitempty"`
	/**
	* Disable/Enable v1 or v2 proxy protocol header for client info insertion
	*/
	Proxyprotocol string `json:"proxyprotocol,omitempty"`
	/**
	* Disable/Enable advanace analytics stats
	*/
	Advancedanalyticsstats string `json:"advancedanalyticsstats,omitempty"`
	/**
	* Set the IP Time to Live (TTL) and Hop Limit value for all outgoing packets from Citrix ADC.
	*/
	Ipttl uint32 `json:"ipttl,omitempty"`

	//------- Read only Parameter ---------;

	Autoscaleoption string `json:"autoscaleoption,omitempty"`

}
