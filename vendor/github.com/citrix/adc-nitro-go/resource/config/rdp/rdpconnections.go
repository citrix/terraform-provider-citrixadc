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

package rdp

/**
* Configuration for active rdp connections resource.
*/
type Rdpconnections struct {
	/**
	* User name for which to display connections.
	*/
	Username string `json:"username,omitempty"`
	/**
	* Terminate all active rdpconnections.
	*/
	All bool `json:"all,omitempty"`

	//------- Read only Parameter ---------;

	Endpointip string `json:"endpointip,omitempty"`
	Endpointport string `json:"endpointport,omitempty"`
	Targetip string `json:"targetip,omitempty"`
	Targetport string `json:"targetport,omitempty"`
	Peid string `json:"peid,omitempty"`

}
