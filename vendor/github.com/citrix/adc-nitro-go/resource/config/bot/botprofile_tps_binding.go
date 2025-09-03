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
* Binding class showing the tps that can be bound to botprofile.
*/
type Botprofiletpsbinding struct {
	/**
	* TPS binding. For each type only binding can be configured. To  update the values of an existing binding, user has to first unbind that binding, and then needs to bind again with new values.
	*/
	Bottps bool `json:"bot_tps,omitempty"`
	/**
	* Type of TPS binding.
	*/
	Bottpstype string `json:"bot_tps_type,omitempty"`
	/**
	* Maximum number of requests that are allowed from (or to) a IP, Geolocation, URL or Host in 1 second time interval.
	*/
	Threshold int `json:"threshold,omitempty"`
	/**
	* Maximum percentage increase in the requests from (or to) a IP, Geolocation, URL or Host in 30 minutes interval.
	*/
	Percentage int `json:"percentage,omitempty"`
	/**
	* One to more actions to be taken if bot is detected based on this TPS binding. Only LOG action can be combined with DROP, RESET, REDIRECT, or MITIGIATION action.
	*/
	Bottpsaction []string `json:"bot_tps_action,omitempty"`
	/**
	* Message to be logged for this binding.
	*/
	Logmessage string `json:"logmessage,omitempty"`
	/**
	* Any comments about this binding.
	*/
	Botbindcomment string `json:"bot_bind_comment,omitempty"`
	/**
	* Enabled or disabled TPS binding.
	*/
	Bottpsenabled string `json:"bot_tps_enabled,omitempty"`
	/**
	* Name for the profile. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.), pound (#), space ( ), at (@), equals (=), colon (:), and underscore (_) characters. Cannot be changed after the profile is added.
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my profile" or 'my profile').
	*/
	Name string `json:"name,omitempty"`


}