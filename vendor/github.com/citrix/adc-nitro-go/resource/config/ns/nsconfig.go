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
* Configuration for system config resource.
*/
type Nsconfig struct {
	/**
	* Configurations will be cleared without prompting for confirmation.
	*/
	Force bool `json:"force,omitempty"`
	/**
	* Types of configurations to be cleared.
		* basic: Clears all configurations except the following:
		- NSIP, default route (gateway), static routes, MIPs, and SNIPs
		- Network settings (DG, VLAN, RHI and DNS settings)
		- Cluster settings
		- HA node definitions
		- Feature and mode settings
		- nsroot password
		* extended: Clears the same configurations as the 'basic' option. In addition, it clears the feature and mode settings.
		* full: Clears all configurations except NSIP, default route, and interface settings.
		Note: When you clear the configurations through the cluster IP address, by specifying the level as 'full', the cluster is deleted and all cluster nodes become standalone appliances. The 'basic' and 'extended' levels are propagated to the cluster nodes.
	*/
	Level string `json:"level,omitempty"`
	/**
	* RBA configurations and TACACS policies bound to system global will not be cleared if RBA is set to NO.This option is applicable only for BASIC level of clear configuration.Default is YES, which will clear rba configurations.
	*/
	Rbaconfig string `json:"rbaconfig,omitempty"`
	/**
	* IP address of the Citrix ADC. Commonly referred to as NSIP address. This parameter is mandatory to bring up the appliance.
	*/
	Ipaddress string `json:"ipaddress,omitempty"`
	/**
	* Netmask corresponding to the IP address. This parameter is mandatory to bring up the appliance.
	*/
	Netmask string `json:"netmask,omitempty"`
	/**
	* VLAN (NSVLAN) for the subnet on which the IP address resides.
	*/
	Nsvlan *int `json:"nsvlan,omitempty"`
	/**
	* Interfaces of the appliances that must be bound to the NSVLAN.
	*/
	Ifnum []string `json:"ifnum,omitempty"`
	/**
	* Specifies that the interfaces will be added as 802.1q tagged interfaces. Packets sent on these interface on this VLAN will have an additional 4-byte 802.1q tag which identifies the VLAN.
		To use 802.1q tagging, the switch connected to the appliance's interfaces must also be configured for tagging.
	*/
	Tagged string `json:"tagged,omitempty"`
	/**
	* The HTTP ports on the Web server. This allows the system to perform connection off-load for any client request that has a destination port matching one of these configured ports.
	*/
	Httpport []int `json:"httpport,omitempty"`
	/**
	* The maximum number of connections that will be made from the system to the web server(s) attached to it. The value entered here is applied globally to all attached servers.
	*/
	Maxconn *int `json:"maxconn,omitempty"`
	/**
	* The maximum number of requests that the system can pass on a particular connection between the system and a server attached to it. Setting this value to 0 allows an unlimited number of requests to be passed.
	*/
	Maxreq *int `json:"maxreq,omitempty"`
	/**
	* The option to control (enable or disable) the insertion of the actual client IP address into the HTTP header request passed from the client to one, some, or all servers attached to the system.
		The passed address can then be accessed through a minor modification to the server.
		l    If cipHeader is specified, it will be used as the client IP header.
		l    If it is not specified, then the value that has been set by the set ns config CLI command will be used as the client IP header.
	*/
	Cip string `json:"cip,omitempty"`
	/**
	* The text that will be used as the client IP header.
	*/
	Cipheader string `json:"cipheader,omitempty"`
	/**
	* The version of the cookie inserted by system.
	*/
	Cookieversion string `json:"cookieversion,omitempty"`
	/**
	* enable/disable secure flag for persistence cookie
	*/
	Securecookie string `json:"securecookie,omitempty"`
	/**
	* The minimum Path MTU.
	*/
	Pmtumin *int `json:"pmtumin,omitempty"`
	/**
	* The timeout value in minutes.
	*/
	Pmtutimeout *int `json:"pmtutimeout,omitempty"`
	/**
	* Port range configured for FTP services.
	*/
	Ftpportrange string `json:"ftpportrange,omitempty"`
	/**
	* Port range for cache redirection services.
	*/
	Crportrange string `json:"crportrange,omitempty"`
	/**
	* Name of the timezone
	*/
	Timezone string `json:"timezone,omitempty"`
	/**
	* The percentage of shared quota to be granted at a time for maxClient
	*/
	Grantquotamaxclient *int `json:"grantquotamaxclient,omitempty"`
	/**
	* The percentage of maxClient to be given to PEs
	*/
	Exclusivequotamaxclient *int `json:"exclusivequotamaxclient,omitempty"`
	/**
	* The percentage of shared quota to be granted at a time for spillover
	*/
	Grantquotaspillover *int `json:"grantquotaspillover,omitempty"`
	/**
	* The percentage of spillover threshold to be given to PEs
	*/
	Exclusivequotaspillover *int `json:"exclusivequotaspillover,omitempty"`
	/**
	* This enabled secure management traffic handling.
	*/
	Securemanagementtraffic string `json:"securemanagementtraffic,omitempty"`
	/**
	* This positive integer identifies Management traffic domain. If not specified, defaults to 4094
	*/
	Securemanagementtd *int `json:"securemanagementtd,omitempty"`
	/**
	* Use this option to do saveconfig for all partitions
	*/
	All bool `json:"all,omitempty"`
	/**
	* Location of the configurations.
	*/
	Config1 string `json:"config1,omitempty"`
	/**
	* Location of the configurations.
	*/
	Config2 string `json:"config2,omitempty"`
	/**
	* Format to display the difference in configurations.
	*/
	Outtype string `json:"outtype,omitempty"`
	/**
	* File that contains the commands to be compared.
	*/
	Template bool `json:"template,omitempty"`
	/**
	* Suppress device specific differences.
	*/
	Ignoredevicespecific bool `json:"ignoredevicespecific,omitempty"`
	/**
	* Option to list all weak passwords (not adhering to strong password requirements). Takes config file as input, if no input specified, running configuration is considered. Command => query ns config -weakpassword  / query ns config -weakpassword /nsconfig/ns.conf
	*/
	Weakpassword bool `json:"weakpassword,omitempty"`
	/**
	* Option to list all passwords changed which would not work when downgraded to older releases. Takes config file as input, if no input specified, running configuration is considered. Command => query ns config -changedpassword / query ns config -changedpassword /nsconfig/ns.conf
	*/
	Changedpassword bool `json:"changedpassword,omitempty"`
	/**
	* configuration File to be used to find weak passwords, if not specified, running config is taken as input.
	*/
	Config string `json:"config,omitempty"`
	/**
	* Full path of config file to be converted to nitro
	*/
	Configfile string `json:"configfile,omitempty"`
	/**
	* Full path of file to store the nitro graph. If not specified, nitro graph is returned as part of the API response.
	*/
	Responsefile string `json:"responsefile,omitempty"`
	/**
	* Using this option will run the operation in async mode and return the job id. The job ID can be used later to track the conversion progress via show ns job <id> Command. This option is mostly useful for API to avoid timeouts for large input configuration
	*/
	Async bool `json:"Async,omitempty"`

	//------- Read only Parameter ---------;

	Message string `json:"message,omitempty"`
	Mappedip string `json:"mappedip,omitempty"`
	Range string `json:"range,omitempty"`
	Svmcmd string `json:"svmcmd,omitempty"`
	Systemtype string `json:"systemtype,omitempty"`
	Primaryip string `json:"primaryip,omitempty"`
	Primaryip6 string `json:"primaryip6,omitempty"`
	Flags string `json:"flags,omitempty"`
	Lastconfigchangedtime string `json:"lastconfigchangedtime,omitempty"`
	Lastconfigsavetime string `json:"lastconfigsavetime,omitempty"`
	Currentsytemtime string `json:"currentsytemtime,omitempty"`
	Systemtime string `json:"systemtime,omitempty"`
	Configchanged string `json:"configchanged,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
	Response string `json:"response,omitempty"`
	Id string `json:"id,omitempty"`

}
