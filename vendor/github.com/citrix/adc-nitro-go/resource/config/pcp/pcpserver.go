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

package pcp

/**
* Configuration for server resource.
*/
type Pcpserver struct {
	/**
	* Name for the PCP server. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore CLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my pcpServer" or my pcpServer).
	*/
	Name string `json:"name,omitempty"`
	/**
	* The IP address of the PCP server.
	*/
	Ipaddress string `json:"ipaddress,omitempty"`
	/**
	* Port number for the PCP server.
	*/
	Port int `json:"port,omitempty"`
	/**
	* pcp profile name
	*/
	Pcpprofile string `json:"pcpprofile,omitempty"`

}
