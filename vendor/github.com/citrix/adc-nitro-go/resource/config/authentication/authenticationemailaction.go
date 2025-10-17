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
* Configuration for Email entity resource.
*/
type Authenticationemailaction struct {
	/**
	* Name for the new email action. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after an action is created.
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my action" or 'my action').
	*/
	Name string `json:"name,omitempty"`
	/**
	* Username/Clientid/EmailID to be used to authenticate to the server.
	*/
	Username string `json:"username,omitempty"`
	/**
	* Password/Clientsecret to use when authenticating to the server.
	*/
	Password string `json:"password,omitempty"`
	/**
	* Address of the server that delivers the message. It is fully qualified fqdn such as http(s):// or smtp(s):// for http and smtp protocols respectively. For SMTP, the port number is mandatory like smtps://smtp.example.com:25.
	*/
	Serverurl string `json:"serverurl,omitempty"`
	/**
	* Content to be delivered to the user. "$code" string within the content will be replaced with the actual one-time-code to be sent.
	*/
	Content string `json:"content,omitempty"`
	/**
	* This is the group that is added to user sessions that match current IdP policy. It can be used in policies to identify relying party trust.
	*/
	Defaultauthenticationgroup string `json:"defaultauthenticationgroup,omitempty"`
	/**
	* Time after which the code expires.
	*/
	Timeout *int `json:"timeout,omitempty"`
	/**
	* Type of the email action. Default type is SMTP.
	*/
	Type string `json:"type,omitempty"`
	/**
	* An optional expression that yields user's email. When not configured, user's default mail address would be used. When configured, result of this expression is used as destination email address.
	*/
	Emailaddress string `json:"emailaddress,omitempty"`

	//------- Read only Parameter ---------;

	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
