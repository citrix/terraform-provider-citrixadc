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

package tm

/**
* Configuration for Form sso action resource.
*/
type Tmformssoaction struct {
	/**
	* Name for the new form-based single sign-on profile. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after an SSO action is created.
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my action" or 'my action').
	*/
	Name string `json:"name,omitempty"`
	/**
	* URL to which the completed form is submitted.
	*/
	Actionurl string `json:"actionurl,omitempty"`
	/**
	* Name of the form field in which the user types in the user ID.
	*/
	Userfield string `json:"userfield,omitempty"`
	/**
	* Name of the form field in which the user types in the password.
	*/
	Passwdfield string `json:"passwdfield,omitempty"`
	/**
	* Expression, that checks to see if single sign-on is successful.
	*/
	Ssosuccessrule string `json:"ssosuccessrule,omitempty"`
	/**
	* Name-value pair attributes to send to the server in addition to sending the username and password. Value names are separated by an ampersand (&) (for example, name1=value1&name2=value2).
	*/
	Namevaluepair string `json:"namevaluepair,omitempty"`
	/**
	* Number of bytes, in the response, to parse for extracting the forms.
	*/
	Responsesize int `json:"responsesize,omitempty"`
	/**
	* Type of processing of the name-value pair. If you specify STATIC, the values configured by the administrator are used. For DYNAMIC, the response is parsed, and the form is extracted and then submitted.
	*/
	Nvtype string `json:"nvtype,omitempty"`
	/**
	* HTTP method used by the single sign-on form to send the logon credentials to the logon server. Applies only to STATIC name-value type.
	*/
	Submitmethod string `json:"submitmethod,omitempty"`

	//------- Read only Parameter ---------;

	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
