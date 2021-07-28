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
* Binding class showing the trapinsertionurl that can be bound to botprofile.
*/
type Botprofiletrapinsertionurlbinding struct {
	/**
	* Bind the trap URL for the configured request URLs. Maximum 30 bindings can be configured per profile.
	*/
	Trapinsertionurl bool `json:"trapinsertionurl,omitempty"`
	/**
	* Request URL regex pattern for which Trap URL is inserted.
	*/
	Bottrapurl string `json:"bot_trap_url,omitempty"`
	/**
	* Enable or disable the request URL pattern.
	*/
	Bottrapurlinsertionenabled string `json:"bot_trap_url_insertion_enabled,omitempty"`
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
	/**
	* Message to be logged for this binding.
	*/
	Logmessage string `json:"logmessage,omitempty"`


}