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
* Binding class showing the pattern that can be bound to policystringmap.
*/
type Policystringmappatternbinding struct {
	/**
	* Character string constituting the key to be bound to the string map. The key is matched against the data processed by the operation that uses the string map. The default character set is ASCII. UTF-8 characters can be included if the character set is UTF-8.  UTF-8 characters can be entered directly (if the UI supports it) or can be encoded as a sequence of hexadecimal bytes '\xNN'. For example, the UTF-8 character 'ue' can be encoded as '\xC3\xBC'.
	*/
	Key string `json:"key,omitempty"`
	/**
	* Character string constituting the value associated with the key. This value is returned when processed data matches the associated key. Refer to the key parameter for details of the value character set.
	*/
	Value string `json:"value,omitempty"`
	/**
	* Comments associated with the string map or key-value pair bound to this string map.
	*/
	Comment string `json:"comment,omitempty"`
	/**
	* Name of the string map to which to bind the key-value pair.
	*/
	Name string `json:"name,omitempty"`


}