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

package appfw

/**
* Configuration for AS settings resource.
*/
type Appfwsettings struct {
	/**
	* Profile to use when a connection does not match any policy. Default setting is APPFW_BYPASS, which sends unmatched connections back to the Citrix ADC without attempting to filter them further.
	*/
	Defaultprofile string `json:"defaultprofile,omitempty"`
	/**
	* Profile to use when an application firewall policy evaluates to undefined (UNDEF). 
		An UNDEF event indicates an internal error condition. The APPFW_BLOCK built-in profile is the default setting. You can specify a different built-in or user-created profile as the UNDEF profile.
	*/
	Undefaction string `json:"undefaction,omitempty"`
	/**
	* Timeout, in seconds, after which a user session is terminated. Before continuing to use the protected web site, the user must establish a new session by opening a designated start URL.
	*/
	Sessiontimeout uint32 `json:"sessiontimeout,omitempty"`
	/**
	* Maximum number of connections per second that the application firewall learning engine examines to generate new relaxations for learning-enabled security checks. The application firewall drops any connections above this limit from the list of connections used by the learning engine.
	*/
	Learnratelimit uint32 `json:"learnratelimit,omitempty"`
	/**
	* Maximum amount of time (in seconds) that the application firewall allows a user session to remain active, regardless of user activity. After this time, the user session is terminated. Before continuing to use the protected web site, the user must establish a new session by opening a designated start URL.
	*/
	Sessionlifetime uint32 `json:"sessionlifetime,omitempty"`
	/**
	* Name of the session cookie that the application firewall uses to track user sessions. 
		Must begin with a letter or number, and can consist of from 1 to 31 letters, numbers, and the hyphen (-) and underscore (_) symbols.
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my cookie name" or 'my cookie name').
	*/
	Sessioncookiename string `json:"sessioncookiename,omitempty"`
	/**
	* Name of an HTTP header that contains the IP address that the client used to connect to the protected web site or service.
	*/
	Clientiploggingheader string `json:"clientiploggingheader,omitempty"`
	/**
	* Cumulative total maximum number of bytes in web forms imported to a protected web site. If a user attempts to upload files with a total byte count higher than the specified limit, the application firewall blocks the request.
	*/
	Importsizelimit uint32 `json:"importsizelimit,omitempty"`
	/**
	* Flag used to enable/disable auto update signatures
	*/
	Signatureautoupdate string `json:"signatureautoupdate,omitempty"`
	/**
	* URL to download the mapping file from server
	*/
	Signatureurl string `json:"signatureurl,omitempty"`
	/**
	* String that is prepended to all encrypted cookie values.
	*/
	Cookiepostencryptprefix string `json:"cookiepostencryptprefix,omitempty"`
	/**
	* Log requests that are so malformed that application firewall parsing doesn't occur.
	*/
	Logmalformedreq string `json:"logmalformedreq,omitempty"`
	/**
	* Enable Geo-Location Logging in CEF format logs.
	*/
	Geolocationlogging string `json:"geolocationlogging,omitempty"`
	/**
	* Enable CEF format logs.
	*/
	Ceflogging string `json:"ceflogging,omitempty"`
	/**
	* Transform multibyte (double- or half-width) characters to single width characters.
	*/
	Entitydecoding string `json:"entitydecoding,omitempty"`
	/**
	* Use configurable secret key in AppFw operations
	*/
	Useconfigurablesecretkey string `json:"useconfigurablesecretkey,omitempty"`
	/**
	* Maximum number of sessions that the application firewall allows to be active, regardless of user activity. After the max_limit reaches, No more user session will be created .
	*/
	Sessionlimit uint32 `json:"sessionlimit,omitempty"`
	/**
	*  flag to define action on malformed requests that application firewall cannot parse
	*/
	Malformedreqaction []string `json:"malformedreqaction,omitempty"`
	/**
	* Flag used to enable/disable ADM centralized learning
	*/
	Centralizedlearning string `json:"centralizedlearning,omitempty"`
	/**
	* Proxy Server IP to get updated signatures from AWS.
	*/
	Proxyserver string `json:"proxyserver,omitempty"`
	/**
	* Proxy Server Port to get updated signatures from AWS.
	*/
	Proxyport int32 `json:"proxyport,omitempty"`

	//------- Read only Parameter ---------;

	Learning string `json:"learning,omitempty"`
	Builtin string `json:"builtin,omitempty"`
	Feature string `json:"feature,omitempty"`

}
