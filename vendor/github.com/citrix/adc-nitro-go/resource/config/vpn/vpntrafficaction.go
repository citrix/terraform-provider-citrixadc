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
* Configuration for VPN traffic action resource.
*/
type Vpntrafficaction struct {
	/**
	* Name for the traffic action. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after a traffic action is created.
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my action" or 'my action').
	*/
	Name string `json:"name,omitempty"`
	/**
	* Protocol, either HTTP or TCP, to be used with the action.
	*/
	Qual string `json:"qual,omitempty"`
	/**
	* Maximum amount of time, in minutes, a user can stay logged on to the web application.
	*/
	Apptimeout int `json:"apptimeout,omitempty"`
	/**
	* Provide single sign-on to the web application.
		NOTE : Authentication mechanisms like Basic-authentication  require the user credentials to be sent in plaintext which is not secure if the server is running on HTTP (instead of HTTPS).
	*/
	Sso string `json:"sso,omitempty"`
	/**
	* Provide hdx proxy to the ICA traffic
	*/
	Hdx string `json:"hdx,omitempty"`
	/**
	* Name of the form-based single sign-on profile. Form-based single sign-on allows users to log on one time to all protected applications in your network, instead of requiring them to log on separately to access each one.
	*/
	Formssoaction string `json:"formssoaction,omitempty"`
	/**
	* Specify file type association, which is a list of file extensions that users are allowed to open.
	*/
	Fta string `json:"fta,omitempty"`
	/**
	* Use the Repeater Plug-in to optimize network traffic.
	*/
	Wanscaler string `json:"wanscaler,omitempty"`
	/**
	* Kerberos constrained delegation account name
	*/
	Kcdaccount string `json:"kcdaccount,omitempty"`
	/**
	* Profile to be used for doing SAML SSO to remote relying party
	*/
	Samlssoprofile string `json:"samlssoprofile,omitempty"`
	/**
	* IP address and Port of the proxy server to be used for HTTP access for this request.
	*/
	Proxy string `json:"proxy,omitempty"`
	/**
	* expression that will be evaluated to obtain username for SingleSignOn
	*/
	Userexpression string `json:"userexpression,omitempty"`
	/**
	* expression that will be evaluated to obtain password for SingleSignOn
	*/
	Passwdexpression string `json:"passwdexpression,omitempty"`

	//------- Read only Parameter ---------;

	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
