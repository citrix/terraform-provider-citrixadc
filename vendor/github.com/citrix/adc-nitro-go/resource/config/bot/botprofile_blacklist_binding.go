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
* Binding class showing the blacklist that can be bound to botprofile.
*/
type Botprofileblacklistbinding struct {
	/**
	* Blacklist binding. Maximum 32 bindings can be configured per profile for Blacklist detection.
	*/
	Botblacklist bool `json:"bot_blacklist,omitempty"`
	/**
	* Type of the black-list entry.
	*/
	Botblacklisttype string `json:"bot_blacklist_type,omitempty"`
	/**
	* Enabled or disbaled black-list binding.
	*/
	Botblacklistenabled string `json:"bot_blacklist_enabled,omitempty"`
	/**
	* Value of the bot black-list entry.
	*/
	Botblacklistvalue string `json:"bot_blacklist_value,omitempty"`
	/**
	* One or more actions to be taken if  bot is detected based on this Blacklist binding. Only LOG action can be combined with DROP or RESET action.
	*/
	Botblacklistaction []string `json:"bot_blacklist_action,omitempty"`
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