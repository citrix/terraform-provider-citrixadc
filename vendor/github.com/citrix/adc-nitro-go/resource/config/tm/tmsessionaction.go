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
* Configuration for TM session action resource.
*/
type Tmsessionaction struct {
	/**
	* Name for the session action. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after a session action is created.
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my action" or 'my action').
	*/
	Name string `json:"name,omitempty"`
	/**
	* Session timeout, in minutes. If there is no traffic during the timeout period, the user is disconnected and must reauthenticate to access intranet resources.
	*/
	Sesstimeout int `json:"sesstimeout,omitempty"`
	/**
	* Allow or deny access to content for which there is no specific authorization policy.
	*/
	Defaultauthorizationaction string `json:"defaultauthorizationaction,omitempty"`
	/**
	* Use single sign-on (SSO) to log users on to all web applications automatically after they authenticate, or pass users to the web application logon page to authenticate to each application individually. Note that this configuration does not honor the following authentication types for security reason. BASIC, DIGEST, and NTLM (without Negotiate NTLM2 Key or Negotiate Sign Flag). Use TM TrafficAction to configure SSO for these authentication types.
	*/
	Sso string `json:"sso,omitempty"`
	/**
	* Use the primary or secondary authentication credentials for single sign-on (SSO).
	*/
	Ssocredential string `json:"ssocredential,omitempty"`
	/**
	* Domain to use for single sign-on (SSO).
	*/
	Ssodomain string `json:"ssodomain,omitempty"`
	/**
	* Allow only an HTTP session cookie, in which case the cookie cannot be accessed by scripts.
	*/
	Httponlycookie string `json:"httponlycookie,omitempty"`
	/**
	* Kerberos constrained delegation account name
	*/
	Kcdaccount string `json:"kcdaccount,omitempty"`
	/**
	* Enable or disable persistent SSO cookies for the traffic management (TM) session. A persistent cookie remains on the user device and is sent with each HTTP request. The cookie becomes stale if the session ends. This setting is overwritten if a traffic action sets persistent cookie to OFF. 
		Note: If persistent cookie is enabled, make sure you set the persistent cookie validity.
	*/
	Persistentcookie string `json:"persistentcookie,omitempty"`
	/**
	* Integer specifying the number of minutes for which the persistent cookie remains valid. Can be set only if the persistent cookie setting is enabled.
	*/
	Persistentcookievalidity int `json:"persistentcookievalidity,omitempty"`
	/**
	* Web address of the home page that a user is displayed when authentication vserver is bookmarked and used to login.
	*/
	Homepage string `json:"homepage,omitempty"`

	//------- Read only Parameter ---------;

	Builtin string `json:"builtin,omitempty"`
	Feature string `json:"feature,omitempty"`

}
