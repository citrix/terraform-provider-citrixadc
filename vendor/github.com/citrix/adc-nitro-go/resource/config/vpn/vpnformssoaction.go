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

package vpn

/**
* Configuration for Form sso action resource.
*/
type Vpnformssoaction struct {
	/**
	* Name for the form based single sign-on profile.
	*/
	Name string `json:"name,omitempty"`
	/**
	* Root-relative URL to which the completed form is submitted.
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
	* Expression that defines the criteria for SSO success. Expression such as checking for cookie in the response is a common example.
	*/
	Ssosuccessrule string `json:"ssosuccessrule,omitempty"`
	/**
	* Other name-value pair attributes to send to the server, in addition to sending the user name and password. Value names are separated by an ampersand (&), such as in name1=value1&name2=value2.
	*/
	Namevaluepair string `json:"namevaluepair,omitempty"`
	/**
	* Maximum number of bytes to allow in the response size. Specifies the number of bytes in the response to be parsed for extracting the forms.
	*/
	Responsesize *int `json:"responsesize,omitempty"`
	/**
	* How to process the name-value pair. Available settings function as follows:
		* STATIC - The administrator-configured values are used.
		* DYNAMIC - The response is parsed, the form is extracted, and then submitted.
	*/
	Nvtype string `json:"nvtype,omitempty"`
	/**
	* HTTP method (GET or POST) used by the single sign-on form to send the logon credentials to the logon server.
	*/
	Submitmethod string `json:"submitmethod,omitempty"`

	//------- Read only Parameter ---------;

	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
