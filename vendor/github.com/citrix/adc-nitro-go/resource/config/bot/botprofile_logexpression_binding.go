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
* Binding class showing the logexpression that can be bound to botprofile.
*/
type Botprofilelogexpressionbinding struct {
	/**
	* Log expression binding.
	*/
	Logexpression bool `json:"logexpression,omitempty"`
	/**
	* Name of the log expression object.
	*/
	Botlogexpressionname string `json:"bot_log_expression_name,omitempty"`
	/**
	* Expression whose result to be logged when violation happened on the bot profile.
	*/
	Botlogexpressionvalue string `json:"bot_log_expression_value,omitempty"`
	/**
	* Enable or disable the log expression binding.
	*/
	Botlogexpressionenabled string `json:"bot_log_expression_enabled,omitempty"`
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