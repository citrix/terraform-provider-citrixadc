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

package system

/**
* Configuration for system session resource.
*/
type Systemsession struct {
	/**
	* ID of the system session about which to display information.
	*/
	Sid uint32 `json:"sid,omitempty"`
	/**
	* Terminate all the system sessions except the current session.
	*/
	All bool `json:"all,omitempty"`

	//------- Read only Parameter ---------;

	Username string `json:"username,omitempty"`
	Logintime string `json:"logintime,omitempty"`
	Logintimelocal string `json:"logintimelocal,omitempty"`
	Lastactivitytime string `json:"lastactivitytime,omitempty"`
	Lastactivitytimelocal string `json:"lastactivitytimelocal,omitempty"`
	Expirytime string `json:"expirytime,omitempty"`
	Numofconnections string `json:"numofconnections,omitempty"`
	Currentconn string `json:"currentconn,omitempty"`
	Clienttype string `json:"clienttype,omitempty"`
	Partitionname string `json:"partitionname,omitempty"`

}
