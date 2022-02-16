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
* Configuration for Dfa authentication action resource.
*/
type Authenticationdfaaction struct {
	/**
	* Name for the DFA action. 
		Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the DFA action is added.
	*/
	Name string `json:"name,omitempty"`
	/**
	* If configured, this string is sent to the DFA server as the X-Citrix-Exchange header value.
	*/
	Clientid string `json:"clientid,omitempty"`
	/**
	* DFA Server URL
	*/
	Serverurl string `json:"serverurl,omitempty"`
	/**
	* Key shared between the DFA server and the Citrix ADC. 
		Required to allow the Citrix ADC to communicate with the DFA server.
	*/
	Passphrase string `json:"passphrase,omitempty"`
	/**
	* This is the default group that is chosen when the authentication succeeds in addition to extracted groups.
	*/
	Defaultauthenticationgroup string `json:"defaultauthenticationgroup,omitempty"`

	//------- Read only Parameter ---------;

	Success string `json:"success,omitempty"`
	Failure string `json:"failure,omitempty"`

}
