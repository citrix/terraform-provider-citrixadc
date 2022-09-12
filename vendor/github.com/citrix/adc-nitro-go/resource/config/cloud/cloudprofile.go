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

package cloud

/**
* Configuration for cloud profile resource.
*/
type Cloudprofile struct {
	/**
	* Name for the Cloud profile. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the profile is created.
	*/
	Name string `json:"name,omitempty"`
	/**
	* Type of cloud profile that you want to create, Vserver or based on Azure Tags
	*/
	Type string `json:"type,omitempty"`
	/**
	* Name for the virtual server. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Can be changed after the virtual server is created.
		CLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my vserver" or 'my vserver'). 
	*/
	Vservername string `json:"vservername,omitempty"`
	/**
	* Protocol used by the service (also called the service type).
	*/
	Servicetype string `json:"servicetype,omitempty"`
	/**
	* IPv4 or IPv6 address to assign to the virtual server.
	*/
	Ipaddress string `json:"ipaddress,omitempty"`
	/**
	* Port number for the virtual server.
	*/
	Port int `json:"port,omitempty"`
	/**
	* servicegroups bind to this server
	*/
	Servicegroupname string `json:"servicegroupname,omitempty"`
	/**
	* The type of bound service
	*/
	Boundservicegroupsvctype string `json:"boundservicegroupsvctype,omitempty"`
	/**
	* The port number to be used for the bound service.
	*/
	Vsvrbindsvcport int `json:"vsvrbindsvcport,omitempty"`
	/**
	* Indicates graceful shutdown of the service. System will wait for all outstanding connections to this service to be closed before disabling the service.
	*/
	Graceful string `json:"graceful,omitempty"`
	/**
	* Time, in seconds, after which all the services configured on the server are disabled.
	*/
	Delay int `json:"delay,omitempty"`
	/**
	* Azure tag name
	*/
	Azuretagname string `json:"azuretagname,omitempty"`
	/**
	* Azure tag value
	*/
	Azuretagvalue string `json:"azuretagvalue,omitempty"`
	/**
	* Azure polling period (in seconds)
	*/
	Azurepollperiod int `json:"azurepollperiod,omitempty"`

}
