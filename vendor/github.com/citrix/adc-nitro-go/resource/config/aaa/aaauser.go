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

package aaa

/**
* Configuration for AAA user resource.
*/
type Aaauser struct {
	/**
	* Name for the user. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the user is added.
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or
		single quotation marks (for example, "my aaa user" or "my aaa user").
	*/
	Username string `json:"username,omitempty"`
	/**
	* Password with which the user logs on. Required for any user account that does not exist on an external authentication server. 
		If you are not using an external authentication server, all user accounts must have a password. If you are using an external authentication server, you must provide a password for local user accounts that do not exist on the authentication server.
	*/
	Password string `json:"password,omitempty"`
	/**
	* Show whether the user is logged in or not.
	*/
	Loggedin bool `json:"loggedin,omitempty"`

}
