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

package aaa

/**
* Configuration for AAA group resource.
*/
type Aaagroup struct {
	/**
	* Name for the group. Must begin with a letter, number, or the underscore character (_), and must consist only of letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at sign (@), equals (=), colon (:), and underscore  characters. Cannot be changed after the group is added.
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or
		single quotation marks (for example, "my aaa group" or 'my aaa group').
	*/
	Groupname string `json:"groupname,omitempty"`
	/**
	* Weight of this group with respect to other configured aaa groups (lower the number higher the weight)
	*/
	Weight int `json:"weight,omitempty"`
	/**
	* Display only the group members who are currently logged in. If there are large number of sessions, this command may provide partial details.
	*/
	Loggedin bool `json:"loggedin,omitempty"`

}
