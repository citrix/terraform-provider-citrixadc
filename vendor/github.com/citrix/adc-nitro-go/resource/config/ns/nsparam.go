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
	Httpport []int `json:"httpport,omitempty"`
	/**
	* Maximum number of connections that will be made from the appliance to the web server(s) attached to it. The value entered here is applied globally to all attached servers.
	*/
	Maxconn int `json:"maxconn"` // Zero is a valid value
	/**
	* Maximum number of requests that the system can pass on a particular connection between the appliance and a server attached to it. Setting this value to 0 allows an unlimited number of requests to be passed. This value is overridden by the maximum number of requests configured on the individual service.
	*/
	Maxreq int `json:"maxreq"` // Zero is a valid value
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
	Cookieversion string `json:"cookieversion"` // Zero is a valid value
	/**
	* Enable or disable secure flag for persistence cookie.
	*/
	Securecookie string `json:"securecookie,omitempty"`
	/**
	* Minimum path MTU value that Citrix ADC will process in the ICMP fragmentation needed message. If the ICMP message contains a value less than this value, then this value is used instead.
	*/
	Pmtumin int `json:"pmtumin,omitempty"`
	/**
	* Interval, in minutes, for flushing the PMTU entries.
	*/
	Pmtutimeout int `json:"pmtutimeout,omitempty"`
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
	* Percentage of shared pool value granted to PE once PE exhausts the local exclusive quota. Where shared pool is the remaining maxclient quota after distribution of exclusive quota to PEs.
		Example: In a 2 PE NetScaler system if configured maxclient is 100 and exclusive quota is 80 percent then each PE will get 40 as local exclusive quota and 20 will be in shared pool. If configured grantQuota is 20 percent, then after exhausting its local exclusive quota PE borrows from shared pool in chunks of 4 i.e. 20 percent of 20.
	*/
	Grantquotamaxclient int `json:"grantquotamaxclient"` // Zero is a valid value
	/**
	* Percentage of maxClient threshold to be divided equally among PEs.
	*/
	Exclusivequotamaxclient int `json:"exclusivequotamaxclient"` // Zero is a valid value
	/**
	* Percentage of shared pool value granted to PE once PE exhausts the local exclusive quota. Where shared pool is the remaining spillover quota after distribution of exclusive quota to PEs.
		Example: In a 2 PE NetScaler system if configured spillover is 100 and exclusive quota is 80 percent then each PE will get 40 as local exclusive quota and 20 will be in shared pool. If configured grantQuota is 20 percent, then after exhausting its local exclusive quota PE borrows from shared pool in chunks of 4 i.e. 20 percent of 20.
	*/
	Grantquotaspillover int `json:"grantquotaspillover,omitempty"`
	/**
	* Percentage of spillover threshold to be divided equally among PEs.
	*/
	Exclusivequotaspillover int `json:"exclusivequotaspillover"` // Zero is a valid value
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
	* The ICA ports on the Web server. This allows the system to perform connection off-load for any client request that has a destination port matching one of these configured ports.
	*/
	Icaports []int `json:"icaports,omitempty"`
	/**
	* Enable or disable the insertion of the client TCP/IP header in TCP payload passed from the client to one, some, or all servers attached to the system. The passed address can then be accessed through a minor modification to the server.
	*/
	Tcpcip string `json:"tcpcip,omitempty"`
	/**
	* VLAN on which the subscriber traffic arrives on the appliance.
	*/
	Servicepathingressvlan int `json:"servicepathingressvlan,omitempty"`
	/**
	* The Secure ICA ports on the Web server. This allows the system to perform connection off-load for any
		client request that has a destination port matching one of these configured ports.
	*/
	Secureicaports []int `json:"secureicaports,omitempty"`
	/**
	* This allow the configuration of management HTTP port.
	*/
	Mgmthttpport int `json:"mgmthttpport,omitempty"`
	/**
	* This allows the configuration of management HTTPS port.
	*/
	Mgmthttpsport int `json:"mgmthttpsport,omitempty"`
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
	Ipttl int `json:"ipttl,omitempty"`

	//------- Read only Parameter ---------;

	Autoscaleoption    string `json:"autoscaleoption,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}
