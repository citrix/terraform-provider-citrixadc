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
* Configuration for Storefront authentication action resource.
*/
type Authenticationstorefrontauthaction struct {
	/**
	* Name for the Storefront Authentication action. 
		Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the profile is created.
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my authentication action" or 'my authentication action').
	*/
	Name string `json:"name,omitempty"`
	/**
	* URL of the Storefront server. This is the FQDN of the Storefront server. example: https://storefront.com/.  Authentication endpoints are learned dynamically by Gateway.
	*/
	Serverurl string `json:"serverurl,omitempty"`
	/**
	* Domain of the server that is used for authentication. If users enter name without domain, this parameter is added to username in the authentication request to server.
	*/
	Domain string `json:"domain,omitempty"`
	/**
	* This is the default group that is chosen when the authentication succeeds in addition to extracted groups.
	*/
	Defaultauthenticationgroup string `json:"defaultauthenticationgroup,omitempty"`

	//------- Read only Parameter ---------;

	Success string `json:"success,omitempty"`
	Failure string `json:"failure,omitempty"`

}
