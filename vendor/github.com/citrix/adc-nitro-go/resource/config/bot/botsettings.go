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

package bot

/**
* Configuration for Bot engine settings resource.
*/
type Botsettings struct {
	/**
	* Profile to use when a connection does not match any policy. Default setting is " ", which sends unmatched connections back to the Citrix ADC without attempting to filter them further.
	*/
	Defaultprofile string `json:"defaultprofile,omitempty"`
	/**
	* Profile to use when the feature is not enabled but feature is licensed.
	*/
	Defaultnonintrusiveprofile string `json:"defaultnonintrusiveprofile,omitempty"`
	/**
	* Name of the JavaScript that the Bot Management feature  uses in response.
		Must begin with a letter or number, and can consist of from 1 to 31 letters, numbers, and the hyphen (-) and underscore (_) symbols.
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my cookie name" or 'my cookie name').
	*/
	Javascriptname string `json:"javascriptname,omitempty"`
	/**
	* Timeout, in seconds, after which a user session is terminated.
	*/
	Sessiontimeout int `json:"sessiontimeout,omitempty"`
	/**
	* Name of the SessionCookie that the Bot Management feature uses for tracking.
		Must begin with a letter or number, and can consist of from 1 to 31 letters, numbers, and the hyphen (-) and underscore (_) symbols.
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my cookie name" or 'my cookie name').
	*/
	Sessioncookiename string `json:"sessioncookiename,omitempty"`
	/**
	* Number of requests to allow without bot session cookie if device fingerprint is enabled
	*/
	Dfprequestlimit int `json:"dfprequestlimit,omitempty"`
	/**
	* Flag used to enable/disable bot auto update signatures
	*/
	Signatureautoupdate string `json:"signatureautoupdate,omitempty"`
	/**
	* URL to download the bot signature mapping file from server
	*/
	Signatureurl string `json:"signatureurl,omitempty"`
	/**
	* Proxy Server IP to get updated signatures from AWS.
	*/
	Proxyserver string `json:"proxyserver,omitempty"`
	/**
	* Proxy Server Port to get updated signatures from AWS.
	*/
	Proxyport int `json:"proxyport,omitempty"`
	/**
	* Enable/disable trap URL auto generation. When enabled, trap URL is updated within the configured interval.
	*/
	Trapurlautogenerate string `json:"trapurlautogenerate,omitempty"`
	/**
	* Time in seconds after which trap URL is updated.
	*/
	Trapurlinterval int `json:"trapurlinterval,omitempty"`
	/**
	* Length of the auto-generated trap URL.
	*/
	Trapurllength int `json:"trapurllength,omitempty"`

	//------- Read only Parameter ---------;

	Builtin string `json:"builtin,omitempty"`
	Feature string `json:"feature,omitempty"`

}
