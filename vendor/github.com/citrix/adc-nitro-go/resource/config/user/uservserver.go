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

package user

/**
* Configuration for virtual server resource.
*/
type Uservserver struct {
	/**
	* Name for the virtual server. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters.
		CLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my vserver" or 'my vserver'). 
	*/
	Name string `json:"name,omitempty"`
	/**
	* User protocol uesd by the service.
	*/
	Userprotocol string `json:"userprotocol,omitempty"`
	/**
	* IPv4 or IPv6 address to assign to the virtual server.
	*/
	Ipaddress string `json:"ipaddress,omitempty"`
	/**
	* Port number for the virtual server.
	*/
	Port int `json:"port,omitempty"`
	/**
	* Name of the default Load Balancing virtual server used for load balancing of services. The protocol type of default Load Balancing virtual server should be a user type.
	*/
	Defaultlb string `json:"defaultlb,omitempty"`
	/**
	* Any comments associated with the protocol.
	*/
	Params string `json:"Params,omitempty"`
	/**
	* Any comments that you might want to associate with the virtual server.
	*/
	Comment string `json:"comment,omitempty"`
	/**
	* Initial state of the user vserver.
	*/
	State string `json:"state,omitempty"`

	//------- Read only Parameter ---------;

	Curstate string `json:"curstate,omitempty"`
	Value string `json:"value,omitempty"`
	Statechangetimesec string `json:"statechangetimesec,omitempty"`
	Statechangetimemsec string `json:"statechangetimemsec,omitempty"`
	Tickssincelaststatechange string `json:"tickssincelaststatechange,omitempty"`
	Nodefaultbindings string `json:"nodefaultbindings,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
