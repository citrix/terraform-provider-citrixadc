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

package policy

/**
* Binding class showing the pattern that can be bound to policypatset.
*/
type Policypatsetpatternbinding struct {
	/**
	* String of characters that constitutes a pattern. For more information about the characters that can be used, refer to the character set parameter.
		Note: Minimum length for pattern sets used in rewrite actions of type REPLACE_ALL, DELETE_ALL, INSERT_AFTER_ALL, and INSERT_BEFORE_ALL, is three characters.
	*/
	String string `json:"String,omitempty"`
	/**
	* The index of the string associated with the patset.
	*/
	Index uint32 `json:"index,omitempty"`
	/**
	* Character set associated with the characters in the string.
		Note: UTF-8 characters can be entered directly (if the UI supports it) or can be encoded as a sequence of hexadecimal bytes '\xNN'. For example, the UTF-8 character 'ue' can be encoded as '\xC3\xBC'.
	*/
	Charset string `json:"charset,omitempty"`
	/**
	* Any comments to preserve information about this patset or a pattern bound to this patset.
	*/
	Comment string `json:"comment,omitempty"`
	/**
	* Indicates that a variable is a built-in (SYSTEM INTERNAL) type.
	*/
	Builtin []string `json:"builtin,omitempty"`
	/**
	* The feature to be checked while applying this config
	*/
	Feature string `json:"feature,omitempty"`
	/**
	* Name of the pattern set to which to bind the string.
	*/
	Name string `json:"name,omitempty"`


}