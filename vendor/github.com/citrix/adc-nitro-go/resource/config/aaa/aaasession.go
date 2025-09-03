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

package aaa

/**
* Configuration for active connection resource.
*/
type Aaasession struct {
	/**
	* Name of the AAA user.
	*/
	Username string `json:"username,omitempty"`
	/**
	* Name of the AAA group.
	*/
	Groupname string `json:"groupname,omitempty"`
	/**
	* IP address or the first address in the intranet IP range.
	*/
	Iip string `json:"iip,omitempty"`
	/**
	* Subnet mask for the intranet IP range.
	*/
	Netmask string `json:"netmask,omitempty"`
	/**
	* Show aaa session associated with given session key
	*/
	Sessionkey string `json:"sessionkey,omitempty"`
	/**
	* Unique number that identifies the cluster node.
	*/
	Nodeid int `json:"nodeid,omitempty"`
	/**
	* Terminate all active AAA-TM/VPN sessions.
	*/
	All bool `json:"all,omitempty"`

	//------- Read only Parameter ---------;

	Publicip string `json:"publicip,omitempty"`
	Publicport string `json:"publicport,omitempty"`
	Ipaddress string `json:"ipaddress,omitempty"`
	Port string `json:"port,omitempty"`
	Privateip string `json:"privateip,omitempty"`
	Privateport string `json:"privateport,omitempty"`
	Destip string `json:"destip,omitempty"`
	Destport string `json:"destport,omitempty"`
	Intranetip string `json:"intranetip,omitempty"`
	Intranetip6 string `json:"intranetip6,omitempty"`
	Peid string `json:"peid,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
