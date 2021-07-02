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
* Binding class showing the natrule that can be bound to netprofile.
*/
type Netprofilenatrulebinding struct {
	/**
	* IPv4 network address on whose traffic you want the Citrix ADC to do rewrite ip prefix.
	*/
	Natrule string `json:"natrule,omitempty"`
	Netmask string `json:"netmask,omitempty"`
	Rewriteip string `json:"rewriteip,omitempty"`
	/**
	* Name of the netprofile to which to bind port ranges.
	*/
	Name string `json:"name,omitempty"`


}