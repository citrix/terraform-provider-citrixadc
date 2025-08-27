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
* Configuration for Authentication profile resource.
*/
type Authenticationauthnprofile struct {
	/**
	* Name for the authentication profile.
		Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the RADIUS action is added.
	*/
	Name string `json:"name,omitempty"`
	/**
	* Name of the authentication vserver at which authentication should be done.
	*/
	Authnvsname string `json:"authnvsname,omitempty"`
	/**
	* Hostname of the authentication vserver to which user must be redirected for authentication.
	*/
	Authenticationhost string `json:"authenticationhost,omitempty"`
	/**
	* Domain for which TM cookie must to be set. If unspecified, cookie will be set for FQDN.
	*/
	Authenticationdomain string `json:"authenticationdomain,omitempty"`
	/**
	* Authentication weight or level of the vserver to which this will bound. This is used to order TM vservers based on the protection required. A session that is created by authenticating against TM vserver at given level cannot be used to access TM vserver at a higher level.
	*/
	Authenticationlevel int `json:"authenticationlevel,omitempty"`

	//------- Read only Parameter ---------;

	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
