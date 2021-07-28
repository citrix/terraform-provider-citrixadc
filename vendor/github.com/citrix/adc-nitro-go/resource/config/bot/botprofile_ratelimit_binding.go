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
* Binding class showing the ratelimit that can be bound to botprofile.
*/
type Botprofileratelimitbinding struct {
	/**
	* Rate-limit binding. Maximum 30 bindings can be configured per profile for rate-limit detection. For SOURCE_IP type, only one binding can be configured, and for URL type, only one binding is allowed per URL, and for SESSION type, only one binding is allowed for a cookie name. To update the values of an existing binding, user has to first unbind that binding, and then needs to bind again with new values.
	*/
	Botratelimit bool `json:"bot_ratelimit,omitempty"`
	/**
	* Rate-limiting type Following rate-limiting types are allowed:
		*SOURCE_IP - Rate-limiting based on the client IP.
		*SESSION - Rate-limiting based on the configured cookie name.
		*URL - Rate-limiting based on the configured URL.
	*/
	Botratelimittype string `json:"bot_rate_limit_type,omitempty"`
	/**
	* Enable or disable rate-limit binding.
	*/
	Botratelimitenabled string `json:"bot_rate_limit_enabled,omitempty"`
	/**
	* URL for the resource based rate-limiting.
	*/
	Botratelimiturl string `json:"bot_rate_limit_url,omitempty"`
	/**
	* Cookie name which is used to identify the session for session rate-limiting.
	*/
	Cookiename string `json:"cookiename,omitempty"`
	/**
	* Maximum number of requests that are allowed in this session in the given period time.
	*/
	Rate int `json:"rate,omitempty"`
	/**
	* Time interval during which requests are tracked to check if they cross the given rate.
	*/
	Timeslice int `json:"timeslice,omitempty"`
	/**
	* One or more actions to be taken when the current rate becomes more than the configured rate. Only LOG action can be combined with DROP, REDIRECT or RESET action.
	*/
	Botratelimitaction []string `json:"bot_rate_limit_action,omitempty"`
	/**
	* Message to be logged for this binding.
	*/
	Logmessage string `json:"logmessage,omitempty"`
	/**
	* Any comments about this binding.
	*/
	Botbindcomment string `json:"bot_bind_comment,omitempty"`
	/**
	* Name for the profile. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.), pound (#), space ( ), at (@), equals (=), colon (:), and underscore (_) characters. Cannot be changed after the profile is added.
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my profile" or 'my profile').
	*/
	Name string `json:"name,omitempty"`


}