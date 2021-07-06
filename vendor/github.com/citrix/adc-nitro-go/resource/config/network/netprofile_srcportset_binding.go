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
* Binding class showing the srcportset that can be bound to netprofile.
*/
type Netprofilesrcportsetbinding struct {
	/**
	* When the source port range is configured and associated with the netprofile bound to a service group, Citrix ADC will choose a port from the range configured for connection establishment at the backend servers.
	*/
	Srcportrange string `json:"srcportrange,omitempty"`
	/**
	* Name of the netprofile to which to bind port ranges.
	*/
	Name string `json:"name,omitempty"`


}