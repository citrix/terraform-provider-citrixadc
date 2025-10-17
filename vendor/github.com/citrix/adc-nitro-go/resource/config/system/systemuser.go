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
* Configuration for system user resource.
*/
type Systemuser struct {
	/**
	* Name for a user. Must begin with a letter, number, or the underscore (_) character, and must contain only alphanumeric, hyphen (-), period (.), hash (#), space ( ), at (@), equal (=), colon (:), and underscore characters. Cannot be changed after the user is added.
		CLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my user" or 'my user').
	*/
	Username string `json:"username,omitempty"`
	/**
	* Password for the system user. Can include any ASCII character.
	*/
	Password string `json:"password,omitempty"`
	/**
	* Whether to use external authentication servers for the system user authentication or not
	*/
	Externalauth string `json:"externalauth,omitempty"`
	/**
	* String to display at the command-line prompt. Can consist of letters, numbers, hyphen (-), period (.), hash (#), space ( ), at (@), equal (=), colon (:), underscore (_), and the following variables:
		* %u - Will be replaced by the user name.
		* %h - Will be replaced by the hostname of the Citrix ADC.
		* %t - Will be replaced by the current time in 12-hour format.
		* %T - Will be replaced by the current time in 24-hour format.
		* %d - Will be replaced by the current date.
		* %s - Will be replaced by the state of the Citrix ADC.
		Note: The 63-character limit for the length of the string does not apply to the characters that replace the variables.
	*/
	Promptstring string `json:"promptstring,omitempty"`
	/**
	* CLI session inactivity timeout, in seconds. If Restrictedtimeout argument of system parameter is enabled, Timeout can have values in the range [300-86400] seconds. If Restrictedtimeout argument of system parameter is disabled, Timeout can have values in the range [0, 10-100000000] seconds. Default value is 900 seconds.
	*/
	Timeout *int `json:"timeout,omitempty"`
	/**
	* Users logging privilege
	*/
	Logging string `json:"logging,omitempty"`
	/**
	* Maximum number of client connection allowed per user
	*/
	Maxsession *int `json:"maxsession,omitempty"`
	/**
	* Allowed Management interfaces to the system user. By default user is allowed from both API and CLI interfaces. If management interface for a user is set to API, then user is not allowed to access NS through CLI. GUI interface will come under API interface
	*/
	Allowedmanagementinterface []string `json:"allowedmanagementinterface,omitempty"`

	//------- Read only Parameter ---------;

	Encrypted string `json:"encrypted,omitempty"`
	Hashmethod string `json:"hashmethod,omitempty"`
	Promptinheritedfrom string `json:"promptinheritedfrom,omitempty"`
	Timeoutkind string `json:"timeoutkind,omitempty"`
	Allowedmanagementinterfacekind string `json:"allowedmanagementinterfacekind,omitempty"`
	Lastpwdchangetimestamp string `json:"lastpwdchangetimestamp,omitempty"`
	Daystoexpirekind string `json:"daystoexpirekind,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
