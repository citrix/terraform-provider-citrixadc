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

package authentication

/**
* Configuration for ADFSProxy Profile resource.
*/
type Authenticationadfsproxyprofile struct {
	/**
	* Name for the adfs proxy profile. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Cannot be changed after the profile is created.
		CLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my push service" or 'my push service'). 
	*/
	Name string `json:"name,omitempty"`
	/**
	* This is the name of an account in directory that would be used to authenticate trust request from ADC acting as a proxy.
	*/
	Username string `json:"username,omitempty"`
	/**
	* This is the password of an account in directory that would be used to authenticate trust request from ADC acting as a proxy.
	*/
	Password string `json:"password,omitempty"`
	/**
	* Fully qualified url of the adfs server.
	*/
	Serverurl string `json:"serverurl,omitempty"`
	/**
	* SSL certificate of the proxy that is registered at adfs server for trust.
	*/
	Certkeyname string `json:"certkeyname,omitempty"`

	//------- Read only Parameter ---------;

	Adfstruststatus string `json:"adfstruststatus,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
