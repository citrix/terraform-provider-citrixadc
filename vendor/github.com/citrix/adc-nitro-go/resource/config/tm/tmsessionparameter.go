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
* Configuration for session parameter resource.
*/
type Tmsessionparameter struct {
	/**
	* Session timeout, in minutes. If there is no traffic during the timeout period, the user is disconnected and must reauthenticate to access the intranet resources.
	*/
	Sesstimeout *int `json:"sesstimeout,omitempty"`
	/**
	* Allow or deny access to content for which there is no specific authorization policy.
	*/
	Defaultauthorizationaction string `json:"defaultauthorizationaction,omitempty"`
	/**
	* Log users on to all web applications automatically after they authenticate, or pass users to the web application logon page to authenticate for each application. Note that this configuration does not honor the following authentication types for security reason. BASIC, DIGEST, and NTLM (without Negotiate NTLM2 Key or Negotiate Sign Flag). Use TM TrafficAction to configure SSO for these authentication types.
	*/
	Sso string `json:"sso,omitempty"`
	/**
	* Use primary or secondary authentication credentials for single sign-on.
	*/
	Ssocredential string `json:"ssocredential,omitempty"`
	/**
	* Domain to use for single sign-on.
	*/
	Ssodomain string `json:"ssodomain,omitempty"`
	/**
	* Kerberos constrained delegation account name
	*/
	Kcdaccount string `json:"kcdaccount,omitempty"`
	/**
	* Allow only an HTTP session cookie, in which case the cookie cannot be accessed by scripts.
	*/
	Httponlycookie string `json:"httponlycookie,omitempty"`
	/**
	* Use persistent SSO cookies for the traffic session. A persistent cookie remains on the user device and is sent with each HTTP request. The cookie becomes stale if the session ends.
	*/
	Persistentcookie string `json:"persistentcookie,omitempty"`
	/**
	* Integer specifying the number of minutes for which the persistent cookie remains valid. Can be set only if the persistence cookie setting is enabled.
	*/
	Persistentcookievalidity *int `json:"persistentcookievalidity,omitempty"`
	/**
	* Web address of the home page that a user is displayed when authentication vserver is bookmarked and used to login.
	*/
	Homepage string `json:"homepage,omitempty"`

	//------- Read only Parameter ---------;

	Name string `json:"name,omitempty"`
	Tmsessionpolicybindtype string `json:"tmsessionpolicybindtype,omitempty"`
	Tmsessionpolicycount string `json:"tmsessionpolicycount,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
