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

package lb

/**
* Binding class showing the servicegroupmember that can be bound to lbvserver.
*/
type Lbvserverservicegroupmemberbinding struct {
	/**
	* The service group name bound to the selected load balancing virtual server.
	*/
	Servicegroupname string `json:"servicegroupname,omitempty"`
	/**
	* IPv4 or IPv6 address to assign to the virtual server.
	*/
	Ipv46 string `json:"ipv46,omitempty"`
	/**
	* Port number for the virtual server.
	*/
	Port int `json:"port,omitempty"`
	/**
	* Protocol used by the service (also called the service type).
	*/
	Servicetype string `json:"servicetype,omitempty"`
	/**
	* Current LB vserver state.
	*/
	Curstate string `json:"curstate,omitempty"`
	/**
	* Weight to assign to the specified service.
	*/
	Weight int `json:"weight,omitempty"`
	/**
	* Dynamic weight
	*/
	Dynamicweight int `json:"dynamicweight,omitempty"`
	/**
	* Encryped Ip address and port of the service that is inserted into the set-cookie http header
	*/
	Cookieipport string `json:"cookieipport,omitempty"`
	/**
	* Use this parameter to specify the cookie name for COOKIE peristence type. It specifies the name of cookie with a maximum of 32 characters. If not specified, cookie name is internally generated.
	*/
	Cookiename string `json:"cookiename,omitempty"`
	/**
	* Vserver Id
	*/
	Vserverid string `json:"vserverid,omitempty"`
	/**
	* Used for displaying the location of bound services.
	*/
	Preferredlocation string `json:"preferredlocation,omitempty"`
	/**
	* Order number to be assigned to the service when it is bound to the lb vserver.
	*/
	Order int `json:"order,omitempty"`
	/**
	* Order in string form assigned to the service when it is bound to the lb vserver.
	*/
	Orderstr string `json:"orderstr,omitempty"`
	/**
	* Name for the virtual server. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Can be changed after the virtual server is created.
		CLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my vserver" or 'my vserver'). 
	*/
	Name string `json:"name,omitempty"`


}