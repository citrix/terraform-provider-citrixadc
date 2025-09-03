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


type Authenticationloginschema struct {
	/**
	* Name for the new login schema. Login schema defines the way login form is rendered. It provides a way to customize the fields that are shown to the user. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after an action is created.
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my action" or 'my action').
	*/
	Name string `json:"name,omitempty"`
	/**
	* Name of the file for reading authentication schema to be sent for Login Page UI. This file should contain xml definition of elements as per Citrix Forms Authentication Protocol to be able to render login form. If administrator does not want to prompt users for additional credentials but continue with previously obtained credentials, then "noschema" can be given as argument. Please note that this applies only to loginSchemas that are used with user-defined factors, and not the vserver factor.
	*/
	Authenticationschema string `json:"authenticationschema,omitempty"`
	/**
	* Expression for username extraction during login. This can be any relevant advanced policy expression.
	*/
	Userexpression string `json:"userexpression,omitempty"`
	/**
	* Expression for password extraction during login. This can be any relevant advanced policy expression.
	*/
	Passwdexpression string `json:"passwdexpression,omitempty"`
	/**
	* The index at which user entered username should be stored in session.
	*/
	Usercredentialindex int `json:"usercredentialindex,omitempty"`
	/**
	* The index at which user entered password should be stored in session.
	*/
	Passwordcredentialindex int `json:"passwordcredentialindex,omitempty"`
	/**
	* Weight of the current authentication
	*/
	Authenticationstrength int `json:"authenticationstrength,omitempty"`
	/**
	* This option indicates whether current factor credentials are the default SSO (SingleSignOn) credentials.
	*/
	Ssocredentials string `json:"ssocredentials,omitempty"`

	//------- Read only Parameter ---------;

	Builtin string `json:"builtin,omitempty"`
	Feature string `json:"feature,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
