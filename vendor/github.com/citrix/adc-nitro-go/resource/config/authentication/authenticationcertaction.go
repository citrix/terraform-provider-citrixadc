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
* Configuration for CERT action resource.
*/
type Authenticationcertaction struct {
	/**
	* Name for the client cert authentication server profile (action).
		Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after certifcate action is created.
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my authentication action" or 'my authentication action').
	*/
	Name string `json:"name,omitempty"`
	/**
	* Enables or disables two-factor authentication.
		Two factor authentication is client cert authentication followed by password authentication.
	*/
	Twofactor string `json:"twofactor,omitempty"`
	/**
	* Client-cert field from which the username is extracted. Must be set to either ""Subject"" and ""Issuer"" (include both sets of double quotation marks).
		Format: <field>:<subfield>.
	*/
	Usernamefield string `json:"usernamefield,omitempty"`
	/**
	* Client-cert field from which the group is extracted.  Must be set to either ""Subject"" and ""Issuer"" (include both sets of double quotation marks).
		Format: <field>:<subfield>
	*/
	Groupnamefield string `json:"groupnamefield,omitempty"`
	/**
	* This is the default group that is chosen when the authentication succeeds in addition to extracted groups.
	*/
	Defaultauthenticationgroup string `json:"defaultauthenticationgroup,omitempty"`

	//------- Read only Parameter ---------;

	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
