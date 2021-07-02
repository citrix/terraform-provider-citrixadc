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
* Configuration for licenseserver resource.
*/
type Nslicenseserver struct {
	/**
	* IP address of the License server.
	*/
	Licenseserverip string `json:"licenseserverip,omitempty"`
	/**
	* Fully qualified domain name of the License server.
	*/
	Servername string `json:"servername,omitempty"`
	/**
	* License server port.
	*/
	Port uint32 `json:"port,omitempty"`
	/**
	* If this flag is used while adding the licenseserver, existing config will be overwritten. Use this flag only if you are sure that the new licenseserver has the required capacity.
	*/
	Forceupdateip bool `json:"forceupdateip,omitempty"`
	/**
	* Unique number that identifies the cluster node.
	*/
	Nodeid uint32 `json:"nodeid,omitempty"`

	//------- Read only Parameter ---------;

	Status string `json:"status,omitempty"`
	Grace string `json:"grace,omitempty"`
	Gptimeleft string `json:"gptimeleft,omitempty"`

}
