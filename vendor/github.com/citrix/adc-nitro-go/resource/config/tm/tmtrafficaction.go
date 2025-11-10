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
* Configuration for TM traffic action resource.
*/
type Tmtrafficaction struct {
	/**
	* Name for the traffic action. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after a traffic action is created.
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my action" or 'my action').
	*/
	Name string `json:"name,omitempty"`
	/**
	* Time interval, in minutes, of user inactivity after which the connection is closed.
	*/
	Apptimeout *int `json:"apptimeout,omitempty"`
	/**
	* Use single sign-on for the resource that the user is accessing now.
	*/
	Sso string `json:"sso,omitempty"`
	/**
	* Name of the configured form-based single sign-on profile.
	*/
	Formssoaction string `json:"formssoaction,omitempty"`
	/**
	* Use persistent cookies for the traffic session. A persistent cookie remains on the user device and is sent with each HTTP request. The cookie becomes stale if the session ends.
	*/
	Persistentcookie string `json:"persistentcookie,omitempty"`
	/**
	* Initiate logout for the traffic management (TM) session if the policy evaluates to true. The session is then terminated after two minutes.
	*/
	Initiatelogout string `json:"initiatelogout,omitempty"`
	/**
	* Kerberos constrained delegation account name
	*/
	Kcdaccount string `json:"kcdaccount,omitempty"`
	/**
	* Profile to be used for doing SAML SSO to remote relying party
	*/
	Samlssoprofile string `json:"samlssoprofile,omitempty"`
	/**
	* Setting to start, stop or reset TM session force timer
	*/
	Forcedtimeout string `json:"forcedtimeout,omitempty"`
	/**
	* Time interval, in minutes, for which force timer should be set.
	*/
	Forcedtimeoutval *int `json:"forcedtimeoutval,omitempty"`
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
