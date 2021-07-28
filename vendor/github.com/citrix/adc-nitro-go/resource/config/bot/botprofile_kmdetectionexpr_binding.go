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
* Binding class showing the kmdetectionexpr that can be bound to botprofile.
*/
type Botprofilekmdetectionexprbinding struct {
	/**
	* Keyboard-mouse based detection binding. For each name, only one binding is allowed. To update the values of an existing binding, user has to first unbind that binding, then needs to bind again with new vlaues. Maximum 30 bindings can be configured per profile.
	*/
	Kmdetectionexpr bool `json:"kmdetectionexpr,omitempty"`
	/**
	* Name of the keyboard-mouse expression object.
	*/
	Botkmexpressionname string `json:"bot_km_expression_name,omitempty"`
	/**
	* JavaScript file for keyboard-mouse detection, would be inserted if the result of the expression is true.
	*/
	Botkmexpressionvalue string `json:"bot_km_expression_value,omitempty"`
	/**
	* Enable or disable the keyboard-mouse based binding.
	*/
	Botkmdetectionenabled string `json:"bot_km_detection_enabled,omitempty"`
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