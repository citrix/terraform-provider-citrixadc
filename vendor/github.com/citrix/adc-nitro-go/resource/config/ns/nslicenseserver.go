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
	Port int `json:"port,omitempty"`
	/**
	* If this flag is used while adding the licenseserver, existing config will be overwritten. Use this flag only if you are sure that the new licenseserver has the required capacity.
	*/
	Forceupdateip bool `json:"forceupdateip,omitempty"`
	/**
	* This paramter indicates type of license customer interested while configuring add/set licenseserver
	*/
	Licensemode string `json:"licensemode,omitempty"`
	/**
	* Username to authenticate with ADM Agent for LAS licensing. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters.
	*/
	Username string `json:"username,omitempty"`
	/**
	* Password to use when authenticating with ADM Agent for LAS licensing.
	*/
	Password string `json:"password,omitempty"`
	/**
	* Device profile is created on ADM and contains the user name and password of the instance(s). ADM will use this info to add the NS for registration
	*/
	Deviceprofilename string `json:"deviceprofilename,omitempty"`
	/**
	* Unique number that identifies the cluster node.
	*/
	Nodeid int `json:"nodeid,omitempty"`

	//------- Read only Parameter ---------;

	Status string `json:"status,omitempty"`
	Grace string `json:"grace,omitempty"`
	Gptimeleft string `json:"gptimeleft,omitempty"`
	Type string `json:"type,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
